import unittest
import sys
import os

sys.path.append(os.path.abspath('..'))

from E_array import main
from testutils import testutils

class TestPractice(unittest.TestCase):
    random_value = testutils.random_string(5)

    def test_append(self):
        list = ["foo", "bar", "baz"]
        result = main.append(list, self.random_value)
        self.assertEqual(result, ["foo", "bar", "baz", self.random_value])

    def test_update(self):
        list = ["foo", "bar", "baz"]
        index = 1
        result = main.update(list, index, self.random_value)
        self.assertEqual(result, ["foo", self.random_value, "baz"])

    def test_remove(self):
        list = ["foo", "bar", self.random_value]
        index = 1
        result = main.remove(list, index)
        self.assertEqual(result, ["foo", self.random_value])

    def test_lookup_by_index(self):
        list = ["foo", "bar", "baz"]
        index = random.randint(0, len(list)-1)
        list[index] = self.random_value
        result = main.lookup_by_index(list, index)
        self.assertEqual(result, self.random_value)


if __name__ == '__main__':
    unittest.main()
