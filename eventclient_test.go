package census

//func TestEcho(t *testing.T) {
//	const test = "This is a test."
//
//	c, err := NewEventClient("", "", "example")
//	if err != nil {
//		t.Fatalf("Failed to create EventClient: %v", err)
//	}
//	defer c.Close()
//
//	var out struct {
//		Test string `json:"test"`
//	}
//	err = c.echo(&out, map[string]interface{}{
//		"test": test,
//	})
//	if err != nil {
//		t.Errorf("Failed to echo: %v", err)
//	}
//
//	if out.Test != test {
//		t.Errorf("Got back anomolous data. Ex: %q Got: %q", test, out.Test)
//	}
//}
