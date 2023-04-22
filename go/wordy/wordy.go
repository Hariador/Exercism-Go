package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PLUS  = "plus"
	MINUS = "minus"
	MULT  = "multiplied"
	DIV   = "divided"
)

const (
	EXPECT_TERM = 0
	EXPECT_OP   = 1
	EXPECT_FILL = 2
)

func Answer(question string) (int, bool) {
	if question[len(question)-1:] != "?" {
		return 0, false
	}
	tokens := strings.Split(question[:len(question)-1], " ")
	if len(tokens) < 3 {
		return 0, false
	}
	if tokens[0] != "What" && tokens[1] != "is" {
		return 0, false
	}
	ans, err := strconv.Atoi(tokens[2])
	fmt.Println(ans)
	if err != nil {
		return 0, false
	}

	step := EXPECT_OP
	var op string
	for _, token := range tokens[3:] {
		switch step {
		case EXPECT_FILL:
			{
				step = EXPECT_TERM
			}
		case EXPECT_OP:
			{
				op = token
				if op == MULT || op == DIV {
					step = EXPECT_FILL
				} else if op == PLUS || op == MINUS {
					step = EXPECT_TERM
				} else {
					return 0, false
				}

			}
		case EXPECT_TERM:
			{
				term, err := strconv.Atoi(token)
				if err != nil {
					return 0, false
				}
				switch op {
				case PLUS:
					{
						ans = ans + term
					}
				case MINUS:
					{
						ans = ans - term
					}
				case MULT:
					{
						ans = ans * term
					}
				case DIV:
					{
						ans = ans / term
					}
				default:
					return 0, false
				}

				step = EXPECT_OP
			}
		}
	}
	if step == EXPECT_TERM {
		return 0, false
	}
	return ans, true
}
