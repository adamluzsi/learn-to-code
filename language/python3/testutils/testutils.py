import random
import string

def random_string(length):
    characters = string.ascii_letters + string.digits  # A-Z, a-z, 0-9
    return ''.join(random.choices(characters, k=length))

def random_integer(max_value):
    return random.randint(0, max_value)
