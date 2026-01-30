package tennis

import (
	"fmt"
	"tennis/game"
)

type Set struct {
	games []game.Game
}

// New erstellt einen neuen Tennis-Satz.
// Der Satz beginnt mit einem neuen Spiel.
func New() Set {
	return Set{games: []game.Game{game.New()}}
}

// CurrentGameNo gibt die Nummer des aktuellen Spiels im Satz zurück.
func (s Set) CurrentGameNo() int {
	return len(s.games)
}

// CurrentGame gibt das aktuelle Spiel im Satz zurück.
func (s Set) CurrentGame() *game.Game {
	return &s.games[s.CurrentGameNo()-1]
}

// ScoreP1 erhöht den Punktestand von Spieler 1 im aktuellen Spiel.
func (s *Set) ScoreP1() {
	s.CurrentGame().ScoreP1()
}

// ScoreP2 erhöht den Punktestand von Spieler 2 im aktuellen Spiel.
func (s *Set) ScoreP2() {
	s.CurrentGame().ScoreP2()
}

// CurrentGamesP1 gibt die Anzahl der gewonnenen Spiele von Spieler 1 im Satz zurück.
func (s Set) CurrentGamesP1() int {
	count := 0
	for _, g := range s.games {
		if g.P1Wins() {
			count++
		}
	}
	return count
}

// CurrentGamesP2 gibt die Anzahl der gewonnenen Spiele von Spieler 2 im Satz zurück.
func (s Set) CurrentGamesP2() int {
	count := 0
	for _, g := range s.games {
		if g.P2Wins() {
			count++
		}
	}
	return count
}

// String gibt den aktuellen Satzstand als String zurück.
// Dabei wird zunächst der Stand der fertigen Spiele angezeigt,
// gefolgt vom Punktestand des aktuellen Spiels.
func (s Set) String() string {
	result := fmt.Sprintf("%d-%d\n", s.CurrentGamesP1(), s.CurrentGamesP2())
	result += fmt.Sprintf("%s\n", s.CurrentGame())
	return result
}
