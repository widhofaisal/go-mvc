package main

import "golang-echo-gorm/route"

func main() {
	e := route.New()

	e.Start(":8000")
}