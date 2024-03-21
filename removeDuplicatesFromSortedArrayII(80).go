func removeDuplicates(nums []int) int {
    index, count := 1, 1
    for i:=1; i<len(nums); i++{
        if nums[i]==nums[i-1]{
            count++
        }else{
            count = 1
        }
        if count<=2{
            nums[index]=nums[i]
            index++
        }
    }
    return index
}