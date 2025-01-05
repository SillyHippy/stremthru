//go:build heavy

package main

import (
	"context"
	"embed"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/SillyHippy/stremthru/internal/db"
)

//go:embed schema.hcl schema.postgres.hcl
var schemaFS embed.FS

func RunSchemaMigration(uri db.ConnectionURI) {
	dsnModifiers := []db.DSNModifier{}
	filename := "schema.hcl"
	switch uri.Dialect {
	case "sqlite":
		dsnModifiers = append(dsnModifiers, func(u *url.URL, q *url.Values) {
			if u.Host != "." {
				return
			}

			if dir, err := os.Getwd(); err == nil {
				u.Host = ""
				u.Path = filepath.Join(dir, u.Path)
				return
			}

			if executable, err := os.Executable(); err == nil {
				u.Host = ""
				u.Path = filepath.Join(filepath.Dir(executable), u.Path)
			}
		})
	case "postgres":
		filename = "schema.postgres.hcl"
	}

	workdir, err := atlasexec.NewWorkingDir(
		func(dir *atlasexec.WorkingDir) error {
			return dir.CreateFile("schema.hcl", func(w io.Writer) error {
				f, err := schemaFS.Open(filename)
				if err != nil {
					return err
				}
				defer f.Close()
				_, err = io.Copy(w, f)
				return err
			})
		},
	)
	if err != nil {
		log.Fatalf("[schema] failed to create directory: %v\n", err)
	}
	defer workdir.Close()

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		log.Fatalf("[schema] failed to initialize: %v\n", err)
	}
	dsn := uri.DSN(dsnModifiers...)
	res, err := client.SchemaApply(context.Background(), &atlasexec.SchemaApplyParams{
		AutoApprove: true,
		To:          "file://schema.hcl",
		URL:         dsn,
	})
	if err != nil {
		log.Fatalf("[schema] failed to run migration: %v\n", err)
	}
	if res.Changes.Error != nil {
		log.Printf("[schema] ERROR:\n\n%s\n\n%s\n\n", res.Changes.Error.Stmt, res.Changes.Error.Text)
	}
	if len(res.Changes.Applied) > 0 {
		log.Printf("[schema] APPLIED:\n\n%s\n\n", strings.Join(res.Changes.Applied, "\n"))
	}
	if len(res.Changes.Pending) > 0 {
		log.Printf("[schema] PENDING:\n\n%s\n\n", strings.Join(res.Changes.Pending, "\n"))
	}
	log.Printf("[schema] migration done\n")
}
