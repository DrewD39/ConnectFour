/* board.go contains the board stuct on which the game will be played.
/  The same board will be used, regardless of who is playing.
/  Author: Drew Davis
*/
package main

 /* the board struct holds the data and operations of a connect4 game board
 /  it handles features and info that are shared by all players
 */
type Board struct {
  board [][]int //necassary? //board [][]int // numeric matrix that holds current board chips
  // board [][]string // string matrix that holds the visual representation of the current board
  whosMove int// an integer representing if it is player 1's or 2's move
  victory bool // a boolean that denotes a victory state with a 1
}

/* newBoard serves as the constructor for Board struct
/ sets the board matrix to be empty, whosMove to 1, and victory to false
*/
func newBoard() Board {
  var emptyMat [][]int = [][]int{
    []int{0,0,0,0,0,0,0},
    []int{0,0,0,0,0,0,0},
    []int{0,0,0,0,0,0,0},
    []int{0,0,0,0,0,0,0},
    []int{0,0,0,0,0,0,0},
    []int{0,0,0,0,0,0,0},
  }
  b := Board{
    board: emptyMat,
    whosMove: 1,
    victory: false,
  }
  return b
}

/* copy makes a deep copy of the entire Board struct and returns the copy
/ Input: N/A
/ Output: N/A
*/
func (b Board) copy() Board {
  newBoardMat := make([][]int, 6)
  for i:= range b.board {
    newBoardMat[i] = make([]int, len(b.board[i]))
    copy(newBoardMat[i], b.board[i])
  }
  newBoard := Board{
    board: newBoardMat,
    whosMove: b.whosMove,
    victory: b.victory,
  }
  return newBoard
}

/* chipAt checks to see which, if any, player has a chip at a specified location
/ Input: 2 ints that represent a position in the matrix of the board
/ Output: int representing which player has a chip at the location
*/
func (b *Board) chipAt(i int, j int) int {
  return b.board[i][j] // row, column ?
}

/* columnHeight returns the current number of chips at the bottom of a column
/ Input: int of column to check
/ Output: int of number of chips in a column
*/
func (b *Board) columnHeight(j int) int {
  for p := 0; p < 6; p++ {
    if (b.board[p][j] == 0) {return p} // return the lowest open spot
  }
  return 6 // if the column is full, return 6
}

/* legalMove checks to see if a certain move is legal.
/ A move is a single column selction where the chip will be dropped
/ Input: int of column to check move in
/ Output: bool telling if move is legit
*/
func (b *Board) legalMove(j int) bool {
  if (j < 0 || 6 < j) {return false} // if unreal column, return false
  if (b.columnHeight(j) == 6) {return false} // if column full, illegal move
  return true // if not illegal, return true
  }

/* gameOver checks for victory of if there aren't any moves left
/ Input: N/A
/ Output: N/A
*/
 func (b *Board) gameOver() bool {
   lastTurn := 0
   if (b.whosMove == 1) {lastTurn = 2
   } else {lastTurn = 1}
   // if the last player to play has 4 in a row, set victory
   if (b.checkFor(lastTurn, 4) > 0) {
     //print(int(math.Abs(float64(b.whosMove-1)))); print("\n")
     //print(b.checkFor(int(math.Abs(float64(b.whosMove-1))), 4))
     b.victory = true
     return true
   }

   for j := 0; j <= 6; j++ {
     if (b.columnHeight(j) < 6) { // if it is possible to play in this column
       return false
     }
   }
   return true // no legit moves left
 }

/* swapTurn switches who's turn it is
/ Can only be player 1 or player 2's turn
/ Input: N/A
/ Output: N/A
*/
func (b *Board) swapTurn() {
  if b.whosMove == 1 {
    b.whosMove = 2
  } else {
    b.whosMove = 1 // switch turns
  }
}

