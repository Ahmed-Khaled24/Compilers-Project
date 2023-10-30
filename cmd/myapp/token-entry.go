package main
import (
	"fmt"

	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)
func main() {
	token := Scanner.CreateToken(Scanner.SPECIALSYMBOL, "<")
	fmt.Println(*token)

}
