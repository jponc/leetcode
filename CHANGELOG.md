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

