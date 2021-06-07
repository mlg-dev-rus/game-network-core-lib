package utils

import (
	"github.com/mlg-dev-rus/game-network-core-lib/model"
)

func GameToSmallModel(src Game) GameSmall {
	var srcSmall GameSmall

	srcSmall.ID = src.ID
	srcSmall.Name = src.Name
	srcSmall.NameTranslation = src.NameTranslation
	srcSmall.Type = src.Type
	srcSmall.DLC = src.DLC
	if src.Platforms.Windows {
		srcSmall.Platforms = append(srcSmall.Platforms, "Windows")
	}
	if src.Platforms.Linux {
		srcSmall.Platforms = append(srcSmall.Platforms, "Linux")
	}
	if src.Platforms.Mac {
		srcSmall.Platforms = append(srcSmall.Platforms, "Mac")
	}
	srcSmall.ReleaseDate = src.ReleaseDate
	srcSmall.Categories = src.Categories
	srcSmall.Genres = src.Genres
	srcSmall.AboutTheGame = src.AboutTheGame
	srcSmall.ShortDescription = src.ShortDescription
	srcSmall.Website = src.Website
	srcSmall.Developers = src.Developers
	srcSmall.Critics = append(srcSmall.Critics, struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Score int    `json:"score"`
		URL   string `json:"url"`
	}{
		"Metacritic",
		"Metacritic",
		src.Metacritic.Score,
		src.Metacritic.URL,
	})
	srcSmall.StoreIds = src.StoreIds
	return srcSmall
}
