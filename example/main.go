package main

import (
	"context"
	einstein "einstein-golang-sdk"
	"einstein-golang-sdk/client/operations"
	"einstein-golang-sdk/models"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	client := einstein.NewEinsteinClient(os.Getenv("x-cq-client-id"))
	resp, err := client.Operations.GetZoneRecommendations(&operations.GetZoneRecommendationsParams{
		ZoneName: "test",
		SiteID:   os.Getenv("site-id"),
		Body: operations.GetZoneRecommendationsBody{
			Rules: models.Rules{},
		},
		Context: ctxTimeout,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.Error())
}
