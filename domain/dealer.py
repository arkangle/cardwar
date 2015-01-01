import random

class Dealer:
    def __init__(self):
        pass
    def deal(self,deck,players):
        while deck.hasCards():
            player = players.nextPlayer()
            card = deck.nextCard()
            player.getHand().addCard(card)
    def shuffle(self,deck):
        random.shuffle(deck.getCards())
