package main

import (
	"os"

	"github.com/mpetason/nanoleaf/api"
)

func main() {
	token := os.Getenv("NANOLEAFTOKEN")
	ip := os.Getenv("NANOLEAFIP")
	port := "16021"

	var n api.NanoLeaf
	n.GetAllInfo(token, ip, port)
}
