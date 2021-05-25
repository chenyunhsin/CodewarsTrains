package kata
import (
"strings"
  "strconv"
)
func isNum(ip string)bool{
  if x, err := strconv.Atoi(ip); err == nil {
    if x>255{
      return false
    }
    return true
  }
  return false
}
func isStartWithZero(ip string)bool{
  if ip=="0"{
    return true
  }
  if string(ip[0]) == "0" ||string(ip[0]) == "-" {
    return false
  }
  return true
}
func Is_valid_ip(ip string) bool {
  ipSplits:=strings.Split(ip,".")
  if len(ipSplits)!=4{
    return false
  }
  
  for i:=0;i<len(ipSplits);i++{
    if !isStartWithZero(ipSplits[i]) || !isNum(ipSplits[i]){
      return false
    }
  }
  return true
}
