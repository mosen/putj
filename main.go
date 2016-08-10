package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/mosen/putj/jss"
	_ "github.com/mosen/putj/jss/handlers/accounts"
	_ "github.com/mosen/putj/jss/handlers/activation"
	_ "github.com/mosen/putj/jss/handlers/buildings"
	_ "github.com/mosen/putj/jss/handlers/categories"
	_ "github.com/mosen/putj/jss/handlers/computercheckin"
	_ "github.com/mosen/putj/jss/handlers/departments"
	_ "github.com/mosen/putj/jss/handlers/distributionpoints"
	_ "github.com/mosen/putj/jss/handlers/netbootservers"
	_ "github.com/mosen/putj/jss/handlers/smtpserver"
	"encoding/json"
	"io/ioutil"
)

func main() {

	var (
		flSetCapture = flag.NewFlagSet("capture", flag.ExitOnError)
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
		fmt.Println("Capture")

		api, err := jss.NewApi(*flUrl, *flUsername, *flPassword)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var state = make(map[string]interface{})
		if err := api.Capture(state); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		stateJson, err := json.Marshal(state)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", stateJson)


	case "enforce":
		fmt.Println("Enforce")

		api, err := jss.NewApi(*flUrl, *flUsername, *flPassword)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		f, err := ioutil.ReadFile("state.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		state := make(map[string]interface{})
		if err := json.Unmarshal(f, &state); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := api.Enforce(state); err != nil {
			fmt.Println(err)
			os.Exit(1)
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
