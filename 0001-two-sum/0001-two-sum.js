/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let result = 0
    let pointer = 0
    for (let i = 0;i<nums.length;i++){
        for (let j=0;j<nums.length;j++){
            if (j == pointer){
                j+=1
            }
            result = nums[pointer]+nums[j]
                if (result == target){
                    return [pointer,j]
                }    
            

        }
        pointer++

    }
}