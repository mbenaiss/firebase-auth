package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/auth/api"
	firebaseauth "github.com/auth/auth/firebase"
	"google.golang.org/api/option"

	"github.com/auth/config"
	"github.com/auth/service"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatalf("unable to init config %+v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	options := []option.ClientOption{
		option.WithCredentialsFile(c.GoogleApplicationCredentials),
	}

	app, err := firebase.NewApp(ctx, nil, options...)
	if err != nil {
		log.Panicf("unable to initialize firebase app: %v", err)
	}

	auth, err := firebaseauth.New(ctx, c.FirebaseAPIKey, app)
	if err != nil {
		log.Panicf("unable to initialize firebase auth: %v", err)
	}

	svc := service.New(auth)

	server := api.NewServer(c.Port, svc)

	err = server.Run()
	if err != nil {
		log.Panicf("unable to run server: %v", err)
	}
}
