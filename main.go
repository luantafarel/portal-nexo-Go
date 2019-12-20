package main

import (
	"fmt"
	"token"
)

func main() {
	token, expires := token.get()
	fmt.Println("Token: ", token, "\n Expires: ", expires)
}
