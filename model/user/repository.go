package user

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
func (r *Repository) GetByID(id int64) (*Entity, error) {
	ent := new(Entity)
	ent.ID = id
	if err := r.RootRepository.GetByPrimary(ent); err != nil {
		return nil, err
	}
	return ent, nil
}

// GetByEmail ...
func (r *Repository) GetByEmail(email string) (*Entity, error) {
	ent := new(Entity)
	if err := r.DB.Where("email = ?", email).First(ent).Error; err != nil {
		return nil, err
	}
	return ent, nil
}

// GetByToken ...
func (r *Repository) GetByToken(token string) (*Entity, error) {
	ent := new(Entity)
	if err := r.DB.Where("token = ?", token).First(ent).Error; err != nil {
		return nil, err
	}
	return ent, nil
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

// FindByIDs ...
func (r *Repository) FindByIDs(ids []int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("id IN (?)", ids).Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ents, nil
}
