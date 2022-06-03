"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""
import lasagna

# TODO: define the 'EXPECTED_BAKE_TIME' constant
lasagna.EXPECTED_BAKE_TIME = 40
# TODO: consider defining the 'PREPARATION_TIME' constant

#       equal to the time it takes to prepare a single layer


# TODO: define the 'bake_time_remaining()' function
def bake_time_remaining(time):
    return EXPECTED_BAKE_TIME - time


# pass


# TODO: define the 'preparation_time_in_minutes()' function
#       and consider using 'PREPARATION_TIME' here
def preparation_time_in_minutes(layers):
    return layers * 2


# TODO: define the 'elapsed_time_in_minutes()' function
def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    return preparation_time_in_minutes(number_of_layers) + (40 - bake_time_remaining(elapsed_bake_time))
