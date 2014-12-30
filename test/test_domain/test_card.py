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
