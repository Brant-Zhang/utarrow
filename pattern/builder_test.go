package pattern

import "testing"

func TestBuild(t *testing.T) {
	sportCar := NewBuilder().Color(BlueColor).Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
	err := sportCar.Drive()
	if err != nil {
		t.Fatal(err)
	}
	familyCar := NewBuilder().Color(RedColor).Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
	err = familyCar.Drive()
	if err != nil {
		t.Fatal(err)
	}
}
