package yadrive

import (
	"log"
	"net/http"
)

func ExampleDownloadPublicLink() {

	client, err := NewClient(http.DefaultClient, BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	link, err := client.DownloadPublicLink("https://yadi.sk/i/-KPBYtDv3PHJ8c", nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(link)
}
