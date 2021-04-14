package main

import (
	"os"

	"github.com/mpetason/nanoleaf/api"
)

func main() {
	var n api.NanoLeaf
	n.Token = os.Getenv("NANOLEAFTOKEN")
	n.IP = os.Getenv("NANOLEAFIP")
	n.Port = "16021"
	n.GetAllInfo()
}
