package main

import (
	"html/template"
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
	Sponsors []sponsor `yaml:"sponsors"`
	ConfSponsors []confsponsor `yaml:"confSponsors"`
	Schedule []conferenceSchedule `yaml:"-"`
	Rooms []string `yaml:"rooms"`
	Timeslots []string `yaml:"timeslots"`
	Events []conferenceEvent  `yaml:"events"`
	SkipMap map[string]bool `yaml:"-"`
	Timetable timetable `yaml:"-"`
	FirstDayOfTerm string `yaml:"term_start"`
	TermLengthWeeks int `yaml:"term_weeks_length"`
	Opportunities opportunities `yaml:"opportunities`
	Opportunity []opportunity `yaml:"-"`
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
	Level string `yaml:"level"`
	Description string `yaml:"description"`
	Opportunities opportunities `yaml:opportunities`
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
	Company     string `yaml:"company"`
	Room        string `yaml:"room"`
	Start       string `yaml:"start"`
	End         string `yaml:"end"`
	RowSpan     int    `yaml:"-"`
}

type conferenceSchedule struct {
	Rooms     []string           `yaml:"rooms"`
	Timeslots []string           `yaml:"timeslots"`
	Events    []conferenceEvent  `yaml:"events"`
	SkipMap   map[string]bool
}

type Week struct {
	Index		int		`yaml:"-"`
	Date 		string	`yaml:"-"`
}

type currentEvent struct {
	Title       string 	`yaml:"title"`
	Room        string 	`yaml:"room"`
	Start       string 	`yaml:"start"`
	End         string	`yaml:"end"`
	Type		string	`yaml:"type"` // academic, sober, drinking
	Week		int		`yaml:"week"`
	Day			string	`yaml:"day"`
}

type timetable struct {
	Weeks			[]Week				`yaml:"-"`
	CurrentEvents	[]currentEvent		`yaml:"-"`
	Days			[]string			`yaml:"-"`
}

type TemplateData struct {
	Config   Config
	Schedule conferenceSchedule
	Timetable timetable
}

type opportunity struct {
	Role string `yaml:"role"`
	Link string `yaml:"link"`
}

type opportunities struct {
	Internships []opportunity `yaml:"internships"`
	Gradroles []opportunity `yaml:"gradroles"`
}