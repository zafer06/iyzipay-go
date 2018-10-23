package iyzipay

import (
    "encoding/json"
)

func ApiTest(options Options) string {
    reqUrl := options.BaseUrl + "/payment/test"
    pkiString := ""

    return connect("GET", reqUrl, options, "", pkiString)
}

func BinNumber(request string, options Options) string {
    var req map[string]interface{}
    json.Unmarshal([]byte(request), &req)

    var pki PkiBuilder
    pki.append("locale", req["locale"].(string))
    pki.append("conversationId", req["conversationId"].(string))
    pki.append("binNumber", req["binNumber"].(string))

    reqUrl := options.BaseUrl + "/payment/bin/check"
    pkiString := pki.getPkiString()

    return connect("POST", reqUrl, options, request, pkiString)
}
