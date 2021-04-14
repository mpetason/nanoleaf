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
	State   State   `json:"state"`
	Token   string
	IP      string
	Port    string
}

type Effects struct {
	EffectsList []string `json:"effectsList"`
	Select      string   `json:"select"`
}

type State struct {
	Brightness Brightness `json:"brightness"`
	On         On         `json:"on"`
}

type Brightness struct {
	Value int `json:"value"`
	Max   int `json:"max"`
	Min   int `json:"min"`
}

type On struct {
	Value bool `json:"value"`
}

func (n NanoLeaf) GetAllInfo() {
	endpoint := "/"
	resp, err := http.Get("http://" + n.IP + ":" + n.Port + "/api/v1/" + n.Token + endpoint)

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
	fmt.Printf("Brightness: %v \n", n.State.Brightness.Value)
	fmt.Printf("On: %v \n", n.State.On.Value)
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
