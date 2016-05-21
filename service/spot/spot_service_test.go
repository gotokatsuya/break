package spot

import (
	"net/url"
	"path"
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

func TestPhotoURL(t *testing.T) {
	u, err := url.Parse("https://irs1.4sqi.net/img/user/")
	if err != nil {
		t.Error(err)
	}
	u.Path = path.Join(u.Path, "/AAA.jpg")
	if u.String() != "https://irs1.4sqi.net/img/user/AAA.jpg" {
		t.Fail()
	}
}
