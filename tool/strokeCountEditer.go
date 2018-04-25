package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please input start, end, and offset")
		return
	}

	start, _ := strconv.Atoi(os.Args[1])
	end, _ := strconv.Atoi(os.Args[2])
	offset, _ := strconv.Atoi(os.Args[3])

	f, err := os.OpenFile("utf8_Stroke.txt", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}

	var buf = make([]string, 0, 20964)
	var last = make([]string, 0, 1)
	var line int = 1

	/* TODO:
	1. read line from file
	2. if line number in range ? yes --> 3, no --> 4
	3. assign new value to the stroke count
	4. write line to buffer, if EOF? yes -->5, no-->1
	5. empty the data of file and write buffer to it.
	*/

	input := bufio.NewScanner(f)
	for input.Scan() {
		if line >= start && line <= end {
			s := StrokeUpdate(input.Text(), offset)
			buf = append(buf, s+"\n")
		} else {
			buf = append(buf, input.Text()+"\n")
		}
		last = append(last[:0], input.Text())
		if line == 20964 {
			break
		} else {
			line++
		}

	}
	//fmt.Printf("%s\n", buf[len(buf)-1])
	//fmt.Printf("%s\n", last[0])
	//buf = append(buf[:len(buf)-1], last[0])

	fmt.Printf("len:%d\n", len(buf))
	f.Seek(0, 0)
	for i, b := range buf {
		if i <= 20964 {
			_, err = f.WriteString(b)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			break
		}
	}
	/* end of TODO */

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

/*
function for update stroke count
*/
func StrokeUpdate(s string, offset int) string {
	ss := strings.Split(s, " ")
	iNew, err := strconv.Atoi(ss[2])
	if err != nil {
		log.Fatal(err)
	}
	iNew = iNew + offset
	sNew := strconv.Itoa(iNew)
	return ss[0] + " " + ss[1] + " " + sNew
}
