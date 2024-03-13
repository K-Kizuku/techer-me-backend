package entity

type User struct {
	ID          string
	Name        string
	IsOrganizer bool
	ImageURL    string
	Message     string
	Skills      map[string]string
	URLs        map[URLs]string
}

type URLs int

const (
	Twitter URLs = iota
	Github
	X
	Discode
	Mastodon
)
