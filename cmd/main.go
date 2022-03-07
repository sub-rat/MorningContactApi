package main

import (
	"github.com/sub-rat/MorningContactApi/internals/server"
)

func main() {
	s := server.GetServer()
	s.Run()
}
