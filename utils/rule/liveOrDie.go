package rule

func SetLiveOrDie(before int, count int) (after int) {
	after = before
	if before == 1 && count < 2 {
		after = 0
	} else if before == 1 && count > 3 {
		after = 0
	} else if before == 1 { // && (count == 2 || count == 3)
		after = 1
	} else if before == 0 && count == 3 {
		after = 1
	}
	return after
}
