package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

const profileUrl = "https://www.instagram.com/%s"

var randomProfiles = []string{
	"vsbg.sexy",
	"asian_girls79",
	"asianhotgirlsdaily",
	"asian_girls_cuties",
	"vnsexy.collection",
	"vietnamesexybabe",
}

type InstagProfileResponse struct {
	EntryData struct {
		ProfilePage []struct {
			Graphql struct {
				User struct {
					EdgeOwnerToTimelineMedia struct {
						Edges []Edge `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
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
	response, err = ExtractJson(response)
	if err != nil {
		return "", err
	}
	result := &InstagProfileResponse{}
	err = json.Unmarshal(response, result)
	if err != nil {
		return "", err
	}
	edges := result.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges
	max := len(edges)
	if max > 100 {
		max = 100
	}
	randomEdge := edges[rand.Intn(max)]
	return randomEdge.Node.DisplayUrl, nil
}
