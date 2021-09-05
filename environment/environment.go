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

// NewEnvironment creates and returns a new instance of type Environment.
// If base_path is left as empty string then the BasePath defaults to the
// location where the programs executable lives. If base_path is not an empty
// string then the path is checked to ensure that it is a directory that can be
// written to.
func NewEnvironment(base_path string) (e Environment) {
    var bin_path string
    var program_path string
    var err error
    // Determine the path to the folder the executable is in.
    bin_path, err = os.Executable()
    if err != nil {
        log.Fatal(err)
    }
    program_path = filepath.Dir(bin_path)

    if base_path != "" {
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
    } else {
        // BasePath defaults to the folder the program lives in.
        e.BasePath = program_path
    }
    e.ProgramPath = program_path

    return
}

// RegisterDir is how a directory is included in the Environment. dir_name is
// how a directory is referenced within the Environment object and also will be
// the name that is given to the directory when it is created, dir_name cannot
// be the same as an already registered dir. parent_dir is the name of another
// (already registered) directory that this new directory will live under.
// The same name provided to RegisterDir is also the return value to simplify
// storing the directory name, since you will be using the name again in the
// future to retrieve the directory's full path (for example).
func (s *Environment) RegisterDir(dir_name string, parent_dir string) (reg_name string) {
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
    reg_name = d.Name
    return
}

// FullPath takes a directory name previously registered using RegisterDir and
// then returns that directory's absolute path.
func (s *Environment) FullPath(dir_name string) (pull_path string) {
    return s.BasePath+s.dirPath(dir_name)
}

// CreateDirs creates all dirs previosly registered with RegisterDir.
// These directories are created in BasePath.
func (s *Environment) CreateDirs() {
    for _, d := range s.dirs {
        os.Mkdir(s.BasePath+s.dirPath(d.Name), 0755)
    }
}

// RegisterDefaultDirs is a quick little shortcut to register a few dirs
// that are created for a lot of different programs. The main parent directory
// is provided by master_dir, all the default directories are all created under
// that. The default directory names are then returned.
func (s *Environment) RegisterDefaultDirs(master_dir string) (archive, data, settings string) {
    s.RegisterDir(master_dir, "")
    archive = s.RegisterDir("archive", master_dir)
    data = s.RegisterDir("data", master_dir)
    settings = s.RegisterDir("settings", master_dir)
    return
}

// NewDefaultEnvironment provides a way to create a generic environment in just
// a single function call. It defaults to creating a main parent directory in
// the user's home directory, then creates the 3 dirs created by
// RegisterDefaultDirs under that main parent dir and finally returns their
// names after the new Environemnt object is returned.
func NewDefaultEnvironment(master_dir string) (e Environment, archive, data, settings string) {
    base_path, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    e = NewEnvironment(base_path)
    archive, data, settings = e.RegisterDefaultDirs(master_dir)
    e.CreateDirs()
    return
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
