package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrewfrench/instagram-api-bypass/pkg/common/extract"
	"math/rand"
)

const profileUrl = "https://www.instagram.com/%s"

var randomProfiles = []string{
	"vietnamsexy.girl",
	"vietnam_sexygirl_asian",
	"vsbg.sexy",
	"asian_girls79",
	"asianhotgirlsdaily",
	"asian_girls_cuties",
}

type InstagProfileResponse struct {
	Graphql struct {
		Hashtag struct {
			EdgeHashtagToMedia struct {
				Edges []Edge `json:"edges"`
			} `json:"edge_hashtag_to_media"`
		} `json:"hashtag"`
	} `json:"graphql"`
}

func GetRandomFromProfile() (string, error) {
	profile := randomProfiles[rand.Intn(len(randomProfiles))]
	return GetFromProfile(profile)
}

func GetFromProfile(profile string) (string, error) {
	url := fmt.Sprintf(profileUrl, profile)
	response, err := Get(url)
	if err != nil {
		return "", err
	}
	response, err = extract.ExtractJson(response)
	if err != nil {
		return "", err
	}
	result := &InstagProfileResponse{}
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
