package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/models"
	"github.com/redmeros/htrade/web/controllers"
	"github.com/redmeros/htrade/web/helpers"
	wm "github.com/redmeros/htrade/web/models"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			helpers.Bad(c, "No auth header", 403)
			c.Abort()
			return
		}
		vals := strings.Split(header, " ")
		if len(vals) != 2 {
			helpers.Bad(c, "Bad header architecture", 403)
			return
		}
		fullTknStr := vals[1]
		claims := &wm.Claims{}

		tkn, err := jwt.ParseWithClaims(fullTknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Web.Secret), nil
		})
		if err != nil {
			helpers.Bad(c, err.Error(), 403)
			c.Abort()
			return
		}

		if !tkn.Valid {
			helpers.Bad(c, "token not valid", 403)
			return
		}

		db, err := controllers.GetDB(c)
		if err != nil {
			helpers.Bad(c, err.Error(), 500)
			c.Abort()
			return
		}
		var user models.User
		if err = db.Where("Username = ?", claims.Username).Find(&user).Error; err != nil {
			helpers.Bad(c, err.Error(), 403)
		}
		c.Set("user", &user)
		c.Next()
	}
}
