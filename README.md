# About
itunes4go is a [iTunes API](http://www.apple.com/itunes/affiliates/resources/documentation/itunes-store-web-service-search-api.html) library for Go.  

# Usage
## go get
`go get github.com/itunes4go`

## import
```
package main

import "github.com/itunes4go"

func main() {
    // Do something.
}
```

## call Search method

```
// Term is a required parameter.
req := itunes4go.Request{Term:"shokran"}
res := itunes4go.Search(req)
fmt.Println(res.Results[0].PreviewUrl)
```
