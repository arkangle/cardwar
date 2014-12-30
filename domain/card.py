class Card:
    def __init__(self,value):
        self.value = value
    def __eq__(self,Other):
        return self.value == Other.value
    def __gt__(self,Other):
        return self.value > Other.value
    def __ge__(self,Other):
        return self.value >= Other.value
