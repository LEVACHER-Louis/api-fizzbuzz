package internal

import "testing"

func TestFizzBuzz(t *testing.T) {
	res := FizzBuzz(16)
	exp := "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16" 
	if res != exp {
		t.Errorf("Devrait avoir %v \n A obtenu %v", exp, res)
	}

	res2 := FizzBuzz(0)
	exp2 := ""
	if res2 != exp2 {
		t.Errorf("Devrait avoir string vide\n A obtenu %v", res2)
	}
}

func TestFizzBuzzMap(t *testing.T) {
	param1 := map[uint]string{2 : "Tres", 5:"Jazz", 3:"Oold"}
	res1, err1 := FizzBuzzMap(16, param1)
	exp1 := "1 Tres Oold Tres Jazz TresOold 7 Tres Oold TresJazz 11 TresOold 13 Tres OoldJazz Tres"
	if res1 != exp1 || err1 != nil {
		t.Errorf("Devrait avoir %v\n A obtenu %v", exp1, res1)
	}

	param2 := map[uint]string{}
	res2, err2 := FizzBuzzMap(7, param2)
	exp2 := "1 2 3 4 5 6 7"
	if res2 != exp2 || err2 != nil {
		t.Errorf("Devrait avoir %v\n A obtenu %v", exp2, res2)
	}

	param3 := map[uint]string{3:"Fizz", 5:"Buzz"}
	res3, err3 := FizzBuzzMap(0, param3)
	exp3 := ""
	if res3 != exp3 || err3 != nil {
		t.Errorf("Devrait avoir %v\n A obtenu %v", exp3, res3)
	}

	param4 := map[uint]string{0:"Test", 7:"Why"}
	res4, err4 := FizzBuzzMap(7, param4)
	exp4 := ""
	if res4 != exp4 || err4 == nil {
		t.Error("Devrait renvoyer erreur")
	}
}
