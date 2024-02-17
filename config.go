package main

import (
	"html/template"
)

type Config struct {
	templates map[string]*template.Template
	Title string `yaml:"title"`
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
}

type execMember struct {
	Name string `yaml:"name"`
	Image string `yaml:"image"`
	Role string `yaml:"role"`
	Enabled bool `yaml:"enabled"`
}

type navbarLink struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}

type achievement struct {
	Title string `yaml:"title"`
	Year string `yaml:"year"`
}
