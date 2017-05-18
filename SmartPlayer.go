/* Player.go contains a substuct of player that can play intelligently
/
/  Author: Drew Davis
*/


package main
//import "fmt"
import "time"
import "math/rand"

type SmartPlayer struct {
  Player
}

/* newSmartPlayer acts as the SmartPlayer constructor
/  name, playerNum, and timeLim are set as specified by the caller
/  a random seed is generated to allow the results of AI vs AI games
/  to be varied
*/
func newSmartPlayer(aName string, aPlayerNum int, maxTime float64, seed int) SmartPlayer {
  p := SmartPlayer{Player{
    name: aName,
    playerNum: aPlayerNum,
    timeLim: maxTime,
  }}
  rand.Seed(int64(seed))
  return p
}


/* evaluate a board situation and return a numerical score
/  higher score == better situation for player 1
/  lower score  == better situation for player 2
/  this evaluate function will can overwritten by specific players
/  Input: a board to be evaluated
/  Output: an int score representing quality of situation
*/
func (p *SmartPlayer) evaluate(b Board) int {
  value1 := 0 // player 1's board value
  value1 += ( b.checkFor(1,1) * 1      )
  value1 += ( b.checkFor(1,2) * 4     )
  value1 += ( b.checkFor(1,3) * 16    )
  value1 += ( b.checkFor(1,4) * 10000 )
  value2 := 0 // player 2's board value
  value2 += ( b.checkFor(2,1) * 1      )
  value2 += ( b.checkFor(2,2) * 4     )
  value2 += ( b.checkFor(2,3) * 16    )
  value2 += ( b.checkFor(2,4) * 10000 )

  if b.whosMove == 2 {
    value1 += 12
  } else {
    value2 += 12
  }

  return value1 - value2
}

/* chooseMove selects a move and returns the int of thee column to drop in
/  the choice is made by using a DFS and evaluating the resulting board
/  it will continue to perform increasingly deep DFS searches until time
/  runs out
/  Input: a board to make choice on
/  Output: int representing which column to play in
*/
func (p SmartPlayer) chooseMove(b Board) int {
  t := time.Now()
  time.Since(t)
  depth := 1
  best := newMoveResult(6,0)
  /// perform deeper DFS until time is gone
  /// if part-way through a DFS when time runs out, finish the DFS first
  for ( (time.Since(t)).Seconds() < float64(p.timeLim) ) {
    //fmt.Printf("\n\nBegin new DFS at with depth: %d", depth)
    best = p.DFS(b, depth) // DFS at depth
    depth++
    if (time.Since(t)).Seconds() > float64(p.timeLim) { break }
  }
  //fmt.Printf("\n\nThe final decision is to play at %d (valued at: %d)",best.move,best.val)
  return best.move

}

/* DFS performs the actual depth-first-search recursively to find the optimal play
/  it assumes the opponent will play optimally and chooses the best
/  move to counteract their plays
/  Input: a board to search on, the depth to do the DFS to
/  Output: a MoveResult pair that represents which move a player should make
*/
func (p *SmartPlayer) DFS(b Board, depth int) MoveResult {
  scrapBoard := b.copy()
  randStartNum := rand.Intn(7) // this should help keep things interesting
  columnNum := -1;
  best := newMoveResult(-1,-1)
  if (depth == 0 || scrapBoard.gameOver()) {
    best.val = p.evaluate(scrapBoard) // return value of current board
    //fmt.Printf(" - This sequence is valued at %d (For player 2)",best.val)
    return best
  } else {

    // evaluate for player 1 (high eval is better)
    if scrapBoard.whosMove == 1 {
      best = MoveResult{move: -1, val: -999999}
      play := newMoveResult(-1, -999999)
      for i := 0; i < 7; i++ { // for each possible move (6)
        randStartNum++; columnNum = randStartNum%7
        if scrapBoard.legalMove(columnNum) {
          //fmt.Printf("\nDFS%d: Player %d plays at %d",depth,scrapBoard.whosMove,i)
          scrapBoard.makeMove(columnNum)
          play = p.DFS(scrapBoard, depth-1) // will return results of optimal next move
          if (play.val > best.val) { // if results are better for player 1
            best.move = columnNum
            best.val = play.val
          }
          scrapBoard = b.copy()
        }
      }
    }
    // evaluate for player 2 (low eval is better)
    if scrapBoard.whosMove == 2 {
      best = MoveResult{move: -1, val: 999999}
      play := newMoveResult(-1, 999999)
      for i := 0; i < 7; i++ { // for each possible move (6)
        randStartNum++; columnNum = randStartNum%7
        if scrapBoard.legalMove(columnNum) {
          //fmt.Printf("\nDFS%d: Player %d plays at %d",depth,scrapBoard.whosMove,i)
          scrapBoard.makeMove(columnNum)
          play = p.DFS(scrapBoard, depth-1) // will return results of optimal next move
          if (play.val < best.val) { // if results are better for player 2
            best.move = columnNum
            best.val = play.val
          }
          scrapBoard = b.copy()
        }
      }
    }

      return best
    }
  }
