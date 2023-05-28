package regexpb

import (
	"fmt"
	"regexp"
)

func Main() {
	r := regexp.MustCompile(`\xff`)
	fmt.Println(r)
}

func doMatch(data []byte) {

}
