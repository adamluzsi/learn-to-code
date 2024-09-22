import unittest
from main import MyClass

class TestMyClassMethods(unittest.TestCase):

    def test_class_initialization(self):
        instance = MyClass()
        self.assertIsInstance(instance, MyClass, 'Object must be an instance of MyClass')

    def test_method_call(self):
        instance = MyClass()
        input_and_output = 'input/output'
        result = instance.method_name(input_and_output)
        self.assertEqual(result, input_and_output, 'Method should process input and return expected output')


if __name__ == '__main__':
    unittest.main()
