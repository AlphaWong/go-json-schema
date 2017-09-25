# Objective
Testing the json schema in golang implementation.

# Run
```go
go get
go run ./main.go
```

# Result
```sh
The document is valid
The document is not valid. see errors :
- gender: Does not match pattern '^(fe)?(male)$'
```

# Tools
1. https://jsonschema.net/#/editor
1. http://json-schema.org/example1.html
1. https://github.com/xeipuuv/gojsonschema
1. https://github.com/b3log/wide

# Reference
1. https://spacetelescope.github.io/understanding-json-schema/structuring.html
1. https://github.com/surajssd/talks/blob/master/golangmeetupNov2016/validate_intro_example3.go
1. https://github.com/json-schema/json-schema/wiki/$ref-traps