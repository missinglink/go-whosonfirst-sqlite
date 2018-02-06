package sqlite

import (
       "database/sql"
       "github.com/whosonfirst/go-whosonfirst-brands"
       "github.com/whosonfirst/go-whosonfirst-geojson-v2"
)

type Database interface {
     Conn() (*sql.DB, error)
     DSN() string
     Close() error
     Lock() error
     Unlock() error
}

type Table interface {
     Name() string
     Schema() string
     InitializeTable(Database) error
     IndexRecord(Database, interface{}) error
}

type FeatureTable interface {
     Table
     IndexFeature(Database, geojson.Feature) error
}

type BrandTable interface {
     Table
     IndexBrand(Database, brands.Brand) error
}


// this is here so we can pass both sql.Row and sql.Rows to the
// ResultSetFunc below (20170824/thisisaaronland)

type ResultSet interface {
	Scan(dest ...interface{}) error
}

type ResultRow interface {
     Row() interface{}
}

type ResultSetFunc func(row ResultSet) (ResultRow, error)
