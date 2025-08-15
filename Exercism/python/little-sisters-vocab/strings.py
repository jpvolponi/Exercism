"""Functions for creating, transforming, and adding prefixes to strings."""


def add_prefix_un(word):
    """Take the given word and add the 'un' prefix.

    :param word: str - containing the root word.
    :return: str - of root word prepended with 'un'.
    """
    return "un" + word
    pass
#print(add_prefix_un("happy"))

def make_word_groups(vocab_words):
    """Transform a list containing a prefix and words into a string with the prefix followed by the words with prefix prepended.

    :param vocab_words: list - of vocabulary words with prefix in first index.
    :return: str - of prefix followed by vocabulary words with
            prefix applied.

    This function takes a `vocab_words` list and returns a string
    with the prefix and the words with prefix applied, separated
     by ' :: '.

    For example: list('en', 'close', 'joy', 'lighten'),
    produces the following string: 'en :: enclose :: enjoy :: enlighten'.
    """
    prefixo = vocab_words[0]
    palavras_com_prefixo = [prefixo + palavra for palavra in vocab_words[1:]]
    lista_resultado = [prefixo] + palavras_com_prefixo
    return ' :: '.join(lista_resultado)
    pass

#print(make_word_groups(['en', 'close', 'joy', 'lighten']))

def remove_suffix_ness(word):
    """Remove the suffix from the word while keeping spelling in mind.

    :param word: str - of word to remove suffix from.
    :return: str - of word with suffix removed & spelling adjusted.

    For example: "heaviness" becomes "heavy", but "sadness" becomes "sad".
    """
    new_word = None
    
    if word[-5] == "i":
        palavara_sem_sufixo = (word[:-5]+"y")
    else:
        palavara_sem_sufixo = word[:-4]
       
    return palavara_sem_sufixo
    
    pass
#print(remove_suffix_ness("heaviness"))

def adjective_to_verb(sentence, index):
    """Change the adjective within the sentence to a verb.

    :param sentence: str - that uses the word in sentence.
    :param index: int - index of the word to remove and transform.
    :return: str - word that changes the extracted adjective to a verb.

    For example, ("It got dark as the sun set.", 2) becomes "darken".
    """

    pass
