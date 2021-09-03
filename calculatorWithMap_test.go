package calculatorWithMap

import (
	"testing"
)

func TestCalculator(t *testing.T){
	cal := NewCalculator()
	if a, ok := cal["add"]; ok {
	 	r := a.Calculate(5, 6)
		if r != 11 {
			t.Errorf("not passed, result must be equal 11")
		}
	} else {
		t.Errorf("Error service not fount in map")
	}

	if a, ok := cal["substract"]; ok {
		r := a.Calculate(7, 6)
	   if r != 1 {
		   t.Errorf("not passed, result must be equal 1")
	   }
   } else {
	   t.Errorf("Error service not fount in map")
   }

   if a, ok := cal["multiple"]; ok {
	r := a.Calculate(7, 6)
   if r != 42 {
	   t.Errorf("not passed, result must be equal 42")
   }
  } else {
   t.Errorf("Error service not fount in map")
  }

  if a, ok := cal["division"]; ok {
	r := a.Calculate(12, 6)
   if r != 2 {
	   t.Errorf("not passed, result must be equal 2")
   }
  } else {
   t.Errorf("Error service not fount in map")
  }
	
}