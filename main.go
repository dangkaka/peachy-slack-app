package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "`wrong input`")
		return
	}
	text := r.FormValue("text")
	imgUrl, err := getInstaImg(text)
	if err != nil {
		fmt.Fprint(w, "`no img found`")
		return
	}
	fmt.Fprintf(w, `{"response_type": "in_channel", "text": "%s"}`, imgUrl)
}

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", handler)

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
