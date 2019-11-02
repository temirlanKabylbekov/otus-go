package main

import (
	"flag"
	"fmt"
)

const READ_ALL_THE__REST_OF_THE_FILE int64 = -1

var (
	from   string
	to     string
	offset int64
	limit  int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", READ_ALL_THE__REST_OF_THE_FILE, "limit to read from file")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	err := Copy(from, to, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
}
