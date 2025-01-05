package easydebrid

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/store"
)

type StoreClientConfig struct {
	HTTPClient *http.Client
	UserAgent  string
}

type StoreClient struct {
	Name   store.StoreName
	client *APIClient
}

func NewStoreClient(config *StoreClientConfig) *StoreClient {
	c := &StoreClient{}
	c.client = NewAPIClient(&APIClientConfig{
		HTTPClient: config.HTTPClient,
		UserAgent:  config.UserAgent,
	})
	c.Name = store.StoreNameEasyDebrid

	return c
}

func (s *StoreClient) GetName() store.StoreName {
	return s.Name
}

type LockedFileLink string

const lockedFileLinkPrefix = "stremthru://store/easydebrid/"

func (l LockedFileLink) encodeData(magnetHash string, fileIdx int) string {
	return core.Base64Encode(magnetHash + ":" + strconv.Itoa(fileIdx))
}

func (l LockedFileLink) decodeData(encoded string) (magnetHash string, fileIdx int, err error) {
	decoded, err := core.Base64Decode(encoded)
	if err != nil {
		return "", 0, err
	}
	magnetHash, fIdx, found := strings.Cut(decoded, ":")
	if !found {
		return "", 0, err
	}
	fileIdx, err = strconv.Atoi(fIdx)
	if err != nil {
		return "", 0, err
	}
	return magnetHash, fileIdx, nil
}

func (l LockedFileLink) create(magnetHash string, fileIdx int) string {
	return lockedFileLinkPrefix + l.encodeData(magnetHash, fileIdx)
}

func (l LockedFileLink) parse() (magnetHash string, fileIdx int, err error) {
	encoded := strings.TrimPrefix(string(l), lockedFileLinkPrefix)
	return l.decodeData(encoded)
}

const magnetIdPrefix = "st:ed:"

func (s *StoreClient) AddMagnet(params *store.AddMagnetParams) (*store.AddMagnetData, error) {
	magnet, err := core.ParseMagnetLink(params.Magnet)
	if err != nil {
		return nil, err
	}
	res, err := s.client.LookupLinkDetails(&LookupLinkDetailsParams{
		Ctx:  params.Ctx,
		URLs: []string{magnet.Link},
	})
	if err != nil {
		return nil, err
	}
	data := &store.AddMagnetData{
		Id:      magnetIdPrefix + magnet.Hash,
		Hash:    magnet.Hash,
		Magnet:  magnet.Link,
		Name:    magnet.Name,
		Status:  store.MagnetStatusUnknown,
		Files:   []store.MagnetFile{},
		AddedAt: time.Now(),
	}
	if len(res.Data.Result) < 1 {
		return data, nil
	}
	detail := res.Data.Result[0]
	if detail.Cached {
		data.Status = store.MagnetStatusDownloaded
		for idx, f := range detail.Files {
			if core.HasVideoExtension(f.Name) {
				data.Files = append(data.Files, store.MagnetFile{
					Idx:  idx,
					Link: LockedFileLink("").create(magnet.Hash, idx),
					Name: f.Name,
					Path: "/" + f.Folder + "/" + f.Name,
					Size: f.Size,
				})
			}
		}
	}
	return data, nil
}

func (s *StoreClient) CheckMagnet(params *store.CheckMagnetParams) (*store.CheckMagnetData, error) {
	magnets := []core.MagnetLink{}
	links := []string{}
	for _, m := range params.Magnets {
		magnet, err := core.ParseMagnetLink(m)
		if err != nil {
			return nil, err
		}
		magnets = append(magnets, magnet)
		links = append(links, magnet.Link)
	}
	res, err := s.client.LookupLinkDetails(&LookupLinkDetailsParams{
		Ctx:  params.Ctx,
		URLs: links,
	})
	if err != nil {
		return nil, err
	}
	data := &store.CheckMagnetData{
		Items: []store.CheckMagnetDataItem{},
	}
	for i, detail := range res.Data.Result {
		magnet := magnets[i]
		item := store.CheckMagnetDataItem{
			Hash:   magnet.Hash,
			Magnet: magnet.Link,
			Status: store.MagnetStatusUnknown,
			Files:  []store.MagnetFile{},
		}
		if detail.Cached {
			item.Status = store.MagnetStatusCached
			for idx, f := range detail.Files {
				item.Files = append(item.Files, store.MagnetFile{
					Idx:  idx,
					Name: f.Name,
					Size: f.Size,
				})
			}
		}
		data.Items = append(data.Items, item)
	}
	return data, nil
}

