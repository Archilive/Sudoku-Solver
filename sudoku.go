package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if !ValidArgs(args) {
		PrintStr("Error")
	} else {
		board := FormatArgs(args)

		Solve(board)
	}
}

// Controle des arguments
func ValidArgs(args []string) bool {
	// Controle le nombre d'argument
	if len(args) != 9 {
		return false
	}
	// Controle la longueur de chaque argument
	for _, arg := range args {
		if len(arg) != 9 {
			return false
		}
	}
	// Si un sudoku à moins de 17 chiffres de base il n'est pas valable
	count := 0
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] != '.' {
				count += 1
			}
		}
	}
	if count <= 17 {
		return false
	}
	return true
}

// creation du tableau avec les valeurs donné
func FormatArgs(s []string) [9][9]int {
	res := [9][9]int{}
	for y, v := range s {
		for x, val := range v {
			if val == '.' {
				continue
			} else if val >= '0' && val <= '9' {
				res[y][x] = int(val - 48)
				// controle de la valeur
				if !IsValid(res, x, y, res[y][x]) {
					PrintStr("Error")
					os.Exit(0)
				}
			} else {
				PrintStr("Error")
				os.Exit(0)
			}
		}
	}
	return res
}

// function print
func PrintStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

// Trouver la valeur vide à remplir
func GetEmpty(board [9][9]int) (int, int) {
	for y, ligne := range board {
		for x := range ligne {
			if board[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

// Different Check ligne/Colonne/box
func IsValid(board [9][9]int, x int, y int, value int) bool {
	//Check ligne en regardant l'egalité des valeurs
	for xBoard, boardValue := range board[y] {
		if boardValue == value && xBoard != x {
			return false
		}
	}
	//Check colonne en regardant l'egalité des valeurs
	for yBoard, ligne := range board {
		if ligne[x] == value && yBoard != y {
			return false
		}
	}

	//Check Box en regardant l'egalité des valeurs
	xBox := (x / 3) * 3
	yBox := (y / 3) * 3
	for yBoard := yBox; yBoard < yBox+3; yBoard++ {
		for xBoard := xBox; xBoard < xBox+3; xBoard++ {
			if board[yBoard][xBoard] == value && xBoard != x {
				return false
			}
		}
	}
	return true
}

func Solve(board [9][9]int) {
	// valeur vide
	targetX, targetY := GetEmpty(board)
	// Quand le sudoku est soit completé soit impossible
	if targetX == -1 && targetY == -1 {
		if SudokuChecker(board) {
			PrintTableau(board)
		} else {
			PrintStr("Error")
		}
	} else {
		for val := 1; val <= 9; val++ { // Quand il est pas completé on essaye les valeurs de 1 à 9
			if IsValid(board, targetX, targetY, val) {
				board[targetY][targetX] = val
				Solve(board)
			}
		}
	}
}

// function printer le tableau
func PrintTableau(board [9][9]int) {
	for _, ligne := range board {
		for _, value := range ligne {
			PrintStr(string(rune(value + 48)))
			PrintStr(" ")
		}
		z01.PrintRune('\n')
	}
	os.Exit(0)
}

// On regarde si le board finale est bon avant de l'imprimer en faisant un check pour toute les valeurs
func SudokuChecker(board [9][9]int) bool {
	for y, ligne := range board {
		for x := range ligne {
			if !IsValid(board, x, y, board[y][x]) {
				return false
			}
		}
	}
	return true
}
