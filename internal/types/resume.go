package types

type Experience struct {
	Title       string
	Company     string
	Location    string
	Period      string
	Description string
	Tags        []string
}

type Skill struct {
	Name        string
	Category    string
	Level       int
	Description string
	YearsOfExp  float32
	Tags        []string
}

type ResumeData struct {
	Experiences []Experience
	Skills      []Skill
}
