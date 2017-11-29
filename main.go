package main

import "github.com/kcmerrill/po/po"

func main() {
	po.NewPo().HTTPServer("8080", "")
}
