/* LearnedPlayer.go contains a substuct of player that was taught
/ to play intelligently through reinforcement learning
/
/  Author: Drew Davis
*/


package main

import (
  "time"
  //"math/rand"
  "os"
  "bufio"
  "fmt"
  "strconv"
)


type LearnedPlayer struct {
  Player
  mapOfLosses map[string]int
  moveSeq *string
}

/* newLearnedPlayer acts as the LearnedPlayer constructor
/  name, playerNum, and timeLim are set as specified by the caller
/  a random seed is generated to allow the results of AI vs AI games
/  to be varied
*/
func newLearnedPlayer(aName string, aPlayerNum int, maxTime float64, moveSeqPtr *string) LearnedPlayer {
  // populate a map with all previous move sequences that lead to a loss
  var aMap map[string]int
  aMap = make(map[string]int)
  file, err := os.Open("player2LosingMoveSeqs.txt")
  if err != nil {
    print("\n\n\nError loading in learned data"); print(err)
  }
  aScanner := bufio.NewScanner(file)
  // put each string from history file into map
  for aScanner.Scan() {
    aMap[aScanner.Text()] = 1
    //fmt.Printf("adding sequence %s to map",aScanner.Text())
  }

  p := LearnedPlayer{Player{
    name: aName,
    playerNum: aPlayerNum,
    timeLim: maxTime,
  }, aMap, moveSeqPtr}

  return p
}

/* updateLossData adds the sequence of moves to the data text file
/  after player 2 loses OR ties a game. In this way, the LearnedPlayer
/  will continually be updating its moves and be improving
*/
func (p *LearnedPlayer) updateLossData() {

  file, err := os.OpenFile("player2LosingMoveSeqs.txt", os.O_RDWR|os.O_APPEND, 0755)
  if err != nil {
    fmt.Printf("Error opening learned data file\n")
    return
  }
  defer file.Close()

  print("Match sequence learned.\n")
  for k := 0; k <= len(*p.moveSeq); k += 2 {
    //print("\nAdded sequence to loss data: ");print((*p.moveSeq)[0:k]);
    // if the sequence is already in the list, do not duplicate
    if p.mapOfLosses[(*p.moveSeq)[0:k]] == 0 {
      _, err = file.WriteString("\n"+(*p.moveSeq)[0:k])
      if err != nil {
        fmt.Printf("\nError updating learned data%s\n",err)
        os.Exit(1)
      }
    }
  }

}

/* evaluate a board situation and return a numerical score
/  higher score == better situation for player 1
/  lower score  == better situation for player 2
/  Input: a board to be evaluated
/  Output: an int score representing quality of situation
*/
func (p *LearnedPlayer) evaluate(b Board, localMoveSeq string) int {
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

  // check to see if this move has lead to defeat previously
  // if it has, avoid move if reasonable
  // (only checks for player 2)
  //  if the turn at the current depth DFS matches a move sequence
  //  that has previously caused player 2 to lose, it will be avoided
  //  through this evaluation. It will only avoid these conditions if
  // reasonable (AKA won't just give up game to avoid them)
  //fmt.Printf("\nmoveSeq being evaluated: %s",localMoveSeq);
  /*if p.mapOfLosses[localMoveSeq] > 0 {
    value1 += 100
    //print(" <- this has gone bad before. Avoid\n")
  }*/

  // check to see if this move has lead to defeat previously
  // if it has, avoid move if reasonable
  // (only checks for player 2)
  //  if any amount of the move sequence matches a move sequence that
  //  has previously caused player 2 to lose, it will be avoided
  //  through this evaluation. It will avoid going through conditions
  // have previously caused defeat, even if it does not anticipate
  // that the players will follow the same route
  for k := 0; k <= len(localMoveSeq); k += 2 {
    if p.mapOfLosses[localMoveSeq[0:k]] > 0 {
      value1 += 20
      //print("\n(A previous losing move sequence has been detected\n")
    }
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
func (p LearnedPlayer) chooseMove(b Board) int {
  t := time.Now()
  time.Since(t)
  depth := 1
  best := newMoveResult(6,0)
  // perform deeper DFS until time is gone
  // if part-way through a DFS when time runs out, finish the DFS first
  for ( (time.Since(t)).Seconds() < float64(p.timeLim) ) {
    //fmt.Printf("\n\nBegin new DFS at with depth: %d", depth)
    //fmt.Printf("search at depth %d complete", depth)
    best = p.DFS(b, depth, *p.moveSeq) // DFS at depth
    depth++
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
func (p *LearnedPlayer) DFS(b Board, depth int, inScrapMoveSeq string) MoveResult {
  //fmt.Printf("\nDFS: %d", depth)
  scrapBoard := b.copy()
  scrapMoveSeq := inScrapMoveSeq
  best := newMoveResult(-1,-1)
  if (depth == 0 || scrapBoard.gameOver()) {
    best.val = p.evaluate(scrapBoard, scrapMoveSeq) // return value of current board
    //fmt.Printf(" - This sequence is valued at %d (For player 2)",best.val)
    return best
  } else {

    // evaluate for player 1 (high eval is better)
    if scrapBoard.whosMove == 1 {
      best = MoveResult{move: -1, val: -999999}
      play := newMoveResult(-1, -999999)
      for i := 0; i < 7; i++ { // for each possible move (6)
        if scrapBoard.legalMove(i) {
          //fmt.Printf("\nDFS%d: Player %d plays at %d",depth,scrapBoard.whosMove,i)
          scrapBoard.makeMove(i)
          scrapMoveSeq = scrapMoveSeq + " " + strconv.Itoa(i)
          play = p.DFS(scrapBoard, depth-1, scrapMoveSeq) // will return results of optimal next move
          if (play.val > best.val) { // if results are better for player 1
            best.move = i
            best.val = play.val
          }
          scrapBoard = b.copy()
          scrapMoveSeq = inScrapMoveSeq
        }
      }
    }
    // evaluate for player 2 (low eval is better)
    if scrapBoard.whosMove == 2 {
      best = MoveResult{move: -1, val: 999999}
      play := newMoveResult(-1, 999999)
      for i := 0; i < 7; i++ { // for each possible move (6)
        if scrapBoard.legalMove(i) {
          //fmt.Printf("\nDFS%d: Player %d plays at %d",depth,scrapBoard.whosMove,i)
          scrapBoard.makeMove(i)
          scrapMoveSeq = scrapMoveSeq + " " + strconv.Itoa(i)
          play = p.DFS(scrapBoard, depth-1, scrapMoveSeq) // will return results of optimal next move
          if (play.val < best.val) { // if results are better for player 2
            best.move = i
            best.val = play.val
          }
          scrapBoard = b.copy()
          scrapMoveSeq = inScrapMoveSeq
        }
      }
    }

      return best
    }
  }
