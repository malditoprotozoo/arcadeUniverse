package main

/*

Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

*/

func adjacentElementsProduct(inputArray[]int) int {
  var products []int
  for i := range inputArray {
    if i + 1 < len(inputArray) {
      product := inputArray[i] * inputArray[i+1]
      products = append(products, product)
    }
  }
	largest := products[0]
  for i := range products {
    if products[i] > largest {
      largest = products[i]
    }
  }
  return largest
}
