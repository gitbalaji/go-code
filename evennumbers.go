package main

import "fmt"

func main(){
	
	numbers :=[6] int {10,20,9,4,2,11}
	nums := [] int {}

	for j:=0; j< 6; j++{
		if numbers[j] % 2==0 {
			 nums = append(nums,numbers[j])
		}
	}
	
	fmt.Println(numbers)
	fmt.Println(nums)
	test(&nums)
	
}

func test(nums *[]int){
	fmt.Println(nums);
}