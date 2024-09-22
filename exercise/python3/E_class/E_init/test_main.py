import unittest
import sys
import os
import random

sys.path.append(os.path.abspath('../..'))

from E_init import main


class TestMyClassInitialization(unittest.TestCase):
    def test_initialization(self):
        test_value = random.randint(1, 100)
        my_obj = main.MyClass(test_value)
        self.assertEqual(my_obj.get_value(), test_value)


if __name__ == '__main__':
    unittest.main()
