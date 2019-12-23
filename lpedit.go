package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ansd/lastpass-go"
)

func main() {
	email := flag.String("email", "", "LastPass email/account")
	password := flag.String("password", "", "LastPass master password")
	output := flag.Bool("out", false, "output LastPass credentials as JSON via STDOUT")
	del := flag.Bool("del", false, "delete LastPass credentials piped through STDIN as JSON")
	in := flag.Bool("in", false, "add or overwrite LastPass records")

	flag.Parse()

	if *email == "" {
		log.Fatalln("the LastPass email/account must be specified")
	}

	if *password == "" {
		log.Fatalln("the LastPass master password must be specified")
	}

	// authenticate with LastPass servers
	client, err := lastpass.NewClient(context.Background(), *email, *password)
	if err != nil {
		log.Fatalln(err)
	}

	if *output {
		accounts, err := client.Accounts(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.SetEscapeHTML(false)
		enc.Encode(accounts)

	} else if *del {
		accounts := slurpAccounts()
		for _, account := range accounts {
			err := client.Delete(context.Background(), account.ID)
			if err != nil {
				log.Fatalf("could not delete account '%s': %v\n", account.Name, err)
			} else {
				fmt.Printf("account '%s': deleted!\n", account.Name)
			}
		}
	} else if *in {
		accounts := slurpAccounts()
		for _, account := range accounts {
			err := client.Add(context.Background(), &account)
			if err != nil {
				log.Fatalf("could not add account '%s': %v\n", account.Name, err)
			}
		}
	} else {
		fmt.Printf("Unrecognized usage!\nUsage:\n\n")
		flag.PrintDefaults()
	}
}

func slurpAccounts() []lastpass.Account {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	var accounts []lastpass.Account
	err = json.Unmarshal(b, &accounts)
	if err != nil {
		log.Fatalln("could not decode JSON")
	}

	return accounts
}
