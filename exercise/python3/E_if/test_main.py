import unittest
import sys
import os

# Adjusting sys.path to import main and testutils
sys.path.append(os.path.abspath('..'))

from E_if import main
from testutils import testutils

class TestSolutions(unittest.TestCase):

    def test_solution1(self):
        self.assertEqual(main.solution1(""), "empty")
        self.assertEqual(main.solution1(testutils.random_string(5)), "not-empty")

    def test_solution2(self):
        self.assertEqual(main.solution2("flash"), "thunder")
        self.assertEqual(main.solution2("foo"), "baz")
        input = testutils.random_string(5)
        self.assertEqual(main.solution2(input), input)

if __name__ == '__main__':
    unittest.main()
