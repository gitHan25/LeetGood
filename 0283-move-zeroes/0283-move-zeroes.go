func moveZeroes(nums []int)  {
    
      var k int = 0
    
    
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[k] = nums[i]
           
            k++
        }
    }
    for i:=k;i<len(nums);i++{
        nums[i]=0
    }

   
}