package middleware

import (
	"strings"

	"github.com/Gocyber-world/navigator-demo/model/common/response"
	"github.com/gin-gonic/gin"

	"errors"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
	Expiretime int
}

func NewJWT(signingKey []byte) *JWT {
	return &JWT{
		SigningKey: signingKey,
	}
}

// claim——载荷部分
// ID 由数据库中自增的数字id混淆为字符串
type CustomClaims struct {
	jwt.StandardClaims
	UserID  string `json:"userId"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// create jwt
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// parse jwt
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err == nil && jwtToken != nil {
		// 检查 JWT 有效性
		if claims, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
			return claims, nil
		}
	}
	if jwtToken == nil {
		return nil, errors.New("couldn't handle this token")
	} else {
		return nil, err
	}
}

// 从token中获取claims信息
func (j *JWT) GetClaimsFromToken(tokenString string) (*CustomClaims, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (j *JWT) JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		authCookie, cookieErr := ctx.Cookie("goc-jwt")

		var token string
		// 同时出现时,Header中的JWT优先级会更高
		if authHeader != "" && len(strings.Split(authHeader, " ")) == 2 {
			token = strings.Split(authHeader, " ")[1]
		} else if authHeader == "" && authCookie != "" && cookieErr == nil {
			token = authCookie
		} else {
			response.UnauthorizedWithMessage("No JWT in header or cookie", ctx)
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(token)
		if err != nil {
			response.UnauthorizedWithMessage("Invalid JWT", ctx)
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
