package main

import (
	"C"
	"api4Deeplx/server"
)

func main() {
	run()
}

//export run
func run() {
	go server.Run()
}
