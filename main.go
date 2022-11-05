package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var info = color.Blue + "[info] " + color.Reset
var error_msg = color.Red + "[error] " + color.Reset
var target_ip string
var api_key string

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

type Data struct {
	IP             string `json:"ip"`
	Noise          bool   `json:"noise"`
	Riot           bool   `json:"riot"`
	Classification string `json:"classification"`
	Name           string `json:"name"`
	Link           string `json:"link"`
	LastSeen       string `json:"last_seen"`
	Message        string `json:"message"`
}

func banner() {
	var author string = "sc4rfurry"
	var go_version string = "1.19"
	var github string = "https://github.com/sc4rfurry"
	banner := figure.NewColorFigure("GreyGoJack", "", "green", true)
	banner.Print()
	println(color.Ize(color.Red, "\tAuthor: "), author)
	println(color.Ize(color.Red, "\tGo: \t"), go_version)
	println(color.Ize(color.Red, "\tGithub: "), github)
	println(color.Ize(color.Blue, "===================================================================================================\n"))
}

func help() {
	var helper string = `
		-- Help for GreyGoJack --
	
	usage: ./main -i [IP_ADDRESS] --api [API_KEY]
---------------------------------------------------------
	Installation:
		- sudo apt-get update && sudo apt-get golang
		- git clone https://github.com/sc4rfurry/GreyGoJack.git
		- cd GreyGoJack
		- go get
		- go build
`
	println(helper)
	os.Exit(0)
}

func checkConnection() bool {
	var url string = "https://google.com"
	println(info + color.Ize(color.Green, "Checking for Internet Access..."))
	r, e := http.Head(url)
	if r.StatusCode == 200 {
		return e == nil && true
	} else {
		log.Fatal(e)
		return false
	}

}

func checkArgs() (string, string) {
	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			help()
			return "nil", "nil"
		} else {
			println("\n" + error_msg + color.Ize(color.Yellow, "Argument Missing...\n"))
			help()
			return "nil", "nil"
		}
	} else if len(os.Args) == 5 {
		if os.Args[1] == "-i" && os.Args[3] == "--api" {
			target_ip = string(os.Args[2])
			api_key = string(os.Args[4])
			return target_ip, api_key
		}
	} else {
		println("\n" + error_msg + color.Ize(color.Yellow, "Argument Missing...\n"))
		help()
		return "nil", "nil"
	}
	return "nil", "nil"
}

func main() {
	var url string
	var result Data
	banner()
	checkArgs()
	// fmt.Println(api_key)
	if checkConnection() {
		println(info + color.Ize(color.Green, "Valid Connection..."))
	} else {
		println(error_msg + color.Ize(color.Yellow, "No Internet Access..."))
	}

	println(info + color.Ize(color.Green, "Getting Data..."))
	time.Sleep(2 * time.Second)
	url = "https://api.greynoise.io/v3/community/" + target_ip
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	req.Header.Set("key", api_key)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		os.Exit(1)
	}
	println(color.Ize(color.Red, "===================================================================================================\n"))
	println(color.Ize(color.Cyan, "\tIP: \t\t"), result.IP)
	// fmt.Printf("\tIP: %s\n", result.IP)
	println(color.Ize(color.Cyan, "\tNoise: \t\t"), result.Noise)
	// fmt.Printf("\tNoise: %t\n", result.Noise)
	println(color.Ize(color.Cyan, "\tRiot: \t\t"), result.Riot)
	// fmt.Printf("\tRiot: %t\n", result.Riot)
	println(color.Ize(color.Cyan, "\tClassification: "), result.Classification)
	// fmt.Printf("\tClassification: %s\n", result.Classification)
	println(color.Ize(color.Cyan, "\tName: \t\t"), result.Name)
	// fmt.Printf("\tName: %s\n", result.Name)
	println(color.Ize(color.Cyan, "\tViz Link: \t"), result.Link)
	// fmt.Printf("\tLink: %s\n", result.Link)
	println(color.Ize(color.Cyan, "\tLast Seen: \t"), result.LastSeen)
	// fmt.Printf("\tLast Seen: %s\n", result.LastSeen)
	println(color.Ize(color.Cyan, "\tMessage: \t"), result.Message)
	// fmt.Printf("\tMessage: %s\n", result.Message)
}
