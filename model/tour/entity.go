package tour

import (
	"time"

	libtime "github.com/gotokatsuya/break/library/time"
	usermodel "github.com/gotokatsuya/break/model/user"
)

// Entity describes user table.
type Entity struct {
	ID        int64     `json:"id"         gorm:"column:id;primary_key"`
	UserID    int64     `json:"user_id"    gorm:"column:user_id"         sql:"not null;index;type:bigint(20)"`
	Name      string    `json:"name"       gorm:"column:name"            sql:"not null;type:varchar(190)"`
	PhotoURL  string    `json:"photo_url"  gorm:"column:photo_url"       sql:"type:varchar(190)"`
	MinLat    float64   `json:"min_lat"    gorm:"column:min_lat;index:idx_tour_min_lat_min_lng" sql:"double(9,7)"`
	MinLng    float64   `json:"min_lng"    gorm:"column:min_lng;index:idx_tour_min_lat_min_lng" sql:"double(10,7)"`
	MaxLat    float64   `json:"max_lat"    gorm:"column:max_lat;index:idx_tour_max_lat_max_lng" sql:"double(9,7)"`
	MaxLng    float64   `json:"max_lng"    gorm:"column:max_lng;index:idx_tour_max_lat_max_lng" sql:"double(10,7)"`
	TakenHour float32   `json:"taken_hour" gorm:"column:taken_hour"      sql:"float"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"      sql:"not null;type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"      sql:"not null;type:datetime"`

	User usermodel.Entity `gorm:"ForeignKey:UserID"`
}

// New ...
func New(userID int64, name, photoURL string, minLat, minLng, maxLat, maxLng float64, takenHour float32) *Entity {
	return &Entity{
		UserID:    userID,
		Name:      name,
		PhotoURL:  photoURL,
		MinLat:    minLat,
		MinLng:    minLng,
		MaxLat:    maxLat,
		MaxLng:    maxLng,
		TakenHour: takenHour,
		CreatedAt: libtime.Now(),
		UpdatedAt: libtime.Now(),
	}
}

// TableName returns the table name
func (e Entity) TableName() string {
	return "tour"
}

// Validate ...
func (e *Entity) Validate() error {
	return nil
}

// Primary returns
func (e *Entity) Primary() (interface{}, []interface{}) {
	return "id = ?", []interface{}{e.ID}
}
