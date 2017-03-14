package utils

import (
	"fmt"
	"testing"
)

func Test_Distance(t *testing.T) {
	lat1 := 28.741171
	long1 := 77.135054

	lat2 := 28.725429
	long2 := 77.144348

	dist := Distance(lat1, long1, lat2, long2)
	if dist > 1971 && dist < 1974 {
		t.Log("Distance test passed")
	} else {
		t.Error("Distance test failed" + fmt.Sprintf("%f", dist))
	}
}

func Test_ValidateLatLong(t *testing.T) {
	lat1 := 28.741171
	long1 := 77.135054

	lat2 := 97.725429
	long2 := 77.144348

	lat3 := 23.725429
	long3 := 182.144348

	err := ValidateLatLong(lat1, long1)
	if err != nil {
		t.Error("ValidateLatLong failed for set 1")
	}

	err = ValidateLatLong(lat2, long2)
	if err == nil {
		t.Error("ValidateLatLong failed for set 2")
	}

	err = ValidateLatLong(lat3, long3)
	if err == nil {
		t.Error("ValidateLatLong failed for set 3")
	}
}
