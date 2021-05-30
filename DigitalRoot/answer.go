package kata
func reduce(n int)int{
  result:=0
   for n>=10 {
    result+=n%10
    n/=10
  }
   result+=n%10
  return result
}
func DigitalRoot(n int) int {
  // ...
 // result :=0
  if n<10{
    return n
  }
  for n>=10 {
   n = reduce(n)
  }
  
  return n
}
