package webhook

import (
	"net/http"
	"net/http/cookiejar"
)

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

const (
	Image_URL = "https://cdn.discordapp.com/attachments/965899789021642752/965899835570016286/DBFF8755-874B-4436-B79A-0C02DDBBEBBA.jpg"
)

var cookieJar, _ = cookiejar.New(nil)
var clientWebhook = &http.Client{
	Jar: cookieJar,
}
