# Golang Dimensions - parser for weight, width values

[![Build Status](https://travis-ci.org/mantyr/dimensions.svg?branch=master)](https://travis-ci.org/mantyr/dimensions) [![GoDoc](https://godoc.org/github.com/mantyr/dimensions?status.png)](http://godoc.org/github.com/mantyr/dimensions) [![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This not stable version

## Installation

    $ go get github.com/mantyr/dimensions

## Example
```GO
package main

import (
    "github.com/mantyr/dimensions"
    "fmt"
)

func main() {
    width := dimensions.NewDimension()
    width.SetDimensionWidth()
    width.SetDefaultType("DEFAULT_TYPE")

    width.Parse("1 179.20m")

    fmt.Println(price.Get())        // print "1179.20"
    fmt.Println(price.GetInt())     // print "1179"
    fmt.Println(price.GetInt64())   // print "1179"
    fmt.Println(price.GetFloat64()) // print "1179.2"
    fmt.Println(price.GetType())    // print "M"        // meter

}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
