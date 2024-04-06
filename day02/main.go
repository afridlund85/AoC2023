package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Game struct {
    id int
    Hands []Hand
}

type Hand struct {
    green int
    red int
    blue int
}

func (g *Game) is_valid(mh Hand) bool{
    for _, h := range g.Hands {
        if h.red > mh.red || h.green > mh.green || h.blue > mh.blue {
            return false
        }
    }
    return true
}

func main() {
    maxHand := Hand{13,12,14}
    input, _ := os.Open("./input")
    defer input.Close()
    res, _ := day02(maxHand, input)
    fmt.Println(res)
    i2, _ := os.Open("./input")
    defer i2.Close()
    res = part2(i2) 
    fmt.Println(res)
}

func part2(input io.Reader) int {
    scanner := bufio.NewScanner(input)
    res := 0
    for scanner.Scan() {
        mh := Hand{0,0,0}
        line := scanner.Text()
        game, _ := gameFromLine(line)
        for _, h := range game.Hands {
            if h.green > mh.green {
                mh.green = h.green
            }
            if h.blue > mh.blue {
                mh.blue = h.blue
            }
            if h.red > mh.red {
                mh.red = h.red
            }
        }
        pow := mh.red * mh.blue * mh.green
        res += pow
    }
    return res
}
func day02(maximum Hand, input io.Reader) (int, error) {
    scanner := bufio.NewScanner(input)
    res := 0
    for scanner.Scan() {
        line := scanner.Text()
        game, _ := gameFromLine(line)
        if game.is_valid(maximum){
            res += game.id
        }
    }
    return res,nil
}

func gameFromLine(line string) (Game, error) {
    g := Game{}
    s := strings.Split(line, ":") 
    id,_ := strconv.Atoi(strings.TrimPrefix(s[0], "Game "))
    g.id = id
    hands := strings.Split(s[1], ";")
    for _,h := range hands {
        hand := Hand{0,0,0}
        marbles := strings.Split(h, ",")
        for _, m := range marbles {
            marble := strings.Split(strings.TrimSpace(m), " ")
            color := marble[1]
            count, _ := strconv.Atoi(marble[0])

            switch color {
            case "green":
                hand.green = count
            case "blue":
                hand.blue = count
            case "red":
                hand.red = count
            }
        }
        g.Hands = append(g.Hands, hand)
    }
    return g, nil
}

