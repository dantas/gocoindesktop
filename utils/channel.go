package utils

func RedirectChannel[C any](input <-chan C, output chan<- C) {
	go func() {
		for err := range input {
			output <- err
		}
	}()
}
