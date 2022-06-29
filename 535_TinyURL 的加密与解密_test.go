package Code

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func Test535(t *testing.T) {
	client := Constructor535()
	tmp := client.encode("http://www.baidu.com")
	t.Log(tmp)
	fina := client.decode(tmp)
	t.Log(fina)

	t.Log(Encode(349051260))
	t.Log(Decode("2bACn"))

}

const k1, k2 = 1117, 1e9 + 7

type Codec535 struct {
	dataBase map[int]string
	urlToKey map[string]int
}

func Constructor535() Codec535 {
	return Codec535{map[int]string{}, map[string]int{}}

}

// Encodes a URL to a shortened URL.
func (c *Codec535) encode(longUrl string) string {
	if key, ok := c.urlToKey[longUrl]; ok {
		return "http://tinyurl.com/" + strconv.Itoa(key)
	}
	key, base := 0, 1
	for _, c := range longUrl {
		key = (key + int(c)*base) % k2
		base = (base * k1) % k2
	}
	for c.dataBase[key] != "" {
		key = (key + 1) % k2
	}
	c.dataBase[key] = longUrl
	c.urlToKey[longUrl] = key
	return "http://tinyurl.com/" + strconv.Itoa(key)

}

// Decodes a shortened URL to its original URL.
func (c *Codec535) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	key, _ := strconv.Atoi(shortUrl[i+1:])
	return c.dataBase[key]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

const CODE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const CODE_LENTH = 62

var EDOC = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "a": 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15, "g": 16, "h": 17, "i": 18, "j": 19, "k": 20, "l": 21, "m": 22, "n": 23, "o": 24, "p": 25, "q": 26, "r": 27, "s": 28, "t": 29, "u": 30, "v": 31, "w": 32, "x": 33, "y": 34, "z": 35, "A": 36, "B": 37, "C": 38, "D": 39, "E": 40, "F": 41, "G": 42, "H": 43, "I": 44, "J": 45, "K": 46, "L": 47, "M": 48, "N": 49, "O": 50, "P": 51, "Q": 52, "R": 53, "S": 54, "T": 55, "U": 56, "V": 57, "W": 58, "X": 59, "Y": 60, "Z": 61}

/**
 * 编码 整数 为 base62 字符串
 */
func Encode(number int) string {
	if number == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for number > 0 {
		round := number / CODE_LENTH
		remain := number % CODE_LENTH
		result = append(result, CODE62[remain])
		number = round
	}
	return string(result)
}

/**
 * 解码字符串为整数
 */
func Decode(str string) int {
	str = strings.TrimSpace(str)
	var result int = 0
	for index, char := range []byte(str) {
		result += EDOC[string(char)] * int(math.Pow(CODE_LENTH, float64(index)))
	}
	return result
}
