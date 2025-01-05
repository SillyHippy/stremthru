package buddy

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/SillyHippy/stremthru/core"
	"github.com/SillyHippy/stremthru/internal/config"
	"github.com/SillyHippy/stremthru/internal/magnet_cache"
	"github.com/SillyHippy/stremthru/internal/peer"
	"github.com/SillyHippy/stremthru/store"
)

var Buddy = NewAPIClient(&APIClientConfig{
	BaseURL: config.BuddyURL,
})

var Peer = peer.NewAPIClient(&peer.APIClientConfig{
	BaseURL: config.PeerURL,
	APIKey:  config.PeerAuthToken,
})

func TrackMagnet(s store.Store, hash string, files []store.MagnetFile, sid string, cacheMiss bool, storeToken string) {
	mcFiles := magnet_cache.Files{}
	if !cacheMiss {
		for _, f := range files {
			mcFiles = append(mcFiles, magnet_cache.File{Idx: f.Idx, Name: f.Name, Size: f.Size})
		}
	}
	magnet_cache.Touch(s.GetName().Code(), hash, mcFiles, sid)

	if config.HasBuddy {
		params := &TrackMagnetCacheParams{
			Store:     s.GetName(),
			Hash:      hash,
			Files:     files,
			CacheMiss: cacheMiss,
			SId:       sid,
		}
		start := time.Now()
		if _, err := Buddy.TrackMagnetCache(params); err != nil {
			core.LogError(fmt.Sprintf("[buddy] failed to track magnet cache for %s:%s, took %v", s.GetName(), hash, time.Since(start)), err)
		} else {
			log.Printf("[buddy] track magnet cache for %s:%s, took %v", s.GetName(), hash, time.Since(start))
		}
	}

	if config.HasPeer {
		params := &peer.TrackMagnetParams{
			StoreName:  s.GetName(),
			StoreToken: storeToken,
			Hash:       hash,
			Files:      files,
			IsMiss:     cacheMiss,
			SId:        sid,
		}
		start := time.Now()
		if _, err := Peer.TrackMagnet(params); err != nil {
			core.LogError(fmt.Sprintf("[buddy:upstream] failed to track magnet cache for %s:%s, took %v", s.GetName(), hash, time.Since(start)), err)
		} else {
			log.Printf("[buddy:upstream] track magnet cache for %s:%s, took %v", s.GetName(), hash, time.Since(start))
		}
	}
}

func CheckMagnet(s store.Store, hashes []string, storeToken string, clientIp string, sid string) (*store.CheckMagnetData, error) {
	if matched, err := regexp.MatchString("^tt[0-9]+(:[0-9]{1,2}:[0-9]{1,3})?$", sid); err != nil || !matched {
		sid = ""
	}

	data := &store.CheckMagnetData{
		Items: []store.CheckMagnetDataItem{},
	}

	mcs, err := magnet_cache.GetByHashes(s.GetName().Code(), hashes, sid)
	if err != nil {
		return nil, err
	}
	mcByHash := map[string]magnet_cache.MagnetCache{}
	for _, mc := range mcs {
		mcByHash[mc.Hash] = mc
	}

	magnetByHash := map[string]core.MagnetLink{}

	staleOrMissingHashes := []string{}
	for _, hash := range hashes {
		magnet, err := core.ParseMagnetLink(hash)
		if err != nil {
			continue
		}
		magnetByHash[magnet.Hash] = magnet
		if mc, ok := mcByHash[magnet.Hash]; ok && !mc.IsStale() {
			item := store.CheckMagnetDataItem{
				Hash:   magnet.Hash,
				Magnet: magnet.Link,
				Status: store.MagnetStatusUnknown,
				Files:  []store.MagnetFile{},
			}
			if mc.IsCached {
				item.Status = store.MagnetStatusCached
				item.Files = mc.Files.ToStoreMagnetFile()
			}
			data.Items = append(data.Items, item)
		} else {
			staleOrMissingHashes = append(staleOrMissingHashes, magnet.Hash)
		}
	}

	if len(staleOrMissingHashes) == 0 {
		return data, nil
	}

	if config.HasBuddy {
		params := &CheckMagnetCacheParams{
			Store:    s.GetName(),
			Hashes:   staleOrMissingHashes,
			ClientIP: clientIp,
		}
		params.SId = sid
		start := time.Now()
		res, err := Buddy.CheckMagnetCache(params)
		duration := time.Since(start)
		if err != nil {
			core.LogError(fmt.Sprintf("[buddy] failed to check magnet, took %v", duration), err)
			return data, nil
		} else {
			log.Printf("[buddy] check manget, took %v\n", duration)
			filesByHash := map[string]magnet_cache.Files{}
			for _, item := range res.Data.Items {
				res_item := store.CheckMagnetDataItem{
					Hash:   item.Hash,
					Magnet: item.Magnet,
					Status: item.Status,
				}
				res_files := []store.MagnetFile{}
				files := magnet_cache.Files{}
				if item.Status == store.MagnetStatusCached {
					for _, f := range item.Files {
						res_files = append(res_files, store.MagnetFile{Idx: f.Idx, Name: f.Name, Size: f.Size})
						files = append(files, magnet_cache.File{Idx: f.Idx, Name: f.Name, Size: f.Size, SId: f.SId})
					}
				}
				res_item.Files = res_files
				data.Items = append(data.Items, res_item)
				filesByHash[item.Hash] = files
			}
			magnet_cache.BulkTouch(s.GetName().Code(), filesByHash, sid)
			return data, nil
		}
	}

	if config.HasPeer {
		params := &peer.CheckMagnetParams{
			StoreName:  s.GetName(),
			StoreToken: storeToken,
		}
		params.Magnets = hashes
		params.ClientIP = clientIp
		params.SId = sid
		start := time.Now()
		res, err := Peer.CheckMagnet(params)
		duration := time.Since(start)
		if err != nil {
			core.LogError(fmt.Sprintf("[buddy:upstream] failed to check magnet, took %v", duration), err)
			return data, nil
		} else {
			log.Printf("[buddy:upstream] check manget, took %v\n", duration)
			filesByHash := map[string]magnet_cache.Files{}
			for _, item := range res.Data.Items {
				files := magnet_cache.Files{}
				if item.Status == store.MagnetStatusCached {
					for _, f := range item.Files {
						files = append(files, magnet_cache.File{Idx: f.Idx, Name: f.Name, Size: f.Size})
					}
				}
				filesByHash[item.Hash] = files
				data.Items = append(data.Items, item)
			}
			magnet_cache.BulkTouch(s.GetName().Code(), filesByHash, sid)
			return data, nil
		}
	}

	for _, hash := range staleOrMissingHashes {
		magnet := magnetByHash[hash]
		item := store.CheckMagnetDataItem{
			Hash:   magnet.Hash,
			Magnet: magnet.Link,
			Status: store.MagnetStatusUnknown,
			Files:  []store.MagnetFile{},
		}
		data.Items = append(data.Items, item)
	}
	return data, nil
}
