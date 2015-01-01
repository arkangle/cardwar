class Player:
    def __init__(self,name,hand):
        self.name = name
        self.Hand = hand
    def getName(self):
        return self.name
    def getHand(self):
        return self.Hand

class Players:
    def __init__(self,players):
        self.players = players
        self.position = -1
        self.count = len(self.players)
    def nextPlayer(self):
        self.position += 1
        if(self.count == self.position):
            self.position = 0
        return self.players[self.position]
