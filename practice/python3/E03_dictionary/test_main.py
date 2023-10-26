import unittest
import sys
import os

# Adjusting sys.path to import main and testutils
sys.path.append(os.path.abspath('..'))

from E03_dictionary import main
from testutils import testutils

class TestDictionaryOperations(unittest.TestCase):

    def test_add_entry(self):
        initial_dict = {"a": 1, "b": 2}
        result_dict = main.add_entry(initial_dict, "c", 3)
        self.assertEqual(result_dict, {"a": 1, "b": 2, "c": 3})

    def test_update_entry(self):
        initial_dict = {"a": 1, "b": 2}
        result_dict = main.update_entry(initial_dict, "b", 3)
        self.assertEqual(result_dict, {"a": 1, "b": 3})

    def test_remove_entry(self):
        initial_dict = {"a": 1, "b": 2, "c": 3}
        result_dict = main.remove_entry(initial_dict, "b")
        self.assertEqual(result_dict, {"a": 1, "c": 3})

    def test_find_value_by_key(self):
        initial_dict = {"a": 1, "b": 2, "c": 3}
        random_key = testutils.random_string(5)
        random_value = testutils.random_string(5)
        extended_dict = {**initial_dict, random_key: random_value}
        found_value = main.find_value_by_key(extended_dict, random_key)
        self.assertEqual(found_value, random_value)


if __name__ == '__main__':
    unittest.main()
