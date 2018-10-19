package channeltesting

import "sync"

func Process(input <-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(1)
	output := make(chan string)

	go func() {
		for str := range input {
			output <- doHeavyOperation(str)
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func doHeavyOperation(str string) string {
	return "(" + str + ")"
}
