package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
)

type tally struct {
	m       map[string]*result
	ranking []string
}

type result struct {
	mp int // MP: Matches Played
	w  int //  W: Matches Won
	d  int //  D: Matches Drawn (Tied)
	l  int //  L: Matches Lost
	p  int //  P: Points
}

// Tally gets an csv List of teams and sports results and returns a nice printed table with the results of an competition
func Tally(r io.Reader, w io.Writer) error {
	t := newTally()

	c := csv.NewReader(r)
	c.Comma = ';'
	c.Comment = '#'
	c.FieldsPerRecord = 3

	for {
		record, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = t.trackResult(record[0], record[1], record[2])
		if err != nil {
			return err
		}
	}
	t.rank()
	t.Write(w)
	return nil
}

// trackResult keeps track of the results
func (t *tally) trackResult(first, second, res string) error {
	winR := result{w: 1, p: 3}
	lossR := result{l: 1}
	drawR := result{d: 1, p: 1}

	switch res {
	case "win":
		t.save(first, winR)
		t.save(second, lossR)
	case "loss":
		t.trackResult(second, first, "win")
	case "draw":
		t.save(first, drawR)
		t.save(second, drawR)
	default:
		return errors.New("Error: Don't know " + res)
	}
	return nil
}

// save is used to save the result of an team
func (t *tally) save(team string, res result) {
	if r, ok := t.m[team]; ok {
		//r.mp += res.mp
		r.w += res.w
		r.d += res.d
		r.l += res.l
		r.p += res.p
	} else { // team not in map
		t.m[team] = &res
	}
	cur := t.m[team]
	cur.mp++
}

// newTally returns a new Tally object
func newTally() tally {
	return tally{m: make(map[string]*result)}
}

// rank is used to generate the ranked list when all data was saved
func (t *tally) rank() {
	points := []int{}
	byPoints := make(map[int][]string)
	for name, r := range t.m {
		points = append(points, r.p)
		if _, ok := byPoints[r.p]; ok {
			byPoints[r.p] = append(byPoints[r.p], name)
		} else { // points not in map
			byPoints[r.p] = []string{name}
		}
	}
	sort.Ints(points)
	seen := make(map[string]bool)
	for i := len(points) - 1; i >= 0; i-- { //ordered by points, descending
		candidates := byPoints[points[i]] //teams with the same score
		sort.Strings(candidates)
		for j := 0; j < len(candidates); j++ { //ordered alphabetically
			team := byPoints[points[i]][j]
			if !seen[team] {
				t.ranking = append(t.ranking, team)
				seen[team] = true
			}
		}
	}
}

// Write is used to present the output in a text file
func (t *tally) Write(w io.Writer) {
	out := fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team")
	for _, name := range t.ranking {
		r := t.m[name]
		out += fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", name, r.mp, r.w, r.d, r.l, r.p)
	}
	b := []byte(out)
	w.Write(b)
}
