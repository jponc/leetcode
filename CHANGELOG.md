# 2020-04-28
## [Merge Two Sorted List](https://leetcode.com/problems/merge-two-sorted-lists/submissions/) [easy]

## [Remove Duplicates From Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) [easy]
```
Identify an index pointer where you placed the last unique element.
Then iterate the array as usual and compare current against value in index pointer.
If they're different, increment index pointer, and set value to current
```

## [Remove Elements](https://leetcode.com/problems/remove-element) [easy]

## [Implement strStr](https://leetcode.com/problems/implement-strstr/) [easy]
```
Faster than 100% of Go solution, and 2.3MB memory less than 100% of Go solutions :p

1. Created a dynamic programming solution O(m*n), but takes so much memory
2. In plaece comparison also, O(m*n)
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

## [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) [easy]
```
s = new stack
for char in string:
  if opening:
    stack.push(closing complementary parentheses against char)
  else:
    if stack.pop() != char:
      return false

return stack.isEmpty()
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


