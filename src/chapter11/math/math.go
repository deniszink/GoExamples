package math


func Average(xs []int) int{ // Average - is a public accessor, average is a private accessor
	total := 0
	for _, x := range xs{
		total += x
	}
	return total / len(xs)
}
