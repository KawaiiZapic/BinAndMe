package main

import (
        "fmt"
        "os/user"
)

func main() {
        usr, err := user.Current()
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println("I'm a rubbish in the bin. --",usr.Name)
}
