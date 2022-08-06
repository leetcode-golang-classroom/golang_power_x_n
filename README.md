# golang_power_x_n

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).

## Examples

**Example 1:**

```
Input: x = 2.00000, n = 10
Output: 1024.00000

```

**Example 2:**

```
Input: x = 2.10000, n = 3
Output: 9.26100

```

**Example 3:**

```
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25

```

**Constraints:**

- `100.0 < x < 100.0`
- `-$2^{31}$ <= n <= $2^{31}$-1`
- $`-10^4$ <= $x^n$ <= $10^4$`

## 解析

實作 pow(x, n) 其中  -100 < x < 100,  `-$2^{31}$ <= n <= $2^{31}$-1`

從數學定義基本上可以透過 O(n) 對 x 做連乘或是連除在 n < 0 時

然而 這題希望是更優化的方式

因為可以發現 假設是連乘

$2^{10}$ = 只要先算好 $2^5$ 然後再把兩個 $2^5$ 相乘即可

這樣的話就可以把 O(n) 優化成 O(logn)

定義 recursive(x, n) 其中 n ≥ 0

if x = 0 時 recursive = 0

if x = 1 或是 n= 0 時 recursive = 1

res = recursive(x, n/2)

res = res * res 

if n % 2 == 1 時, recursive = res * x

否則 recursive = res

## 程式碼
```go
package sol

func myPow(x float64, n int) float64 {
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	var Recursive func(x float64, n int) float64
	Recursive = func(x float64, n int) float64 {
		n = abs(n)
		if x == 0 {
			return 0
		}
		if x == 1 || n == 0 {
			return 1
		}
		res := Recursive(x, n/2)
		res = res * res
		if n%2 == 1 {
			res *= x
		}
		return res
	}
	result := Recursive(x, n)
	if n < 0 {
		return 1 / result
	}
	return result
}

```
## 困難點

1. 要知道 divide and conquer 的作法來做優化

## Solve Point

- [x]  定義一個 recursive method 其輸入值是 x float64, n int
- [x]  當 x = 0  recursive = 0
- [x]  當 x = 1 或是 n = 0 時 , recursive = 1
- [x]  計算 res = recursive(x, n/2) , res = res * res
- [x]  if n % 2 == 1 , res *= x
- [x]  回傳 res
- [x]  if  n ≥ 0 回傳 res  否則回傳 1/res