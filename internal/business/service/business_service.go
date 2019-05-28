package service

import (
	"github.com/PeppyS/black-api/internal/business/model"
	"github.com/jmoiron/sqlx"
)

type BusinessService struct {
	db *sqlx.DB
}

func NewBusinessService(db *sqlx.DB) *BusinessService {
	return &BusinessService{db}
}

func (bs *BusinessService) GetByID(ID string) (model.Business, error) {
	business := model.Business{}
	err := bs.db.Get(&business, "SELECT * FROM business WHERE id = $1", ID)

	return business, err
}

func (bs *BusinessService) List() ([]model.Business, error) {
	businesses := []model.Business{}
	err := bs.db.Select(&businesses, "SELECT * FROM business")

	return businesses, err
}
