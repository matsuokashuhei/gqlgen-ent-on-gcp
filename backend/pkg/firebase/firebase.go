package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func initializeAppDefault() (*firebase.App, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}
	return app, nil
}

func CreateUser(ctx context.Context, email string, password string) (*auth.UserRecord, error) {
	app, err := initializeAppDefault()
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return nil, err
	}
	params := (&auth.UserToCreate{}).Email(email).EmailVerified(false).Password(password)
	user, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
		return nil, err
	}
	return user, nil
}
