func reverseWords(s string) string {
    s = " " + s
    
    var res strings.Builder
    
    right := len(s)-1
    
    for left := len(s)-2; left >= 0; left-- {
        if s[left] == ' ' && s[left+1] != ' ' {
            if res.Len() > 0 {
                res.WriteString(" ")
            }
            
            res.WriteString(s[left+1:right+1])
        }
        
        if s[left] != ' ' && s[left+1] == ' ' {
            right = left
        }
    }
    
    return res.String()
}