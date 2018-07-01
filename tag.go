package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const tagUrl = "https://www.instagram.com/explore/tags/%s/?__a=1"

type InstagTagResponse struct {
	Graphql struct {
		Hashtag struct {
			EdgeHashtagToMedia struct {
				Edges []Edge `json:"edges"`
			} `json:"edge_hashtag_to_media"`
		} `json:"hashtag"`
	} `json:"graphql"`
}

func GetFromTag(tag string) (string, error) {
	url := fmt.Sprintf(tagUrl, tag)
	response, err := Get(url)
	if err != nil {
		return "", err
	}
	result := &InstagTagResponse{}
	err = json.Unmarshal(response, result)
	if err != nil {
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
