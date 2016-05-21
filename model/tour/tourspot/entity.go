package tourspot

import (
	"time"

	libtime "github.com/gotokatsuya/break/library/time"
)

// Entity describes user table.
type Entity struct {
	TourID    int64     `json:"tour_id"      gorm:"column:tour_id;primary_key" sql:"not null;type:bigint(20)"`
	SpotID    int64     `json:"spot_id"      gorm:"column:spot_id;primary_key" sql:"not null;type:bigint(20)"`
	CreatedAt time.Time `json:"created_at"   gorm:"column:created_at"          sql:"not null;type:datetime"`
}

// New ...
func New(tourID, spotID int64) *Entity {
	return &Entity{
		TourID:    tourID,
		SpotID:    spotID,
		CreatedAt: libtime.Now(),
	}
}

// TableName returns the table name
func (e Entity) TableName() string {
	return "tour_spot"
}

// Validate ...
func (e *Entity) Validate() error {
	return nil
}

// Primary returns
func (e *Entity) Primary() (interface{}, []interface{}) {
	return "tour_id = ? AND spot_id = ?", []interface{}{e.TourID, e.SpotID}
}
