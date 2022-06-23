// Copyright 2009 The go9p Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"flag"

	"github.com/gbraad/go9p"
)

var addr = flag.String("addr", ":5640", "Bind address")
var path = flag.String("path", "/", "Path to share")
var debug = flag.Int("debug", 0, "Print debug messages")

func main() {
	fmt.Print("go9P UFS server\n")
	flag.Parse()
	fmt.Printf(" Bind address:\t%s\n", *addr)
	fmt.Printf(" Shared path:\t%s\n", *path)
	go9p.StartServer(*addr, *debug, *path)
}
