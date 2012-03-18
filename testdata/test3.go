package main

import (
       "fmt"
)

func main() {
       for {
               var s string
               _,err := fmt.Scanln(&s)
		if err != nil {
		}
               _,err = fmt.Scanln(&s)
		if err != nil {
		}
               fmt.Scanln(&s)
       }
}

