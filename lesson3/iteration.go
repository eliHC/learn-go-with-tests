package main

func Repeat(char string, numberOfIterations int) string {
	var repeated string
	for i := 0; i < numberOfIterations; i++ {
		repeated += char
	}

	return repeated
}
