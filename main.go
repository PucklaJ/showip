/***********************************************************************************
 *                         This file is part of showip
 *                    https://github.com/PucklaMotzer09/showip
 ***********************************************************************************
 * Copyright (c) 2022 PucklaMotzer09
 *
 * This software is provided 'as-is', without any express or implied warranty.
 * In no event will the authors be held liable for any damages arising from the
 * use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute it
 * freely, subject to the following restrictions:
 *
 * 1. The origin of this software must not be misrepresented; you must not claim
 * that you wrote the original software. If you use this software in a product,
 * an acknowledgment in the product documentation would be appreciated but is
 * not required.
 *
 * 2. Altered source versions must be plainly marked as such, and must not be
 * misrepresented as being the original software.
 *
 * 3. This notice may not be removed or altered from any source distribution.
 ************************************************************************************/

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
