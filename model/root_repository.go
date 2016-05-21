package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/net/context/database"
)

// RootRepository ...
type (
	RootRepository struct {
		Ctx    echo.Context
		DB     *gorm.DB
		Entity Entity
	}

	Entity interface {
		Validate() error
		Primary() (interface{}, []interface{})
	}

	CreateRepository interface {
		Create(Entity) error
	}

	ReadRepository interface {
		GetByPrimary(Entity) error
	}

	UpdateRepository interface {
		UpdateByPrimary(Entity) error
	}

	DeleteRepository interface {
		Delete(Entity) error
	}
)

// NewRootRepository creates new RootRepository
func NewRootRepository(ctx echo.Context, ent Entity) *RootRepository {
	return &RootRepository{
		Ctx:    ctx,
		DB:     database.Get(ctx),
		Entity: ent,
	}
}

// Create returns inserted entity if no error
func (r *RootRepository) Create(ent Entity) error {
	if err := ent.Validate(); err != nil {
		return err
	}
	if err := r.DB.Create(ent).Error; err != nil {
		return err
	}
	return nil
}

// GetByPrimary ...
func (r *RootRepository) GetByPrimary(ent Entity) error {
	if err := r.DB.Where(ent.Primary()).First(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *RootRepository) UpdateByPrimary(ent Entity) error {
	if err := ent.Validate(); err != nil {
		return err
	}
	if err := r.DB.Where(ent.Primary()).Save(ent).Error; err != nil {
		return err
	}
	return nil
}

func (r *RootRepository) Delete(ent Entity) error {
	if err := ent.Validate(); err != nil {
		return err
	}
	if err := r.DB.Delete(ent).Error; err != nil {
		return err
	}
	return nil
}
