package main

/*import (          PRUEBA 1
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Name string `json:"name"`
}

func main() {
	engine := gin.New()
	engine.POST("/test", func(context *gin.Context) {
		body := Body{}

		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		context.JSON(http.StatusAccepted, &body)
	})
	engine.Run(":3000")
}*/
//Prueba 2
import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
