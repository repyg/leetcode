func majorityElement(nums []int) int {
    ans, count := 0, 0
    for _, num := range nums{
        if count==0{ans = num}
        if ans == num{
            count++
        }else{
            count--
        }
    }
    return ans
}