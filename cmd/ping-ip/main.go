package main

import pingHelper "github.com/RajabovIlyas/ping-ip/internal/app/ping/helpers"

func main() {
	ping := pingHelper.NewPing()
	err := ping.Run()
	if err != nil {
		panic(err)
	}
}
