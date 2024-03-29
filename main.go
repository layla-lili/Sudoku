package main

import (
	"fmt"
	"os"
)

/*
declare a 2 dimensional array, that will represent our 9*9 sudoku board.
initially all cells will contain zeros.
*/
var sudoku = [9][9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

/*
The draw func will print all rows of the sudoku board.
*/
func draw() {
	//loop through each row
	for _, row := range sudoku {
		//loop through each cell of a row
		for _, cell := range row {
			//print cell value
			fmt.Print(cell)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

/*
The populateSliceFromUserInput func will take the arguments passed to the program,
validate them and populate the sudoku array with the unsolved puzzle.

***
validation conditions:
- 9 args are passed
- each arg contains 9 c
- only numbers are passed
- empty cells are represented by '.'

***
in case validation fails program will panic and exit
*/
func populateSliceFromUserInput() bool {
	//store the passed arguments in an slice.
	args := os.Args[1:]
	//validate args length to be 9
	if len(args) != 9 {
		// panic("Error: Missing arguments")
		return false
	}
	//loop through args and validate each c. case validation false panic, else populate position and continue.
	for i, arg := range args {
		//check arg to be length of 9, else panic
		if len(arg) != 9 {
			// panic("Error: Argument missing cells")
			return false
		}
		//loop through each arg c
		for j, c := range arg {
			//c == ., continue as this means cell is empty
			if c == '.' {
				continue
			} else if c > '0' && c <= '9' {
				//c is > 0, <= 9, a valid input. alter at position
				//convert to int
				nb := int(c - '0')
				//add to sudoku board
				sudoku[i][j] = nb
			} else {
				//any other input will panic
				// panic("Error: Invalid input")
				return false
			}
		}
	}
	//after population done, validate sudoku board
	if !validateRows() || !validateColumns() || !validateSquares() {
		return false
	}
	return true
}

/*
The validateRows func will loop through each row and each cell of that row to validate that a number is not occured more than once.
*/
func validateRows() bool {
	//loop through rows
	for i := 0; i < 9; i++ {
		//create a map for each row to keep track of numbers that already occured in a row cells
		row := make(map[int]bool)
		//loop through row cells
		for j := 0; j < 9; j++ {
			//get the cell value
			val := sudoku[i][j]
			//if 0, empty cell then skip
			if val == 0 {
				continue
			}
			//if val is already occured
			//panic and exit
			if _, ok := row[val]; ok {
				// panic("A value has occured more than once in a row")
				return false
			}
			//else set row of key value to true
			row[val] = true
		}
	}
	//return true if validated
	return true
}

/*
The validateColumns func will loop through each column and each cell of that column to validate that a number is not occured more than once.
*/
func validateColumns() bool {
	//loop through columns
	for j := 0; j < 9; j++ {
		//create a map for each column to keep track of numbers that already occured in a column cells
		col := make(map[int]bool)
		//loop through column cells
		for i := 0; i < 9; i++ {
			//get the cell value
			val := sudoku[i][j]
			//if 0, empty cell then skip
			if val == 0 {
				continue
			}
			//if val is already occured
			//panic and exit
			if _, ok := col[val]; ok {
				// panic("A value has occured more than once in a column")
				return false
			}
			//else set col of key value to true
			col[val] = true
		}
	}
	//return true if validated
	return true
}

/*
The validateSquares func will loop through each 3*3 square to validate that a number is not occured more than once in a square.
*/
func validateSquares() bool {
	for x := 0; x < 9; x += 3 {
		for y := 0; y < 9; y += 3 {
			square := make(map[int]bool)
			for i := x; i < x+3; i++ {
				for j := y; j < y+3; j++ {
					val := sudoku[i][j]
					if val == 0 {
						continue
					}
					if _, ok := square[val]; ok {
						return false
					}
					square[val] = true
				}
			}
		}
	}
	return true
}

/*
code starts by populating a slice from user input

	The code then draws the board and checks if there is a solution to the puzzle.
	If there is, it prints "solution found!"
	and stops drawing; otherwise, it prints "no solution found!
*/
func solve() {
	if populateSliceFromUserInput() {
		// draw()
		if sudokuSolver(0, 0) {
			// fmt.Println("solution found!")
			draw()
		} else {
			fmt.Println("no solution found!")
		}
	} else {
		fmt.Println("Error")
	}
}

/* The code will return true if the x and y coordinates are not already in a vertical, horizontal or square position.
 
*/
func canPut(x int, y int, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

/*
The code is trying to check if the value is in a Vertical column.
The code does this by iterating through all of the rows and checking if the value is in that row.
If it is, then it returns true; otherwise, it returns false.
The code checks to see if the value is already in the Vertical position.
The code iterates through a range of integers and checks each one against the value.
*/
func alreadyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if sudoku[i][x] == value {
			return true
		}
	}
	return false
}

/*
The code is trying to check if the value is in a horizontal row.
The code does this by iterating through all of the rows and checking if the value is in that row.
If it is, then it returns true; otherwise, it returns false.
The code checks to see if the value is already in the horizontal position.
The code iterates through a range of integers and checks each one against the value.
*/
func alreadyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if sudoku[y][i] == value {
			return true
		}
	}
	return false
}

// The code is used to check if the value of a given square is already in the range of values for that row and column.
// first parameter x represents the row number,
// while the second parameter  y represents the column number.
func alreadyInSquare(x int, y int, value int) bool {
	//"sx", it's referring to how many spaces over from our starting point on one side (in this case 3).
	//"sy", it's referring to how many spaces over from our starting point on one side (in this case 3).
	sx, sy := int(x/3)*3, int(y/3)*3
	//checking if the x and y coordinates are within 3 squares of each other.
	//also checks if any of the values in the list have been duplicated
	//dy refers to how many times through all possible combinations of numbers do we want to iterate?
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if sudoku[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

// calculates the nextX and nextY values by adding 1 to x until it reaches 9, then dividing y by 9
func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y

	if nextX == 0 {
		//y will be set to 1 instead of being incremented by 1 because there's no point in incrementing when you're already at your maximum value.
		nextY = y + 1
	}
	return nextX, nextY
}

// The code is a function that takes two integers as parameters and returns true if the puzzle can be solved, false otherwise.
func sudokuSolver(x int, y int) bool {
	if y == 9 {
		return true
	}
	if sudoku[y][x] != 0 {
		return sudokuSolver(next(x, y))
	} else {
		for i := range [9]int{} {
			var v = i + 1
			if canPut(x, y, v) {
				sudoku[y][x] = v
				if sudokuSolver(next(x, y)) {
					return true
				}
				sudoku[y][x] = 0
			}
		}
		return false
	}

}

func main() {
	solve()
}
