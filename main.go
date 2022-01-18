package main

import (
    "bufio"
    "fmt"
    "log"
    "strings"
    "os"
)

type Operator uint8

const (
    Present Operator = iota
    NotPresent Operator = iota
    MaxOccurances Operator = iota
    MinOccurances Operator = iota
)

type Condition struct {
    letter rune
    operator Operator
    value int
}


func (c Condition) Matches(word string) bool {
    if c.operator == MaxOccurances {
        // occurs no more than the given number of times?
        return strings.Count(word, string(c.letter)) <= c.value
    } else if c.operator == MinOccurances {
        // occurs at least the given number of times?
        return strings.Count(word, string(c.letter)) >= c.value
    } else {
        var present bool
        if c.value == -1 {
            // present anywhere?
            present = strings.Contains(word, string(c.letter))
        } else {
            // present at given position?
            i := 0
            for _, char := range word {
                if i == c.value && char == c.letter {
                    present = true
                }
                i++
            }
        }
        if present {
            return c.operator == Present
        } else {
            return c.operator == NotPresent
        }
    }
}

var initialConditions = []Condition{
    {'s', Present, -1},
    {'a', Present, -1},
    {'r', Present, -1},
    {'e', Present, -1},
    {'t', Present, -1},
}

/*
var initialConditions = []Condition{
    {'a', NotPresent, -1},
    {'t', NotPresent, -1},
    {'e', NotPresent, -1},
    {'s', NotPresent, -1},
    {'r', NotPresent, 0},
    {'r', Present, -1},
    {'w', NotPresent, -1},
    {'l', NotPresent, -1},
    {'d', NotPresent, -1},
    {'o', Present, -1},
    {'o', NotPresent, 1},
    {'r', NotPresent, 2},
    {'r', Present, 1},
    {'o', Present, 2},
    {'g', NotPresent, -1},
    {'u', NotPresent, -1},
    {'p', Present, -1},
    {'p', NotPresent, 4},
    {'p', Present, 0},
    {'o', MaxOccurances, 1},
    {'f', NotPresent, -1},
}
*/

func readWords(filename string) []string {
    f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)

    var words []string
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), ",")
        if len(s) == 2 && len(s[0]) == 5 && s[1] != "count" {
            words = append(words, s[0])
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return words
}

func findGuess(words []string, guesses map[string]bool, conditions []Condition) string {
    for i := 0; i < len(words); i++ {
        word := words[i]
        _, guessedBefore := guesses[word]
        if !guessedBefore && matchesAll(conditions, word) {
            return word
        }
    }
    return ""
}

func matchesAll(conditions []Condition, word string) bool {
    for i := 0; i < len(conditions); i++ {
        if (!conditions[i].Matches(word)) {
            if (word == "humph") {
                fmt.Printf("Not maching humph with condition %d\n", i)
            }
            return false
        }
    }
    return true
}

func mostCommonLetter(words []string, position int) rune {
    var countMap = map[rune]int{}
    for i := 0; i < len(words); i++ {
        word := words[i]
        j := 0
        for _, char := range word {
            if j == position {
                count, ok := countMap[char]
                if ok {
                    countMap[char] = count + 1
                } else {
                    countMap[char] = 1
                }
            }
            j++
        }
    }
    var bestLetter rune
    var bestCount int
    for key, value := range countMap {
        if value > bestCount {
            bestCount = value
            bestLetter = key
        }
    }
    return bestLetter
}

func main() {
    words := readWords("unigram_freq.csv")
    fmt.Printf("Most common first letter: %s\n", string(mostCommonLetter(words, 0)))
    fmt.Printf("Most common second letter: %s\n", string(mostCommonLetter(words, 1)))
    fmt.Printf("Most common third letter: %s\n", string(mostCommonLetter(words, 2)))
    fmt.Printf("Most common fourth letter: %s\n", string(mostCommonLetter(words, 3)))
    fmt.Printf("Most common fifth letter: %s\n", string(mostCommonLetter(words, 4)))
    guesses := map[string]bool{}
    guesses["huynh"] = true // TODO: allow user to reject guesses
    nextGuess := findGuess(words, guesses, initialConditions)
    fmt.Printf("Next guess: %s\n", nextGuess)
    guesses[nextGuess] = true
}
