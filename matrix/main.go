package main

import "fmt"

func main(){
	var oneD = [5]int{10,20,30,40,50}

	fmt.Println("---ONE DIMENSION MATRIX---")
	for i := 0; i<len(oneD);i++{
		fmt.Println(oneD[i])
	}
	for index, value := range oneD{
		fmt.Printf("Index is %d and value is %d\n",index, value)
	}
	
	var threeD = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println("---THREE DIMENSIONAL MATRIX---")
	for dimension:=0; dimension< 3;dimension++{
		for element:=0;element<3;element++{
			fmt.Printf("%d, ",threeD[dimension][element])
		}
		fmt.Printf("\n")
	} 
}