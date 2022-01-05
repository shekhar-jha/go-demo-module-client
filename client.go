package main

import (
	"github.com/shekhar-jha/go-demo-module/mod1"
	"github.com/shekhar-jha/go-demo-module/mod2"
	"github.com/shekhar-jha/go-demo-module/mod3"
	"fmt"
)

func main() {
	fmt.Println("Entering main...")
	fmt.Println("Calling mod1.Mod1(). Result:", mod1.Mod1())
	fmt.Println("Calling mod2.Mod2(). Result:", mod2.Mod2())
	fmt.Println("Calling mod3.Mod3(). Result:", mod3.Mod3())
	fmt.Println("Exiting main")
}
