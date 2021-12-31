//参考 Hystrix 实现一个滑动窗口计数器

package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

var LimitQueue map[string][]int64
var ok bool

//滑动窗口限流
func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	if LimitQueue == nil {
		LimitQueue = make(map[string][]int64)
	}
	if _, ok = LimitQueue[queueName]; !ok {
		LimitQueue[queueName] = make([]int64, 0)
	}
	//队列未满
	if uint(len(LimitQueue[queueName])) < count {
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
		return true
	}
	//队列满了,取出最早访问的时间
	earlyTime := LimitQueue[queueName][0]
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if currTime-earlyTime <= timeWindow {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		LimitQueue[queueName] = LimitQueue[queueName][1:]
		LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
	}
	return true
}

func limitIpFreq(c *gin.Context, timeWindow int64, count uint) bool {
	ip := c.ClientIP()
	key := "limit:" + ip
	if LimitFreqSingle(key, count, timeWindow) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Current IP frequently visited",
		})
		return false
	}
	return true
}

func main() {

}
