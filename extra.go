package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var i int = 0
	func() {
		con := bufio.NewScanner(os.Stdin)
		for con.Scan() {
			i += 1
			input := con.Text()
			//fmt.Println(input, i)
			io.WriteString(os.Stdout, input)

		}
		if err := con.Err(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
}
