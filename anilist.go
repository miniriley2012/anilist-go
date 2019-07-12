package anilist

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var url = "https://graphql.anilist.co"

// SetURL changes the default url to newURL in case the AniList url has changed
func SetURL(newURL string) { url = newURL }

// recursively builds fields of a struct
func buildFields(f reflect.StructField) string {
	var str string
	if name, ok := f.Tag.Lookup("name"); ok {
		str += name
	} else {
		str += f.Name
	}

	if t := f.Type; t.Kind() == reflect.Struct {
		str += strings.ReplaceAll(f.Tag.Get("params"), " ", "") + "{"
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			str += buildFields(field)

			if i != t.NumField()-1 && field.Type.Kind() != reflect.Struct {
				str += " "
			}
		}
		str += "}"
	}

	return str
}

// Queries AniList using scheme and writes the response to result
func Query(schema interface{}, result interface{}) error {
	query, queryOK := reflect.TypeOf(schema).FieldByNameFunc(func(fieldName string) bool { return strings.ToLower(fieldName) == "query" })
	variables := reflect.ValueOf(schema).FieldByNameFunc(func(fieldName string) bool { return strings.ToLower(fieldName) == "variables" })

	if t := reflect.TypeOf(result); t.Kind() != reflect.Ptr {
		return errors.New("the second argument must be a pointer")
	}

	if !queryOK {
		return errors.New("missing Query in struct")
	}

	if !variables.IsValid() {
		return errors.New("missing Variables in struct")
	}

	body := buildFields(query)

	b, err := json.Marshal(struct {
		Query     string      `json:"query"`
		Variables interface{} `json:"variables"`
	}{Query: body, Variables: variables.Interface()})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	things, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(things, &result); err != nil {
		return err
	}

	return nil
}
