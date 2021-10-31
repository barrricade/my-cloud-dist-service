// 默认实现的 ping api

package routes

import (
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
	// ur := services.UserRepo{RepoBase: nil}
	// services.UserRepo.add(c)
	// response.JSON(c, data)
	return
}
