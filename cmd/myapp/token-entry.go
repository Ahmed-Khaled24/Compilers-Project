package main
import (
	"fmt"

	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)
func main() {
	// token := Scanner.CreateToken(Scanner.SPECIALSYMBOL, "=")
	// fmt.Println(*token)

	// Scanner.Fsm.Scan("read 3;    if  0 < x   then     fact  := 1;repeat fact  := fact *  x; x  := x  -  1until  x  =  0;write  fact  end")
	Scanner.Fsm.Scan(`{ Sample program in TINY language – computes factorial}
	read 3;   {input an integer } 
	if  0 < x   then     {  don’t compute if x <= 0 }
	   fact  := 1;
	   repeat 
		  fact  := fact *  x;
		   x  := x  -  1
	   until {rerer} x  =  0;
	   write  fact  
	end`)
	fmt.Println(Scanner.Fsm.TokenList)
	
}
