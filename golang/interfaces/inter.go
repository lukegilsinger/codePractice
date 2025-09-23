package main

import (
	"fmt"
	"strings"
)

type DataProcessor2 interface {
	Process2(data string) strings
}