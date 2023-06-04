package engine

import (
	"net/http"
	"net/http/cookiejar"
)

type Ticket struct {
	Offers []Offers `json:"offers,omitempty"`
}

type Offers struct {
	ID         string `json:"id,omitempty"`
	PriceLevel int    `json:"priceLevel,omitempty"`
	PriceType  int    `json:"priceType,omitempty"`
	Type       string `json:"type,omitempty"`
}

//WEBHOOK MESSAGE STRUCTURE//

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Embeds struct {
	Author      Author    `json:"author"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Fields      []Fields  `json:"fields"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Image       Image     `json:"image"`
	Footer      Footer    `json:"footer"`
}
type Top struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []Embeds `json:"embeds"`
}

var cookieJar, _ = cookiejar.New(nil)
var clientWebhook = &http.Client{
	Jar: cookieJar,
}
