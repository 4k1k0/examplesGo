package service

import (
	"config/config"
	"fmt"
)

func Start() {
	conf := config.GetConfig()
	fmt.Printf("Hello %s. You are %d years old.\n", conf.Name, conf.Age)
	foo()
	bar()
}

func foo() {
	conf := config.GetConfig()
	fmt.Printf("Hello from foo> %s | %d\n", conf.Name, conf.Age)
}

func bar() {
	conf := config.GetConfig()
	fmt.Printf("Hello from bar> %d | %s\n", conf.Age, conf.Name)
}
