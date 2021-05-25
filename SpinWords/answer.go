package kata
import(
 "strings"
)
func isMoreThanFive(str string)bool{
  if len(str)>=5{
    return true
  }
  return false
}
func spin(str string)string{
  result:=""
  for i:=len(str)-1;i>=0;i--{
    result+=string(str[i])
  }
  return result
}

func SpinWords(str string) string {
  //one or more
  //more than five
  result:=""
  strSplits := strings.Split(str, " ")
  for i:=0;i<len(strSplits);i++{
    target:= strSplits[i]
    if isMoreThanFive(target){
      target = spin(target)
    }
    if i==len(strSplits)-1{
      result+=target
    }else{
      result+=target+" "
    }
    
  }
  return result
}// SpinWords
