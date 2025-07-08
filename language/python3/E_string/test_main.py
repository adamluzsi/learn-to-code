import unittest
import random
import string
import sys
import os

sys.path.append(os.path.abspath('..'))

from E_string import main
from testutils import testutils

class TestStringMethods(unittest.TestCase):

    def test_concatenate_strings(self):
        str1 = testutils.random_string(5)
        str2 = testutils.random_string(5)
        result = main.concatenate_strings(str1, str2)
        self.assertEqual(result, str1 + str2)

    def test_find_substring(self):
        haystack = testutils.random_string(30)
        needle = haystack[10:20]
        result = main.find_substring(haystack, needle)
        self.assertEqual(result, 10)

    def test_convert_to_uppercase(self):
        s = testutils.random_string(5).lower()
        result = main.convert_to_uppercase(s)
        self.assertEqual(result, s.upper())

    def test_slice_string(self):
        s = testutils.random_string(30)
        start = random.randint(0, 15)
        end = random.randint(start, 30)
        result = main.slice_string(s, start, end)
        self.assertEqual(result, s[start:end])

    def test_replace_substring(self):
        s = testutils.random_string(30)
        old = s[5:15]
        new = testutils.random_string(10)
        result = main.replace_substring(s, old, new)
        self.assertIn(new, result)
        self.assertEqual(result.count(old), 0)


if __name__ == '__main__':
    unittest.main()
