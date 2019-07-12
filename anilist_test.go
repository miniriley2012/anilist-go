package anilist

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Test using Query to Query by Title
func TestQuery(t *testing.T) {
	var schema struct {
		Query struct {
			Media struct {
				countryOfOrigin string
				ID              int `name:"id"`
				Title           struct {
					English string `name:"english"`
					Native  string `name:"native"`
					Romaji  string `name:"romaji"`
				} `name:"title"`
				Description string `name:"description"`
			} `params:"(search: $title, type: ANIME)"`
		} `name:"query" params:"($title: String)"`
		Variables struct {
			Title string `json:"title"`
		}
	}

	schema.Variables.Title = "Death Note"

	// If result is of type interface{} then it will be filled with a map[string]interface{} representing the JSON
	var result struct {
		Data struct {
			Media struct {
				CountryOfOrigin string
				Id              int
				Title           struct {
					English string
					Native  string
					Romaji  string
				}
				Description string
			}
		}

		Errors []interface{}
	}

	if err := Query(schema, &result); err != nil {
		t.Fatal(err)
	}

	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(b))
}

// Example of using the API
func ExampleQuery() {
	var schema struct {
		Query struct {
			Media struct {
				countryOfOrigin string
				ID              int `name:"id"`
				Title           struct {
					English string `name:"english"`
					Native  string `name:"native"`
					Romaji  string `name:"romaji"`
				} `name:"title"`
				Description string `name:"description"`
			} `params:"(search: $title, type: ANIME)"`
		} `name:"query" params:"($title: String)"`
		Variables struct {
			Title string `json:"title"`
		}
	}

	schema.Variables.Title = "Shigatsu wa Kimi no Uso"

	// If result is of type interface{} then it will be filled with a map[string]interface{} representing the JSON
	var result struct {
		Data struct {
			Media struct {
				CountryOfOrigin string
				Id              int
				Title           struct {
					English string
					Native  string
					Romaji  string
				}
				Description string
			}
		}

		Errors []interface{}
	}

	if err := Query(schema, &result); err != nil {
		// Handle error
	}
}
