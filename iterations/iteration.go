package iterations

func Repeat(char string, times int) (repetitions string) {
	for i := 0; i < times; i++ {
		repetitions += char
	}
	return
}
