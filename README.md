# googleauthenticator
Two- / Multi- Factor Authenication (2FA / MFA) for Golang

[![codecov](https://codecov.io/gh/dev-templates/googleauthenticator/branch/main/graph/badge.svg)](https://codecov.io/gh/dev-templates/googleauthenticator)
[![Build Status](https://github.com/dev-templates/googleauthenticator/workflows/build/badge.svg)](https://github.com/dev-templates/googleauthenticator)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/dev-templates/googleauthenticator)
[![go.mod](https://img.shields.io/github/go-mod/go-version/dev-templates/googleauthenticator)](go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/dev-templates/googleauthenticator)](https://goreportcard.com/report/github.com/dev-templates/googleauthenticator)
[![LICENSE](https://img.shields.io/github/license/dev-templates/googleauthenticator)](LICENSE)

## Installation

To get the package, execute:

```bash
go get github.com/dev-templates/googleauthenticator
```

## Usage
```go
package main

import (
    "fmt"

	"github.com/dev-templates/googleauthenticator"
)

func main {
    // generate key
	formattedKey := googleauthenticator.GenerateKey()
	authenticator := googleauthenticator.NewAuthenticator("issuer", "xxx@gmail.com", formattedKey)
    // generate uri for show
    uri := authenticator.GenerateTotpUri()
    fmt.Println(uri)
    // verify token
	passcode := "<from input>"
	if authenticator.VerifyToken(passcode) {
        // ok
    }
}
```