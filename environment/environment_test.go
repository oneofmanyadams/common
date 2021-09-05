package environment

import (
    "fmt"
    "log"
    "os"
)

func ExampleNew() {
    // Creates directories in the user's home directory.
    // This example the user's home directory is "C:\Users\Me"
    home_path, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    env := New(home_path)
    env_base_path := env.RegisterDir("base", "")
    env_data_path := env.RegisterDir("data", "base")
    env_stng_path := env.RegisterDir("settings", "base")
    env.CreateDirs()
    fmt.Println(env_base_path)
    fmt.Println(env_data_path)
    fmt.Println(env_stng_path)
    // Output: C:\Users\Me\base
    // C:\Users\Me\base\data
    // C:\Users\Me\base\settings
}
