package environment

import (
    "fmt"
)

func ExampleNew() {
    env := New("C:\\Users\\Me")
    env_maf := env.RegisterDir("mainAppFolder", "")
    env_data := env.RegisterDir("data", "mainAppFolder")
    env_stng := env.RegisterDir("settings", "mainAppFolder")
    env.CreateDirs()
    fmt.Println(env.FullPath(env_maf))
    fmt.Println(env.FullPath(env_data))
    fmt.Println(env.FullPath(env_stng))
    // Output: C:\Users\Me\mainAppFolder
    // C:\Users\Me\mainAppFolder\data
    // C:\Users\Me\mainAppFolder\settings
}

func ExampleNewDefault() {
    // assuming the user's HomeDir is C:\Users\Me
    env, env_archive, env_data, env_settings := NewDefault("mainAppFolder")
    fmt.Println(env.FullPath(env_archive))
    fmt.Println(env.FullPath(env_data))
    fmt.Println(env.FullPath(env_settings))
    // Output: C:\Users\Me\mainAppFolder\archive
    // C:\Users\Me\mainAppFolder\data
    // C:\Users\Me\mainAppFolder\settings
}
