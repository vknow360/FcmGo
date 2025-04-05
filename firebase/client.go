package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase(serviceAccountJSON []byte) error {
	opt := option.WithCredentialsJSON(serviceAccountJSON)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}
	App = app
	return nil
}
