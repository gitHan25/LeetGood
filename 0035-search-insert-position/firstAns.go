func searchInsert(nums []int, target int) int {
    var found int
    for i:=0;i<len(nums);i++{
        if nums[i] == target{
            return i
        }else if nums[i]<target  && i+1 <len(nums)+1 {
            found = i+1
        }
    }
    return found
    
}
