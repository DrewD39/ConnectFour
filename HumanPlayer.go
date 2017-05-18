/* HumanPlayer.go contains a substuct of player which allows people to play
/
/  Author: Drew Davis
*/


package main
import "fmt"

/* HumanPlayer struct is a substruct of player that allows a person to play
*/
type HumanPlayer struct {
  Player
}

/* newHumanPlayer acts as the HumanPlayer constructor
/  it sets the name and playerNum based on the call
*/
func newHumanPlayer(aName string, aPlayerNum int) HumanPlayer {
  p := HumanPlayer{Player{
    name: aName,
    playerNum: aPlayerNum,
  }}
  return p
}


/* let user choose a move and return the int of the column to drop in
/  Input: a board to be evaluated
/  Output: an int representing which column to play in
*/
func (p HumanPlayer) chooseMove(b Board) int {
  aMove := -1
  // request user make a move
  print("\nHuman, please select which column to play in (0-6)\n")
  fmt.Scan(&aMove)
  return aMove
}
