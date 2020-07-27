package main

import "log"

func main(){
    log.Println("default")

    log.SetFlags(0)

    log.Println("setflag0")
    log.SetFlags(1)
    log.Println("setflag1")
    log.SetFlags(2)
    log.Println("setflag2")
    log.SetFlags(3)
    log.Println("setflag3")
    log.SetFlags(4)
    log.Println("setflag4")
}
