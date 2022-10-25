package gobrew

import (
	"fmt"
	"os/exec"
)

type BrewPackage struct {
	name         string
	descriptiong string
}

// https://docs.brew.sh/Querying-Brew

func Initialize() map[string]BrewPackage {
	// Initialize initializes our current idea of the homebrew database
	//
	// This will pull all currently installed packages and their information
	// and store it in an internal representation
	db := make(map[string]BrewPackage)

	return db
}

func getPackages() {
	// Get all currently installed packages.
	//
	// Can be filtered by manually installed as it gets unwieldy
	// and complicated otherwise

	// `brew leaves` vs `brew list`
}

func getDescription() {
	// Get the description for each package
	//
	// Will be displayed with each displayed package so the user knows what
	// they're looking at.

}

func callLs() {
	// Testing os.exec
	ls := exec.Command("/bin/ls", "/Users/savino/")
	lsOutput, err := ls.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(lsOutput))
}
