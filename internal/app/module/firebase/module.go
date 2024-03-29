package firebaseModule

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Module represents the Firebase module.
type Module struct {
	FirebaseApp *firebase.App
}

// NewModule creates a new Firebase Module instance.
func NewModule() *Module {
	ctx := context.Background()
	credentialPath := "firebase-credentials.json"
	opt := option.WithCredentialsFile(credentialPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
		return nil
	}

	return &Module{
		FirebaseApp: app,
	}
}
