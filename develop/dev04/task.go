package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	words := []string{"пятак", "листок", "слиток", "столик", "пятка", "тяпка"}
	out := Find(words)
	for _, o := range out {
		fmt.Println(o)
	}
}

func Find(words []string) map[string][]string {
	buf := make(map[string]int8)
	out := make(map[string][]string)
	for _, word := range words {
		s := strings.ToLower(word)
		if _, ok := buf[s]; ok == false {
			buf[s] = 1
			sortS := sortString(s)
			if _, ok := out[sortS]; ok == false {
				out[sortS] = make([]string, 0)
			}
			out[sortS] = append(out[sortS], s)
		}
	}
	return out
}

func sortString(word string) string {
	wordRunes := []rune(word)
	sort.Slice(wordRunes, func(i, j int) bool {
		return wordRunes[i] < wordRunes[j]
	})

	wordSorted := string(wordRunes)
	return wordSorted
}
