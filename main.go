// Copyright © 2016-2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// This is the XETH control daemon for Platina's Mk1 TOR switch.
// Build it with,
//	go build
//	zip drivers vnet-platina-mk1
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/platinasystems/buildinfo"
	"github.com/platinasystems/fe1"
	fe1a "github.com/platinasystems/firmware-fe1a"
	"github.com/platinasystems/redis"
	vnetfe1 "github.com/platinasystems/vnet/devices/ethernet/switch/fe1"
	yaml "gopkg.in/yaml.v2"
)

const usage = `
usage:	vnet-platina-mk1
	vnet-platina-mk1 install
	vnet-platina-mk1 [show] {version, buildinfo, license, patents}`

var ErrUsage = errors.New(usage[1:])

func main() {
	var err error
	f := mk1Main
	stub := func() error { return nil }
	show := false

	redis.DefaultHash = "platina-mk1"
	vnetfe1.AddPlatform = fe1.AddPlatform
	vnetfe1.Init = fe1.Init

	for _, arg := range os.Args[1:] {
		switch strings.TrimLeft(arg, "-") {
		case "install":
			f = stub
			if show {
				err = ErrUsage
			} else {
				err = install()
			}
		case "show":
			show = true
		case "version":
			f = stub
			fmt.Println(buildinfo.New().Version())
		case "buildinfo":
			f = stub
			fmt.Println(buildinfo.New())
		case "copyright", "license":
			f = stub
			err = marshalOut(licenses())
		case "patents":
			f = stub
			err = marshalOut(patents())
		case "h", "help", "usage":
			fmt.Println(usage[1:])
			return
		default:
			err = fmt.Errorf("%q unknown", arg)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	if err = f(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func marshalOut(m map[string]string) error {
	b, err := yaml.Marshal(m)
	if err == nil {
		os.Stdout.Write(b)
	}
	return err
}

func licenses() map[string]string {
	return map[string]string{
		"fe1":  fe1.License,
		"fe1a": fe1a.License,

		"vnet-platina-mk1": License,
	}
}

func patents() map[string]string {
	return map[string]string{
		"fe1": fe1.Patents,
	}
}
