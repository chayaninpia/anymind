package modules

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DbConnect() (*sqlx.DB, error) {

	dataSoruce := `host='localhost' port=5432 user='chayanin' password='anymind' dbname='bitcoin' sslmode=false`
	dx, err := sqlx.Connect(`postgress`, dataSoruce)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return dx, nil
}
