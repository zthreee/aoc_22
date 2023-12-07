package main

//
//import (
//	"bufio"
//	"fmt"
//	"log"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//	"time"
//)
//
//func main() {
//	startTime := time.Now().UnixMilli()
//
//	result := solve("./2023/day4/input_2.txt")
//	fmt.Println("Result:", result)
//
//	duration := time.Now().UnixMilli() - startTime
//	fmt.Printf("Duration: %vms\n", duration)
//}
//
//func solve(path string) int {
//	file, err := os.Open(path)
//	if err != nil {
//		log.Panic(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanLines)
//
//	hands := []Hand{}
//	for scanner.Scan() {
//		line := scanner.Text()
//		//fmt.Println("line", line)
//
//		hand := parseInput(line)
//		//fmt.Println("hand", hand)
//
//		hands = append(hands, hand)
//	}
//
//	sort.Slice(hands, func(i, j int) bool {
//		if hands[i].Combo == hands[j].Combo {
//			for i, v1 := range hands[i].Card {
//				v2 := rune(hands[j].Card[i])
//				if v1 == v2 {
//					continue
//				}
//
//				return cardPower(v1) < cardPower(v2)
//			}
//
//			return true
//		} else {
//			return hands[i].Combo < hands[j].Combo
//		}
//	})
//
//	total := 0
//	for rank := 1; rank <= len(hands); rank++ {
//		hand := hands[rank-1]
//		total += rank * hand.Bid
//	}
//
//	return total
//}
//
//type Hand struct {
//	Card  string
//	Bid   int
//	Combo Combo
//}
//
//type Combo int
//
//const (
//	FIVE_OF_A_KIND  Combo = 7
//	FOUR_OF_A_KIND  Combo = 6
//	FULL_HOUSE      Combo = 5
//	THREE_OF_A_KIND Combo = 4
//	TWO_PAIR        Combo = 3
//	ONE_PAIR        Combo = 2
//	HIGH_CARD       Combo = 1
//)
//
//func parseInput(line string) Hand {
//	strs := strings.Split(line, " ")
//	card := strs[0]
//	bid, err := strconv.Atoi(strs[1])
//	if err != nil {
//		log.Panic(err)
//	}
//
//	jokers := 0
//	m := map[rune]int{}
//	for _, v := range card {
//		if v == 'J' {
//			jokers++
//			continue
//		}
//
//		m[v]++
//	}
//
//	combo := HIGH_CARD
//	for k1, v := range m {
//		switch {
//		case v == 5:
//			combo = FIVE_OF_A_KIND
//		case v == 4:
//			combo = FOUR_OF_A_KIND
//			if jokers == 1 {
//				combo = FIVE_OF_A_KIND
//			}
//		case v == 3:
//			combo = THREE_OF_A_KIND
//			if jokers == 2 {
//				combo = FIVE_OF_A_KIND
//			} else if jokers == 1 {
//				combo = FOUR_OF_A_KIND
//			} else {
//				for k2, v2 := range m {
//					if k1 == k2 {
//						continue
//					}
//
//					if v2 == 2 {
//						combo = FULL_HOUSE
//						break
//					}
//				}
//			}
//		case v == 2:
//			combo = ONE_PAIR
//			if jokers == 3 {
//				combo = FIVE_OF_A_KIND
//			} else if jokers == 2 {
//				combo = FOUR_OF_A_KIND
//			} else {
//				for k2, v2 := range m {
//					if k1 == k2 {
//						continue
//					}
//
//					if v2 == 3 {
//						combo = FULL_HOUSE
//						break
//					} else if v2 == 2 {
//						combo = TWO_PAIR
//						if jokers == 1 {
//							combo = FULL_HOUSE
//						}
//						break
//					}
//				}
//			}
//
//			if combo == ONE_PAIR && jokers == 1 {
//				combo = THREE_OF_A_KIND
//			}
//		}
//
//		if combo != HIGH_CARD {
//			break
//		}
//	}
//
//	if combo == HIGH_CARD {
//		switch jokers {
//		case 4, 5:
//			combo = FIVE_OF_A_KIND
//		case 3:
//			combo = FOUR_OF_A_KIND
//		case 2:
//			combo = THREE_OF_A_KIND
//		case 1:
//			combo = ONE_PAIR
//		}
//	}
//
//	return Hand{
//		Card:  card,
//		Bid:   bid,
//		Combo: combo,
//	}
//}
//
//func cardPower(c rune) int {
//	switch c {
//	case 'A':
//		return 14
//	case 'K':
//		return 13
//	case 'Q':
//		return 12
//	case 'J':
//		return 1
//	case 'T':
//		return 10
//	default:
//		return int(c - '0')
//	}
//}
//
//func toComboStr(combo Combo) string {
//	switch combo {
//	case FIVE_OF_A_KIND:
//		return "FIVE_OF_A_KIND"
//	case FOUR_OF_A_KIND:
//		return "FOUR_OF_A_KIND"
//	case FULL_HOUSE:
//		return "FULL_HOUSE"
//	case THREE_OF_A_KIND:
//		return "THREE_OF_A_KIND"
//	case TWO_PAIR:
//		return "TWO_PAIR"
//	case ONE_PAIR:
//		return "ONE_PAIR"
//	case HIGH_CARD:
//		return "HIGH_CARD"
//	default:
//		return "UNKNOWN"
//	}
//}
