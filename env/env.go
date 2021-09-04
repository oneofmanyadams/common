// Package env provides some boiler plate funtionality for easily
// creating and maintaining environment information for Go programs.
package env

import (
    _ "embed"
    "encoding/json"
    "log"
    "os"
    "path/filepath"
)

//go:embed env.json
var json_env []byte

// Environment stores path information in a standardized
// and centralized spot. This should not be created directly
// and instead be initialized with NewEnvironment().
type Environment struct {
    Archive string
    Data string
    Settings string
    ProgramPath string
    BasePath string
}

// NewEnvironment creates a new instance of type Environment
// and populates Archive, Data, and Settings with names
// that are predetermined by this package's env.json file.
// It also populates ProgramPath with the full path to
// the fodler that the executable lives in and duplicates
// that path to BasePath.
func NewEnvironment() (e Environment, err error) {
    // Load default folder locations from env.json file.
    err = json.Unmarshal(json_env, &e)
    if err != nil {
        log.Fatal(err)
        return e, err
    }
    // Populate local system information.
    var bin_path string

    // Path to executable.
    bin_path, err = os.Executable()
    if err != nil {
        log.Fatal(err)
        return e, err
    }

    // Path to folder the executable is in, also
    // defaults BasePath to the same location.
    e.ProgramPath = filepath.Dir(bin_path)
    e.BasePath = e.ProgramPath
    return e, err
}

// SetBasePath allows the implementing program to override where
// the environment stores, creates, and uses it's folders (like
// Data, Settings, etc..). The default location for these
// is the same place that the executable lives, so overriding
// that can be usefull in keeping a "bin" folder clear (for
// example).
func (e *Environment) SetBasePath(full_path string) {
    e.BasePath = full_path
}

// CreateDirs creates the Archive, Data, and Settings folders
// in the location defined by BasePath. If BasePath does not
// exist it is created as well.
func (e *Environment) CreateDirs() {
    slash := string(os.PathSeparator) // lol, idk, looks slightly nicer
    var err error
    // Create BasePath directory
    err = os.MkdirAll(e.BasePath, 0777)
    if err != nil {
        log.Fatal(err)
        return
    }
    // Create archive directory
    os.Mkdir(e.BasePath+slash+e.Archive, 0777)
    // Create data directory
    os.Mkdir(e.BasePath+slash+e.Data, 0777)
    // Create settings directory
    os.Mkdir(e.BasePath+slash+e.Settings, 0777)
}

// RegisterDir allows the implementing program to add additional
// "custom" subdirectories to one of the default directories (
// settings/user for example). parent_dir needs to be one
// of the default directories. name is how this custom direcotry
// will be referenced. path is a relative path from the parent_dir.
// !! TBD if this will actually be needed or not. KISS maybe
// applies here?
func (e *Environment) RegisterDir(parent_dir string, name string, path string) error {
}
