package common

import "math"

// haversine formula to calculate the distance between two points on the Earth
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	var (
		rad  = math.Pi / 180
		R    = 6371e3 // Earth's radius in meters
		dLat = (lat2 - lat1) * rad
		dLon = (lon2 - lon1) * rad
		a    = math.Sin(dLat/2)*math.Sin(dLat/2) +
			math.Cos(lat1*rad)*math.Cos(lat2*rad)*
				math.Sin(dLon/2)*math.Sin(dLon/2)
		c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	)
	return R * c // Distance in meters
}
