/*
模拟一个dao层, 使用errors.Wrap来包装db.go中抛出的错误
*/

package src

import (
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
)

type dao struct {
	db *db
}

// 创建dao对象
func NewDAO(db *db) *dao {
	return &dao{
		db: db,
	}
}

// 数据库中，通过ID查找用户
func (d *dao) FindUserByID(id int) (map[string]interface{}, error) {
	syntax := "SELECT id, name FROM user WHERE id = " + strconv.Itoa(rand.Intn(10))
	row, err := d.db.Query(syntax)
	if errors.Cause(err) == sql.ErrNoRows {
		err = errors.Wrapf(err, "dao: %v", syntax)
	} else {
		errors.Wrap(err, "dao: unknown error")
	}
	return row, err
}
