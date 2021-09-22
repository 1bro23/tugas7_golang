package main

import "fmt"
import "reflect"
import "runtime"

var a int = 10
var b string = "see"

func main(){
  runtime.GOMAXPROCS(2)
  go reflect1()
  reflect2()
}

func reflect1 (){
  var typea = reflect.ValueOf(a).Type()
  fmt.Println("Variabel pertama bertipe",typea)
}
func reflect2 (){
  var typeb = reflect.ValueOf(b).Type()
  fmt.Println("Variabel kedua bertipe",typeb)
}
