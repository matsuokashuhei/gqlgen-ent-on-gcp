package middleware

import (
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/matsuokashuhei/landin/internal/auth"
	"github.com/matsuokashuhei/landin/internal/repositories"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		app, err := firebase.NewApp(c, nil)
		if err != nil {
			log.Printf("error initializing App: %v\n", err)
			panic(err)
		}
		client, err := app.Auth(c)
		if err != nil {
			log.Printf("error initializing Auth: %v\n", err)
			panic(err)
		}

		h := c.Request.Header.Get("Authorization")
		var t string
		if s := strings.Split(h, "Bearer "); len(s) > 1 {
			t = s[1]
		} else {
			return
		}

		log.Printf("t: %s", t)

		token, err := client.VerifyIDToken(c, t)
		if err != nil {
			log.Printf("VerifyIDToken returned the error: %v\n", err)
			return
		}
		repository := repositories.NewUserRepository(db)
		user, err := repository.FindByAuth(token.UID)
		if err != nil {
			log.Printf("FindByAuth returned the error: %v\n", err)
			return
		}
		log.Printf("user: %v\n", user)
		c.Set(auth.USER_KEY, user)
		c.Next()
	}
}
