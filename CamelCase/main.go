// https://www.codewars.com/kata/587731fda577b3d1b0001196/train/gopackage main
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CamelCase("camel case method"))
}

func CamelCase(s string) string {
	return strings.ReplaceAll(strings.Title(s), " ", "")
}
