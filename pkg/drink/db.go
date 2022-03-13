package drink

import (
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"github.com/nicjohnson145/mixer-service/pkg/common"
)

var ModelStruct = sqlbuilder.NewStruct(new(Model))

const (
	TableName = "drink"
)

type Model struct {
	ID             int64  `db:"id"`
	Name           string `db:"name" fieldtag:"required_insert"`
	Username       string `db:"username" fieldtag:"required_insert"`
	PrimaryAlcohol string `db:"primary_alcohol" fieldtag:"required_insert"`
	PreferredGlass string `db:"preferred_glass" fieldtag:"required_insert"`
	Ingredients    string `db:"ingredients" fieldtag:"required_insert"`
	Instructions   string `db:"instructions" fieldtag:"required_insert"`
	Notes          string `db:"notes" fieldtag:"required_insert"`
}

func getByID(id int64, db *sql.DB) (*Model, error) {
	sb := ModelStruct.SelectFrom(TableName)
	sb.Where(sb.Equal("id", id))

	sql, args := sb.Build()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasRow := rows.Next()
	if !hasRow {
		return nil, common.ErrNotFound
	}

	var drink Model
	err = rows.Scan(ModelStruct.Addr(&drink)...)
	if err != nil {
		return nil, err
	}

	return &drink, nil
}

func getByNameAndUsername(name string, username string, db *sql.DB) (*Model, error) {
	sb := ModelStruct.SelectFrom(TableName)
	sb.Where(
		sb.Equal("name", name),
		sb.Equal("username", username),
	)

	sql, args := sb.Build()
	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasRow := rows.Next()
	if !hasRow {
		return nil, common.ErrNotFound
	}

	var drink Model
	err = rows.Scan(ModelStruct.Addr(&drink)...)
	if err != nil {
		return nil, err
	}

	return &drink, nil
}

func create(d Model, db *sql.DB) (int64, error) {
	sql, args := ModelStruct.InsertIntoForTag(TableName, "required_insert", d).Build()
	rows, err := db.Exec(sql, args...)
	if err != nil {
		return -1, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
