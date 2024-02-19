package main 

import "fmt"

var var1 int = 10

const var2 string = "sham"


func main(){

    var3 := 25
    fmt.Printf("The value of var3 is %v and the type is %T\n",var3,var3)
    fmt.Printf("The value of var2 is %v and the type is %T\n",var2,var2)
    fmt.Printf("The value of var1 is %v and the type is %T",var1,var1)
}
