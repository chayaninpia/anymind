package modules

import (
	"fmt"

	"xorm.io/xorm"
)

func DbConnect() (*xorm.Engine, error) {

	dataSoruce := `host='localhost' port=9998 user='chayanin' password='anymind' dbname='bitcoin' sslmode=false`
	dx, err := xorm.NewEngine(`postgress`, dataSoruce)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return dx, nil
}
