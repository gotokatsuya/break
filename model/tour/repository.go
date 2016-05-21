package tour

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/model"
)

// Repository ...
type Repository struct {
	*model.RootRepository
}

// NewRepository ...
func NewRepository(ctx echo.Context) *Repository {
	return &Repository{
		RootRepository: model.NewRootRepository(ctx, new(Entity)),
	}
}

// FindByUserID ...
func (r *Repository) FindByUserID(userID int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("user_id = ?", userID).Preload("User").Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ents, nil
}

// FindByLatLngRange ...
func (r *Repository) FindByLatLngRange(lat, lng float64) ([]Entity, error) {
	var ents []Entity
	minLatLngQuery := "min_lat BETWEEN (? - 0.05) AND (? + 0.05) AND min_lng BETWEEN (? - 0.05) AND (? + 0.05)"
	maxLatLngQuery := "max_lat BETWEEN (? - 0.05) AND (? + 0.05) AND max_lng BETWEEN (? - 0.05) AND (? + 0.05)"
	if err := r.DB.Where(fmt.Sprintf("(%s) OR (%s)", minLatLngQuery, maxLatLngQuery), lat, lat, lng, lng, lat, lat, lng, lng).Preload("User").Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ents, nil
}

// Create inserts the entity
func (r *Repository) Create(ent *Entity) (*Entity, error) {
	if err := r.RootRepository.Create(ent); err != nil {
		return nil, err
	}
	return ent, nil
}

// Update ...
func (r *Repository) Update(ent *Entity) (*Entity, error) {
	if err := r.RootRepository.UpdateByPrimary(ent); err != nil {
		return nil, err
	}
	return ent, nil
}
