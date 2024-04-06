package main

import (
	"strings"
	"testing"
)

func TestDay02(t *testing.T){
    s := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
    input := strings.NewReader(s)
    maxHand := Hand{13,12,14}

    want := 8
    got, err := day02(maxHand, input)

    if got != want || err != nil {
        t.Fatalf("got %d want %d", got, want)
    }
}

func TestPart2(t *testing.T){
    s := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
    input := strings.NewReader(s)
    want := 2286
    got := part2(input)
    if got != want {
        t.Fatalf("got %d want %d", got, want)
    }
}
