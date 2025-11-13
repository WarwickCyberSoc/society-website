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
	Posts []post `yaml:"posts"`
	Resources []resource `yaml:"resources"`
	Sponsors []sponsor `yaml:"sponsors"`
	GoldSponsors []goldSponsor `yaml:"goldSponsors"`
	SilverSponsors []silverSponsor `yaml:"silverSponsors"`
	ConfSponsors []confsponsor `yaml:"confSponsors"`
	Schedule []conferenceSchedule `yaml:"-"`
	Rooms []string `yaml:"rooms"`
	Timeslots []string `yaml:"timeslots"`
	Events []conferenceEvent  `yaml:"events"`
	SkipMap map[string]bool `yaml:"-"`
	Timetable timetable `yaml:"-"`
	TimetableDays []string `yaml:"days_of_the_week"`
	TimetableWeeks []Week `yaml:"weeks_term1_25"`
	TimetableCurrentEvents []currentEvent `yaml:"current_events"`
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
	Opportunities []opportunities `yaml:opportunities`
}

type goldSponsor struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
	Image string `yaml:"image"`
	Description string `yaml:"description"`
}

type silverSponsor struct {
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
	Index		int		`yaml:"index"`
	Date 		string	`yaml:"date"`
}

type currentEvent struct {
	Title       string 	`yaml:"title"`
	Room        string 	`yaml:"room"`
	Start       string 	`yaml:"start"`
	End         string	`yaml:"end"`
	Type		string	`yaml:"type"`
	Sober		bool	`yaml:"sober"`
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