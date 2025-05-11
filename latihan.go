package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, j, compalier, idx, c int
	var b bool 
	b = false
	fmt.Scan(&n)
	
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i =0; i < n; i++{
		for j = i+1; j < n; j++{
			if A[i] > A[j]{
				idx = j
				b = true
			}
		}
		if b == true{
			A[idx] = compalier
			A[i] = A[idx]
			A[idx] = compalier
			c++
			b = false
		}
	}
	fmt.Print(c)
}