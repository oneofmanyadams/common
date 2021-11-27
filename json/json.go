// Package json is some super simple boiler plate functionality for 
// reading and writting structs as JSON data.
// This is kind of lame, but I find myself constantly writting this stuff
// over and over again, might as well just do it once.
package json

import (
	"encoding/json"
	"log"
	"os"
)

// Load attempts to load a json file located at file_path into s.
// s MUST be a pointer!
func Load(s interface{}, file_path string) {
    d, err := os.ReadFile(file_path)
    if err != nil {
        log.Fatal(err)
        return
    }
    err = json.Unmarshal(d, s)
    if err != nil {
        log.Fatal(err)
        return
    }
}

// Save takes s and saves it as json encoded data to file_path.
func Save(s interface{}, file_path string) {
    d, err := json.MarshalIndent(s, "", "   ")
    if err != nil {
        log.Fatal(err)
        return
    }
    err = os.WriteFile(file_path, d, 0666)
    if err != nil {
        log.Fatal(err)
        return
    }
}