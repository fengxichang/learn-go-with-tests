package iteration

func Repeat(c string, repeatCount int) string {
	repeated := ""
	for i :=0; i<repeatCount; i++ {
		repeated += c
	}
	return repeated
}
