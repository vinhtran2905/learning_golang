package channeltesting

import "testing"

func TestProcess(t *testing.T) {
	//GIVEN
	input := make(chan string)
	defer close(input)

	done := make(chan bool)
	defer close(done)

	go func() {
		input <- "input string"
		done <- true
	}()

	//WHEN
	output := Process(input)
	<-done // blocks until the input write routing is finished

	//THEN
	expected := "(input string)"
	actual := <-output
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
