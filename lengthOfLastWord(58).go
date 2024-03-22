func lengthOfLastWord(s string) int {
    if len(s)==0{
        return 0
    }
    lenLastWorld :=0
    for i:=len(s)-1; i>=0; i--{
        if string(s[i])!=" "{
            lenLastWorld++
        }else if lenLastWorld>0{
            return lenLastWorld
        }
    }
    return lenLastWorld
}