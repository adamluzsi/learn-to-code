import unittest
import random
import sys
import os

sys.path.append(os.path.abspath('..'))

from E_serialization import main
from testutils import testutils


class TestSerialization(unittest.TestCase):

    def test_serialization(self):
        number = random.randint(1, 42)
        text = testutils.random_string(10)
        data = main.Data(number, text)
        encoded = main.marshal(data)
        self.assertIsInstance(encoded, str)
        got = main.unmarshal(encoded)
        self.assertIsInstance(got, main.Data)
        self.assertEqual(data.number, got.number)
        self.assertEqual(data.text, got.text)


if __name__ == '__main__':
    unittest.main()
