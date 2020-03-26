package PythagoreanTriple

//CheckTriple will 
func CheckTriple(arr [3]string) bool {
	return IsTriple(arr[0], arr[1], arr[2])
}

//Square will square a given integer
func Square(num int) int {
	return num * num 
}

//IsTriple will determine if the array is a Pythagorean Triple 
func IsTriple(a int, b int, c int) bool {
	if Square(c) == (Square(a) + Square(b)){
		return true
	}
	return false 
}