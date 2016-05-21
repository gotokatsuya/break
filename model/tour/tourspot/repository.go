package tourspot

import (
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

// GetByID ...
func (r *Repository) GetByID(tourID, spotID int64) (*Entity, error) {
	ent := new(Entity)
	ent.TourID = tourID
	ent.SpotID = spotID
	if err := r.RootRepository.GetByPrimary(ent); err != nil {
		return nil, err
	}
	return ent, nil
}

// FindByTourID ...
func (r *Repository) FindByTourID(tourID int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("tour_id = ?", tourID).Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
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
