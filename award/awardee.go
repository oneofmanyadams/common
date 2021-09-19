package award

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
)

type Awardee interface {
    GetId() string
    GetAwards() []Award
}

////////////////////////////////////////////
////////////////////////////////////////////
////////////////////////////////////////////
func LoadAwardee(s Awardee, awardee_path string) error {
    if s.GetId() == "" {
        return errors.New("Awardee Id cannot be empty string.")
    }
    af := awardee_path+string(os.PathSeparator)+s.GetId()+".json"
    LoadAwardeeAwards(s, awardee_path)
    LoadAwardeeFile(s, af)
    return nil
}
func SaveAwardee(s Awardee, awardee_path string) error {
    if s.GetId() == "" {
        return errors.New("Awardee Id cannot be empty string.")
    }
    af := awardee_path+string(os.PathSeparator)+s.GetId()+".json"

    // This is currently not at the right level, it needs to go one lower
    SaveAwardeeAwards(s, awardee_path)


    SaveAwardeeFile(s, af)
    return nil
}

////////////////////////////////////////////
////////////////////////////////////////////
////////////////////////////////////////////
func LoadAwardeeAwards(s Awardee, awardee_path string) {
    // open awardee_path dir
    files, err := ioutil.ReadDir(awardee_path)
    if err != nil {
    	log.Fatal(err)
    }

    // read all .json files.
    for _, f := range files {
        // load json data into an award struct, append to s Awardee
        fmt.Println(f.Name())
    }
}

func SaveAwardeeAwards(s Awardee, awardee_path string) {
    // for _, award := range s.GetAwards() {
    //     award.SaveTo(awardee_path)
    // }
}

////////////////////////////////////////////
////////////////////////////////////////////
////////////////////////////////////////////
func LoadAwardeeFile(s Awardee, file_path string) {
    // open file_path
    file, err := os.Open(file_path)
    if err != nil {
    	log.Fatal(err)
    }
    defer file.Close()
    // read file_path json data
    json_data, err := io.ReadAll(file)
    if err != nil {
    	log.Fatal(err)
    }
    // take read json data and load it into the soon-to-be returned Awardee.
    err = json.Unmarshal(json_data, s)
    if err != nil {
        log.Fatal(err)
    }
}

func SaveAwardeeFile(s Awardee, file_path string) {
    // Marshal Configuration
    d, err := json.MarshalIndent(s, "", "   ")
    if err != nil {
        log.Fatal(err)
        return
    }
    // write Marshaled Awardee data
    err = os.WriteFile(file_path, d, 0666)
    if err != nil {
        log.Fatal(err)
        return
    }
}
