package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func getInstaImg(text string) (string, error) {
	var imgUrl string
	var err error
	if len(text) <= 1 {
		imgUrl, err = GetRandomFromProfile()
	} else if text[:1] == "@" {
		imgUrl, err = GetFromProfile(text[1:])
	} else if text[:1] == "#" {
		imgUrl, err = GetFromTag(text[1:])
	}
	return imgUrl, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "`Wrong input`")
		return
	}

	//strip string
	limitedChannels := strings.Replace(os.Getenv("LIMITED_CHANNELS"), " ", "", -1)
	if limitedChannels != "" {
		limitedChannelsArr := strings.Split(limitedChannels, ",")
		if !contains(limitedChannelsArr, r.FormValue("channel_name")) {
			fmt.Fprint(w, "`No permission`")
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	text := r.FormValue("text")
	imgUrl, err := getInstaImg(text)
	if err != nil {
		fmt.Fprint(w, "`No img found`")
		return
	}
	fmt.Fprintf(w, `{"response_type": "in_channel", "attachments": [{"fields": [{"title": "%s"}],"image_url": "%s"}]}`, text, imgUrl)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", handler)

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func contains(arr []string, str string) bool {
	for _, value := range arr {
		if value == str {
			return true
		}
	}
	return false
}
