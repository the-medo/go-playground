package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return sum
		}
		if err != nil {
			fmt.Println("Error:", err)
		}

		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Error:", err)
		}

		sum += num
	}
}

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"} //817,974,962,928,1154
	sum := 0

	for i := 0; i < len(files); i++ {
		fileName := files[i]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		rd := bufio.NewReader(file)

		calculate := func(fileName string) {
			fmt.Println("Reading...", fileName)
			fileSum := sumFile(*rd)
			sum += fileSum
		}

		go calculate(fileName)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(sum)

}
