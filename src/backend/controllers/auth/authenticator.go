package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	rsakey "kube_web/apipem"
	"kube_web/controllers/base"
	"kube_web/models"
	"kube_web/models/response/errors"
	"kube_web/util/hack"
	"kube_web/util/logs"
)

// Authenticator provides interface to authenticate user credentials.
type Authenticator interface {
	// Authenticate ...
	Authenticate(m models.AuthModel) (*models.User, error)
}

var registry = make(map[string]Authenticator)

// Register add different authenticators to registry map.
func Register(name string, authenticator Authenticator) {
	if _, dup := registry[name]; dup {
		logs.Info("authenticator: %s has been registered", name)
		return
	}
	registry[name] = authenticator
}

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
	c.Mapping("CurrentUser", c.CurrentUser)
}

type LoginResult struct {
	Token string `json:"token"`
}

// @Title Login
// @Description user login db || ldap
// @Param	body	body 	models.AuthModel	true		"The app content"
// @Success 200 return id success
// @Failure 403 body is empty
// @router /login [post]
func (c *AuthController) Login() {
	var (
		authMdole models.AuthModel
		user      *models.User
		err       error
	)
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &authMdole)
	if err != nil {
		logs.Error("get body error. %v", err)
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Ctx.Output.Body(hack.Slice("Invalid param"))
		return
	}
	if authMdole.UserType != "ldap" {
		dbauth := DBAuth{}
		user, err = dbauth.Authenticate(authMdole)
		fmt.Println(user, authMdole, err)
		if err != nil {
			logs.Error("用户或密码验证失败. %v", err)
			c.Ctx.Output.SetStatus(http.StatusBadRequest)
			c.Ctx.Output.Body(hack.Slice("用户或密码验证失败"))
			return
		}
	} else {
		ldap := LDAPAuth{}
		user, err = ldap.Authenticate(authMdole)
		if err != nil {
			logs.Error("用户或密码验证失败. %v", err)
			c.Ctx.Output.SetStatus(http.StatusBadRequest)
			c.Ctx.Output.Body(hack.Slice("用户或密码验证失败"))
			return
		}
	}
	now := time.Now()
	user.LastIp = c.Ctx.Input.IP()
	user.LastLogin = &now
	user, err = models.UserModel.EnsureUser(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body(hack.Slice(err.Error()))
		return
	}

	// default token exp time is 3600s.
	expSecond := beego.AppConfig.DefaultInt64("TokenLifeTime", 86400)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		// 签发者
		"iss": "wayne",
		// 签发时间
		"iat": now.Unix(),
		"exp": now.Add(time.Duration(expSecond) * time.Second).Unix(),
		"aud": user.Name,
	})

	apiToken, err := token.SignedString(rsakey.RsaPrivateKey)
	if err != nil {
		logs.Error("create token form rsa private key  error.", rsakey.RsaPrivateKey, err.Error())
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body(hack.Slice(err.Error()))
		return
	}

	loginResult := LoginResult{
		Token: apiToken,
	}
	c.Data["json"] = base.Result{Data: loginResult}
	c.ServeJSON()
}

func (c *AuthController) Logout() {

}

// @Title CurrentUser
// @Description Get user info
// @Param	header	Header 	string	true		"Authorization"
// @Success 200 return id success
// @Failure 403 body is empty
// @router /currentuser [get]
func (c *AuthController) CurrentUser() {
	c.Controller.Prepare()
	authString := c.Ctx.Input.Header("Authorization")
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		logs.Info("AuthString invalid:", authString)
		c.CustomAbort(http.StatusUnauthorized, "Token Invalid ! ")
	}
	tokenString := kv[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return rsakey.RsaPublicKey, nil
	})
	errResult := errors.ErrorResult{}
	switch err.(type) {
	case nil: // no error
		if !token.Valid { // but may still be invalid
			errResult.Code = http.StatusUnauthorized
			errResult.Msg = "Token Invalid ! "
		}

	case *jwt.ValidationError: // something was wrong during the validation
		errResult.Code = http.StatusUnauthorized
		errResult.Msg = err.Error()

	default: // something else went wrong
		errResult.Code = http.StatusInternalServerError
		errResult.Msg = err.Error()
	}

	if err != nil {
		c.CustomAbort(errResult.Code, errResult.Msg)
	}

	claim := token.Claims.(jwt.MapClaims)
	aud := claim["aud"].(string)
	user, err := models.UserModel.GetUserDetail(aud)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, err.Error())
	}

	c.Data["json"] = base.Result{Data: user}
	c.ServeJSON()
}
