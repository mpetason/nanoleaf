package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type NanoLeaf struct {
	Name    string  `json:"name"`
	Effects Effects `json:"effects"`
}

type Effects struct {
	EffectsList []string `json:"effectsList"`
	Select      string   `json:"select"`
}

func (n NanoLeaf) GetAllInfo(token string, ip string, port string) {
	endpoint := "/"
	resp, err := http.Get("http://" + ip + ":" + port + "/api/v1/" + token + endpoint)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &n)

	for _, j := range n.Effects.EffectsList {
		fmt.Println(j)
	}

	fmt.Printf("Selected: %v \n", n.Effects.Select)
}

func GetStatus(token string, ip string, port string) {
	endpoint := "/state/on"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetBrightness(token string, ip string, port string) {
	endpoint := "/state/brightness"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetHue(token string, ip string, port string) {
	endpoint := "/state/hue"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetSaturation(token string, ip string, port string) {
	endpoint := "/state/sat"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetColorTemp(token string, ip string, port string) {
	endpoint := "/state/ct"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetColorMode(token string, ip string, port string) {
	endpoint := "/state/colorMode"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func ListEffects(token string, ip string, port string) {
	endpoint := "/effects/effectsList"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetGlobalOrientation(token string, ip string, port string) {
	endpoint := "/panelLayout/globalOrientation"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func GetLayout(token string, ip string, port string) {
	endpoint := "/panelLayout/layout"
	fmt.Println(apiBuilder(token, ip, port, endpoint))
}

func apiBuilder(token string, ip string, port string, endpoint string) string {
	resp, err := http.Get("http://" + ip + ":" + port + "/api/v1/" + token + endpoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
