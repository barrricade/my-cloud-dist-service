// 默认实现的 ping api

package routes

import (
	"errors"

	"github.com/axiaoxin-com/pink-lady/models"
	"github.com/axiaoxin-com/pink-lady/routes/response"
	"github.com/axiaoxin-com/pink-lady/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User godoc
// @Summary 默认的 user 接口
// @Description 添加用户
// @Tags User
// @Accept json
// @Produce json
// @param data body models.User true "data"
// @Success 200 {object} response.Response
// @Router /cloud/add_user [post]
func add_user(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "Not a valid From")
		return
	}
	// todo判断账户和手机号是否已存在
	if err := services.DB().First(&user, "account = ?", user.Account).Error; err == nil {
		response.Respond(c, -1, "该账户已存在", errors.New("该账户已存在"))
		return
	}
	if err := services.DB().First(&user, "phone = ?", user.Phone).Error; err == nil {
		response.Respond(c, -1, "该手机号已存在", errors.New("该手机号已存在"))
		return
	}
	// todo 如果不存在则密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)
	services.DB().Create(&user)
	response.JSON(c, "创建成功")
	return
}