func (s *StoreClient) GenerateLink(params *store.GenerateLinkParams) (*store.GenerateLinkData, error) {
	magnetHash, fileIdx, err := LockedFileLink(params.Link).parse()
	if err != nil {
		return nil, err
	}
	magnet, err := core.ParseMagnetLink(magnetHash)
	if err != nil {
		return nil, err
	}
	res, err := s.client.GenerateLink(&GenerateLinkParams{
		Ctx:      params.Ctx,
		URL:      magnet.Link,
		ClientIP: params.ClientIP,
	})
	if err != nil {
		return nil, err
	}
	if len(res.Data.Files) < fileIdx+1 {
		err := UpstreamErrorWithCause(&ResponseContainer{
			Err: string(core.ErrorCodeNotFound),
		})
		err.StatusCode = http.StatusNotFound
		return nil, err
	}
	file := res.Data.Files[fileIdx]
	data := &store.GenerateLinkData{
		Link: file.URL,
	}
	return data, nil
}

func (s *StoreClient) GetMagnet(params *store.GetMagnetParams) (*store.GetMagnetData, error) {
	magnet, err := core.ParseMagnetLink(strings.TrimPrefix(params.Id, magnetIdPrefix))
	if err != nil {
		return nil, err
	}
	res, err := s.client.LookupLinkDetails(&LookupLinkDetailsParams{
		Ctx:  params.Ctx,
		URLs: []string{magnet.Link},
	})
	if err != nil {
		return nil, err
	}
	if len(res.Data.Result) < 1 || !res.Data.Result[0].Cached {
		err := UpstreamErrorWithCause(&ResponseContainer{
			Err: string(core.ErrorCodeNotFound),
		})
		err.StatusCode = http.StatusNotFound
		return nil, err
	}
	detail := res.Data.Result[0]
	data := &store.GetMagnetData{
		Id:      params.Id,
		Hash:    magnet.Hash,
		Name:    magnet.Name,
		Status:  store.MagnetStatusDownloaded,
		Files:   []store.MagnetFile{},
		AddedAt: time.Now(),
	}
	for idx, f := range detail.Files {
		if core.HasVideoExtension(f.Name) {
			data.Files = append(data.Files, store.MagnetFile{
				Idx:  idx,
				Link: LockedFileLink("").create(magnet.Hash, idx),
				Name: f.Name,
				Path: "/" + f.Folder + "/" + f.Name,
				Size: f.Size,
			})
		}
	}
	return data, nil
}

func (s *StoreClient) GetUser(params *store.GetUserParams) (*store.User, error) {
	res, err := s.client.GetUserDetails(&GetUserDetailsParams{
		Ctx: params.Ctx,
	})
	if err != nil {
		return nil, err
	}
	data := &store.User{
		Id:                 string(res.Data.Id),
		Email:              "",
		SubscriptionStatus: store.UserSubscriptionStatusExpired,
	}
	if res.Data.PaidUntil > time.Now().Unix() {
		data.SubscriptionStatus = store.UserSubscriptionStatusPremium
	}
	return data, nil
}

func (s *StoreClient) ListMagnets(params *store.ListMagnetsParams) (*store.ListMagnetsData, error) {
	return &store.ListMagnetsData{
		Items:      []store.ListMagnetsDataItem{},
		TotalItems: 0,
	}, nil
}

func (s *StoreClient) RemoveMagnet(params *store.RemoveMagnetParams) (*store.RemoveMagnetData, error) {
	return &store.RemoveMagnetData{
		Id: params.Id,
	}, nil
}
