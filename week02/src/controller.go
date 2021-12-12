/*
模拟路由调用dao
*/

package src

type controller struct {
	dao *dao
}

func NewController(dao *dao) *controller {
	return &controller{
		dao: dao,
	}
}

// 数据库中，通过用户ID查找用户姓名
func (c *controller) FindUserNameByID(id int) (string, error) {
	var name string
	row, err := c.dao.FindUserByID(id)
	if err != nil {
		return name, err
	}
	name = row["name"].(string)
	return name, err
}
