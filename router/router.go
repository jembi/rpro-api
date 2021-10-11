package router

import (
	"jembi/rpro-api/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFlowFlowsRS() {

	router := gin.Default()
	router.GET("/flowsflow", getFlowsFlow)
	router.Run("localhost:8087")

}

func getFlowsFlow(c *gin.Context) {
	ff, err := data.GetFlowsFlowList()

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, ff)
}
