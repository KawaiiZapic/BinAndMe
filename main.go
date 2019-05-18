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
        var name = usr.Name
        var nameline = "    =@@@^"
        //TODO:Count name length and insert into Line
        fmt.Println("               ,]OOOOOOOOOOO]`              ")
        fmt.Println("             ,@@@@@@@@@@@@@@@@@^            ")
        fmt.Println("            @@@@`            @@@@`          ")
        fmt.Println(",OOOOOOOOOO@@@@@OOOOOOOOOOOOOO@@@@OOOOOOOOOO")
        fmt.Println(" @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
        fmt.Println("    =@@@^                           =@@@/   ")
        fmt.Println(nameline)
        fmt.Println("    =@@@^                           =@@@^   ")
        fmt.Println("    =@@@^   ,/OO`   ,/OO`   ,/OO`   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^   =@@@^   =@@@^   =@@@^   =@@@^   ")
        fmt.Println("    =@@@^    ,[[     ,[[     ,[[    =@@@^   ")
        fmt.Println("    =@@@^                           =@@@^   ")
        fmt.Println("    =@@@^                           =@@@^   ")
        fmt.Println("     @@@@\\]]]]]]]]]]]]]]]]]]]]]]]]]/@@@@`   ")
        fmt.Println("      @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@`    ")

}
