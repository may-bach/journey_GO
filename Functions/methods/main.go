package main

import "fmt"

type Player struct {
	name   string
	team   string
	points int
}                              // type structure

type PlayerScore int           // type variable

func (p PlayerScore) isHighScoring() bool {
	return p >= 30
}

func (p Player) showStats() {
	fmt.Printf("Player: %s (%s) -- Points: %d\n", p.name, p.team, p.points)
}

func (p *Player) scorePoints(pts int) {
	p.points += pts
	fmt.Printf("  ...%s scores %d! New total: %d\n", p.name, pts, p.points)
}

func main() {
	ball := Player{
		name:   "lamelo",
		team:   "hornets",
		points: 28,
	}

	ball.showStats()

	fmt.Println("Lamelo shoots, AND HE SCORES")
	ball.scorePoints(3)

	ballpts := PlayerScore(ball.points)

	ball.showStats()
	if(ballpts.isHighScoring()){
		fmt.Printf("%s is a high scorer")
	}
}