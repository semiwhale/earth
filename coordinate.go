package earth

import (
	"errors"
	"strconv"
	"strings"
	"fmt"
)

// GCJLocation ...
type GCJLocation struct {
	Latitude  float64
	Longitude float64
}

// ToWGS ...
func (l GCJLocation) ToWGS() WGSLocation {
	lat, lon := GCJtoWGS(l.Latitude, l.Longitude)
	return WGSLocation{
		Latitude:  lat,
		Longitude: lon,
	}
}

// ToWGSExact ...
func (l GCJLocation) ToWGSExact() WGSLocation {
	lat, lon := GCJtoWGSExact(l.Latitude, l.Longitude)
	return WGSLocation{
		Latitude:  lat,
		Longitude: lon,
	}
}

// WGSLocation ...
type WGSLocation struct {
	Latitude  float64
	Longitude float64
}

// ToGCJ ...
func (l WGSLocation) ToGCJ() GCJLocation {
	lat, lon := WGStoGCJ(l.Latitude, l.Longitude)
	return GCJLocation{
		Latitude:  lat,
		Longitude: lon,
	}
}

// ParseCoordinate ...
func ParseCoordinate(content string) (float64, float64, error) {
	a := strings.SplitN(content, ",", 2)
	if len(a) != 2 {
		return 0, 0, errors.New("invalid coordinate")
	}

	lon, err := strconv.ParseFloat(a[0], 64)
	if err != nil {
		return 0, 0, err
	}

	lat, err := strconv.ParseFloat(a[1], 64)
	if err != nil {
		return 0, 0, err
	}

	return lon, lat, nil
}

// FormatCoordinate ...
func FormatCoordinate(longitude, latitude float64) string {
	return fmt.Sprintf("%f,%f", longitude, latitude)
}
