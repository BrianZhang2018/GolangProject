package api

import "encoding/json"

type MyError struct {
	ID            string `json:"id,omitempty"`
	DeveloperText string `json:"developerText,omitempty"`
	StatusCode    int    `json:"statusCode,omitempty"`
}

func (ae *MyError) Error() string {
	jsonE, _ := json.Marshal(ae)
	return string(jsonE)
}

func (ae *MyError) WithDeveloperText(text string) error {
	ae.DeveloperText = text
	return ae
}
