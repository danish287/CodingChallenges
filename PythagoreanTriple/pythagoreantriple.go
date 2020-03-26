//Package pythagoreantriple is a porgram that checks if an array of three integers 
package pythagoreantriple

//CheckTriple will check if the given array is 
func CheckTriple(arr [3]int) bool {
	return IsTriple(arr[0], arr[1], arr[2])
}

//Square will square a given integer and return the result
func Square(num int) int {
	return num * num 
}

//IsTriple will check if a^2 + b^ 2 = c^2, if so, it will return T/F
func IsTriple(a int, b int, c int) bool {
	if Square(c) == (Square(a) + Square(b)){
		return true
	}
	return false 
}
