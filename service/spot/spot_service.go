package spot

import (
	"net/url"
	"path"

	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/foursquare"
	libmath "github.com/gotokatsuya/break/library/math"
	spotmodel "github.com/gotokatsuya/break/model/spot"
	tourspotmodel "github.com/gotokatsuya/break/model/tour/tourspot"
	parameter "github.com/gotokatsuya/break/parameter/spot"
)

func getSpotsByFoursquare(userID, visitTime int64, lat, lng float64) ([]*spotmodel.Entity, error) {
	latStr, lngStr := libmath.RoundLatLngToStr(lat, lng)
	response, err := foursquare.GetExplore(latStr, lngStr)
	if err != nil {
		return nil, err
	}
	venues := response.GetVenues()
	spots := make([]*spotmodel.Entity, len(venues))
	for i, venue := range venues {
		spot := spotmodel.New(userID, visitTime)
		lat, lng, err := libmath.RoundSpotLatLng(venue.Location.Lat, venue.Location.Lng)
		if err != nil {
			return nil, err
		}
		spot.SetVenue(
			venue.Name,
			venue.Location.State,
			venue.Location.City,
			venue.Location.Address,
			lat,
			lng,
		)
		for _, category := range venue.Categories {
			spot.CategoryName = category.Name
			break
		}
		for _, group := range venue.Photos.Groups {
			for _, item := range group.Items {
				u, err := url.Parse(item.Prefix)
				if err != nil {
					return nil, err
				}
				u.Path = path.Join(u.Path, item.Suffix)
				spot.PhotoURL = u.String()
				break
			}
		}
		spots[i] = spot
	}
	return spots, nil
}

func CreateSpot(ctx echo.Context, req *parameter.CreateRequest) error {
	for _, spotLog := range req.SpotLogs {
		//TODO Should use go routine.
		spots, err := getSpotsByFoursquare(req.UserID, spotLog.VisitTime, spotLog.Lat, spotLog.Lng)
		if err != nil {
			return err
		}
		spotRepository := spotmodel.NewRepository(ctx)
		for _, spot := range spots {
			if _, err := spotRepository.Create(spot); err != nil {
				return err
			}
		}
	}
	return nil
}

func GetSpot(ctx echo.Context, req *parameter.GetRequest) (*parameter.GetResponse, error) {
	spotRepository := spotmodel.NewRepository(ctx)
	var (
		spots []spotmodel.Entity
		err   error
	)
	switch {
	case req.HasTourID():
		tourSpotRepository := tourspotmodel.NewRepository(ctx)
		tourSpots, err := tourSpotRepository.FindByTourID(req.TourID)
		spotIDs := make([]int64, len(tourSpots))
		for i, tourSpot := range tourSpots {
			spotIDs[i] = tourSpot.SpotID
		}
		spots, err = spotRepository.FindByIDs(spotIDs)
		if err != nil {
			return nil, err
		}
	default:
		spots, err = spotRepository.FindByUserIDAndVisitRange(req.UserID, req.StartVisitTime, req.EndVisitTime)
		if err != nil {
			return nil, err
		}
	}
	res := parameter.NewGetResponse()
	if err := res.Convert(spots); err != nil {
		return nil, err
	}
	return res, nil
}
