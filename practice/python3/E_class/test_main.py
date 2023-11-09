import unittest
import sys
import os

# Adjusting sys.path to import main and testutils
sys.path.append(os.path.abspath('..'))

from E_class import main
from testutils import testutils

class TestSolutions(unittest.TestCase):

    def test_solution(self):
        random_value = testutils.random_string(5)
        instance = main.solution(random_value)
        self.assertIsInstance(instance, main.MyClass)
        self.assertEqual(instance.my_field, random_value)

if __name__ == '__main__':
    unittest.main()
