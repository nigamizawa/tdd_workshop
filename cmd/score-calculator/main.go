package main

type frame struct {
	firstRoll  int
	secondRoll int
}

type frames []frame

func calculate_score(frames frames) int {
	total := 0
	for _, frame := range frames {
		total = frame.firstRoll + frame.secondRoll

		// todo: ストライク時の判定
		// sumStrike()

		// todo: スペア時の判定
		// sumSpare()
	}

	return total
}
