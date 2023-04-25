package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Team struct {
	name                string
	wins, losses, draws int
}

type teamSlice []Team

func (t teamSlice) Len() int {
	return len(t)
}

func (t teamSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t teamSlice) Less(i, j int) bool {
	if t[i].points() != t[j].points() {
		return t[i].points() > t[j].points()
	}
	return t[i].name < t[j].name
}

func (t *Team) matchCount() int {
	return t.wins + t.losses + t.draws
}

func (t *Team) points() int {
	return t.wins*3 + t.draws
}

func (t *Team) toString() string {
	return fmt.Sprintf("|  %v |  %v |  %v |  %v |  %v\n", t.matchCount(), t.wins, t.draws, t.losses, t.points())
}

func Tally(reader io.Reader, writer io.Writer) error {
	results, err := parseInput(reader)
	if err != nil {
		return err
	}
	r := make(teamSlice, 0, len(results))
	for _, v := range results {
		r = append(r, v)
	}
	sort.Sort(r)
	if _, err := writer.Write([]byte(getBanner(31))); err != nil {
		return err
	}

	for _, result := range r {
		if _, err := writer.Write([]byte(fmt.Sprintf("%-31v", result.name))); err != nil {
			return err
		}
		if _, err := writer.Write([]byte(result.toString())); err != nil {
			return err
		}
	}

	return nil
}

func getBanner(pad int) string {
	sprintString := "%-" + strconv.Itoa(pad) + "v| MP |  W |  D |  L |  P\n"
	return fmt.Sprintf(sprintString, "Team")
}

func parseInput(reader io.Reader) (map[string]Team, error) {
	teams := make(map[string]Team)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		match := strings.Split(scanner.Text(), ";")
		if len(match) == 2 {
			return nil, errors.New("parsing error")
		}
		if len(match) == 3 {
			switch match[2] {
			case "win":
				{
					recordResult(teams, match[0], 1, 0, 0)
					recordResult(teams, match[1], 0, 1, 0)
				}
			case "loss":
				{
					recordResult(teams, match[0], 0, 1, 0)
					recordResult(teams, match[1], 1, 0, 0)
				}
			case "draw":
				{
					recordResult(teams, match[0], 0, 0, 1)
					recordResult(teams, match[1], 0, 0, 1)
				}
			default:
				return nil, errors.New("invalid match result")
			}
		}
	}

	return teams, nil
}

func recordResult(teams map[string]Team, team string, w, l, d int) {
	if t, exists := teams[team]; exists {
		t.wins += w
		t.losses += l
		t.draws += d
		teams[team] = t
	} else {
		teams[team] = Team{name: team, wins: w, losses: l, draws: d}
	}
}
