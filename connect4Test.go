package main

import (
  "os"
  "strconv"
)


func main() {
  // set up sequence string to be shared by multiple different objects
  aString := ""
  moveSeqPtr := &aString

  // Only need this seedForRandom if using a player ir SmartPlayer
  seedForRandom := 1
  if len(os.Args) > 1 {
    seedForRandom, _ = strconv.Atoi(os.Args[1])
  }

  // Player1 can be dumbAI, smartAI, or human
  //play1 := newHumanPlayer("Drew",1)
  //play1 := newPlayer("mr 1", 1)
  play1 := newSmartPlayer("Socrates",1, .1, seedForRandom)


  // Player2 can be dumbAI
  //play2 := newPlayer("mr 2", 2, 10.0, seedForRandom)
  //play2 := newHumanPlayer("Drew",2)
  //play2 := newSmartPlayer("Pluto",1, .005, seedForRandom)
  play2 := newLearnedPlayer("my student",1, .1, moveSeqPtr)

  // create match and play a game
  aMatch := newMatch(moveSeqPtr)
  aMatch.playGame(play1, play2)

	}
