# GO-MM-NRC-Validator


[![Build Status](https://travis-ci.org/winhtaikaung/Rabbit-Go.svg?branch=master)](https://travis-ci.org/winhtaikaung/Rabbit-Go)

[![GO Report](https://goreportcard.com/badge/github.com/winhtaikaung/Rabbit-Go)](https://goreportcard.com/report/github.com/winhtaikaung/Rabbit-Go)

**Myanmar NRC Validator for GoLang that validates the NRC format using township data**

## Installation

install using go get

`go get github.com/ezysign/go-mm-nrc-validator/validator`

## Usage

playground link

https://play.golang.org/p/W9mx5cGZFb2

```
package main

import (
	"fmt"

	"github.com/ezysign/go-mm-nrc-validator/validator"
)

func main() {
	v := validator.NewNRCValidator()
	fmt.Println(v.ExtractNrcInfo("၁၂/မဂတ(နိုင်)၀၉၄၁၈၆"))
	fmt.Println(v.VerifyNRC("၄/ကပလ(နိုင်)၀၉၄၁၈၆"))
	fmt.Println(v.ValidateNRC("၄/ကပလ(နိုင်)၀၉၄၁၈၆"))
	fmt.Println(v.GetMatchNRC("၄/ကပလ(နိုင်)၀၉၄၁၈၆"))

}

```

## License

MIT