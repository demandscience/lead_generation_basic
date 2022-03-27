package data

import (
	"lead_generation_basic/data/sql"
)

// Sample sample data
type Sample struct {
	Pk        int     `db:"pk"`
	Value     string  `db:"value"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
	DeletedAt *string `db:"deleted_at"`
}

// GetAll get all data
func GetAll() ([]*Sample, error) {
	var result []*Sample

	err := sql.Db.Select(&result, "SELECT * FROM sample")
	if err != nil {
		return nil, err
	}

	return result, nil
}
