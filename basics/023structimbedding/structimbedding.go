/*
 * Go supports embedding of structs and interfaces 
 * to express a more seamless composition of types
 */

 package main
 import "fmt"

 type base struct {
     num int
 }

 type container struct {
     base 
     str string
 }

 func (b base) describe() string {
     return fmt.Sprintf("base with num=%v", b.num)
 }

 func (b base) greet() {
     fmt.Println("Hello")
 }


 func main() {
     co := container{
         base: base{
             num: 1,
         },
         str: "some name",
     }

     fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
     fmt.Println("also num:", co.base.num)
     fmt.Println("describe:", co.describe())

     co.base.greet()
     co.greet()


     type describer interface {
         describe() string
     }

     var d describer = co
     fmt.Println("describer:", d.describe())
 }
