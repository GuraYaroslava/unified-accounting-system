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

func ArrayInterfaceToString(array []interface{}, length int) []string {
    result := make([]string, length)
    for i, v := range array {
        result[i] = v.(string)
    }
    return result
}
