package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices class")

	var scores = []int{1, 2, 3}
	fmt.Println("Slice = ", scores)

	// adding new value to slices
	scores = append(scores, 34, 56, 78)
	fmt.Println("Slice = ", scores)

	// using make()
	// here 2 is default size but we increase it by using append
	highscore := make([]int, 2)
	highscore[0] = 167
	highscore[1] = 100
	fmt.Println("Highscore = ", highscore)

	highscore = append(highscore, 40, 20, 400)
	fmt.Println("New Highscores = ", highscore)

	fmt.Println(sort.IntsAreSorted(highscore))
	sort.Ints(highscore)
	fmt.Println("New Sorted Highscores = ", highscore)


	// slicing the slice
	// it means from index 1 to 3 , if blank means from zero or to end
	highscore = append(highscore[1:3])
	fmt.Println(highscore)

	var courses = []int{1, 2, 3, 4, 5}
	index := 2
	courses = append(courses[:index], courses[index + 1: ]... )
	fmt.Println("Courses = ", courses)


}
