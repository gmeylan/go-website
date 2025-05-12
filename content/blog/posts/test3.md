---
published-on: 2025-03-05
title: Test 3
author: Grégoire Meylan
tags: Go, devOps, xxx
slug: test3
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