package route

import (
	"github.com/gin-gonic/gin"
	"github.com/laioncorcino/students/controller"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/api/v1/students", controller.GetAll)
	r.GET("/api/v1/students/:studentID", controller.GetById)
	r.DELETE("/api/v1/students/:studentID", controller.Delete)
	r.POST("/api/v1/students", controller.Create)
	_ = r.Run(":8001")
}
