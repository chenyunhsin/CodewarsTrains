package kata
import "strings"
func replace (s1 string)(int, string){
  cnt:=0
  result:=s1
  target :=s1[0]
   localCnt:=0
  localCnt+= strings.Count(result, string(target))
  result = strings.ReplaceAll(result, string(target), "")
    if 65<=target &&target<=90{   
      localCnt+= strings.Count(result, string(target+32))   
      result = strings.ReplaceAll(result, string(target+32), "")
    }else  if 97<=target &&target<=122{
      localCnt+= strings.Count(result, string(target-32))
      result = strings.ReplaceAll(result, string(target-32), "")
    }
  if localCnt>1{
        cnt =1
      }
  return cnt,result
}
func duplicate_count(s1 string) int {
    //your code goes here
    cnt:=0
  target := s1
  for len(target)>0{
    res1,res2 := replace(target)
    cnt += res1
    target = res2
  }
    return cnt //Instead of returning '1', you have to return integer 'i' as answer of solution.  
    }

