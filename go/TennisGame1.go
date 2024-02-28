package tenniskata

// score should be attached to a player
type tennisGame1 struct {
	m_score1    int
	m_score2    int
	player1Name string
	player2Name string
}

// 2 string
func NewTennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name,
	}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.m_score1 += 1
	} else {
		game.m_score2 += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""

	if game.m_score1 == game.m_score2 {
		score = onEqualScore(game)
	} else if game.m_score1 >= 4 || game.m_score2 >= 4 {
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
	player1Score := numToScoreString(game.m_score1)
	player2Score := numToScoreString(game.m_score2)

	return player1Score + "-" + player2Score
}

func onPotentialWinner(game *tennisGame1) string {
	minusResult := game.m_score1 - game.m_score2

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
	switch game.m_score1 {
	case 0:
		return "Love-All"
	case 1:
		return "Fifteen-All"
	case 2:
		return "Thirty-All"
	default:
		return "Deuce"
	}
}
