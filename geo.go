package redis

import "github.com/go-redis/redis"

// Geo represents a Redis GeoHash structure.
type Geo struct {
	client *redis.Client
}

// GeoLocation is used with GeoAdd to add geospatial location.
type GeoLocation = redis.GeoLocation

// GeoPos is used with GeoPos to add geospatial position.
type GeoPos = redis.GeoPos

// GeoRadiusQuery is used with GeoRadius to query geospatial index.
type GeoRadiusQuery = redis.GeoRadiusQuery

// GeoAdd add one or more geospatial items in the geospatial index represented using a sorted set
func (g *Geo) GeoAdd(key string, geoLocation ...*GeoLocation) (int64, error) {
	resp := g.client.GeoAdd(key, geoLocation...)
	return resp.Result()
}

// GeoDist returns the distance between two members of a geospatial index
func (g *Geo) GeoDist(key string, member1, member2, unit string) (float64, error) {
	resp := g.client.GeoDist(key, member1, member2, unit)
	return resp.Result()
}

// GeoHash returns members of a geospatial index as standard geohash strings
func (g *Geo) GeoHash(key string, members ...string) ([]string, error) {
	resp := g.client.GeoHash(key, members...)
	return resp.Result()
}

// GeoPos returns longitude and latitude of members of a geospatial index
func (g *Geo) GeoPos(key string, members ...string) ([]*GeoPos, error) {
	resp := g.client.GeoPos(key, members...)
	return resp.Result()
}

// GeoRadius query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point
func (g *Geo) GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error) {
	resp := g.client.GeoRadius(key, longitude, latitude, query)
	return resp.Result()
}

// GeoRadiusByMember query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member
func (g *Geo) GeoRadiusByMember(key, member string) ([]GeoLocation, error) {
	resp := g.client.GeoRadiusByMember(key, member, nil)
	return resp.Result()
}
