package main

import (
	"fmt"
	"github.com/luantafarel/portal-nexo-go/gettoken"
)

func main() {
	token, expires := get.token()
	fmt.Println("Token: ", token, "\n Expires: ", expires)
}
