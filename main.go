package main
import (
	"time"
	"fmt"
	_ "github.com/Lineblocs/go-helpers"
)

func main() {

	for ;; {
		fmt.Println("Running billing commands")
		time.Sleep(2 * time.Second)
	}
}