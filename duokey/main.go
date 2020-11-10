package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"gitlab.com/sepior/go-tsm-sdk/sdk/tsm"
)

// Configure with your TSM cluster here
const credentials string = `
[
  {
    "username": "user1",
    "password": "X3MjwZDK0l",
    "url": "https://duokey1.tsm.sepior.net"
  },
  {
    "username": "user1",
    "password": "MZ9sVhXPRe",
    "url": "https://duokey2.tsm.sepior.net"
  },
  {
    "username": "user1",
    "password": "iZY9JQrwZz",
    "url": "https://duokey3.tsm.sepior.net"
  }
]

`

func main() {
	// Create client from credentials
	tsmClient := parseCredentials(credentials)
	rsaClient := tsm.NewRSAClient(tsmClient)

	keyID, err := rsaClient.Keygen(2048)

	if err != nil {
		fmt.Println("error generating key: ", err)
		os.Exit(1)
	}
	fmt.Printf("Generated key: ID=%s\n", keyID)
}

func parseCredentials(credentials string) tsm.Client {
	var parsed []map[string]string
	err := json.Unmarshal([]byte(credentials), &parsed)
	if err != nil {
		panic(err)
	}

	var nodes []tsm.Node
	for _, c := range parsed {
		parsedURL, err := url.Parse(c["url"])
		if err != nil {
			fmt.Println("Could not parse credentials: ", err)
			os.Exit(1)
		}
		nodes = append(nodes, tsm.NewURLNode(*parsedURL, &tsm.PasswordAuthenticator{
			Username: c["username"],
			Password: c["password"],
		}))
	}

	return tsm.NewClient(nodes)
}
