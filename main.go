package main

import (
    "blokchain/account"
    "fmt"
)

func main() {
    myBlockchain, _, err := account.InitBlockchain(1000000)
    fmt.Println(err)
    fmt.Println(myBlockchain)
}
