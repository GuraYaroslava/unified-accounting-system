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

func ArrayStringToInterface(array []string) []interface{} {
    result := make([]interface{}, len(array))
    for i, v := range array {
        result[i] = interface{}(v)
    }
    return result
}

func IsExist(array []string, value string) bool {
    for _, v := range array {
        if v == value {
            return true
        }
    }
    return false
}
