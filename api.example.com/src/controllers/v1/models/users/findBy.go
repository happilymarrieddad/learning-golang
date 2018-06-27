package users

import (
	"github.com/go-xorm/xorm"
)

func FindBy(db *xorm.Engine, limit int, offset int) (users Users, err error) {
	err = db.
		Limit(limit, offset).
		Find(&users)

	return
}
