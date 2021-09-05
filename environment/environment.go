// Package env provides some boiler plate funtionality for easily
// creating and maintaining environment information for Go programs.
package environment

import (
    "log"
    "os"
    "path/filepath"
)

// Environment stores path information in a standardized and centralized spot.
// This should not be created directly and instead be initialized
// with NewEnvironment().
type Environment struct {
    ProgramPath string
    BasePath string
    dirs []dir
}

type dir struct {
    Name string
    Parent string
}

// New creates and returns a new instance of type Environment.
// If base_path is left as empty string then the BasePath defaults to the
// location where the programs executable lives. If base_path is not an empty
// string then the path is checked to ensure that it is a directory that can be
// written to.
func New(base_path string) (e Environment) {
    // Determine the path to the folder the executable is in.
    bin_path, err := os.Executable()
    if err != nil {
        log.Fatal(err)
    }
    e.ProgramPath = filepath.Dir(bin_path)

    if base_path == "" {
        // BasePath defaults to the folder the program lives in.
        e.BasePath = e.ProgramPath
    } else {
        // Determine if the base_path provided is valid.
        var base_path_stat os.FileInfo
        base_path_stat, err = os.Stat(base_path)
        if err != nil {
            log.Fatal(err)
        }
        if base_path_stat.IsDir() == false {
            log.Fatal("base_path provided to NewEnvironment is not a directory.")
        }
        if base_path_stat.Mode().Perm() < 0755 {
            log.Fatal("base_path provided to NewEnvironment has incompatible permissions.")
        }
        e.BasePath = base_path
    }
    return
}

// AddDir is how a directory is included in the Environment. dir_name is
// how a directory is referenced within the Environment object and also will be
// the name that is given to the directory when it is created, dir_name cannot
// be the same as an already added dir and can also not be an empty string.
// parent_dir is the name of another (already added) directory that this new
// directory will live under. AddDir then returns the full path of dir_name.
func (s *Environment) AddDir(dir_name string, parent_dir string) (dir_path string) {
    // Don't allow empty string as a dir name. This breaks stuff.
    if dir_name == "" {
        log.Fatal("Directory name cannot be empty string.")
    }

    // Don't register if a directory of the same name already exists.
    if s.dirIsRegistered(dir_name) == true {
        log.Fatal("Directory "+dir_name+" already registered.")
    }

    // Verify parent_dir is a valid registered directory.
    if s.isValidParentDir(parent_dir) == false {
        log.Fatal("Directory "+dir_name+" using unregistered parent "+parent_dir+".")
    }

    var d dir
    d.Name = dir_name
    d.Parent = parent_dir

    s.dirs = append(s.dirs, d)

    // Return the name of the newly registered dir.
    dir_path = s.FullPath(d.Name)
    return
}

// FullPath takes a directory name previously registered using AddDir and
// then returns that directory's absolute path.
func (s *Environment) FullPath(dir_name string) string {
    return s.BasePath+s.dirPath(dir_name)
}

// CreateDirs creates all dirs previosly registered with AddDir.
// These directories are created in BasePath.
func (s *Environment) CreateDirs() {
    for _, d := range s.dirs {
        os.Mkdir(s.BasePath+s.dirPath(d.Name), 0755)
    }
}

////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////

// dirIsRegistered is a helper function to simplify checking if a Directory
// is registered based on the directory's name.
func (s *Environment) dirIsRegistered(name string) bool {
    for _, c := range s.dirs {
        if c.Name == name {
            return true
        }
    }
    return false
}
// isValidParentDir is a helper function to simplify checking if a Directory
// is already registered and is a valid parent to use.
func (s *Environment) isValidParentDir(parent_dir string) bool {
    if parent_dir == "" {
        return true
    }
    return s.dirIsRegistered(parent_dir)
}

// dirPath is a helper function designed to build out a direco's complete
// path to the Environemnt BasePath.
func (s *Environment) dirPath(dir_name string) (dir_path string) {
    for _, c := range s.dirs {
        if c.Name == dir_name {
            return s.dirPath(c.Parent)+string(os.PathSeparator)+dir_name
        }
    }
    return dir_name
}
