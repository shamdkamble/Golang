package main 

import (

"fmt"
"math/rand"

)
var var1 int = 10

const var2 string = "sham"


func main(){

   x:= rand.Intn(10)
   y:= rand.Intn(10)
   
   switch{
   
   case x<4 && y<4:
   
   fmt.Printf("The value %v and %v is less than 4 \n",x,y)
   
   case x>6 && y>6:
   fmt.Printf("The value %v and %v is greater than 6\n",x,y)
   
   
   
   case x>=4 && x<=6:
   fmt.Printf("The value %v is >= 4 and <=6 ",x)
   
   case y!=5 :
   fmt.Printf("The value %v is not = 5 ",y)
   
   default:
   fmt.Println("None of the  cases met")   
  }
}
