package service

import (
	"encoding/json"
	"fmt"
	"github.com/TISUnion/most-simple-mcd/constant"
	json_struct "github.com/TISUnion/most-simple-mcd/json-struct"
	"github.com/TISUnion/most-simple-mcd/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	TOKEN_HEADER_NAME = "X-Token"
)

func RegisterRouter() {
	router := GetGinServerInstanceRouter()
	// 用户登陆
	router.POST("/user/login", userLogin)
	// 用户注销
	router.POST("/user/logout", userLogout)
	// 需要登陆才能请求的接口
	v1 :=
		router.Group("/api/v1")
	{
		// 登陆验证中间件
		v1.Use(func(c *gin.Context) {
			token := c.GetHeader(TOKEN_HEADER_NAME)
			dbtoken := GetFromDatabase(constant.DEFAULT_TOKEN_DB_KEY)
			if token == dbtoken && dbtoken != "" {
				c.Next()
			} else {
				c.JSON(http.StatusOK, getResponse(constant.TOKEN_FAILED, constant.TOKEN_FAILED_MESSAGE, ""))
				c.Abort()
			}
		})
		// 获取用户信息
		v1.GET("/user/info", getInfo)
		// websocket实时监听服务端耗费资源
		v1.GET("/server/resources/listen/:serverId", serversResourcesListen)

	}
}

// 用户登录
func userLogin(c *gin.Context) {
	var reqInfo json_struct.AdminUser
	if err := c.BindJSON(&reqInfo); err != nil {
		c.JSON(http.StatusOK, getResponse(constant.HTTP_PARAMS_ERROR, constant.HTTP_PARAMS_ERROR_MESSAGE, ""))
		return
	}
	var adminObj json_struct.AdminUser
	adminJson := GetFromDatabase(constant.DEFAULT_ACCOUNT_DB_KEY)
	if adminJson == "" {
		adminObj = *setDefaultAccount()
	} else {
		if err := json.Unmarshal([]byte(adminJson), &adminObj); err != nil {
			c.JSON(http.StatusOK, getResponse(456789, fmt.Sprintf("%v", err), ""))
			return
		}
	}
	if reqInfo.Account == adminObj.Account && utils.Md5(reqInfo.Password) == adminObj.Password {
		token := GetFromDatabase(constant.DEFAULT_TOKEN_DB_KEY)
		if token == "" {
			token = utils.Md5(fmt.Sprintf("%v%s", time.Now().UnixNano(), reqInfo.Password))
			SetWiteTTLFromDatabase(constant.DEFAULT_TOKEN_DB_KEY, token, constant.DEFAULT_TOKEN_DB_KEY_EXPIRE)
		}
		c.JSON(http.StatusOK, getResponse(constant.HTTP_OK, "", json_struct.UserToken{Token: token}))
		return
	}
	c.JSON(http.StatusOK, getResponse(constant.PASSWORD_ERROR, constant.PASSWORD_ERROR_MESSAGE, ""))
}

// 用户信息获取
func getInfo(c *gin.Context) {
	var adminObj json_struct.AdminUser
	adminJson := GetFromDatabase(constant.DEFAULT_ACCOUNT_DB_KEY)
	_ = json.Unmarshal([]byte(adminJson), &adminObj)
	adminObj.Password = ""
	c.JSON(http.StatusOK, getResponse(constant.HTTP_OK, "", adminObj))
}

// 用户注销
func userLogout(c *gin.Context) {
	SetFromDatabase(constant.DEFAULT_TOKEN_DB_KEY, "")
	c.JSON(http.StatusOK, getResponse(constant.HTTP_OK, "", ""))
}

// 服务端消耗资源监听 TODO
func serversResourcesListen(c *gin.Context) {
	serverId, ok := c.Params.Get("serverId")
	if serverId == "" || !ok {

	}
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

// 修改

// 设置初始账号密码
func setDefaultAccount() *json_struct.AdminUser {
	pwd := utils.Md5(constant.DEFAULT_PASSWORD)
	adminObj := &json_struct.AdminUser{
		Nickname: constant.DEFAULT_ACCOUNT,
		Account:  constant.DEFAULT_ACCOUNT,
		Password: pwd,
		Roles:    nil,
	}
	adminJson, _ := json.Marshal(adminObj)
	SetFromDatabase(constant.DEFAULT_ACCOUNT_DB_KEY, string(adminJson))
	return adminObj
}

// 返回数据格式化
func getResponse(code int, message string, data interface{}) gin.H {
	responseData := make(gin.H)
	responseData["code"] = code
	responseData["message"] = message
	responseData["data"] = data
	return responseData
}
