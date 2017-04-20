package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    wordCountMap := make(map[string]int)
    words := strings.Fields(s)
    for i := 0;i < len(words);i++{
        word := words[i]
        wordCountMap[word] = wordCountMap[word] + 1
    }
    return wordCountMap
}

func main() {
    wc.Test(WordCount)
}

