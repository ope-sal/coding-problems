package main

import "fmt"

// given an int array, this function returns a corresponding array in which each 
// element is a product of all the the numbers in the original array except
// the number of the corresponding index.
// e.g [1,2,3,4,5] returns [120,60,40,30,24] and
// [3,2,1] returns [2,3,6] --> [2*1,3*1,3*2]

// very naive
func product(s []int) []int{
	var result []int
	for i, _ := range s{
		prod := 1
		for j, num := range s{
			if i != j{
				prod *= num
			}
		}
		result = append(result, prod)
	}
	return result
}

// optimized using division
func productDiv(s []int) []int{
	var result []int 
	overall_prod := 1
	for  _, v := range s{
		overall_prod *= v 
	}
	for _, v := range s{
		result = append(result, overall_prod/v)
	}
	return result
}

// without division - you are not allowed to use divsion
func productNoDiv (s []int) []int{
	pre_prods, post_prods := []int{1}, []int{1}
	var result []int

	for i := 0; i < len(s)-1; i++ {
		pre_prods = append(pre_prods, pre_prods[i]*s[i])
		post_prods = append(post_prods, post_prods[i] * s[len(s) - i - 1])
	}

	for i := 0; i < len(s); i++ {
		result = append(result, pre_prods[i] * post_prods[len(s) - i -1] )
	}

	return result


}

func main() {
	fmt.Println("helloWorld")
}
