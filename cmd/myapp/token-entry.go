package main
import (
	"fmt"

	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)
func main() {
	Scanner.Fsm.Scan("if x = 5 ;")
	fmt.Println(Scanner.Fsm.TokenList)
	
}
