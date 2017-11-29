package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

type User struct {
	Name    string  `json:"name"`
	Gender  string  `json:"gender"`
	Balance float32 `json:"balance"`
}

const SchemaPath = "schema.json"

var simpleLoader gojsonschema.JSONLoader

func init() {
	simpleLoader = gojsonschema.NewReferenceLoader("file:///" + os.Getenv("GOPATH") + "/" + SchemaPath)
}

func main() {
	// Test 1
	tmpUser := User{
		Name:    "alpha",
		Gender:  "male",
		Balance: 3.14,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	IsValidJSON(string(b.Bytes()), simpleLoader)

	// Test 2
	tmpUser = User{
		Name:    "alpha",
		Gender:  "ET",
		Balance: 3.14,
	}
	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	IsValidJSON(string(b.Bytes()), simpleLoader)

	// Test 3
	tmpUser = User{
		Name:   "alpha",
		Gender: "male",
	}
	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	fmt.Println()
	IsValidJSON(string(b.Bytes()), simpleLoader)

}

func IsValidJSON(s string, sl gojsonschema.JSONLoader) {
	subject := gojsonschema.NewStringLoader(s)
	result, err := gojsonschema.Validate(sl, subject)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
		return
	}

	fmt.Printf("The document is not valid. see errors :\n")
	for _, desc := range result.Errors() {
		fmt.Printf("- %s\n", desc)
	}

}

func getCWD() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
