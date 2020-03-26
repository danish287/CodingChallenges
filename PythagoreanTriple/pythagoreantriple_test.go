package pythagoreantriple

import ( 
	"testing"
)

func TestCheckTriple(t *testing.T){
	arr := [3]int{3,4,5}
	if CheckTriple(arr) == false {
		t.Error("Test Failed: array is a Pythagorean Triple")
	}
}

func TestNonTriple(t *testing.T){
	arr := [3]int{5,4,3}
	if CheckTriple(arr) {
		t.Error("Test Failed: array is NOT a Pythagorean Triple")
	}
}