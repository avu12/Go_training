package repositories

import (
	"net/http"

	"../../domain/repositories"
	"../../services"
	"../../utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	err := c.ShouldBindJSON(&request)

	if err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, apierr := services.RepositoryService.CreateRepo(request)
	if apierr != nil {
		c.JSON(apierr.Status(), apierr)
		return
	}

	c.JSON(http.StatusCreated, result)

}
