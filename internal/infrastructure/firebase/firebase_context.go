package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

var FirebaseClient *auth.Client

func InitFirebase() *auth.Client {
	opt := option.WithCredentialsFile("C:\\Go_Tutorial\\firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("Error initializing firebase app")
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		fmt.Println("Error initializing firebase auth")
	}

	return client
}
