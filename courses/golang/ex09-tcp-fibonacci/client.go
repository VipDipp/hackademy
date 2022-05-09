package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type Resp struct {
	num  *big.Int
	time time.Duration
}

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8080")
	errorHandler(err)
	defer con.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		str, err := reader.ReadString('\n')
		input, err := strconv.Atoi(strings.TrimSuffix(str, "\r\n"))
		errorHandler(err)

		encoder := json.NewEncoder(con)
		encod := encoder.Encode(input)
		errorHandler(encod)

		var num *big.Int
		var time time.Duration

		decoder := json.NewDecoder(con)
		decod := decoder.Decode(&num)
		decod = decoder.Decode(&time)
		errorHandler(decod)

		fmt.Println(time, num)
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
