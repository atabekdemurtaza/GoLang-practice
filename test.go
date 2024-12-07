package main

//	func add(x float64, y float64) float64 {
//		return math.Sqrt(float64(x*x + y*y))
//	}
//
//	func main() {
//		fmt.Println(add(1, 2))
//	}

//func swap(x, y string) (string, string) {
//	return y, x
//}
//
//func main() {
//	a, b := swap("hello", "world")
//	fmt.Println(a, b)
//}

//import "fmt"
//
//func split(sum int) (x, y int) {
//	x = sum * 4 / 9
//	y = sum - x
//	return
//}
//
//func main() {
//	fmt.Println(split(17))
//}

//import "fmt"
//
//var c, python, java bool
//
//func main() {
//	var i int
//	fmt.Println(i, c, python, java)
//}

//import "fmt"
//
//var i, j int = 1, 2
//
//func main() {
//	var c, python, java = true, false, "no!"
//	k := 3
//	fmt.Println(i, j, c, python, java, k)
//}

//import "fmt"
//
//func main() {
//	var i int
//	var j float32
//	var b bool
//	var s string
//	fmt.Printf("%v %v %v %q\n", i, j, b, s)
//}

//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	var x, y int = 3, 4
//	var f float64 = math.Sqrt(float64(x*x + y*y))
//	var z uint = uint(f)
//	fmt.Println(x, y, z, f)
//
//	v := 42 // change me!
//	fmt.Printf("v is of type %T\n", v)
//}

//func main() {
//	sum := 1
//	for sum < 1000 {
//		sum += sum
//	}
//	fmt.Println(sum)
//}

//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x) + "i"
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}
//
//func main() {
//	fmt.Println(sqrt(2))
//}

//import (
//	"fmt"
//	"math"
//)
//
//func pow(x, n, lim float64) float64 {
//	if v := math.Pow(x, n); v < lim {
//		return v
//	} else {
//		fmt.Printf("%g >= %g\n", v, lim)
//	}
//	// cant use v here though
//	return lim
//}
//func main() {
//	fmt.Println(pow(3, 2, 10))
//	fmt.Println(pow(3, 3, 10))
//}

import (
	'fmt'
	'math'
)

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}
