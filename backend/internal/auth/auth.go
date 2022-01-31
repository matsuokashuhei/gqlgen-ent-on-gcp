package auth

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

const USER_KEY = "user"

func CreateFirebaseUser(ctx context.Context, email string, password string) (*auth.UserRecord, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Printf("error initializing App: %v\n", err)
		panic(err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Printf("error initializing Auth: %v\n", err)
		panic(err)
	}
	params := (&auth.UserToCreate{}).Email(email).EmailVerified(false).Password(password)
	user, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
		return nil, err
	}
	return user, nil
}

func Error() error {
	return fmt.Errorf("access defnied")
}

// func CreateUser(ctx context.Context, email string, password string) (*models.User, error) {
// 	fUser, err := CreateFirebaseUser(ctx, email, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	user := &models.User{
// 		Name:        strings.Split(fUser.Email, "@")[0],
// 		FirebaseUID: fUser.UID,
// 	}
// 	repository := repositories.NewUserRepository(r.DB)
// 	_, err = repository.Create(user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
