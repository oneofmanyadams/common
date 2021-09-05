// Package configuration provides a simple and easy to implement way for creating,
// loading, and updating user defined settings in a JSON file.
package configuration

import (
    "encoding/json"
    "io"
    "os"
)

type Configuration struct {
    Confs []Conf
}

type Conf struct {
    Name string
    Value string
}


func New(config_path string) (c Configuration) {
    // Check if configuration already exists in provided path.
    //      If it doesn't create it based on embeded defaults.

    // Load configuration from JSON data.
    c.Load(config_path)
    return
}

// Load takes a file_path to a JSON file that correlates to a Configuration.
// It then loads that data into the calling Configuration object.
func (s *Configuration) Load(file_path string) {
    // open file_path
    file, err := os.Open(config_path)
    if err != nil {
    	log.Fatal(err)
    }
    defer file.Close()
    // read file_path json data
    json_data, err := io.ReadAll(file)
    if err != nil {
    	log.Fatal(err)
    }
    // take read json data and load it into the soon-to-be returned Configuration
    err := json.Unmarshal(json_data, s)
    if err != nil {
        log.Fatal(err)
    }
}

func (s *Configuration) Save(file_path string) {

}
