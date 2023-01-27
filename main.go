package main

import (
	"encoding/json"
	"fmt"
	json2 "github.com/goccy/go-json"
	"io/ioutil"
	"os"
)

func UnmarshalGoEncoding() {
	jsonFile, err := os.Open("1000.json")
	if err != nil {
		fmt.Println("no such file")
		return

	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}

	err = json.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Println("unmarshal error")
		return
	}
}

func UnmarshalGoJson() {
	jsonFile, err := os.Open("1000.json")
	if err != nil {
		fmt.Println("no such file")
		return

	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
	err = json2.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Println("unmarshal error")
		return
	}

}

func UnmarshalAndMarshalGoEncoding() {
	jsonFile, err := os.Open("1000.json")
	if err != nil {
		fmt.Println("no such file")
		return

	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
	err = json.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Println("unmarshal error")
		return
	}
	_, err = json.Marshal(res)
	if err != nil {
		fmt.Println("marshal error")
		return
	}

}

func UnmarshalAndMarshalGoJson() {
	jsonFile, err := os.Open("1000.json")
	if err != nil {
		fmt.Println("no such file")
		return

	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
	err = json2.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Println("unmarshal error")
		return
	}
	_, err = json2.Marshal(res)
	if err != nil {
		fmt.Println("marshal error")
		return
	}

}

func main() {
	UnmarshalGoEncoding()
	UnmarshalGoJson()
	UnmarshalAndMarshalGoEncoding()
	UnmarshalAndMarshalGoJson()
}
