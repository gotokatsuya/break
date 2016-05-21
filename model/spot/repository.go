package spot

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

// FindByUserID ...
func (r *Repository) FindByUserID(userID int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("user_id = ?", userID).Order("visit_time ASC").Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ents, nil
}

// FindByUserIDAndVisitRange ...
func (r *Repository) FindByUserIDAndVisitRange(userID, startVisitTime, endVisitTime int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("user_id = ? AND visit_time BETWEEN ? AND ?", userID, startVisitTime, endVisitTime).Order("visit_time ASC").Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
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

// FindByIDsOrderVisitAsc ...
func (r *Repository) FindByIDs(ids []int64) ([]Entity, error) {
	var ents []Entity
	if err := r.DB.Where("id IN (?)", ids).Order("visit_time ASC").Find(&ents).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ents, nil
}
