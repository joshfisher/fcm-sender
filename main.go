package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	opt := option.WithCredentialsFile("./live-162320-firebase-adminsdk-l6wvs-a6c2d39768.json")
	config := &firebase.Config{ProjectID: "live-162320"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	message := &messaging.Message{
		Topic: "all",
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "5",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Alert: &messaging.ApsAlert{
						Title: "YOOOO",
						Body:  "YOU DA BEST",
					},
				},
			},
		},
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully sent message: ", response)
}
