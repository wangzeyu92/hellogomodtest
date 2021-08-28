package hello

//
//go mod tidy
//go get rsc.io/quote@v  可以指定版本
import "rsc.io/quote"
func Hello() string {
	//return "Hello, world."
	return quote.Hello()
}
