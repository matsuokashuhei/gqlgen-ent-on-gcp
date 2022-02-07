package middleware

import (
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func Auth(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		app, err := firebase.NewApp(c, nil)
		if err != nil {
			log.Printf("error initializing App: %v\n", err)
			panic(err)
		}
		auth, err := app.Auth(c)
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

		token, err := auth.VerifyIDToken(c, t)
		if err != nil {
			log.Printf("VerifyIDToken returned the error: %v\n", err)
			return
		}
		repository := repositories.NewUserRepository(client)
		user, err := repository.FindByAuth(c, token.UID)
		if err != nil {
			log.Printf("err: %v\n", err)
			return
		}
		log.Printf("user: %v\n", user)
		c.Set("user", user)
		c.Next()
	}
}

func AuthError() error {
	return fmt.Errorf("access defnied")
}
