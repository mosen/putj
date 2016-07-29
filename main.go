package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/mosen/putj/jss"
)

func main() {

	var (
		flSetCapture = flag.NewFlagSet("capture", flag.ExitOnError)
		flSetEnforce = flag.NewFlagSet("enforce", flag.ExitOnError)
		flUrl = flSetCapture.String("url", getEnvDefault("PUTJ_URL", "http://localhost:8080"), "JSS url")
		flUsername = flSetCapture.String("username", getEnvDefault("PUTJ_USERNAME", "admin"), "JSS username")
		flPassword = flSetCapture.String("password", getEnvDefault("PUTJ_PASSWORD", "password"), "JSS password")
	)

	if len(os.Args) == 1 {
		fmt.Println("usage: putj <command> [<args>]")
		fmt.Println("Available subcommands:")
		fmt.Println("\tcapture\t\tCapture the state of the JSS as JSON to stdout.")
		fmt.Println("\tenforce\t\tEnforce the state of a given JSON file.")
		return
	}

	switch os.Args[1] {
	case "capture":
		if err := flSetCapture.Parse(os.Args[2:]); err != nil {
			fmt.Println("Capture")

			api, err := jss.NewApi(*flUrl, *flUsername, *flPassword)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			api.Capture()

		}
	case "enforce":
		if err := flSetEnforce.Parse(os.Args[2:]); err != nil {
			fmt.Println("enforce")
		}
	}

}

func anyEmpty(flags ...*string) bool {
	for _, flValue := range flags {
		if *flValue == "" {
			return true
		}
	}

	return false
}


func getEnvDefault(key string, defaultValue string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defaultValue
	} else {
		return envValue
	}
}
