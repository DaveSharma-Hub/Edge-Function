package server
import (
	"github.com/gin-gonic/gin"
	"github.com/DaveSharma-Hub/Edge-Function/src/edge-function/server/handler"
	"net/http"
)

func attachHandler(engine *gin.Engine, restType string, endpoint string, fn gin.HandlerFunc){
	switch restType {
		case "GET":
			engine.GET(endpoint, fn);
			return;
		case "POST":
			engine.POST(endpoint, fn);
			return;
		default:
			return;
	}
}

func Start(){
	r := gin.Default();
	attachHandler(r, "GET", "/test", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"body": "Hello World",
		});
	});
	attachHandler(r, "POST", "/executeFunction", func(c *gin.Context){
		// get the docker image
		// run the docker image
		// send output back
		handler.ExecuteFunctionHandler(c);
	});
	r.Run(":3000");
}
