package main

import "fmt"

type Developer struct {
	Name              string
	MainLanguage      string
	YearsOfExperience int
}

func (d *Developer) update(blogText string) {
	fmt.Printf("%s has been notified of a new blog post: %s\n", d.Name, blogText)
}
func (d *Developer) getID() string {
	return d.Name
}
