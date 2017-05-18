/* Player.go contains the super struct which defines players of all types.
/  It provides the ubiquitous features and supplies the simplest
/  definition for chooseMove
/  Author: Drew Davis
*/

package main
import "math/rand"

/* Player is the superclass of all possible players
/ and is under the AnyPlayers interface
*/
type Player struct {
  name string // a string holding the name of the player
  playerNum int // holds player number (1 or 2)
  timeLim float64 // maximum time a player has per turn
  randSeed int
}

/* newPlayer serves as the Player constructor
/ it sets name, playerNum, and timeLim as specified by its call
/ it also generates a random seed for chooseMove
*/
func newPlayer(aName string, aPlayerNum int, maxTime float64, seed int) Player {
  p := Player{
    name: aName,
    playerNum: aPlayerNum,
    timeLim: maxTime,
  }
  rand.Seed(int64(seed))
  return p
}

/* MoveResult struct acts a tuple, holding a move and its evaluation
*/
type MoveResult struct {
  move int
  val int
}

/* newMoveResult serves as the MoveResult constructor
/ it sets move and val as specified by its call
*/
func newMoveResult(aMove int, aVal int) MoveResult {
  m := MoveResult{
    move: aMove,
    val: aVal,
  }
  return m
}


/* Choose a move and return the int of thee column to drop ins
/ the move is randomly chosen (using provided seed in constructor)
/ this method is overidden by more advanced subclasses
/ Input: a board
/ Output: int representing which column to play in
*/
func (p Player) chooseMove(b Board) int {
  // randomly select a move
  aMove := -1
  for (!b.legalMove(aMove)) {
    aMove = rand.Intn(7) // randomly choose a move between 0 and 6
  }
  return aMove
}
