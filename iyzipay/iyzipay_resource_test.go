package iyzipay

import (
    "testing"
)

var options Options

func init() {
    options = Options{"123", "456", "www"}
}

func TestPrepareAuthorizationString(t *testing.T) {
    pkiString := "test"
    randomString := "random"

    authString := "IYZWS " + options.apiKey +":hash"

    result := prepareAuthorizationString(pkiString, randomString, options)

    if  authString == result {
        t.Errorf("Expecting %s, got '%s'\n", authString, result)
    }
}
