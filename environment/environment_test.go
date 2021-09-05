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
    env_maf := env.RegisterDir("base", "")
    env_data := env.RegisterDir("data", "base")
    env_stng := env.RegisterDir("settings", "base")
    env.CreateDirs()
    fmt.Println(env.FullPath(env_maf))
    fmt.Println(env.FullPath(env_data))
    fmt.Println(env.FullPath(env_stng))
    // Output: C:\Users\Me\base
    // C:\Users\Me\base\data
    // C:\Users\Me\base\settings
}
