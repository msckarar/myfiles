package main

import "fmt"
var x,y,s,d,c,i int
var sum[5] int
func testcase(d int){
  s=0
  if (d>0 && d<=100){
    fmt.Scan(&x)
  sum[i]= integer(x)
    i++
    testcase(d-1)
  }
}
func integer(x int )int{
  if(x>0 && x<=100){
    fmt.Scan(&y)
    if(y>0 && x<=100){
      s=s+y*y
    }
    integer(x-1)
  }
  return s
}
func printsum( ){
  if(c<i){
    fmt.Println(sum[c])
    c++
    printsum( )
  }
}
func main() {
  i=0
  c=0
  fmt.Scan(&d)
  testcase(d)
  printsum( )
}
