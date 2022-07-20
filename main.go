package main

import (
	"fmt"
	"os"

	"github.com/PucklaMotzer09/showip/lib"
	flag "github.com/jessevdk/go-flags"
)

var flags struct {
	Local      bool `short:"l" long:"local" description:"Wether to print the local IP Address instead of the public one"`
	All        bool `short:"a" long:"all" description:"Print all found local addresses"`
	Interfaces bool `short:"i" long:"interfaces" description:"Print the local addresses with the corresponding interfaces"`
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

		if flags.All {
			for _, i := range ips {
				if flags.Interfaces {
					fmt.Println(fmt.Sprint(i.Interface.Name, ": ", i.IP.String()))
				} else {
					fmt.Println(i.IP.String())
				}
			}
		} else {
			if flags.Interfaces {
				fmt.Println(fmt.Sprint(ips[0].Interface.Name, ": ", ips[0].IP.String()))
			} else {
				fmt.Println(ips[0].IP.String())
			}
		}
	} else {
		ip, err := lib.GetPublicIPv4Address()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(ip.String())
	}
}
