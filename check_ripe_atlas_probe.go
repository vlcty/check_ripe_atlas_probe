package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Status struct {
	Name  string    `json:"name"`
	Since time.Time `json:"since"`
}

type ApiResult struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

func ExitUnknown(message string) {
	fmt.Println(message)
	os.Exit(3)
}

func main() {
	probeID := flag.Uint("probe", 0, "The probe id")
	flag.Parse()

	if *probeID == 0 {
		ExitUnknown("No probe ID given")
	} else if *probeID > 70000 {
		ExitUnknown("Probe IDs over 70000 don't exist!")
	}

	response, err := http.Get(fmt.Sprintf("https://atlas.ripe.net/api/v2/probes/%d", *probeID))

	if err != nil {
		ExitUnknown("Was not able to contact API")
	} else {
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			result := ApiResult{}

			if decodeErr := json.NewDecoder(response.Body).Decode(&result); decodeErr != nil {
				ExitUnknown("Malformed JSON from API")
			} else {
				var prefix string
				var exitCode int

				switch result.Status.Name {
				case "Disconnected", "Abandoned":
					prefix = "CRITICAL"
					exitCode = 2

				default:
					prefix = "OK"
					exitCode = 0
				}

				name := "nameless"

				if len(result.Description) > 0 {
					name = result.Description
				}

				fmt.Printf("%s - Probe %d (\"%s\") is %s since %.1f hours\n",
					prefix, result.ID, name, strings.ToLower(result.Status.Name),
					time.Now().In(time.UTC).Sub(result.Status.Since).Hours())

				os.Exit(exitCode)
			}
		} else {
			ExitUnknown(fmt.Sprintf("UNKNOWN - Status was %s. Expected 200 OK\n", response.Status))
		}
	}
}
