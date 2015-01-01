import unittest
from domain.deck import Deck,Hand
from domain.card import Card
from domain.dealer import Dealer
from domain.player import Player,Players

class TestDealer(unittest.TestCase):
    def testDeal(self):
        deck = Deck.generate(4,2)
        players = Players([Player("Bob",Hand()), Player("Sue",Hand())])
        dealer = Dealer()
        dealer.deal(deck,players)
        bob = players.nextPlayer()
        bob_hand = bob.getHand()
        self.assertEqual(Card(1),bob_hand.nextCard())
        self.assertEqual(Card(2),bob_hand.nextCard())
        self.assertEqual(Card(3),bob_hand.nextCard())
        self.assertEqual(Card(4),bob_hand.nextCard())
        sue = players.nextPlayer()
        sue_hand = sue.getHand()
        self.assertEqual(Card(1),sue_hand.nextCard())
        self.assertEqual(Card(2),sue_hand.nextCard())
        self.assertEqual(Card(3),sue_hand.nextCard())
        self.assertEqual(Card(4),sue_hand.nextCard())
