package main

import "rest_api/controller"

func main() {
	controller.RegisterControllers()
	// originalSlice := []string{"test1", "test2", "test3", "test4", "test5", "test6"}
	// sliceToInsert := []string{"slice1", "slice2"}

	// //sliceOne := originalSlice[0:2]
	// //sliceTwo := originalSlice[2:len(originalSlice)]
	// firstList := make([]string, 2)
	// secondList := make([]string, 4)
	// copy(firstList, originalSlice[0:2])
	// copy(secondList, originalSlice[2:len(originalSlice)])

	// joinedArray := append(firstList, sliceToInsert...)
	// joinedArray = append(joinedArray, secondList...)

	// fmt.Println(firstList)
	// fmt.Println(secondList)
	// fmt.Println(joinedArray)
	// //joinedSlice := append(tmp, sliceTwo...)

	// //	fmt.Println(tmp)
	// //fmt.Println(sliceOne)
	// //	fmt.Println(sliceTwo)
	// //fmt.Println(joinedSlice)

}
