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
	schemaLoader := gojsonschema.NewReferenceLoader("file:///home/ubuntu/workspace/schema.json")
	schema, _ := gojsonschema.NewSchema(schemaLoader)

	tmpUser := User{
		Name:    "alpha",
		Gender:  "male",
		Balance: 3.14,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)

	IsValidJSON(string(b.Bytes()), schema)

	tmpUser = User{
		Name:    "alpha",
		Gender:  "ET",
		Balance: 3.14,
	}
	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	IsValidJSON(string(b.Bytes()), schema)

	tmpUser = User{
		Name:   "alpha",
		Gender: "male",
	}
	b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(tmpUser)
	fmt.Println(b)
	IsValidJSON(string(b.Bytes()), schema)

}

func IsValidJSON(s string, schema *gojsonschema.Schema) {
	subject := gojsonschema.NewStringLoader(s)
	result, err := schema.Validate(subject)
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
