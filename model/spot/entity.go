package spot

import (
	"time"

	libtime "github.com/gotokatsuya/break/library/time"
)

// Entity describes user table.
type Entity struct {
	ID           int64     `json:"id"            gorm:"column:id;primary_key"`
	UserID       int64     `json:"user_id"       gorm:"column:user_id;index:idx_spot_user_id_visit_time"         sql:"not null;type:bigint(20)"`
	VisitTime    int64     `json:"visit_time"    gorm:"column:visit_time;index:idx_spot_user_id_visit_time"      sql:"not null;bigint(20)"`
	Name         string    `json:"name"          gorm:"column:name"            sql:"not null;type:varchar(190)"`
	PhotoURL     string    `json:"photo_url"     gorm:"column:photo_url"       sql:"type:varchar(190)"`
	CategoryName string    `json:"category_name" gorm:"column:category_name"   sql:"type:varchar(190)"`
	State        string    `json:"state"         gorm:"column:state"           sql:"type:varchar(190)"`
	City         string    `json:"city"          gorm:"column:city"            sql:"type:varchar(190)"`
	Address      string    `json:"address"       gorm:"column:address"         sql:"type:varchar(190)"`
	Lat          float64   `json:"lat"           gorm:"column:lat"             sql:"double(9,7)"`
	Lng          float64   `json:"lng"           gorm:"column:lng"             sql:"double(10,7)"`
	CreatedAt    time.Time `json:"created_at"    gorm:"column:created_at"      sql:"not null;type:datetime"`
	UpdatedAt    time.Time `json:"updated_at"    gorm:"column:updated_at"      sql:"not null;type:datetime"`
}

// New ...
func New(userID, visitTime int64) *Entity {
	return &Entity{
		UserID:    userID,
		VisitTime: visitTime,
		CreatedAt: libtime.Now(),
		UpdatedAt: libtime.Now(),
	}
}

func (e *Entity) SetVenue(name, state, city, address string, lat, lng float64) {
	e.Name = name
	e.State = state
	e.City = city
	e.Address = address
	e.Lat = lat
	e.Lng = lng
}

// TableName returns the table name
func (e Entity) TableName() string {
	return "spot"
}

// Validate ...
func (e *Entity) Validate() error {
	return nil
}

// Primary returns
func (e *Entity) Primary() (interface{}, []interface{}) {
	return "id = ?", []interface{}{e.ID}
}
