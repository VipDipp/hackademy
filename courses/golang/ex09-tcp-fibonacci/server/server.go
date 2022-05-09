package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(time.Hour, time.Hour)

type Resp struct {
	num  *big.Int
	time time.Duration
}

func main() {
	dstream, err := net.Listen("tcp", ":8085")
	ErrorHandler(err)

	defer dstream.Close()

	for {
		con, err := dstream.Accept()
		ErrorHandler(err)
		handle(con)
	}
}

func fibo(num int) Resp {
	start := time.Now()
	output, found := c.Get(string(num))
	if found {
		return Resp{
			num:  output.(*big.Int),
			time: time.Since(start),
		}
	}

	buf := string(num)
	a := big.NewInt(0)
	b := big.NewInt(1)
	for num != 0 {
		a.Add(a, b)
		a, b = b, a
		num--
	}

	c.Set(buf, a, cache.NoExpiration)
	var out = Resp{
		num:  a,
		time: time.Since(start),
	}

	fmt.Println(out.num, out.time)
	return out
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func handle(con net.Conn) {
	for {
		var num int
		decoder := json.NewDecoder(con)
		decod := decoder.Decode(&num)
		ErrorHandler(decod)
		fmt.Println(num)

		response := fibo(num)
		encoder := json.NewEncoder(con)
		encod := encoder.Encode(response.num)
		encod = encoder.Encode(response.time)
		ErrorHandler(encod)
	}
}
