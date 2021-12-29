/*

Given two strings s and t of lengths m and n respectively, return the minimum window substring of s
such that every character in t (including duplicates) is included in the window.

If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.

*/
/*
//edc  edc
Example 1:

Input: s = "ADOBECODEBANC", t = "CDEE"
Output: "ECODE"
Explanation: The result "ECODE" includes 'C', 'D', and 2 'E's from string t.

//1，遍历s,记录有包含在t的，比如E; 查看个数>=2,形成一个集合（字串），在集合看里面的元素是不是有包含t里面的所有字符 最坏：n的n次
//2，申请一个数组，比如tmp [255]int,每个字符对应tmp里面的一个数据。map[int]int,遍历到最后，记录次数。
//3, 标记一个头指针和尾指针，从左到右去遍历字符串s,当匹配到有t中的字符且前面并做记录到 type x里面（类型待定），满足就停止，头指针往右移动？
//头指针往右移动，

Example 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.


Example 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

*/

package main

import "fmt"

func main() {

}
func solution(s, t string) (k string) {
	head := []byte(s)
	tail := []byte(s)
	cmp := []byte(t)
	lenc := len(cmp)
	mapCmp := make(map[byte]int)
	mapCmpOk := make(map[byte]int)
	lens := len(s)
	for i := 0; i < lenc; i++ {
		mapCmp[cmp[i]] = 1
	}
	fmt.Println(head, tail, lens)
	log := 0

	for i := 0; i < lens; i++ {
		if 1 == mapCmp[head[i]] {
			log++
		}
	}

	return k
}
