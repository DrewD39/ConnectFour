/* connect4Match.go contains a the Connect4Match struct
/  which holds the data of a single game of connect4, including two
/  players and a board
/
/  Author: Drew Davis
*/

package main

import "strconv"
//import "fmt"
//import "time"

/* the Connect4Match struct holds the data and operations of a connect4 match
/  it handles the players and the game board
*/
type Connect4Match struct {
  move int
  moveSequence *string // a string that tracks all moves made in the game
}

/* newMatch serves as the constructor for Connect4Match
/  it sets the timeLimit based on the caller
*/
func newMatch(moveSeqPtr *string) Connect4Match {
  c := Connect4Match{
    move: 0,
    moveSequence: moveSeqPtr,
  }
  return c
}

/* playGame simulates a game between two players (AI or human)
/  the game starts with player 1 and alternates turns until a player
/  wins or no more moves can be made
/  Input: 2 player interfaces, which will play each other
/  Output: N/A
*/
func (match *Connect4Match) playGame(p1 AnyPlayer, p2 LearnedPlayer) {
  // create 2 board objects to protect the official board from player classes
  officialBoard := newBoard()
  scrapBoard := officialBoard.copy()
  //officialBoard.display()

  // while there are still moves left, and no one has won
  for (officialBoard.gameOver() == false) {
    scrapBoard = officialBoard.copy()
    //print("\nPlayer "); print(officialBoard.whosMove); print("'s turn")

    // allow current player to make a move, based on their chooseMove function
    if (officialBoard.whosMove == 1) {match.move = p1.chooseMove(scrapBoard)
    } else {match.move = p2.chooseMove(scrapBoard)}
    //print("\nPlayer "); print(officialBoard.whosMove); print(" plays in column "); print(match.move);
    *match.moveSequence = *match.moveSequence + " " + strconv.Itoa(match.move)
    //fmt.Printf("\nOfficial move sequence: %s",*match.moveSequence)

    //update board and present it in console
    officialBoard.makeMove(match.move)

    //officialBoard.display()
    //time.Sleep(2 * time.Second)
    //time.Sleep(0)
    //print("\n\n")
  }
//fmt.Printf("\nFinal match sequence: %s",*match.moveSequence)
  // once the game is over, determine who (if any) the winner was and display it
  if (officialBoard.victory == true && officialBoard.whosMove == 2) {
    // player 1 won
    print("\nPlayer 1 wins!\n")
    p2.updateLossData()
  }
  if (officialBoard.victory == true && officialBoard.whosMove == 1) {
    // player 2 won
    print("\nPlayer 2 wins!\n")
  }
  if (officialBoard.victory == false) {
    // no more avilable moves, but no winner
    print("\nNobody wins :(\n")
    // a draw isn't good enough, learn to improve from it as well
    p2.updateLossData()
  }


}
