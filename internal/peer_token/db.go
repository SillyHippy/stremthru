package peer_token

import (
	"github.com/SillyHippy/stremthru/internal/db"
)

const TableName = "peer_token"

type PeerToken struct {
	Id        string
	Name      string
	CreatedAt db.Timestamp
}

func IsValid(token string) (bool, error) {
	if token == "" {
		return false, nil
	}

	id := ""
	row := db.QueryRow("SELECT id FROM "+TableName+" WHERE id = ?", token)
	if err := row.Scan(&id); err != nil {
		return false, err
	}

	return id == token, nil
}
