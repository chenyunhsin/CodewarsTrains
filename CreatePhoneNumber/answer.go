package kata
import ("fmt"
       // "strings"
       )
func CreatePhoneNumber(numbers [10]uint) string {
  result:="("
  for i:=0;i<10;i++{
    str := fmt.Sprint(numbers[i])
    result+=str
    if(i==2){
      result+=") "
    }
    if(i==5){
      result+="-"
    }
  }
 // result = strings.Replace(result, " ", "", -1)
  return result
}
