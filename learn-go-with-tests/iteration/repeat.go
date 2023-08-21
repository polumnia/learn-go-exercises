package iteration

func Repeat(character string, times int) string {
	// explicitly declating variable
	// variables declared without a corresponding initialization are zero-valued
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
