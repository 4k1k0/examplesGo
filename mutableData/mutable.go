package main

var response = "foo"

func Foo() string {
	return response
}

func Bar(s string) bool {
	return response == s
}
