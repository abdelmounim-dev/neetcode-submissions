func productExceptSelf(nums []int) []int {
    l := len(nums)
    res := make([]int, l)
    prefix, postfix := 1,1
    for i,v := range nums {
        res[i] = prefix
        prefix *= v
    }
    
    for i := l-1; i >= 0; i--{
        res[i] *= postfix
        postfix *= nums[i]
    }
    return res
}