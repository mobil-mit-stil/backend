package algorithm

import (
	"backend/storage"
	"math"
)

// the following distance function is heavily inspired by the following gist:
// https://gist.github.com/hotdang-ca/6c1ee75c48e515aec5bc6db6e3265e49
// it has been modified to fit our use case

// the original notice:
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:::                                                                         :::
//:::  This routine calculates the distance between two points (given the     :::
//:::  latitude/longitude of those points). It is based on free code used to  :::
//:::  calculate the distance between two locations using GeoDataSource(TM)   :::
//:::  products.                                                              :::
//:::                                                                         :::
//:::  Definitions:                                                           :::
//:::    South latitudes are negative, east longitudes are positive           :::
//:::                                                                         :::
//:::  Passed to function:                                                    :::
//:::    lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)  :::
//:::    lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)  :::
//:::    optional: unit = the unit you desire for results                     :::
//:::           where: 'M' is statute miles (default, or omitted)             :::
//:::                  'K' is kilometers                                      :::
//:::                  'N' is nautical miles                                  :::
//:::                                                                         :::
//:::  Worldwide cities and other features databases with latitude longitude  :::
//:::  are available at https://www.geodatasource.com                         :::
//:::                                                                         :::
//:::  For enquiries, please contact sales@geodatasource.com                  :::
//:::                                                                         :::
//:::  Official Web site: https://www.geodatasource.com                       :::
//:::                                                                         :::
//:::          Golang code James Robert Perih (c) All Rights Reserved 2018    :::
//:::                                                                         :::
//:::           GeoDataSource.com (C) All Rights Reserved 2017                :::
//:::                                                                         :::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	radLat1 := float64(math.Pi * lat1 / 180)
	radLat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radTheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radLat1)*math.Sin(radLat2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Cos(radTheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	dist = dist * 1609.344

	return dist
}

func NearestPoint(location storage.LocationLongLat, route []storage.LocationLongLat) (storage.LocationLongLat, float64) {
	var dist float64
	dist = math.MaxFloat64
	var bestPoint storage.LocationLongLat
	for _, point := range route {
		newDist := distance(location.Lat, location.Long, point.Lat, point.Long)
		if newDist < dist {
			dist = newDist
			bestPoint = point
		}
	}
	return bestPoint, dist
}

func NearestTolerablePoint(location storage.LocationLongLat, route []storage.LocationLongLat, tolerance int32) (storage.LocationLongLat, bool) {
	bestPoint, dist := NearestPoint(location, route)
	return bestPoint, int32(dist) < tolerance
}
