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
	schemaLoader := gojsonschema.NewReferenceLoader("file:///home/ubuntu/workspace/src/schema.json")
	subject := gojsonschema.NewStringLoader(string(b.Bytes()))
	result, err := gojsonschema.Validate(schemaLoader, subject)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
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
