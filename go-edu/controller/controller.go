package controller

import (
	"github.com/gin-gonic/gin"
	"go-edu/tools/redis"
	"go.uber.org/zap"
)

func HelloHandler(c *gin.Context) {
	zap.L().Info("req hello")
	//zap.L().Error("测试错误")
	//jenkins.Getjoball()

	//var data []*model.RoomGoods
	//err := mysql.GetDB().Model(model.RoomGoods{}).First(&data).Error
	//if err != nil && err != gorm.ErrEmptySlice {
	//	zap.L().Error(err.Error())
	//}
	//fmt.Println("--------------------------------------", data[0])
	//fmt.Println(mysql.GetDB().First(&data).RowsAffected) //获取数据总条数
	//
	//err = redis.SetStr("resource:", strconv.Itoa(int(data[0].ID)), time.Duration(30)*time.Second)
	//if err != nil {
	//	zap.L().Error(err.Error())
	//}
	a := redis.GetStr("resource:")
	//fmt.Println(a)
	c.JSON(200, gin.H{"code": 0, "res": a})
}
