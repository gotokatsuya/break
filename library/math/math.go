package math

import (
	"fmt"
	"math"
	"strconv"
)

func RoundLatLng(lat, lng float64) (float64, float64, error) {
	latStr, lngStr := RoundLatLngToStr(lat, lng)
	return ParseLatLng(latStr, lngStr)
}

func RoundLatLngToStr(lat, lng float64) (string, string) {
	return fmt.Sprintf("%.5f", Round(lat, 5)), fmt.Sprintf("%.5f", Round(lng, 5))
}

func RoundSpotLatLng(lat, lng float64) (float64, float64, error) {
	latStr, lngStr := RoundSpotLatLngToStr(lat, lng)
	return ParseLatLng(latStr, lngStr)
}

func RoundSpotLatLngToStr(lat, lng float64) (string, string) {
	return fmt.Sprintf("%.7f", Round(lat, 7)), fmt.Sprintf("%.7f", Round(lng, 7))
}

func RoundTakenHour(hour float64) (float32, error) {
	hourStr := RoundTakenHourToStr(hour)
	parsedHour, err := strconv.ParseFloat(hourStr, 32)
	if err != nil {
		return 0, err
	}
	return float32(parsedHour), nil
}

func RoundTakenHourToStr(hour float64) string {
	return fmt.Sprintf("%.2f", Round(hour, 2))
}

func ParseLatLng(latStr, lngStr string) (float64, float64, error) {
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return 0, 0, err
	}
	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		return 0, 0, err
	}
	return lat, lng, nil
}

func Round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
