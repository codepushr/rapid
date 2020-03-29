package rapid

import (
	"github.com/globalsign/mgo"
	"github.com/jinzhu/gorm"
)

// DB type
type DB struct {
	GORM *gorm.DB
	MGO  *mgo.Database
}
