/*
 * Complete the function that takes two ints(a, b where a < b)
 * and return an array of all integers between the input 
 * parameters, including them.
 *
 * from https://www.codewars.com/kata/55ecd718f46fba02e5000029
 */

package main

import "fmt"

func between(arr []int, a, b int) []int {
    for i:= 0; (a + i) <= b; i++{
        arr = append(arr, a + i)
    }
    return arr
}

func main() {
     arr := []int{}
     arr = between(arr, 1, 10)

     fmt.Println(arr)
}
