func maxProfit(prices []int) int {
    minPrice, benefit := prices[0], 0
    for i:=1; i<len(prices); i++{
        if prices[i]<minPrice{
            minPrice = prices[i]
        }else if prices[i]-minPrice > benefit{
            benefit = prices[i]-minPrice
        }
    }
    return benefit
}