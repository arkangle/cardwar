class Deck:
    def __init__(self):
        self.cards = []
    def addCard(self,card):
        self.cards.append(card)
    def getCards(self):
        return self.cards

