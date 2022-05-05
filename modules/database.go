package modules

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func DbConnect() (*xorm.Engine, error) {

	dataSoruce := `host='localhost' port='5432' user='postgres' password='admin' dbname='bitcoin' sslmode=disable`
	dx, err := xorm.NewEngine(`postgres`, dataSoruce)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return dx, nil
}
