package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service ServiceInterface
}

func RegisterRoutes(r *gin.Engine, service ServiceInterface) {
	resource := &resource{
		service: service,
	}
	r.GET("/users", resource.Query)
	r.POST("/users", resource.Create)
	r.GET("/users/:id", resource.Get)
	r.PUT("/users/:id", resource.Put)
	r.DELETE("/users/:id", resource.Delete)
}

func (resource *resource) Query(c *gin.Context) {

}

func (resource *resource) Create(c *gin.Context) {
	user := User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := resource.service.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Created User",
		"data":    user,
	})
}

func (resource *resource) Get(c *gin.Context) {

}

func (resource *resource) Put(c *gin.Context) {

}

func (resource *resource) Delete(c *gin.Context) {

}
