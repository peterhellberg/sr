sr (api v2)
==========

Go library for [Sveriges Radio API](http://sverigesradio.se/api/documentation/v2/index.html)

[![GoDoc](https://godoc.org/github.com/peterhellberg/sr?status.svg)](https://godoc.org/github.com/peterhellberg/sr)
[![Build Status](https://travis-ci.org/peterhellberg/sr.svg?branch=master)](https://travis-ci.org/peterhellberg/sr)

## Installation

```bash
go get -u github.com/peterhellberg/sr
```

## Example usage

Getting the toplist for the current day.

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/sr"
)

func main() {
	sr := sr.NewClient(nil)

	shows, err := sr.Toplist.GetDay()

	if err == nil {
		for i, show := range shows {
			fmt.Println(i+1, show.Title)
		}
	}
}
```

## Services

### Channels
### Episodes
### News
### Playlists
### Program Categories
### Programs
### Scheduled Episodes
### Sport
### Toplist

## Utils

Developed using [JSON-to-Go by](http://mholt.github.io/json-to-go/) by Matthew Holt.

## License

> *The MIT License (MIT)*
>
> Copyright (c) 2014 [Peter Hellberg](http://c7.se/)
>
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in all
> copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
> SOFTWARE.
