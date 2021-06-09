package utils

import (
	. "github.com/mlg-dev-rus/game-network-core-lib/model"
)

const Windows = "Windows"
const Linux = "Linux"
const Mac = "Mac"

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
	srcSmall.Categories = src.Categories
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
	for _, value := range srcSmall.Platforms {
		if value == Windows {
			src.Platforms[Windows] = true
		} else {
			src.Platforms[Windows] = false
		}
		if value == Linux {
			src.Platforms[Linux] = true
		} else {
			src.Platforms[Linux] = false
		}
		if value == Mac {
			src.Platforms[Mac] = true
		} else {
			src.Platforms[Mac] = false
		}
	}
	src.ReleaseDate = srcSmall.ReleaseDate
	src.Categories = srcSmall.Categories
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
	for i := 0; i < len(src); i++ {
		srcSmall[i] = GameToSmallModel(src[i])
	}
	return srcSmall
}

func GameToBigModelArr(srcSmall []GameSmall) []Game {
	var src []Game
	for i := 0; i < len(srcSmall); i++ {
		src[i] = GameToBigModel(srcSmall[i])
	}
	return src
}
