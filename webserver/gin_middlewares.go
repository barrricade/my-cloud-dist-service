package webserver

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/axiaoxin-com/goutils"
	"github.com/axiaoxin-com/logging"
	"github.com/axiaoxin-com/ratelimiter"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// GinBasicAuth gin 的基础认证中间件
// 加到 gin app 的路由中可以对该路由添加 basic auth 登录验证
// 传入 username 和 password 对可以替换默认的 username 和 password
func GinBasicAuth(args ...string) gin.HandlerFunc {
	username := viper.GetString("basic_auth.username")
	password := viper.GetString("basic_auth.password")
	if len(args) == 2 {
		username = args[0]
		password = args[1]
	}
	logging.Debug(nil, "Basic auth username:"+username+" password:"+password)
	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}

// GinRecovery gin recovery 中间件
// save err in context and abort with recoveryHandler
func GinRecovery(
	recoveryHandler ...func(c *gin.Context, status int, data interface{}, err error, extraMsgs ...interface{}),
) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			status := c.Writer.Status()
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				status = http.StatusInternalServerError
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if brokenPipe {
					// save err in context
					c.Error(fmt.Errorf("Broken pipe: %v\n%s", err, string(debug.Stack())))
					if len(recoveryHandler) > 0 {
						c.Abort()
						recoveryHandler[0](c, status, nil, errors.New(http.StatusText(status)))
						return
					}
					c.AbortWithStatus(status)
					return
				}

				// save err in context when panic
				c.Error(fmt.Errorf("Recovery from panic: %v\n%s", err, string(debug.Stack())))
			}

			// 状态码大于 400 的以 status handler 进行响应
			if status >= 400 {
				// 有 handler 时使用 handler 返回
				if len(recoveryHandler) > 0 {
					c.Abort()
					recoveryHandler[0](c, status, nil, errors.New(http.StatusText(status)))
					return
				}
				// 否则只返回状态码
				c.AbortWithStatus(status)
				return
			}
		}()
		c.Next()
	}
}

// GinLogMiddleware 日志中间件
// 可根据实际需求自行修改定制
func GinLogMiddleware() gin.HandlerFunc {
	logging.CtxLoggerName = logging.Ctxkey("ctx_logger")
	logging.TraceIDKeyname = logging.Ctxkey("trace_id")
	logging.TraceIDPrefix = "logging_"
	loggerMiddleware := logging.GinLoggerWithConfig(logging.GinLoggerConfig{
		SkipPaths:           viper.GetStringSlice("logging.access_logger.skip_paths"),
		SkipPathRegexps:     viper.GetStringSlice("logging.access_logger.skip_path_regexps"),
		EnableDetails:       viper.GetBool("logging.access_logger.enable_details"),
		EnableContextKeys:   viper.GetBool("logging.access_logger.enable_context_keys"),
		EnableRequestHeader: viper.GetBool("logging.access_logger.enable_request_header"),
		EnableRequestForm:   viper.GetBool("logging.access_logger.enable_request_form"),
		EnableRequestBody:   viper.GetBool("logging.access_logger.enable_request_body"),
		EnableResponseBody:  viper.GetBool("logging.access_logger.enable_response_body"),
		SlowThreshold:       viper.GetDuration("logging.access_logger.slow_threshold") * time.Millisecond,
		TraceIDFunc:         nil,
		InitFieldsFunc:      nil,
		Formatter:           nil,
	})
	return loggerMiddleware
}

// GinRatelimitMiddleware 限频中间件
// 需先实现对应的 TODO ，可根据实际需求自行修改定制
func GinRatelimitMiddleware() gin.HandlerFunc {
	limiterConf := ratelimiter.GinRatelimiterConfig{
		// TODO: you should implement this function by yourself
		LimitKey: ratelimiter.DefaultGinLimitKey,
		// TODO: you should implement this function by yourself
		LimitedHandler: ratelimiter.DefaultGinLimitedHandler,
		// TODO: you should implement this function by yourself
		TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
			// 每1秒填充1个token，桶容量为100（1秒最多100次请求）
			return time.Second * 1, 100
		},
	}

	// 根据 viper 中的配置信息选择限流类型
	var limiterMiddleware gin.HandlerFunc
	limiterType := strings.ToLower(viper.GetString("ratelimiter.type"))
	logging.Info(nil, "enable ratelimiter with type: "+limiterType)
	if strings.HasPrefix(limiterType, "redis.") {
		which := strings.TrimPrefix(limiterType, "redis.")
		if rdb, err := goutils.RedisClient(which); err != nil {
			panic("redis ratelimiter does not work. get redis client error:" + err.Error())
		} else {
			limiterMiddleware = ratelimiter.GinRedisRatelimiter(rdb, limiterConf)
		}
	} else {
		limiterMiddleware = ratelimiter.GinMemRatelimiter(limiterConf)
	}
	return limiterMiddleware
}

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "hanyun"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, //seconds
		Path:   "/",
	})
	return store
}
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
	session.Set("captcha", captchaId)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if captchaId := session.Get("captcha"); captchaId != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
