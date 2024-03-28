package main

import (
	"math/rand"
	apiserver "postcard/apiserver/internal"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
	apiserver.NewAPIServer().Run()

}
