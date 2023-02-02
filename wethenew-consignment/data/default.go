package data

// Info struct
type Info struct {
	Results []struct {
		ID int `json:"id"`
	} `json:"results"`
}

// WTN struct
type WTN struct {
	Results []Results `json:"results"`
}
type Results struct {
	Type  string   `json:"type"`
	Brand string   `json:"brand"`
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Sizes []string `json:"sizes"`
	Image string   `json:"image"`
}
