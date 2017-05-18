/* PlayerInterface.go contains the interface that represents
/ all types of players
/
/  Author: Drew Davis
*/

package main

/* the AnyPlayer interface allows all subclasses of player to be
/ called by board functions
*/
type AnyPlayer interface {
  chooseMove(Board) int
}
