// 默认实现的 ping api

package routes

import (
	"fmt"

	"github.com/axiaoxin-com/pink-lady/models"
	"github.com/axiaoxin-com/pink-lady/routes/response"
	"github.com/axiaoxin-com/pink-lady/services"
	"github.com/gin-gonic/gin"
)

// User godoc
// @Summary 默认的 user 接口
// @Description 添加用户
// @Tags User
// @Accept json
// @Produce json
// @param data body models.User true "id"
// @Success 200 {object} response.Response
// @Router /cloud/add_user [post]
func add_user(c *gin.Context) {
	fmt.Println(c, "what is c!!!!!")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "Not a Story")
		return
	}
	fmt.Println(&user, "what is user!!!!!!!")
	// services.add(&user)
	services.DB().Create(&user)
	// ur := services.UserRepo{RepoBase: nil}
	// services.UserRepo.add(c)
	response.JSON(c, user)
	return
}
