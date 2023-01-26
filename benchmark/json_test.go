package benchmark

import (
	"fmt"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type Classmates struct {
	Id        int64  `json:"id"`
	FirstNme  string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
}

func TestMarshalAndUnmarshalJson(t *testing.T) {
	jsonFile, err := os.Open("MOCK_DATA.json")
	if err != nil {
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
	json.Unmarshal(byteValue, &res)

}
