import unittest
from domain.player import Player,Players
from domain.deck import Hand

class TestPlayer(unittest.TestCase):
    def testGetName(self):
        player = Player("Bob",Hand())
        self.assertEqual(player.getName(),"Bob")
    def testGetHand(self):
        hand = Hand()
        player = Player("Bob",hand)
        self.assertEqual(player.getHand(),hand)

class TestPlayers(unittest.TestCase):
    def testNextPlayer(self):
        player_list = []
        player_list.append(Player("Bob",Hand()))
        player_list.append(Player("Sue",Hand()))
        player_list.append(Player("Rick",Hand()))
        players = Players(player_list)
        self.assertEqual("Bob",players.nextPlayer().getName())
        self.assertEqual("Sue",players.nextPlayer().getName())
        self.assertEqual("Rick",players.nextPlayer().getName())
        self.assertEqual("Bob",players.nextPlayer().getName())
        self.assertEqual("Sue",players.nextPlayer().getName())
        self.assertEqual("Rick",players.nextPlayer().getName())
        self.assertEqual("Bob",players.nextPlayer().getName())


