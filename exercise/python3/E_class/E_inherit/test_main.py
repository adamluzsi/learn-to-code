import unittest
import sys
import os

sys.path.append(os.path.abspath('..'))

from E_inherit import main

class TestInheritance(unittest.TestCase):

    def test_dog(self):
        dog = main.Dog()
        self.assertEqual(dog.speak(), 'Woof')
        dog.walk()

    def test_cat(self):
        cat = main.Cat()
        self.assertEqual(cat.speak(), 'Meow')
        cat.walk()

if __name__ == '__main__':
    unittest.main()
