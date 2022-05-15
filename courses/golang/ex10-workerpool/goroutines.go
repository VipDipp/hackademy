package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Run(poolSize int) {
	jobs := make(chan float64, poolSize)
	i := 1
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input := string(scan.Bytes())
		f, _ := strconv.ParseFloat(input, 64)
		jobs <- f
		if i <= poolSize {
			go work(i, jobs)
			i++
		}
	}
	close(jobs)
}

func work(id int, jobs <-chan float64) {
	fmt.Printf("worker:%d spawning\n", id)
	for job := range jobs {
		fmt.Printf("worker:%d sleep:%.1f\n", id, job)
		time.Sleep(time.Duration(int(job*1000)) * time.Millisecond)
	}
	fmt.Printf("worker:%d stopping\n", id)
}
