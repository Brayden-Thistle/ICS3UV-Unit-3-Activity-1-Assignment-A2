/*
* @author Brayden Thistle
* @version 1.0.0
* @date 2025-11-14
* @fileoverfiew this program displays the area of a circle
*/

package main

// creating "math" along with "fmt" to calculate the cube root of 1000 to get your answer
import (
"fmt"
"math"
)


func main () {
	//calculation and answer
	fmt.Println("The length and width and height of a cube with a volume of 1000 mm³ is", (math.Cbrt(1000)), "mm")

	fmt.Println("\nDone")
}