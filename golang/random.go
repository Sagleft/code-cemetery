package main

func getRandomInt(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(reasons))
}
