package main

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	name string `gorm:"unique"`
}
type Users struct {
	gorm.Model
	Name string
}

// equals
type Users struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

var u Users

// type users struct {
// 	gorm.Model
// }

func initMigrate() {

}

func main() {
	initMigrate()
}
