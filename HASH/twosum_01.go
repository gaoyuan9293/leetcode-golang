//METHOD 1

func twoSum(nums []int, target int) []int {
    var result []int
LOOP1:
    for i := 0 ; i < len(nums); i++ {
        for j := i+1 ; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                result = append(result, i, j)
                break LOOP1
            }
        }
    }

    return result 
}

//METHOD 2 -BETTER
func twoSum(nums []int, target int) []int {
    numsMap := map[int]int{}
    for i, v := range nums {
        if j, ok := numsMap[target-v]; ok  {
            return []int{j, i}
        }
        numsMap[v] = i
    }
    return nil
}