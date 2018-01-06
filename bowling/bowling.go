package bowling

import "errors"

type Game struct {
	previousRolls []int
}

type frame int

func NewGame() *Game {
	return &Game{}
}

// Roll is called each time the player rolls a ball.
// The argument is the number of pins knocked down.
func (g *Game) Roll(pin int) error {
	if 0 <= pin && pin <= 10 {
		g.previousRolls = append(g.previousRolls, pin)
		return nil
	}
	return errors.New("number of pin not possible")

}

// Score  is called only at the very end of the game.
// It returns the total score for that game.
func (g *Game) Score() (int, error) {
	score := 0
	var cur frame
	if len(g.previousRolls) < 12 {
		return 0, errors.New("an unstarted game cannot be scored")
	}

	for i := 0; i < 10; i++ {
		switch {
		case isStrike(cur, g.previousRolls):
			bonus, err := strikeBonus(cur, g.previousRolls)
			if err != nil {
				return 0, err
			}
			score += 10 + bonus
			cur++
		case isSpare(cur, g.previousRolls):
			bonus, err := spareBonus(cur, g.previousRolls)
			if err != nil {
				return 0, err
			}
			score += 10 + bonus
			cur += 2
		default:
			sum, err := sumOfRolls(cur, g.previousRolls)
			if err != nil {
				return 0, err
			}
			score += sum
			cur += 2
		}
	}
	return score, nil
}

func isStrike(f frame, rolls []int) bool {
	return rolls[f] == 10
}

func isSpare(f frame, rolls []int) bool {
	sum, err := sumOfRolls(f, rolls)
	if err != nil {
		return false
	}
	return sum == 10
}

func strikeBonus(f frame, rolls []int) (int, error) {
	if int(f) < len(rolls) {
		sum, err := sumOfRolls(f+1, rolls)
		if err != nil {
			return 0, err
		}
		return sum, nil
	}
	return 0, errors.New("foo")
}

func spareBonus(f frame, rolls []int) (int, error) {
	if int(f+2) < len(rolls) {
		return rolls[f+2], nil
	}
	return 0, errors.New("foo")
}

func sumOfRolls(f frame, rolls []int) (int, error) {
	if int(f+1) < len(rolls) {
		sum := rolls[f] + rolls[f+1]
		if sum <= 20 {
			return sum, nil
		}
	}
	return 0, errors.New("Sum to high")
}
