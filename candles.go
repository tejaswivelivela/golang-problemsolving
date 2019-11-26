//if your niece is turning 4 years old, and the cake will have 4 candles of height 4, 4, 1, 3, 
//she will be able to blow out 2 candles successfully, since the tallest candles are of height 4 and there are 2 such candles.

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

func birthdayCakeCandles(ar []int32) int32 {

    sort.Slice(ar[:], func(i, j int) bool {
        return ar[i] < ar[j]
    })

    var total int32
    max := ar[len(ar)-1]

    for _,v := range ar{
        if v == max{
            total = total + 1
        }
    }
    
    return total
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(arCount); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := birthdayCakeCandles(ar)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
