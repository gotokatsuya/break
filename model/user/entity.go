package user

import (
	"fmt"
	"time"

	"github.com/gotokatsuya/break/config"
	"github.com/gotokatsuya/break/library/crypto"
	libtime "github.com/gotokatsuya/break/library/time"
)

const (
	Active   = 1
	Inactive = 2
)

// Entity describes user table.
type Entity struct {
	ID           int64     `json:"id"                 gorm:"column:id;primary_key"`
	Name         string    `json:"name"               gorm:"column:name"             sql:"not null;type:varchar(190)"`
	Email        string    `json:"email"              gorm:"column:email"            sql:"not null;unique_index;type:varchar(190)"`
	PhotoURL     string    `json:"photo_url"          gorm:"column:photo_url"        sql:"type:varchar(190)"`
	ActiveStatus int       `json:"active_status"      gorm:"column:active_status"    sql:"not null;type:tinyint(3)"`
	Token        string    `json:"token"              gorm:"column:token"            sql:"not null;unique_index;type:varchar(190)"`
	CreatedAt    time.Time `json:"created_at"         gorm:"column:created_at"       sql:"not null;type:datetime"`
	UpdatedAt    time.Time `json:"updated_at"         gorm:"column:updated_at"       sql:"not null;type:datetime"`
}

// New ...
func New(name, email, photoURL string) *Entity {
	return &Entity{
		Name:         name,
		Email:        email,
		PhotoURL:     photoURL,
		ActiveStatus: Active,
		CreatedAt:    libtime.Now(),
		UpdatedAt:    libtime.Now(),
	}
}

func (e *Entity) RenewToken() error {
	switch {
	case config.DebugMode():
		e.Token = fmt.Sprintf("debug-%d", e.ID)
	case config.TestMode():
		e.Token = fmt.Sprintf("test-%d", e.ID)
	default:
		randString, err := crypto.GenerateRandom(32)
		if err != nil {
			return err
		}
		e.Token = randString
	}
	return nil
}

// TableName returns the table name
func (e Entity) TableName() string {
	return "user"
}

// Validate ...
func (e *Entity) Validate() error {
	return nil
}

// Primary returns
func (e *Entity) Primary() (interface{}, []interface{}) {
	return "id = ?", []interface{}{e.ID}
}
