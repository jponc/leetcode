# 2020-04-26
## [Palindrome Number](https://leetcode.com/problems/palindrome-number/) [easy]
```
str = string value of number
i = 0
j = len(str) - 1

iterate until they don't overlap:
  if str[i] != str[j], return false
  i++; j--
return true
```

## [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) [easy]
```
Set possible characters to hashmap
initialise output to 0

iterate all characters from roman string
  check value against hashmap and add it to accumulator
```

## [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) [easy]
```
Create an index pointer, then iterate and check if they have the same character using a set
```

# 2020-04-25
## [Two Sum](https://leetcode.com/problems/two-sum/) [easy]
```
Find the complementary number when iterating the array.
Store the complementary number to hashmap if it's not there. with k: complementary number, v: index
Else: Return [complementaryIdx, currentIdx]
```

## [Reverse Integer](https://leetcode.com/problems/reverse-integer/) [easy]
```
```


