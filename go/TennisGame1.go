package tenniskata

// score should be attached to a player
type tennisGame1 struct {
	p1 player
	p2 player
}

type player struct {
	name  string
	score Score
}

// 2 string
func NewTennisGame1(player1Name, player2Name string) TennisGame {
	return &tennisGame1{
		p1: player{
			name: player1Name,
		},
		p2: player{
			name: player2Name,
		},
	}
}

func (p *player) IncrementScore() {
	p.score++
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.p1.IncrementScore()
	} else {
		game.p2.IncrementScore()
	}
}

func (game *tennisGame1) GetScore() string {
	if game.p1.score == game.p2.score {
		return game.p1.score.StringOnEquality()
	}

	if game.p1.score.IsMax() || game.p2.score.IsMax() {
		return onPotentialWinner(game.p1.score, game.p2.score)
	}

	return game.wtfunc()
}

type Score int

const MaxRoundScore Score = 4

func (s Score) StringOnEquality() string {
	switch s {
	case 0, 1, 2:
		return s.String() + "-All"
	default:
		return "Deuce"
	}
}

func (s Score) String() string {
	switch s {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	default:
		panic("ooups")
		return "FAIL"
	}
}

func (s Score) IsMax() bool {
	return s >= MaxRoundScore
}

func (game tennisGame1) wtfunc() string {
	return game.p1.score.String() + "-" + game.p2.score.String()
}

func onPotentialWinner(scoreP1, scoreP2 Score) string {
	minusResult := scoreP1 - scoreP2

	switch {
	case minusResult == 1:
		return "Advantage player1"
	case minusResult == -1:
		return "Advantage player2"
	case minusResult >= 2:
		return "Win for player1"
	default:
		return "Win for player2"
	}
}
