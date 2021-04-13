package main

import (
	"os"

	"github.com/mpetason/nanoleaf/api"
)

func main() {
	token := os.Getenv("NANOLEAFTOKEN")
	ip := os.Getenv("NANOLEAFIP")
	port := "16021"

	// api.ListEffects(token, ip, port)
	// api.GetAllInfo(token, ip, port)
	api.GetLayout(token, ip, port)
}
