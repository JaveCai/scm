/* 
stroke count mapping 
*/

package scm

import(
	"os"
	"strings"
	"strconv"
	_ "fmt"
	"bufio"
)

var count = make(map[string]int)

func init(){
	f, err := os.Open("utf8_Stroke.txt")
	if err != nil{
		panic(err)
	}
	var line int = 1
	input := bufio.NewScanner(f)
	for input.Scan() {
		ss := strings.Split(input.Text(), " ")
		c, err:=strconv.Atoi(ss[2])
		if err!=nil{
			panic(err)
		}
		count[ss[1]] = c
		if line == 20964 {
			break
		} else {
			line++
		}

	}

}

func GetStrokeCount(s string) int{
	for _, r:= range s{
		c,ok:=count[string(r)]
		if !ok {
			return 0
		}else{
			return c
		}
		break
	}
	return 0

}