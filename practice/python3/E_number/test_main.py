import unittest
import sys
import os

sys.path.append(os.path.abspath('..'))

from E_number import main

class TestMain(unittest.TestCase):

    def test_add_numbers(self):
        result = main.add_numbers(2, 3)
        self.assertEqual(result, 5)

    def test_subtract_numbers(self):
        result = main.subtract_numbers(5, 3)
        self.assertEqual(result, 2)

    def test_multiply_numbers(self):
        result = main.multiply_numbers(2, 3)
        self.assertEqual(result, 6)

    def test_divide_numbers(self):
        result = main.divide_numbers(6, 3)
        self.assertEqual(result, 2)


if __name__ == '__main__':
    unittest.main()
