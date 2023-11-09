import unittest
import sys
import os

# Adjusting sys.path to import main and testutils
sys.path.append(os.path.abspath('..'))

from E03_dictionary import main
from testutils import testutils

class TestDictionaryOperations(unittest.TestCase):

    random_key = testutils.random_string(5)
    random_value = testutils.random_integer(42)

    def test_add_entry(self):
        initial_dict = {"a": 1, "b": 2}
        result_dict = main.add_entry(initial_dict, self.random_key, self.random_value)
        self.assertEqual(result_dict, {"a": 1, "b": 2, self.random_key: self.random_value})

    def test_update_entry(self):
        initial_dict = {"a": 1, "b": 2}
        result_dict = main.update_entry(initial_dict, "b", self.random_value)
        self.assertEqual(result_dict, {"a": 1, "b": self.random_value})

    def test_remove_entry(self):
        initial_dict = {"a": 1, "b": 2, "c": 3}
        result_dict = main.remove_entry(initial_dict, "b")
        self.assertEqual(result_dict, {"a": 1, "c": 3})

    def test_find_value_by_key(self):
        initial_dict = {"a": 1, "b": 2, "c": 3}
        extended_dict = {**initial_dict, self.random_key: self.random_value}
        found_value = main.find_value_by_key(extended_dict, self.random_key)
        self.assertEqual(found_value, self.random_value)


if __name__ == '__main__':
    unittest.main()
