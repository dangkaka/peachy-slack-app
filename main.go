package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"math/rand"
	"os"
)

const url = "https://www.instagram.com/explore/tags/%s/?__a=1"

var randomTags = []string{"vietnamesegirl", "dantocmongto"}

var tags = map[string]string {
	"vnbabe": "vietnamesegirl",
}

type InstagTagResponse struct {
	Graphql struct {
		Hashtag struct {
			EdgeHashtagToMedia struct {
				Edges []Edge `json:"edges"`
			} `json:"edge_hashtag_to_media"`
		} `json:"hashtag"`
	} `json:"graphql"`
}

type Edge struct {
	Node struct {
		DisplayUrl string `json:"display_url"`
	} `json:"node"`
}

func getInstaImg(tag string) (string, error) {
	if tag == "" {
		tag = randomTags[rand.Intn(len(randomTags))]
	}

	rs, err := http.Get(fmt.Sprintf(url, tag))

	if err != nil {
		return "", err
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		return "", err
	}

	var result InstagTagResponse

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return "", err
	}

	edges := result.Graphql.Hashtag.EdgeHashtagToMedia.Edges
	max := len(edges)
	if max > 100 {
		max = 100
	}
	randomEdge := edges[rand.Intn(max)]

	return randomEdge.Node.DisplayUrl, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, `{"error": "wrong input"}`)
		return
	}
	tag := r.FormValue("text")
	instaTag := ""
	if tag != "" {
		if val, ok := tags[tag]; ok {
			instaTag = val
		} else {
			fmt.Fprint(w, `{"error": "no tag found"}`)
			return
		}
	}
	imgUrl, err := getInstaImg(instaTag)
	if err != nil {
		fmt.Fprint(w, `{"error": "no img found"}`)
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
