package masta_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"spiderden.org/masta"
)

func ExampleRegisterApp() {
	app, err := masta.RegisterApp(context.Background(), &masta.AppConfig{
		Server:     "https://arachnid.town",
		ClientName: "client-name",
		Scopes:     "read write follow",
		Website:    "https://spiderden.org/projects/masta",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("client-id    : %s\n", app.ClientID)
	fmt.Printf("client-secret: %s\n", app.ClientSecret)
}

func ExampleClient() {
	c := masta.NewClient(&masta.Config{
		Server:       "https://arachnid.town",
		ClientID:     "client-id",
		ClientSecret: "client-secret",
	})
	err := c.Authenticate(context.Background(), "your-email", "your-password")
	if err != nil {
		log.Fatal(err)
	}
	timeline, err := c.GetTimelineHome(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(timeline) - 1; i >= 0; i-- {
		fmt.Println(timeline[i])
	}
}

func ExamplePagination() {
	c := masta.NewClient(&masta.Config{
		Server:       "https://arachnid.town",
		ClientID:     "client-id",
		ClientSecret: "client-secret",
	})
	var followers []*masta.Account
	var pg masta.Pagination
	for {
		fs, err := c.GetAccountFollowers(context.Background(), "1", &pg)
		if err != nil {
			log.Fatal(err)
		}
		followers = append(followers, fs...)
		if pg.MaxID == "" {
			break
		}
		time.Sleep(10 * time.Second)
	}
	for _, f := range followers {
		fmt.Println(f.Acct)
	}
}
