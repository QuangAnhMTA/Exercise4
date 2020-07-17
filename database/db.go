package database

import (
	"log"

	"xorm.io/xorm"
)

type Db struct {
	Engine *xorm.Engine
}

func (d *Db) ConnectBD() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", "root:1@tcp(0.0.0.0:3306)/test")
	if err != nil {
		return nil, err
	}
	d.Engine = engine
	log.Println("sucess")
	return d.Engine, nil
}

func (d *Db) CreateTable() error {
	var err error
	err = d.Engine.CreateTables(User{})
	err = d.Engine.CreateTables(Point{})
	err = d.Engine.Sync2(new(User))
	err = d.Engine.Sync2(new(Point))
	if err != nil {
		return nil
	}
	return nil
}
