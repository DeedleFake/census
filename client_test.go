package census_test

import (
	"."
	"testing"
)

func TestClient(t *testing.T) {
	c := &census.Client{
		ServiceID: "example",
		Game:      "ps2",
	}

	num, err := c.Count("character",
		census.SearchOption("name.first_lower", "deedlefaketr"),
	)
	if err != nil {
		t.Fatal(err)
	}
	if num != 1 {
		t.Fatalf("Expected %v\nGot %v", 1, num)
	}
}
