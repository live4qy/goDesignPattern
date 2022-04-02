package Factory

import "testing"

func TestFactory(t *testing.T){
	NewRestaurant("R1").GetFood()
	NewRestaurant("R2").GetFood()
}





