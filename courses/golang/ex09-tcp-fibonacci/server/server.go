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

type resp struct {
	num  *big.Int
	time time.Duration
}

func main() {
	dstream, err := net.Listen("tcp", ":8080")
	ErrorHandler(err)

	defer dstream.Close()

	for {
		con, err := dstream.Accept()
		ErrorHandler(err)
		println("here")
		go handle(con)
	}
}

func fibo(num int) resp {
	start := time.Now()
	output, found := c.Get(string(num))
	if found {
		return resp{
			num:  output.(*big.Int),
			time: time.Since(start),
		}
	}

	println("here2")
	buf := string(num)
	a := big.NewInt(0)
	b := big.NewInt(1)
	for num != 0 {
		a.Add(a, b)
		a, b = b, a
		num--
	}
	c.Set(buf, a, cache.NoExpiration)
	var out = resp{
		num:  a,
		time: time.Since(start),
	}
	println(out.num, out.time)
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
		println("here1")
		decoder := json.NewDecoder(con)
		decod := decoder.Decode(&num)
		ErrorHandler(decod)

		resp := fibo(num)

		encoder := json.NewEncoder(con)
		encod := encoder.Encode(resp)
		ErrorHandler(encod)
	}
}
