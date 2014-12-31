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
    def testNextCard(self):
        deck = Deck()
        cards = deck.getCards()
        deck.addCard(Card(1))
        deck.addCard(Card(2))
        deck.addCard(Card(3))
        self.assertEqual(len(cards),3)
        self.assertEqual(deck.nextCard(),Card(1))
        self.assertEqual(len(cards),2)
        self.assertEqual(deck.nextCard(),Card(2))
        self.assertEqual(len(cards),1)
        self.assertEqual(deck.nextCard(),Card(3))
        self.assertEqual(len(cards),0)
    def testGenerate(self):
        deck = Deck.generate()
        cards = deck.getCards()
        self.assertEqual(len(cards),44)
        i = 0
        for c in range(1,12):
            for e in range(1,5):
                self.assertEqual(cards[i].value,c)
                i += 1

