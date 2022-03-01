package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I     N
A   L S   I G
Y A   H R
P     I
*/
func Test_conver(t *testing.T) {
	Convey("Test_conver", t, func() {
		Convey("PAYPALISHIRING 3", func() {
			So(StringShouldEqual(convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR"), ShouldBeTrue)
		})
		Convey("PAYPALISHIRING 4", func() {
			So(StringShouldEqual(convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI"), ShouldBeTrue)
		})
		Convey("A 1", func() {
			So(StringShouldEqual(convert("A", 1), "A"), ShouldBeTrue)
		})
		Convey("AB 1", func() {
			So(StringShouldEqual(convert("AB", 1), "AB"), ShouldBeTrue)
		})
	})
}
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	zMax := 2*numRows - 2
	zMap := make([]string, numRows)
	for i, v := range s {
		//vertical
		if i%zMax >= 0 && i%zMax < numRows {
			zMap[i%zMax%numRows] = zMap[i%zMax%numRows] + string(v)
		} else { //Slope
			zMap[numRows-((i%zMax)+1)%numRows-1] = zMap[numRows-((i%zMax)+1)%numRows-1] + string(v)
		}
	}
	ret := ""
	for _, vLine := range zMap {
		ret += vLine
	}
	return ret
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	list := make([]string, numRows)
	i := 0
	flag := -1
	for _, v := range s {
		list[i] += string(v)
		if i == 0 || i == numRows-1 {
			flag *= -1
		}
		i += flag
	}
	res := ""
	for _, v := range list {
		res += v
	}
	return res
}
