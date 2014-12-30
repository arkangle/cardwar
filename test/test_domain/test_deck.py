import unittest
from domain.deck import Deck
from domain.card import Card

class TestDeck(unittest.TestCase):
    def testGetCard(self):
        deck = Deck()
        self.assertEqual(deck.getCards(),[])
        deck.addCard(Card(1))
        self.assertEqual(deck.getCards(),[Card(1)])
    def testAddCard(self):
        deck = Deck()
        deck.addCard(Card(1))
        deck.addCard(Card(2))
        deck.addCard(Card(3))
        self.assertEqual(deck.getCards(),[Card(1),Card(2),Card(3)])
