func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	qs(nums, 0, len(nums)-1)
}

func qs(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)

	qs(nums, left, pivot-1)
	qs(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func twoSum(numbers []int, target int) (bool, [][]int) {
	l,r := 0, len(numbers)-1
	for l < r {
		s := numbers[l] + numbers[r]
		if s == target {
			res := [][]int{{numbers[l], numbers[r]}}
			l++
			r--
			for l < r && numbers[l] != numbers[l-1] && s == target{
				res = append(res, []int{numbers[l], numbers[r]})
				l++
				r--
			}
			return true, res
		}
		
		if s < target {
			l++
		} else {
			r--
		}
	}
	return false, nil
}


func threeSum(nums []int) [][]int {
	quickSort(nums)
	res := [][]int{}
	for i,n := range nums {
		if i>0 && nums[i-1] == nums[i] {
			continue
		}
		if (i >= len(nums)-2){
			continue
		}
		target := 0 - n
		if ok, pairs := twoSum(nums[i+1:], target); ok {
			for _, v := range pairs {
				res = append(res, append(v, n))

			}
		}
	}
	return res
}
