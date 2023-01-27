package benchmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	json2 "github.com/goccy/go-json"
	"io/ioutil"
)

func UnmarshalEncoding(fileName string) {

	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	var res map[string]interface{}

	err = json.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Println("unmarshal error")
		return
	}

}

func NewDecoderEncoding(fileName string) {
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	var res map[string]interface{}
	err = json.NewDecoder(bytes.NewReader(byteValue)).Decode(&res)
	if err != nil {
		fmt.Println("decode error")
		return
	}
}
func NewDecoderGoJson(fileName string) {
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	var res map[string]interface{}
	err = json2.NewDecoder(bytes.NewReader(byteValue)).Decode(&res)
	if err != nil {
		fmt.Println("decode error")
		return
	}
}
func UnmarshalGoJson(fileName string) {

	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	var res map[string]interface{}
	err = json2.Unmarshal(byteValue, &res)
	if err != nil {
		fmt.Println("unmarshal error")
		return
	}

}

func UnmarshalAndMarshalEncoding(fileName string) {

	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
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

func UnmarshalAndMarshalGoJson(fileName string) {

	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}

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
