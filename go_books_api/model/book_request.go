package model

import "errors"

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (br *BookRequest) Validate() error {
	if br.Title == "" {
		return errors.New("title is required")
	}
	if br.Author == "" {
		return errors.New("author is required")
	}
	if br.Year == 0 {
		return errors.New("year is required")
	}

	return nil
}
