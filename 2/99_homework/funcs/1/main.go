package main

import ("fmt"; "strconv")

type memoizeFunction func(int, ...int) interface{}

var num = map[string]int{
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

var numInv = map[int]string{
    1000: "M",
    900:  "CM",
    500:  "D",
    400:  "CD",
    100:  "C",
    90:   "XC",
    50:   "L",
    40:   "XL",
    10:   "X",
    9:    "IX",
    5:    "V",
    4:    "IV",
    1:    "I",
}

var maxTable = []int{
    1000,
    900,
    500,
    400,
    100,
    90,
    50,
    40,
    10,
    9,
    5,
    4,
    1,
}

var fibonacci memoizeFunction = func (n int, args ...int) interface{} {
	if n < 2 {
		return 1
	}
	curr := 1
	prev := 1

	for i := 0; i < n - 2; i++ {
		tmp := curr
		curr = prev + curr
		prev = tmp
	}

	return curr;
}

var romanForDecimal memoizeFunction = func (n int, args ...int) interface{} {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

func memoize(function memoizeFunction) memoizeFunction {
    var cache = make(map[string]interface{})

    return func(n int, args ...int) interface{} {
        stringifiedArgs := stringifyArgs(n, args...)

        if (cache[stringifiedArgs] != nil) {
            return cache[stringifiedArgs]
        }

        cache[stringifiedArgs] = function(n, args...)
        return cache[stringifiedArgs]
    }
}

func stringifyArgs(n int, args ...int) string {
    res := strconv.Itoa(n)

    for i := 0; i < len(args); i++ {
        res += strconv.Itoa(args[i])
    }

    return res
}

func initialize() {
    fibonacci = memoize(fibonacci)
    romanForDecimal = memoize(romanForDecimal)
}

func main() {
    initialize()
    fmt.Println("Fibonacci(45) =", fibonacci(45).(int))

	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}
