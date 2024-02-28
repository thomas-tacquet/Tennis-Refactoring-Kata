package tenniskata

// score should be attached to a player
type tennisGame1 struct {
	// m_score1    int
	// m_score2    int
	// player1Name string
	// player2Name string
	p1 player
	p2 player
}

type player struct {
	name  string
	score int
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
	score := ""

	if game.p1.score == game.p2.score {
		score = onEqualScore(game)
	} else if game.p1.score >= 4 || game.p2.score >= 4 {
		score = onPotentialWinner(game)
	} else {
		score = game.wtfunc()
	}
	return score
}

func numToScoreString(score int) string {
	switch score {
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

func (game tennisGame1) wtfunc() string {
	player1Score := numToScoreString(game.p1.score)
	player2Score := numToScoreString(game.p2.score)

	return player1Score + "-" + player2Score
}

func onPotentialWinner(game *tennisGame1) string {
	minusResult := game.p1.score - game.p2.score

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

func onEqualScore(game *tennisGame1) string {
	switch game.p1.score {
	case 0:
		return numToScoreString(game.p1.score) + "-All"
	case 1:
		return numToScoreString(game.p1.score) + "-All"
	case 2:
		return numToScoreString(game.p1.score) + "-All"
	default:
		return "Deuce"
	}
}
