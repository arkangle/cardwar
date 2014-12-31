from domain.card import Card
import random

class Deck:
    @staticmethod
    def generate(count=11,each=4):
        deck = Deck()
        for c in range(1,count+1):
            for e in range(1,each+1):
                deck.addCard(Card(c))
        return deck
    def __init__(self):
        self.cards = []
    def addCard(self,card):
        self.cards.append(card)
    def getCards(self):
        return self.cards
    def nextCard(self):
        return self.cards.pop(0)
    def shuffle(self):
        random.suffle(self.cards)

