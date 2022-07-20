package main

import (
	"fmt"
	"os"

	"github.com/PucklaMotzer09/showip/lib"
	flag "github.com/jessevdk/go-flags"
)

var flags struct {
	Local bool `short:"l" long:"local" description:"Wether to print the local IP Address instead of the public one"`
}

func main() {
	parser := flag.NewParser(&flags, flag.HelpFlag)
	parser.Usage += "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		if !flag.WroteHelp(err) {
			fmt.Fprintln(os.Stderr, "Arguments: ", err)
			parser.WriteHelp(os.Stdout)
		} else {
			fmt.Print(err)
		}
		os.Exit(1)
	}

	if flags.Local {
		ips, err := lib.GetLocalIPv4Address()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(ips[0].IP.String())
	} else {
		ip, err := lib.GetPublicIPv4Address()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(ip.String())
	}
}
