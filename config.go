package main

import (
	"html/template"
	"github.com/Masterminds/sprig/v3"
)

type Config struct {
	templates map[string]*template.Template
	Title string `yaml:"title"`
	Description string `yaml:"description"`
	DiscordURL string `yaml:"discordURL"`
	InstagramURL string `yaml:"instagramURL"`
	TwitterURL string `yaml:"twitterURL"`
	GithubURL string `yaml:"githubURL"`
	SocietyURL string `yaml:"societyURL"`
	Email string `yaml:"email"`
	Exec []execMember `yaml:"exec"`
	Navbar []navbarLink `yaml:"navbar"`
	Achievements []achievement `yaml:"achievements"`
	GoogleCalendarAPIKey string `yaml:"googleCalendarAPIKey"`
	GoogleCalendarID string `yaml:"googleCalendarID"`
	GoogleCalendarURL string `yaml:"googleCalendarURL"`
	Posts []post `yaml:"posts"`
	Resources []resource `yaml:"resources"`
	Sponsors []sponsor `yaml:"sponsors"`
	ConfSponsors []confsponsor `yaml:"confSponsors"`
	ConferenceSchedule []conferenceSchedule `yaml:"conferenceSchedule"`

}

type post struct {
	Title string `yaml:"title"`
	Author string `yaml:"author"`
	Link string `yaml:"link"`
}

type resource struct {
	Title string `yaml:"title"`
	Link string `yaml:"link"`
	Description string `yaml:"description"`
}

type execMember struct {
	Name string `yaml:"name"`
	Role string `yaml:"role"`
}

type navbarLink struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}

type achievement struct {
	Title string `yaml:"title"`
	Year string `yaml:"year"`
}

type sponsor struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
	Image string `yaml:"image"`
	Description string `yaml:"description"`
}

type confsponsor struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
	Image string `yaml:"image"`
	Description string `yaml:"description"`
}

type conferenceEvent struct {
	Title       string `yaml:"title"`
	Speaker     string `yaml:"speaker"`
	Room        string `yaml:"room"`
	Start       string `yaml:"start"`
	End         string `yaml:"end"`
}


type conferenceSchedule struct {
	Rooms     []string           `yaml:"rooms"`
	Timeslots []string           `yaml:"timeslots"`
	Events    []conferenceEvent  `yaml:"events"`
}

tmpl := template.New("misc0nfig.tmpl").Funcs(sprig.FuncMap())
tmpl, err := tmpl.ParseFiles("misc0nfig.tmpl")
if err != nil {
    log.Fatal(err)
}
