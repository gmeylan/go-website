package types

import "time"

type BlogPost struct {
	Title   string
	Date    time.Time
	Content string
	Slug    string
	Raw     map[string]string
	Author  string
	Tags    []string
}

type TagInfo struct {
	Posts []BlogPost
	Count int
	Name  string
}
