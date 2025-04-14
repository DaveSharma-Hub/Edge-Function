package handler;

import (
	"github.com/gin-gonic/gin"
	"github.com/DaveSharma-Hub/Edge-Function/src/edge-function/executionFunction"
	"net/http"
)

type ExecuteFunctionPostBodyParams struct {
	DockerImageUri string `json:"dockerImageUri"`;
	MaxCpuUsage float32 `json:"cpu"`;
	MaxMemoryUsage int `json:"memory"`;
	Timeout int `json:"timeout"`;
};

func ExecuteFunctionHandler(c *gin.Context){
	var input ExecuteFunctionPostBodyParams;
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		});
		return;
	}
	output, outputErr := executeFunction.ExecuteFunction(input.DockerImageUri, input.Timeout, input.MaxCpuUsage, input.MaxMemoryUsage);
	// for now a 30 second timeout

	if (outputErr != nil) {
		c.JSON(400, gin.H{
			"error": outputErr.Error(),
		});
		return;
	}

	c.JSON(http.StatusOK, gin.H{
		"output": output,
	});
}