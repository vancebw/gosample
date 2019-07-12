package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gosample/api/constants"
	"gosample/config"
	"net/http"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Setup() {
	jwtSecret = []byte(config.AppSetting.JwtSecret)
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = constants.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = constants.ERROR_AUTH
		} else {
			_, err := ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = constants.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = constants.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != constants.SUCCESS {
			RespondJSON(http.StatusUnauthorized, c, code, data)
			c.Abort()
			return
		}

		c.Next()
	}
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-sample",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
