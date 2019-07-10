package main

import (
        "fmt"
        "strings"
        "runtime"
        "os/user"
        "os/exec"
)
func exeSysCommand(cmdStr string) string {
    cmd := exec.Command(cmdStr)
    opBytes, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
        return ""
    }
    return string(opBytes)
}
func main() {
        var name = ""
        if (runtime.GOOS!="windows"){
                whoname := exeSysCommand("whoami")
                whoname = strings.Replace(whoname, "\n", "", -1)
                whoname = strings.Replace(whoname, "\r", "", -1)
                name = "|"+whoname+"|"
        }else{
                usr, err := user.Current()
                if err != nil {
                        fmt.Println(err)
                }
                name = "|"+usr.Name+"|"
        }
        var nameline = "    =@@@^"
        var namelen = strings.Count(name,"")-1
        const spacelen = 27
        var flen = 0
        var blen = 0
        if (namelen % 2 != 0){
                flen = (spacelen - namelen)/2 - 1
                blen = flen + 1

        }else{
                flen = (spacelen - 1 - namelen)/2 - 1
                blen = flen + 1
        }
        for a:=0;a <= flen;a++ {
                nameline =nameline + " "
        }
        nameline = nameline + name
        for a:=0;a <= blen;a++ {
                nameline =nameline + " "
        }
        nameline = nameline + "=@@@/   "
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
