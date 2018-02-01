package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-names/tags"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
)

type NamesFullTextSearchTable struct {
	sqlite.Table
	name string
}

func NewNamesFullTextSearchTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewNamesFullTextSearchTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewNamesFullTextSearchTable() (sqlite.Table, error) {

	t := NamesFullTextSearchTable{
		name: "names",
	}

	return &t, nil
}

func (t *NamesFullTextSearchTable) Name() string {
	return t.name
}

func (t *NamesFullTextSearchTable) Schema() string {

     	sql := `CREATE VIRTUAL TABLE documents_search USING fts4(id, placetype, country, language, extlang, privateuse, name);`

	return fmt.Sprintf(sql)
}

func (t *NamesFullTextSearchTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *NamesFullTextSearchTable) IndexFeature(db sqlite.Database, f geojson.Feature) error {

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	id := f.Id()

	sql := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, t.Name())

	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	pt := f.Placetype()
	co := whosonfirst.Country(f)

	names := whosonfirst.Names(f)

	for tag, names := range names {

		lt, err := tags.NewLangTag(tag)

		if err != nil {
			return err
		}

		for _, n := range names {

			if err != nil {
				return err
			}

			sql := fmt.Sprintf(`INSERT INTO %s (
	    			id, placetype, country,
				language, extlang,
	    			privateuse,
				name,
			) VALUES (
	    		  	?, ?, ?,
				?, ?,
				?,
				?
			)`, t.Name())

			stmt, err := tx.Prepare(sql)

			if err != nil {
				return err
			}

			defer stmt.Close()

			_, err = stmt.Exec(id, pt, co, lt.Language(), lt.ExtLang(), lt.PrivateUse(), n)

			if err != nil {
				return err
			}

		}
	}

	return tx.Commit()
}
