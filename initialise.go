package main

import (
	"github.com/simonmartyr/toogoodtogogo"
	"log"
	"net/http"
)

func Initialise(username string) {
	client := toogoodtogo.New(
		&http.Client{},
		toogoodtogo.WithUsername(username),
	)
	authErr := client.Authenticate()
	if authErr != nil {
		log.Fatal(authErr)
		return
	}
	log.Println("Authentication successful, creating config")

	favorites, getErr := client.GetFavorites(0, 0, 500, 0, 50)
	if getErr != nil {
		log.Fatal(getErr)
		return
	}
	var favList []Item
	for _, x := range favorites.MobileBucket.Items {
		favList = append(favList, Item{
			Name:         x.DisplayName,
			ItemName:     x.Item.Name,
			ItemId:       x.Item.ItemId,
			Notify:       true,
			LastNotified: "",
		})
	}
	credentials := client.GetCredentials()
	config := Config{
		EmailConfig: EmailConfig{
			To:      username,
			Account: "gmail",
		},
		Credentials: TooGoodToGoCredentials{
			UserId:       credentials.UserId,
			Email:        username,
			AccessToken:  credentials.AccessToken,
			RefreshToken: credentials.RefreshToken,
			Cookie:       credentials.Cookie,
		},
		Items: favList,
	}
	config.WriteConfig()
	log.Println("Config created")
}
