package utils

import (
	"fmt"
	. "github.com/mlg-dev-rus/game-network-core-lib/model"
	"strconv"
)

const Windows = "windows"
const Linux = "linux"
const Mac = "mac"

func GameToSmallModel(src Game) GameSmall {
	var srcSmall GameSmall

	srcSmall.ID = src.ID
	srcSmall.Name = src.Name
	srcSmall.NameTranslation = src.NameTranslation
	srcSmall.Type = src.Type
	srcSmall.DLC = src.DLC
	for key, value := range src.Platforms {
		if value {
			srcSmall.Platforms = append(srcSmall.Platforms, key)
		}
	}
	srcSmall.ReleaseDate = src.ReleaseDate

	for i := 0; i < len(src.Categories); i++ {
		srcSmall.Categories = append(srcSmall.Categories, struct {
			ID          string `json:"id" bson:"id"`
			Description string `json:"description" bson:"description"`
		}{
			strconv.Itoa(src.Categories[i].ID),
			src.Categories[i].Description,
		})
	}

	srcSmall.Genres = src.Genres
	srcSmall.AboutTheGame = src.AboutTheGame
	srcSmall.ShortDescription = src.ShortDescription
	srcSmall.Website = src.Website
	srcSmall.Developers = src.Developers
	srcSmall.Critics = append(srcSmall.Critics, struct {
		ID    string `json:"id" bson:"id"`
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score" bson:"score"`
		URL   string `json:"url" bson:"url"`
	}{
		"Metacritic",
		"Metacritic",
		src.Metacritic.Score,
		src.Metacritic.URL,
	})
	srcSmall.StoreIds = src.StoreIds
	return srcSmall
}

func GameToBigModel(srcSmall GameSmall) Game {
	var src Game

	src.ID = srcSmall.ID
	src.Name = srcSmall.Name
	src.NameTranslation = srcSmall.NameTranslation
	src.Type = srcSmall.Type
	src.DLC = srcSmall.DLC
	src.Platforms = make(map[string]bool)
	for _, value := range srcSmall.Platforms {
		switch value {
		case Windows:
			src.Platforms[Windows] = true
		case Linux:
			src.Platforms[Linux] = true
		case Mac:
			src.Platforms[Mac] = true
		}
	}
	src.ReleaseDate = srcSmall.ReleaseDate

	for i := 0; i < len(srcSmall.Categories); i++ {
		intID, err := strconv.Atoi(srcSmall.Categories[i].ID)
		if err != nil {
			fmt.Print(err)
		}
		src.Categories = append(src.Categories, struct {
			ID          int    `json:"id" bson:"id"`
			Description string `json:"description" bson:"description"`
		}{
			intID,
			srcSmall.Categories[i].Description,
		})
	}

	src.Genres = srcSmall.Genres
	src.AboutTheGame = srcSmall.AboutTheGame
	src.ShortDescription = srcSmall.ShortDescription
	src.Website = srcSmall.Website
	src.Developers = srcSmall.Developers

	for i := 0; i < len(srcSmall.Critics); i++ {
		src.Metacritic.Score = srcSmall.Critics[i].Score
		src.Metacritic.URL = srcSmall.Critics[i].URL
	}

	src.StoreIds = srcSmall.StoreIds
	return src
}

func GameToSmallModelArr(src []Game) []GameSmall {
	var srcSmall []GameSmall
	for _, value := range src {
		srcSmall = append(srcSmall, GameToSmallModel(value))
	}
	return srcSmall
}

func GameToBigModelArr(srcSmall []GameSmall) []Game {
	var src []Game
	for _, value := range srcSmall {
		src = append(src, GameToBigModel(value))
	}
	return src
}
