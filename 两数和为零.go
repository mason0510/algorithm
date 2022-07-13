package algorithm


func twoSum(nums []int, target int) []int {
	length:=len(nums)
	if length <=0 {
		return nil
	}
	for i := 0; i < length; i++ {
		for j := i+1; j <length ; j++ {
			if nums[i]+nums[j]==target{
				return []int{i,j}
			}
		}
	}

	return nil
}
