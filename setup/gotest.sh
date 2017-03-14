mkdir -p ~/goprojects/src/test/
cat > ~/goprojects/src/test/hello.go <<'endmsg1989'
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go is working fine")
}
endmsg1989

go run ~/goprojects/src/test/hello.go
