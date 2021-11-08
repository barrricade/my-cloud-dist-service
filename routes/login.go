package routes

import (
	"errors"
	"fmt"

	"github.com/axiaoxin-com/pink-lady/models"
	"github.com/axiaoxin-com/pink-lady/routes/response"
	"github.com/axiaoxin-com/pink-lady/services"
	"github.com/axiaoxin-com/pink-lady/webserver"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type login struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	// Captcha  string `form:"captcha" json:"captcha" binding:"required"`
	Captcha string `form:"captcha" json:"captcha"`
}

// User godoc
// @Summary 默认的 user 接口
// @Description 用户登录接口
// @Tags Login
// @Accept json
// @Produce json
// @param data body login true "data"
// @Success 200 {object} response.Response
// @Router /cloud/login [post]
func user_login(c *gin.Context) {
	var user models.User
	var loginVals login
	if err := c.BindJSON(&loginVals); err != nil {
		c.JSON(400, "必填字段未填写")
		return
	}
	// todo判断账户是否存在
	if err := services.DB().First(&user, "account = ?", loginVals.Account).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		response.Respond(c, -1, "该账户尚未注册", errors.New("该账户尚未注册"))
		return
	} else if err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password)); err != nil {
			// fmt.Println(user.Password, string(hashedPassword))
			response.Respond(c, -1, "密码错误", errors.New("密码错误"))
			return
		}
	} else {
		response.ErrJSON(c, nil, err)
		return
	}
	token, err := webserver.GenToken(user)
	if err != nil {
		response.ErrJSON(c, nil, err)
		return
	}
	responseUser := map[string]interface{}{
		"user":  &user,
		"token": token,
	}
	fmt.Println(token)
	response.JSON(c, &responseUser)
	return
}

// Login godoc
// @Summary 默认的 captcha 接口
// @Description 用户登录接口
// @Tags Login
// @Accept json
// @Produce json
// @param v path number true "data"
// @Success 200 {object} response.Response
// @Router /cloud/captcha [get]
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(length) == 1 {
		l = length[0]
	}
	if len(length) == 2 {
		w = length[1]
	}
	if len(length) == 3 {
		h = length[2]
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	fmt.Println(l, w, h, "what is captchaid")
	session.Set("captcha", captchaId)
	_ = session.Save()
	// _ = webserver.Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
	response.JSON(c, "ch")
	return
}
