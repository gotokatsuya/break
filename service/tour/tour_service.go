package tour

import (
	"github.com/labstack/echo"

	libmath "github.com/gotokatsuya/break/library/math"
	libtime "github.com/gotokatsuya/break/library/time"
	spotmodel "github.com/gotokatsuya/break/model/spot"
	tourmodel "github.com/gotokatsuya/break/model/tour"
	tourspotmodel "github.com/gotokatsuya/break/model/tour/tourspot"
	parameter "github.com/gotokatsuya/break/parameter/tour"
)

func getMinMaxLatLng(spots []spotmodel.Entity) (float64, float64, float64, float64) {
	spotsSize := len(spots)
	if spotsSize == 0 {
		return 0, 0, 0, 0
	}
	var (
		minLat float64 = spots[0].Lat
		minLng float64 = spots[0].Lng
		maxLat float64 = spots[0].Lat
		maxLng float64 = spots[0].Lng
	)
	for _, spot := range spots {
		if minLat > spot.Lat {
			minLat = spot.Lat
		}
		if minLng > spot.Lng {
			minLng = spot.Lng
		}
		if maxLat < spot.Lat {
			maxLat = spot.Lat
		}
		if maxLng < spot.Lng {
			maxLng = spot.Lng
		}
	}
	return minLat, minLng, maxLat, maxLng
}

func getTakenHour(spots []spotmodel.Entity) float32 {
	spotsSize := len(spots)
	if spotsSize >= 2 {
		return 0
	}
	firstVisitTime := libtime.CreateTimeFromRecordTime(spots[0].VisitTime)
	endVisitTime := libtime.CreateTimeFromRecordTime(spots[spotsSize-1].VisitTime)
	duration := endVisitTime.Sub(firstVisitTime)
	hours, err := libmath.RoundTakenHour(duration.Hours())
	if err != nil {
		return 0
	}
	return hours
}

func CreateTour(ctx echo.Context, req *parameter.CreateRequest) error {
	spotRepository := spotmodel.NewRepository(ctx)
	spots, err := spotRepository.FindByIDs(req.SpotIDs)
	if err != nil {
		return err
	}
	minLat, minLng, maxLat, maxLng := getMinMaxLatLng(spots)
	takenHour := getTakenHour(spots)
	tourRepository := tourmodel.NewRepository(ctx)
	tour, err := tourRepository.Create(tourmodel.New(req.UserID, req.Name, req.PhotoURL, minLat, minLng, maxLat, maxLng, takenHour))
	if err != nil {
		return err
	}
	tourSpotRepository := tourspotmodel.NewRepository(ctx)
	for _, spot := range spots {
		if _, err := tourSpotRepository.Create(tourspotmodel.New(tour.ID, spot.ID)); err != nil {
			return err
		}
	}
	return nil
}

func GetTour(ctx echo.Context, req *parameter.GetRequest) (*parameter.GetResponse, error) {
	tourRepository := tourmodel.NewRepository(ctx)
	var (
		tours []tourmodel.Entity
		err   error
	)
	switch {
	case req.HasUserID():
		tours, err = tourRepository.FindByUserID(req.UserID)
		if err != nil {
			return nil, err
		}
	default:
		tours, err = tourRepository.FindByLatLngRange(req.Lat, req.Lng)
		if err != nil {
			return nil, err
		}
	}
	res := parameter.NewGetResponse()
	if err := res.Convert(tours); err != nil {
		return nil, err
	}
	return res, nil
}
