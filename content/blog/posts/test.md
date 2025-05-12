---
published-on: 1988-06-11
title: Blogging Like a Boss
author: Grégoire Meylan
tags: Go, devOps, kkk
slug: test
---


# Hello world

## Testouille

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```