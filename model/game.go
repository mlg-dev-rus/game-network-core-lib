package model

//Game - game information
type Game struct {
	ID                 string            `json:"id" bson:"_id"`
	Name               string            `json:"name" bson:"name"`
	NameTranslation    map[string]string `json:"translation_name" bson:"translation_name"`
	Type               string            `json:"type" bson:"type"`
	MainGame           string            `bson:"main_game" bson:"main_game"`
	DLC                []int             `json:"dlc" bson:"dlc"`
	Platforms          map[string]bool   `json:"platforms" bson:"platforms"`
	ReleaseDate        ReleaseDate       `json:"release_date" bson:"release_date"`
	Categories         []Category        `json:"categories" bson:"categories"`
	Genres             []Gener           `json:"genres" bson:"genres"`
	AboutTheGame       string            `json:"about_the_game" bson:"about_the_game"`
	ShortDescription   string            `json:"short_description" bson:"short_description"`
	Website            string            `json:"website" bson:"website"`
	Developers         []string          `json:"developers" bson:"developers"`
	Metacritic         Metacritic        `json:"metacritic" bson:"metacritic"`
	StoreIds           map[string]string `json:"store_ids" bson:"store_ids"`
	SupportedLanguages string            `json:"supported_languages"`

	//price
	IsFree    bool  `json:"is_free" bson:"is_free"`
	LastPrice int64 `json:"last_price" bson:"last_price"`

	//images
	Screenshots []Screenshot `json:"screenshots" bson:"screenshots"`
	Background  string       `json:"background" bson:"background"`
	HeaderImage string       `json:"header_image" bson:"header_image"`
}

//ReleaseDate дата релиза игры
type ReleaseDate struct {
	ComingSoon bool  `json:"coming_soon" bson:"coming_soon"`
	Date       int64 `json:"date" bson:"date"`
}

//Metacritic данные о критики с метакритикс
type Metacritic struct {
	Score int    `json:"score" bson:"score"`
	URL   string `json:"url" bson:"url"`
}

//Category категории игры
type Category struct {
	ID          int    `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
}

//Screenshot скрины игры из steam
type Screenshot struct {
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

//Gener особенности
type Gener struct {
	ID          string `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
}

type Price struct {
	Cost   int64  `bson:"cost" json:"cost"`
	Date   int64  `bson:"date" json:"date"`
	GameID string `bson:"game_id" json:"game_id"`
}

//todo переделать вместе с Кириллом
type GameSmall struct {
	ID              string            `json:"id" bson:"id"`
	Name            string            `json:"default_name" bson:"name"`
	NameTranslation map[string]string `json:"translation_name" bson:"name_translation"`
	Type            string            `json:"type" bson:"type"`
	MainGame        int               `bson:"main_game" bson:"main_game"`
	DLC             []int             `json:"dlc" bson:"dlc"`
	Platforms       []string          `json:"platforms" bson:"platforms"`
	ReleaseDate     ReleaseDate       `json:"release_date" bson:"release_date"`
	Categories      []struct {
		ID          string `json:"id" bson:"id"`
		Description string `json:"description" bson:"description"`
	} `json:"categories" bson:"categories"`
	Genres           []Gener  `json:"genres" bson:"genres"`
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
