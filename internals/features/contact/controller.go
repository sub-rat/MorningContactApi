package contact

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MorningContactApi/pkg/utils"
)

type resource struct {
	service ServiceInterface
}

func RegisterRoutes(r *gin.Engine, service ServiceInterface) {
	resource := &resource{
		service: service,
	}
	r.GET("/users/:id/contacts", resource.Query)
	r.POST("users/:id/contacts", resource.Create)
	r.GET("/contacts/:id", resource.Get)
	r.PUT("/contacts/:id", resource.Put)
	r.DELETE("/contacts/:id", resource.Delete)
}

func (resource *resource) Query(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	page, limit, err := utils.Pagination(c)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	userList, err := resource.service.Query(page*limit, limit, c.Query("q"), uint(id))
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "contacts list",
		"data":    userList,
	})
}

func (resource *resource) Create(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	contact := Contact{}
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	contact.UserID = uint(id)
	contact, err := resource.service.Create(&contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Created Contact",
		"data":    contact,
	})
}

func (resource *resource) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	contact, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get contact",
		"data":    contact,
	})
}

func (resource *resource) Put(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	updateContact := Contact{}
	if err := c.BindJSON(&updateContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contact, err := resource.service.Update(uint(id), &updateContact)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update the contact",
		"data":    contact,
	})

}

func (resource *resource) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = resource.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"message": "successfully Deleted",
	})
}
