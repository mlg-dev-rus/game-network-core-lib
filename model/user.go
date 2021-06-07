package model

//Game - game information
type Game struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	NameTranslation struct {
		RU string `json:"ru"`
		DE string `json:"de"`
	} `json:"translation_name"`
	Type      string `json:"type"`
	DLC       []int  `json:"dlc"`
	Platforms struct {
		Windows bool `json:"windows"`
		Mac     bool `json:"mac"`
		Linux   bool `json:"linux"`
	} `json:"platforms"`
	ReleaseDate struct {
		ComingSoon bool `json:"coming_soon"`
		Date       int  `json:"date"`
	} `json:"release_date"`
	Categories []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"categories"`
	Genres []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"genres"`
	AboutTheGame     string   `json:"about_the_game"`
	ShortDescription string   `json:"short_description"`
	Website          string   `json:"website"`
	Developers       []string `json:"developers"`
	Metacritic       struct {
		Score int    `json:"score"`
		URL   string `json:"url"`
	} `json:"metacritic"`
	StoreIds struct {
		Steam string `json:"steam"`
	} `json:"store_ids"`
}

type GameSmall struct {
	ID              string `json:"id"`
	Name            string `json:"default_name"`
	NameTranslation struct {
		RU string `json:"ru"`
		DE string `json:"de"`
	} `json:"translation_name"`
	Type        string   `json:"type"`
	DLC         []int    `json:"dlc"`
	Platforms   []string `json:"platforms"`
	ReleaseDate struct {
		ComingSoon bool `json:"coming_soon"`
		Date       int  `json:"date"`
	} `json:"release_date"`
	Categories []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"categories"`
	Genres []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"genres"`
	AboutTheGame     string   `json:"about_the_game"`
	ShortDescription string   `json:"short_description"`
	Website          string   `json:"website"`
	Developers       []string `json:"developers"`
	Critics          []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Score int    `json:"score"`
		URL   string `json:"url"`
	} `json:"critics"`
	StoreIds struct {
		Steam string `json:"steam"`
	} `json:"store_ids"`
}
