package kata

func BreakChocolate(n, m int) int {
  // math goes here
  if n<=1 && m <=1{
    return 0
  }
  if (m-1)+(n-1)*m<=0{
    return 0
  }
  return (m-1)+(n-1)*m
}
