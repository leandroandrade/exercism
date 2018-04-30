package letter

type Count map[rune]int

func Frequency(text string) Count {
	countWords := Count{}
	for _, word := range text {
		countWords[word]++
	}
	return countWords
}

func ConcurrentFrequency(texts []string) Count {
	queue := make(chan Count)

	for _, text := range texts {
		go func(text string) {
			queue <- Frequency(text)
		}(text)
	}

	total := Count{}
	for range texts {
		for letter, freq := range <-queue {
			total[letter] += freq
		}
	}
	return total

}
