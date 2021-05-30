package kata
import "strings"
func replace(str string) string{
  target:=str
  target = strings.ReplaceAll(target,"()","")
  target = strings.ReplaceAll(target,"[]","")
  target = strings.ReplaceAll(target,"{}","")
  return target
}
func ValidBraces(str string) bool {
  target := str
  for i:=0;i<len(str)/2;i++{
    target = replace(target)
  }
  if len(target)==0{
    return true
  }
  return false
  
}
