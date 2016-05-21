package spot

import (
	"testing"
	"time"

	"github.com/gotokatsuya/break/config"
)

func init() {
	config.Load()
}

func TestGetSpotsByFoursquare(t *testing.T) {
	res, err := getSpotsByFoursquare(int64(1), time.Now().Unix(), float64(40.7), float64(-74))
	if err != nil {
		t.Error(err)
	}
	if len(res) == 0 {
		t.Fail()
	}
}
