package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello, World! Successfully signed with Certum Cloud Code Signing on %s/%s!\n", runtime.GOOS, runtime.GOARCH)
}
