package kata

func FindOdd(seq []int) int {
    target:=0
  for i:=0;i<len(seq);i++{
    target = seq[i]
    times:=0
    for j:=0;j<len(seq);j++ {
      if target == seq[j]{
        times++
      }
    }
    if times%2==1{
      return target
    }
  }
  return -1
}
