package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name  string
	Rank  int
	Suit  string
	Bonus int
}

type Deck struct {
	Cards []Card
}

type Card struct {
	Rank  int
	Suit  string
	Value int
}

type Game struct {
	Players []Player
	Deck    Deck
	Tricks  []Trick
}

type Trick struct {
	Leader  Player
	Cards   []Card
	Winner  Player
	Bonus   int
	Bonuses []Bonus
}

type Bonus struct {
	Rank  int
	Suit  string
	Value int
}

func main() {
	// Initialize players
	players := []Player{
		{"Player 1", 1, "Hearts", 0},
		{"Player 2", 2, "Spades", 0},
		{"Player 3", 3, "Diamonds", 0},
		{"Player 4", 4, "Clubs", 0},
	}

	// Remove ranks 2-5 from deck
	deck := Deck{
		Cards: []Card{
			{1, "Hearts", 1},
			{2, "Hearts", 2},
			{3, "Hearts", 3},
			{4, "Hearts", 4},
			{5, "Hearts", 5},
			{6, "Hearts", 6},
			{7, "Hearts", 7},
			{8, "Hearts", 8},
			{9, "Hearts", 9},
			{10, "Hearts", 10},
			{11, "Hearts", 11},
			{12, "Hearts", 12},
			{13, "Hearts", 13},
			{14, "Hearts", 14},
			{15, "Hearts", 15},
			{16, "Hearts", 16},
			{17, "Hearts", 17},
			{18, "Hearts", 18},
			{19, "Hearts", 19},
			{20, "Hearts", 20},
			{21, "Hearts", 21},
			{22, "Hearts", 22},
			{23, "Hearts", 23},
			{24, "Hearts", 24},
			{25, "Hearts", 25},
			{26, "Hearts", 26},
			{27, "Hearts", 27},
			{28, "Hearts", 28},
			{29, "Hearts", 29},
			{30, "Hearts", 30},
			{31, "Hearts", 31},
			{32, "Hearts", 32},
			{33, "Hearts", 33},
			{34, "Hearts", 34},
			{35, "Hearts", 35},
			{36, "Hearts", 36},
			{37, "Hearts", 37},
			{38, "Hearts", 38},
			{39, "Hearts", 39},
			{40, "Hearts", 40},
			{41, "Hearts", 41},
			{42, "Hearts", 42},
			{43, "Hearts", 43},
			{44, "Hearts", 44},
			{45, "Hearts", 45},
			{46, "Hearts", 46},
			{47, "Hearts", 47},
			{48, "Hearts", 48},
			{49, "Hearts", 49},
			{50, "Hearts", 50},
		},
	}

	// Shuffle deck
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})

	// Deal cards to players
	for i := 0; i < len(players); i++ {
		players[i].Rank = deck.Cards[i].Rank
		players[i].Suit = deck.Cards[i].Suit
	}

	// Remove cards from deck
	deck.Cards = deck.Cards[len(players):]

	// Initialize game
	game := Game{
		Players: players,
		Deck:    deck,
		Tricks:  []Trick{},
	}

	// Play game
	for {
		// Select dealer
		dealer := game.Players[rand.Intn(len(game.Players))]

		// Lead player is player to left of dealer
		leadPlayer := game.Players[dealer.Rank-1]

		// Dealer shuffles cards
		rand.Shuffle(len(game.Deck.Cards), func(i, j int) {
			game.Deck.Cards[i], game.Deck.Cards[j] = game.Deck.Cards[j], game.Deck.Cards[i]
		})

		// Dealer deals cards to players
		for i := 0; i < len(game.Players); i++ {
			game.Players[i].Rank = game.Deck.Cards[i].Rank
			game.Players[i].Suit = game.Deck.Cards[i].Suit
		}

		// Remove cards from deck
		game.Deck.Cards = game.Deck.Cards[len(game.Players):]

		// Bidding
		for i := 0; i < len(game.Players); i++ {
			player := game.Players[i]
			if player.Rank == 1 {
				player.Bid = "PASS"
			} else {
				player.Bid = "HEARTS"
			}
		}

		// SUIT selected is now TRUMP
		trumpSuit := "HEARTS"

		// Final cards are dealt out to each player
		for i := 0; i < len(game.Players); i++ {
			player := game.Players[i]
			player.Cards = []Card{
				{player.Rank, player.Suit, player.Value},
				{player.Rank + 1, player.Suit, player.Value + 1},
				{player.Rank + 2, player.Suit, player.Value + 2},
			}
		}

		// New trick
		for i := 0; i < len(game.Players); i++ {
			player := game.Players[i]
			trick := Trick{
				Leader: player,
				Cards:  []Card{},
			}
			game.Tricks = append(game.Tricks, trick)
		}

		// Play tricks
		for i := 0; i < len(game.Tricks); i++ {
			trick := game.Tricks[i]
			for j := 0; j < len(trick.Cards); j++ {
				card := trick.Cards[j]
				if card.Suit == trumpSuit {
					trick.Winner = trick.Leader
					break
				}
			}
		}

		// Score players
		for i := 0; i < len(game.Players); i++ {
			player := game.Players[i]
			player.Score = player.Rank + player.Bonus
		}

		// Check for game over
		if game.Score >= 500 {
			break
		}
	}

	// Print final scores
	fmt.Println("Final scores:")
	for i := 0; i < len(game.Players); i++ {
		player := game.Players[i]
		fmt.Printf("%s: %d\n", player.Name, player.Score)
	}
}
