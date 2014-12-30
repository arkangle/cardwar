import unittest
from domain.card import Card

class TestCard(unittest.TestCase):
    def testCardEqual(self):
        card1 = Card(1)
        card2 = Card(1)
        self.assertEqual(card1,card2)
    def testCardNotEqual(self):
        card1 = Card(1)
        card2 = Card(2)
        self.assertNotEqual(card1,card2)
    def testCardGt(self):
        card1 = Card(1)
        card2 = Card(2)
        self.assertTrue(card2>card1)
    def testCardLt(self):
        card1 = Card(1)
        card2 = Card(2)
        self.assertTrue(card1<card2)
    def testCardGtE(self):
        card1 = Card(1)
        card2 = Card(2)
        card3 = Card(2)
        self.assertTrue(card2>=card1)
        self.assertTrue(card2>=card3)
    def testCardLtE(self):
        card1 = Card(1)
        card2 = Card(2)
        card3 = Card(2)
        self.assertTrue(card1<=card2)
        self.assertTrue(card2<=card3)
