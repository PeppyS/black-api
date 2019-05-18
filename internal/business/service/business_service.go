package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type BusinessService struct {
	db *sqlx.DB
}

type Business struct {
	ID                string         `db:"id"`
	Description       string         `db:"description"`
	DisplayName       string         `db:"display_name"`
	DisplayAddress    string         `db:"display_address"`
	AddressLine1      string         `db:"address_line_1"`
	AddressLine2      sql.NullString `db:"address_line_2"`
	AddressCity       string         `db:"address_city"`
	AddressState      string         `db:"address_state"`
	AddressPostalCode string         `db:"address_postal_code"`
}

func NewBusinessService(db *sqlx.DB) *BusinessService {
	return &BusinessService{db}
}

func (bs *BusinessService) GetByID(ID string) (Business, error) {
	business := Business{}
	err := bs.db.Get(&business, "SELECT * FROM business WHERE id = $1", ID)

	return business, err
}

func (bs *BusinessService) List() ([]Business, error) {
	businesses := []Business{}
	err := bs.db.Select(&businesses, "SELECT * FROM business")

	return businesses, err
}
