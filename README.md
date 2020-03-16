# Forpay SDK for Go

[![API Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://api.forpay.pro/docs/overview)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/forpay/forpay-sdk-go/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/forpay/forpay-sdk-go.svg?branch=master)](https://travis-ci.org/github/forpay/forpay-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/forpay/forpay-sdk-go)](https://goreportcard.com/badge/github.com/forpay/forpay-sdk-go)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/4a6b410411514f90a8734635b7c39a15)](https://www.codacy.com/gh/forpay/forpay-sdk-go?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=forpay/forpay-sdk-go&amp;utm_campaign=Badge_Grade)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fforpay%2Fforpay-sdk-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fforpay%2Fforpay-sdk-go?ref=badge_shield)

## Installing

Use `go get` to retrive the SDK to add it to your `GOPATH` workspace, or project's go module dependencies.

```
go get github.com/forpay/forpay-sdk-go/forpay
```

To update the SDK use go get -u to retrieve the latest version of the SDK.

```
go get -u github.com/forpay/forpay-sdk-go/forpay
```

## Quick Start

``` golang
package main

import (
    "fmt"

    "github.com/forpay/forpay-sdk-go/forpay"
    "github.com/forpay/forpay-sdk-go/forpay/response"
)

var client *forpay.Client

func main() {
    var err error

    // Prepare required variables.
    appID := "00000000"
    keyID := "00000000000000000000000000000000"
    keyFile := "./priv.pem"

    // Create forpay sdk client instance.
    client, err = forpay.NewClientWithRSA(appID, keyID, keyFile)
    if err != nil {
        panic(err)
    }

    // Prepare required request parameters.
    currencyID := uint16(1)
    walletID := uint64(8888888888888888)
    // Create request using `CreateXXXRequest` funcs.
    req := forpay.CreateGetBalanceRequest(walletID, currencyID)

    // Send request and receive response.
    resp, err := client.GetBalance(req)
    if err != nil {
        // Check if the error is `*response.Error`.
        if e, ok := err.(*response.Error); ok {
            fmt.Println(e.Error())

            // You can handle specific errors according to the following info.
            // fmt.Println(e.HTTPStatus)
            // fmt.Println(e.IsBusinessFailed())
            // fmt.Println(e.Code)
            // fmt.Println(e.Msg)
            // fmt.Println(e.SubCode)
            // fmt.Println(e.SubMsg)
        } else {
            fmt.Println(err)
        }

        return
    }

    // Handle response.
    fmt.Println(resp.Data.CurrencyID)
    fmt.Println(resp.Data.Available)
    fmt.Println(resp.Data.Frozen)
    fmt.Println(resp.Data.Locked)
}
```

## License

This SDK is distributed under the Apache License, Version 2.0.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fforpay%2Fforpay-sdk-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fforpay%2Fforpay-sdk-go?ref=badge_large)