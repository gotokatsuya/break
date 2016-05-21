package foursquare

import (
	"encoding/gob"
	"fmt"

	"github.com/evalphobia/eurekache"
	"github.com/evalphobia/eurekache/memorycache"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/service/venues"

	"github.com/gotokatsuya/break/config"
)

var cache *eurekache.Eurekache

func init() {
	gob.Register(&venues.ExploreResponse{})
	initCache()
}

func initCache() {
	maxCacheItemSize := 1000     // max allocated caches
	expiredTTL := 60 * 60 * 1000 // 60 minutes (millisecond)
	mc := memorycache.NewCacheTTL(maxCacheItemSize)
	mc.SetTTL(int64(expiredTTL))

	cache = eurekache.New()
	cache.SetCacheSources([]eurekache.Cache{mc})
}

func ClearCache() {
	initCache()
}

func GetExplore(lat, lng string) (*venues.ExploreResponse, error) {
	latLng := fmt.Sprintf("%s,%s", lat, lng)
	cachedRes := new(venues.ExploreResponse)
	if ok := cache.Get(latLng, cachedRes); ok {
		return cachedRes, nil
	}
	foursquareConfig := config.GetFoursquare()
	client := dispatcher.NewClientWithParam(foursquareConfig.ClientID, foursquareConfig.ClientSecret)
	req := venues.NewExploreRequest()
	req.LatLng = latLng
	res, err := venues.Explore(client, req)
	if err != nil {
		return nil, err
	}
	cache.Set(latLng, res)
	return res, nil
}
