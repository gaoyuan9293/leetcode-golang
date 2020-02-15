//METHOD 1
func findLHS(nums []int) int {
	var LenSlice []int
	for i := 0; i < len(nums); i++ {
		var total, count, count1 int
		for j := 0; j < len(nums); j++ {
			if nums[j]-nums[i] == 1 {
				count++
			} else if nums[j] == nums[i] {
				count1++
			} else {
				continue
			}
		}

		if count == 0 || count1 == 0 {
			total = 0
		} else {
			total = count + count1
		}
		LenSlice = append(LenSlice, total)
	}
	sort.SliceStable(LenSlice, func(i, j int) bool {
		return LenSlice[i] > LenSlice[j]
	})

	if len(LenSlice) == 0 {
		return 0
	} else {
		return LenSlice[0]
	}
}

//METHOD 2 -BETTER
func findLHS(nums []int) int {
    numsMap := make(map[int]int)
    for _, v := range nums {
        count := 1
        if _, ok := numsMap[v]; !ok {
            numsMap[v] = count
        } else {
            numsMap[v]++
        }
    }

    LenSlice  := []int{}
    for k, v := range numsMap{
        var length int
        if v1, ok := numsMap[k+1]; ok {
            length = v +  v1
        }
        LenSlice = append(LenSlice, length)
    }
    sort.SliceStable(LenSlice, func(i, j int) bool{
        return LenSlice[i] > LenSlice[j]
    })

    if len(LenSlice) == 0 {
        return 0
    } else {
        return LenSlice[0]
    }
}
