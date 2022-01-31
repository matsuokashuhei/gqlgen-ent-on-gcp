package resolver

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matsuokashuhei/landin/internal/auth"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func IsAuthenticated(ctx context.Context) bool {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		log.Print(err)
		return false
	}
	user, exists := gc.Get(auth.USER_KEY)
	log.Printf("user: %v, exists: %t", user, exists)
	if exists == false {
		return false
	}
	return true
}
