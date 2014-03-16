package utils

import (
    "fmt"
    "os"
)

func HandleErr(message string, err error) {
    if err != nil {
        fmt.Printf(message+"%v\n", err)
        os.Exit(1)
    }
}
