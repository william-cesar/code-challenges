/*
[2, 1, 3, 5, 3, 2] -> {2:1, 1:1, 3:1, 5:1}
*/

func solution(a []int) int {
    kvMap := map[int]int{}
    
    for _, num := range a {
       if _, ok := kvMap[num]; ok {
            return num
        }
        
        kvMap[num] = 1
    }
     
    return -1
}
