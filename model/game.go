package model

//Game - game information
type Game struct {
	ID              string            `json:"id" bson:"id"`
	Name            string            `json:"name" bson:"name"`
	NameTranslation map[string]string `json:"translation_name" bson:"translation_name"`
	Type            string            `json:"type" bson:"type"`
	MainGame        int               `bson:"main_game" bson:"main_game"`
	DLC             []int             `json:"dlc" bson:"dlc"`
	Platforms       map[string]bool   `json:"platforms" bson:"platforms"`
	ReleaseDate     struct {
		ComingSoon bool `json:"coming_soon" bson:"coming_soon"`
		Date       int  `json:"date" bson:"date"`
	} `json:"release_date" bson:"release_date"`
	Categories []struct {
		ID          int    `json:"id" bson:"id"`
		Description string `json:"description" bson:"description"`
	} `json:"categories" bson:"categories"`
	Genres []struct {
		ID          string `json:"id" bson:"id"`
		Description string `json:"description" bson:"description"`
	} `json:"genres" bson:"genres"`
	AboutTheGame     string   `json:"about_the_game" bson:"about_the_game"`
	ShortDescription string   `json:"short_description" bson:"short_description"`
	Website          string   `json:"website" bson:"website"`
	Developers       []string `json:"developers" bson:"developers"`
	Metacritic       struct {
		Score int    `json:"score" bson:"score"`
		URL   string `json:"url" bson:"url"`
	} `json:"metacritic" bson:"metacritic"`
	StoreIds map[string]string `json:"store_ids" bson:"store_ids"`
}

type GameSmall struct {
	ID              string            `json:"id" bson:"id"`
	Name            string            `json:"default_name" bson:"name"`
	NameTranslation map[string]string `json:"translation_name" bson:"name_translation"`
	Type            string            `json:"type" bson:"type"`
	MainGame        int               `bson:"main_game" bson:"main_game"`
	DLC             []int             `json:"dlc" bson:"dlc"`
	Platforms       []string          `json:"platforms" bson:"platforms"`
	ReleaseDate     struct {
		ComingSoon bool `json:"coming_soon" bson:"coming_soon"`
		Date       int  `json:"date" bson:"date"`
	} `json:"release_date" bson:"release_date"`
	Categories []struct {
		ID          string `json:"id" bson:"id"`
		Description string `json:"description" bson:"description"`
	} `json:"categories" bson:"categories"`
	Genres []struct {
		ID          string `json:"id" bson:"id"`
		Description string `json:"description" bson:"description"`
	} `json:"genres" bson:"genres"`
	AboutTheGame     string   `json:"about_the_game" bson:"about_the_game"`
	ShortDescription string   `json:"short_description" bson:"short_description"`
	Website          string   `json:"website" bson:"website"`
	Developers       []string `json:"developers" bson:"developers"`
	Critics          []struct {
		ID    string `json:"id" bson:"id"`
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score" bson:"score"`
		URL   string `json:"url" bson:"url"`
	} `json:"critics" bson:"critics"`
	StoreIds map[string]string `json:"store_ids" bson:"store_ids"`
}
