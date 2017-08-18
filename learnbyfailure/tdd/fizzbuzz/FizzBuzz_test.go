package fizzbuzz

import "testing"

func Test_fizzBuzzShouldSayNumber(t *testing.T) {
	actual := Say(1)

	if actual != "1" {
		t.Errorf("Expected : '%d', but = [%v]",1,actual,);
	}
}
