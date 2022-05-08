package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

type resp struct {
	num  *big.Int
	time time.Duration
}

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8080")
	errorHandler(err)
	defer con.Close()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input, err := strconv.ParseInt(scan.Text(), 10, 64)
		errorHandler(err)

		encoder := json.NewEncoder(con)
		encod := encoder.Encode(input)
		errorHandler(encod)

		var msg resp
		decoder := json.NewDecoder(con)
		decod := decoder.Decode(&msg)
		errorHandler(decod)

		fmt.Printf("%s %d\n", msg.time, msg.num)
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
