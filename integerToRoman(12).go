import "strings"

type IntRoman struct {
	integer int
	numeral string
}

func newIntRoman(integer int, numeral string) IntRoman {
	return IntRoman{integer, numeral}
}

var list = []IntRoman{
	newIntRoman(1000, "M"),
	newIntRoman(900, "CM"),
	newIntRoman(500, "D"),
	newIntRoman(400, "CD"),
	newIntRoman(100, "C"),
	newIntRoman(90, "XC"),
	newIntRoman(50, "L"),
	newIntRoman(40, "XL"),
	newIntRoman(10, "X"),
	newIntRoman(9, "IX"),
	newIntRoman(5, "V"),
	newIntRoman(4, "IV"),
	newIntRoman(1, "I"),
}

func intToRoman(num int) string {
    var res string
    for _, itr := range list {
        m := num / itr.integer
        if m >= 1 {
            r := int(m)
            num -= r * itr.integer
            res += strings.Repeat(itr.numeral, r)
        }
    }
    return res    
}