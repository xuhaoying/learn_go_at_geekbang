/*
模拟一个数据库驱动对象，内含一个Query方法，随机抛出sql.ErrNoRows或者未知错误或者为nil的error
*/

package src

import (
	"database/sql"
	"math/rand"

	"github.com/pkg/errors"
)

// 定义一个错误
var ErrUnknow = errors.New("sql: unkonwn error")

// 一个假的数据库对象
type db struct{}

// 创建一个假的数据库对象
func NewDB() *db {
	return &db{}
}

/*
在假的数据库对象中模拟一个查询方法，可能会引起3种结果:
 - 返回id和name
 - 导致sql.ErrNoRow
 - 导致未知的错误
*/
func (d *db) Query(syntax string) (map[string]interface{}, error) {
	id := rand.Intn(10)
	status := rand.Intn(10)
	switch {
	case status < 3:
		return nil, sql.ErrNoRows
	case status >= 6:
		return map[string]interface{}{
			"id":   id,
			"name": "user",
		}, nil
	default:
		return nil, ErrUnknow
	}
}
