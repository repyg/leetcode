func hIndex(citations []int) int {
    sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
    hIndex := 0
    for i, num := range citations{
        if i+1<=num{
            hIndex = i+1
        }else{
            return hIndex
        }
    }
    return hIndex
}