package location

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	SetLocation() gin.HandlerFunc
	ResetLocation() gin.HandlerFunc
}

type locationHandler struct {
	locationService Service
}

func (handler *locationHandler) SetLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Location string `json:"city"`
		}
		var requestBody Request
		err := c.ShouldBind(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid json",
			})
			return
		}

		handler.locationService.SetLocation(requestBody.Location)

		c.JSON(http.StatusOK, gin.H{
			"message": "location set",
		})
	}
}

func (handler *locationHandler) ResetLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.locationService.ResetLocation()
		c.JSON(http.StatusOK, gin.H{
			"message": "location reset",
		})
	}
}

func NewLocationHandler(locationService Service) Handler {
	return &locationHandler{
		locationService: locationService,
	}
}
