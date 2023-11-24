package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := NewPlayingCardDeck()
	card := deck.RandomCard()
	playCard, ok := card.(*PlayingCard)
	if ok {
		ReadCard(playCard)
	} else {
		fmt.Println("The card type is invalid")
	}

}

// PlayingCard making playingCard
type PlayingCard struct {
	Suit string
	Rank string
}

func NewPlayingCard(suit string, rank string) *PlayingCard {
	return &PlayingCard{Suit: suit, Rank: rank}
}
func ReadCard(playCard *PlayingCard) {
	fmt.Printf("%s of %s\n", playCard.Rank, playCard.Suit)
}

//--------------------------------------------------------------------

// TradeCard making tradeCard
type TradeCard struct {
	Point string
	Sign  string
}

func NewTradeCard(point string, sign string) *TradeCard {
	return &TradeCard{Point: point, Sign: sign}
}
func ReadTradeCard(tradeCard *TradeCard) {
	fmt.Printf("%s of %s\n", tradeCard.Point, tradeCard.Sign)
}

//--------------------------------------------------------------------

type Deck struct {
	cards []interface{}
}

// NewPlayingCardDeck creating new deck
func NewPlayingCardDeck() (deck *Deck) {
	suits := []string{"Diamonds", "Spades", "Clubs", "Hearts"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Q", "K", "J"}
	deck = &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.AddCard(NewPlayingCard(suit, rank))
		}
	}
	deck.AddCard(NewTradeCard("15", "Sword"))
	return
}

func (deck *Deck) AddCard(card interface{}) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) RandomCard() interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cardIdx := r.Intn(len(deck.cards))
	return deck.cards[cardIdx]
}
