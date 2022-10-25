package main

import (
	"fmt"

	"github.com/sjossi/gobrewcleanup/cli"
	"github.com/sjossi/gobrewcleanup/gobrew"
)

func questionnaire() {
	// Goes through the list of installed packages and prompts for action
	//
	// Allows to (r)emove package.
}

func main() {
	// GoBrewCleanup will list all locally installed packages and allows for
	// interactive review of those packages and allows to have them uninstalled

	fmt.Println("[+] Starting gobrewcleanup")

	// TODO: can we build a database asynchronously, so we can start the
	// questionnaire as soon as the first results come in? Good practice for
	// goroutines maybe.
	db := gobrew.Initialize()
	fmt.Println(db)

	// depending on performance, maybe cache results? can be sketchy as soon as
	// information gets out of date though
	cli.Questionnaire()
}
