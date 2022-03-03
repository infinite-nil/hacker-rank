# Plus minus

|Problem|Submission|
|-|-|
|[URL](https://www.hackerrank.com/challenges/one-month-preparation-kit-plus-minus/problem)|[URL](https://www.hackerrank.com/challenges/one-month-preparation-kit-plus-minus/submissions/code/259091380)|


Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with places after the decimal.

Note: This challenge introduces precision problems. The test cases are scaled to six decimal places, though answers with absolute error of up to 10^4 are acceptable.

**Example**:

`arr = [1,1,0,-1,-1]`

There are elements, two positive, two negative and one zero. Their ratios are `2/5 = 0.400000`, `2/5 = 0.400000` and `1/5 = 0.200000`. Results are printed as:

| ratios   |
| -------- |
| 0.400000 |
| 0.400000 |
| 0.200000 |

**Constraints**

```
0 < N ≤ 100
-100 ≤ arr[i] ≤ 100
```

**Output Format**

Print the following lines, each to decimals:

proportion of positive values
proportion of negative values
proportion of zeros

**Sample Input**

```
STDIN           Function
-----           --------
6               arr[] size n = 6
-4 3 -9 0 4 1   arr = [-4, 3, -9, 0, 4, 1]
```

**Sample Output**

```
0.500000
0.333333
0.166667
```
