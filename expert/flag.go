package main

import (
	"flag"
	"fmt"
)

func main() {
	hostname := flag.String("host", "localhost", "Put your host")
	user := flag.String("user", "root", "Put your database user")
	pw := flag.String("pass", "password", "Put your database password")

	flag.Parse()

	fmt.Println(*hostname, *user, *pw)
}
