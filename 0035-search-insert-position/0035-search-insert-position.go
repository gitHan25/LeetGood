func searchInsert(nums []int, target int) int {
   
    var left int = 0
    var right int = len(nums)-1
    var found int = len(nums)
    
    for left<=right{
        mid:=(left+right)/2
        if nums[mid]==target{
            return mid
        }
        if nums[mid]<target{
            left = mid + 1
        }
        if nums[mid]>target{
            found = mid
            right = mid - 1
        }
    }
    return found
    
}