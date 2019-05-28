package model

import "database/sql"

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
