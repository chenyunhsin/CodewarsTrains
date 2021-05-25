package kata
func Cnt(item rune,target,target2 rune,arr []rune)int{
   times:=0
  for j:=0;j<len(arr);j++{
   
      if arr[j]==target || arr[j]==target2{
        times++
        
      }
    }
  return times
}
func CntOdd(item rune,target rune,arr []rune)int{
   times:=0
  for j:=0;j<len(arr);j++{
   
      if arr[j]==target {
        times++
        
      }
    }
  return times
}
func FirstNonRepeating(str string) string {
  //your code here
  arr:=[]rune(str)
  for i:=0;i<len(arr);i++{
    target:=arr[i]
    target2:=arr[i]
    times:=0
    if target<97{
      target2=target+32
    }else{
      target2=target-32
    }
    
    if target<65 || target>122 {
      times= CntOdd(arr[i],target,arr)
    }else if target>90 &&target<97{
      times= CntOdd(arr[i],target,arr)
      
    }else {
       times = Cnt(arr[i],target,target2,arr)
   
    }
    if times==1 {
      return string(rune(target))
    }
    
  }
  return ""
}