/* makeMove performs a single column drop for a player.
/ After performing move, player turn is switched
/ Input: column to attempt move in
/ Output: N/A
*/
func (b *Board) makeMove(j int) {
  if b.legalMove(j) {
    b.board[b.columnHeight(j)][j] = b.whosMove // drop a chip at column j
    b.swapTurn()
  } else {
    //fmt.Printf("\nERROR: Illegal move attempted (%d)! Turn not switched.",j)
  }
}

/* display prints the connect4 board to the console
/ Input: N/A
/ Output: N/A
*/
func (b *Board) display() {
  print("\n")
  for i := 5; i >= 0; i-- { // for each row
    for j := 0; j < 7; j++ { // for each column
      print(" | "); print(b.board[i][j])
    }
    print(" |\n")
  }
}

/* checkFor surveys a board for specified conditions in rows, columns, and diagonals
/ Input: Player number and # in a line to check for
/ output: number of full conditions found
*/
func (b *Board) checkFor(playerNum int, numInLine int) int {
  matches := 0
  matches += b.checkStraight(playerNum, numInLine, "rows")
  matches += b.checkStraight(playerNum, numInLine, "cols")
  matches += b.checkDiag(playerNum, numInLine)
  return matches
}

/* checkStraight surveys a board for specified conditions in rows/columns
/ Input: Player number, # in a line to check for, direction to check
/ output: number of full conditions found
*/
func (b *Board) checkStraight(playerNum int, numInLine int, direction string) int {
  streak := 0
  matches := 0
  outLoopLim := 0
  inLoopLim := 0
  chipHere := 0
  if direction == "rows" {
    outLoopLim = 6 // outer loop is rows
    inLoopLim = 7 // inner loop is columns
  } else if direction == "cols" {
    outLoopLim = 7 // outer loop is columns
    inLoopLim = 6 // inner loop is rows
  } else {print("Error! Invalid search direction"); return -1}
  for i := 0; i < outLoopLim; i++ {
    for j := 0; j < inLoopLim; j++ {
      if(direction == "rows") {
        chipHere = b.chipAt(i,j)
      } else {chipHere = b.chipAt(j,i)}
      if(chipHere == playerNum) { // if chip matches player
        streak++                       // increase streak size by 1
        } else {streak = 0}            // else reset streak
      if(streak == numInLine) {
        streak = 0
        matches++
      }
    }
    streak = 0 // move on to new row, reset streak count
  }
  return matches
}


/* checkDiag surveys a board for specified conditions in diagonals
/ Input: Player number, # in a line to check for
/ output: number of full conditions found
*/
func (b *Board) checkDiag(playerNum int, numInLine int) int {
  streak := 0
  matches := 0
  col := 0
  row := 0
  for j := -5; j < 7; j++ { // for each column

    // check downleft-to-upright diagonals
    col = j; row = 0
    for {
      if( (0 <= col && col < 7) && (0 <= row && row < 6) ) { // if valid coordinates
        if( b.chipAt(row,col) == playerNum) { // if sought chip is here
          streak++ // add 1 to streak
          if(streak == numInLine) { // if full streak met
            matches++
            streak = 0
          }
        } else { // if valid coordinate, but not desired chip
          streak = 0 // reset streak
        }
      } else if (col >= 7 || row >= 6) { // if reached end of diagonal
        streak = 0 // reset streak
        break
      } // if not end of diag, but not yet legit coordinates, do nothing
      col++
      row++
    }

    // check upleft-to-downright diagonals
    col = j; row = 5
    for {
      if( (0 <= col && col < 7) && (0 <= row && row < 6) ) { // if valid coordinates
        if( b.chipAt(row,col) == playerNum) { // if sought chip is here
          streak++ // add 1 to streak
          if(streak == numInLine) { // if full streak met
            matches++
            streak = 0
          }
        } else { // if valid coordinate, but not desired chip
          streak = 0 // reset streak
        }
      } else if (col >= 7 || row >= 6) { // if reached end of diagonal
        streak = 0 // reset streak
        break
      } // if not end of diag, but not yet legit coordinates, do nothing
      col++
      row--
    }

  } // all columns (-5 through 6) checked
  return matches
}
