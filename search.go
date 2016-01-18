package itunes4go

import (
    "encoding/json"
    "github.com/google/go-querystring/query"
    "io/ioutil"
    "log"
    "net/http"
)

const url = "http://ax.itunes.apple.com/WebObjects/MZStoreServices.woa/wa/wsSearch?"

func Search(req Request) *Response {
    v, err := query.Values(req)
    if err != nil {
        log.Fatalln(err.Error())
    }

    res, err := http.DefaultClient.Get(url + v.Encode())
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer res.Body.Close()

    buf, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalln(err.Error())
    }

    resp := new(Response)
    if err = json.Unmarshal(buf, &resp); err != nil {
        log.Fatalln(err.Error())
    }

    return resp
}