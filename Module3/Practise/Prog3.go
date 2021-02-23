package main

import "fmt"

/**
	i/p : Number of the month
	o/p : Season the month belongs to
 */

func Season(month int) string {
	var ans string
	switch  {
	case month == 1 || month == 2 || month == 12:
		ans = "Winter"
	case month >=3 && month <= 5:
		ans = "Spring"
	case month >=6 && month <= 8:
		ans = "Summer"
	case month >=9 && month <=11:
		ans = "Autumn"
	}
	return ans
}

func main(){
	fmt.Println(Season(1))
	fmt.Println(Season(3))
}
