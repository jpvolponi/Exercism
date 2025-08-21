"""Functions for tracking poker hands and assorted card tasks.

Python list documentation: https://docs.python.org/3/tutorial/datastructures.html
"""


def get_rounds(number):
    """Create a list containing the current and next two round numbers.

    :param number: int - current round number.
    :return: list - current round and the two that follow.
    """
    rounds_list = [number, number + 1, number + 2]
    return rounds_list
    pass
#print(get_rounds(27))

def concatenate_rounds(rounds_1, rounds_2):
    """Concatenate two lists of round numbers.

    :param rounds_1: list - first rounds played.
    :param rounds_2: list - second set of rounds played.
    :return: list - all rounds played.
    """
    
    rounds_1.extend(rounds_2)
    return rounds_1
    pass
#print(concatenate_rounds([27, 28, 29], [35, 36]))


def list_contains_round(rounds, number):
    """Check if the list of rounds contains the specified number.

    :param rounds: list - rounds played.
    :param number: int - round number.
    :return: bool - was the round played?
    """
    contains = False
    for item in rounds:
        #print(item)
        if item == number:
            contains = True   
    return contains
    pass
#print(list_contains_round([27, 28, 29, 35, 36], 30))
#print(list_contains_round([27, 28, 29, 35, 36], 29))


def card_average(hand):
    """Calculate and returns the average card value from the list.

    :param hand: list - cards in hand.
    :return: float - average value of the cards in the hand.
    """
    if hand == []:
        return 0
    else:
        return sum(hand)/len(hand)
    pass
#print(card_average([5, 6, 7]))

def approx_average_is_average(hand):
    """Return if the (average of first and last card values) OR ('middle' card) == calculated average.

    :param hand: list - cards in hand.
    :return: bool - does one of the approximate averages equal the `true average`?
    """
    mediana = hand[(int(len(hand)/2))]
    media = card_average(hand)
    primeiro_e_ultimo = [hand[0], hand[len(hand)-1]]
    media_primeiro_e_ultimo = card_average(primeiro_e_ultimo)
    """print("Mediana:", mediana)
    print("Media:", media)
    print("Primeiro e ultimo:", primeiro_e_ultimo)
    print("Media Primeiro e ultimo:", media_primeiro_e_ultimo)"""

    if mediana == media or media_primeiro_e_ultimo == media:
        return True
    else:
        return False
pass


def average_even_is_average_odd(hand):
    """Return if the (average of even indexed card values) == (average of odd indexed card values).

    :param hand: list - cards in hand.
    :return: bool - are even and odd averages equal?
    """
    pares = []
    impares = []
    for index, item in enumerate(hand):
        if index % 2 == 0:
            pares.append(item)
        else:
            impares.append(item)
    if card_average(pares) == card_average(impares):
        return True
    else:
        return False

    pass



def maybe_double_last(hand):
    """Multiply a Jack card value in the last index position by 2.

    :param hand: list - cards in hand.
    :return: list - hand with Jacks (if present) value doubled.
    """
    new_hand = hand[:len(hand)-1]
    if hand[len(hand)-1] == 11:
        new_hand.append(22)
        return new_hand
    else:
        return hand

    pass
