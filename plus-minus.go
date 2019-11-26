package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func plusMinus(arr []int32) {

    var positive,negative,zero int
    total := float64(len(arr))

    for _,v := range arr{
         if v > 0{
            positive = positive +1
        }
        if v < 0{
            negative = negative +1
        }
        if v == 0{
            zero = zero +1
        }
    }

    p := fmt.Sprintf("%6f", float64(positive)/total)
    n := fmt.Sprintf("%6f", float64(negative)/total)
    z := fmt.Sprintf("%6f", float64(zero)/total)
    fmt.Println(p)
    fmt.Println(n)
    fmt.Println(z)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
