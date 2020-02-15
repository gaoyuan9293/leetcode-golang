//Version1
func containsDuplicate(nums []int) bool {
    numsMap  := map[int]int{} 
    var res bool
    for _, v := range nums {
        numsMap[v] = 0
    }
    if len(nums) != len(numsMap) {
        res = true
    } else {
        res = false
    }
    return res 
}

//Version2-Better

func containsDuplicate(nums []int) bool {
    numsMap  := map[int]int{} 
    for k, v := range nums {
        if _, ok := numsMap[v]; ok {
            return true
        } 
        numsMap[v] = k
    }
    return false
}