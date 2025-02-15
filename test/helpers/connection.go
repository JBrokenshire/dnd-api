package helpers

import (
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

func MockDb() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, err := gorm.Open(mocket.DriverName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
