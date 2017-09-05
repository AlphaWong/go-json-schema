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

func main() {
	tmpUser := User{
		Name:    "alpha",
		Gender:  "male",
		Balance: 3.14,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	IsValidJSON(string(b.Bytes()), "file:///home/ubuntu/workspace/schema.json")

	tmpUser = User{
		Name:    "alpha",
		Gender:  "ET",
		Balance: 3.14,
	}
	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	IsValidJSON(string(b.Bytes()), "file:///home/ubuntu/workspace/schema.json")
}

func IsValidJSON(s string, path string) {
	schemaLoader := gojsonschema.NewReferenceLoader(path)
	subject := gojsonschema.NewStringLoader(s)
	result, err := gojsonschema.Validate(schemaLoader, subject)
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
