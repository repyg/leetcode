type RandomizedSet struct {
    Val []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{Val: []int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
    if slices.Contains(this.Val, val) {
        return false
    }

    this.Val = append(this.Val, val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if !(slices.Contains(this.Val, val)) {
        return false
    }

    index := slices.Index(this.Val, val)
    this.Val = slices.Delete(this.Val, index, index+1)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    index := rand.Intn(len(this.Val))
    return this.Val[index]
}
