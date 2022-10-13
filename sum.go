package sums

// Erwartet zwei Zahlen from und to.
// Liefert die Summe aller Zahlen, die zwischen from und to liegen.
func Sum(from, to int) int {
	result := 0
	for i := from; i <= to; i++ {
		result += i
	}
	return result
}
