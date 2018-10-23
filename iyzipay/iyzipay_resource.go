package iyzipay

import (
    "time"
    "io/ioutil"
    "net/http"
    "crypto/sha1"
    "encoding/base64"
    "strings"
)

const Version = "0.2.0"

type Options struct {
    ApiKey string
    SecretKey string
    BaseUrl string
}

func connect(method string, url string, options Options, request string, pkiString string) string {
    var randomString string = time.Now().String()
    var authString = prepareAuthorizationString(pkiString, randomString, options)

    payload := strings.NewReader(request)

    req, _ := http.NewRequest(method, url, payload)
    req.Header.Add("accept", "application/json")
    req.Header.Add("content-type", "application/json")
    req.Header.Add("authorization", authString)
    req.Header.Add("x-iyzi-rnd", randomString)
    req.Header.Add("cache-control", "no-cache")

    res, _ := http.DefaultClient.Do(req)
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)

    return string(body)
}

func prepareAuthorizationString(pkiString string, randomString string, o Options) string {
    var h = sha1.New()
    h.Write([]byte(o.ApiKey + randomString + o.SecretKey + pkiString))
    var sha1 []byte = h.Sum(nil)
    return "IYZWS " + o.ApiKey +":"+ base64.StdEncoding.EncodeToString(sha1)
}
