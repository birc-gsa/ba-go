package gsa

// Borderarray computes the border array over the string x. The border
// array ba will at index i have the length of the longest proper border
// of the string x[:i+1], i.e. the longest non-empty string that is both
// a prefix and a suffix of x[:i+1].
func Borderarray(x string) []int {
	ba := make([]int, len(x))
	// FIXME
	return ba
}

// StrictBorderarray computes the strict border array over the string x.
// This is almost the same as the border array, but ba[i] will be the
// longest proper border of the string x[:i+1] such that x[ba[i]] != x[i].
func StrictBorderarray(x string) []int {
	ba := make([]int, len(x))
	// FIXME
	return ba
}
