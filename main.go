package main

import (
    "os"
    "fmt"
    "iyzipay-go/iyzipay"
)

func main() {
    var options iyzipay.Options
    options.ApiKey = os.Getenv("API_KEY")
    options.SecretKey = os.Getenv("SECRET_KEY")
    options.BaseUrl = "https://sandbox-api.iyzipay.com"

    apiTest(options)
    binNumber(options);
}

func binNumber(options iyzipay.Options)  {
    request := `{
        "locale": "tr",
        "conversationId": "123456789",
        "binNumber": "542119"
    }`
	result := iyzipay.BinNumber(request, options)
	fmt.Println("BinNumber: ", result)
}

func apiTest(options iyzipay.Options)  {
    result := iyzipay.ApiTest(options)
    fmt.Println("ApiTest: ", result)
}
