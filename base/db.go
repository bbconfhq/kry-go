package base

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kry-go/model/entity"
)

type Connection struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type Db struct {
	Db   *gorm.DB
	Conn *Connection
}

func (d *Db) SetConnection(conn *Connection) {
	d.Conn = &Connection{
		User:     conn.User,
		Password: conn.Password,
		Name:     conn.Name,
		Host:     conn.Host,
		Port:     conn.Port,
	}
}

func (d *Db) Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Conn.User,
		d.Conn.Password,
		d.Conn.Host,
		d.Conn.Port,
		d.Conn.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	d.Db = db
	d.MigrateDatabase()
}

func (d *Db) MigrateDatabase() {
	err := d.Db.AutoMigrate(
		&entity.Contest{},
		&entity.Problem{},
		&entity.Submission{},
		&entity.Tag{},
		&entity.Testcase{},
		&entity.User{})

	if err != nil {
		panic(err.Error())
	}
}
