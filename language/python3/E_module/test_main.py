import unittest
import sys
import os
import random
from types import ModuleType
from unittest.mock import patch

sys.path.append(os.path.abspath('..'))

from E_module import main
from testutils import testutils


class TestModules(unittest.TestCase):
    random_value = testutils.random_string(5)

    def test_import_module(self):
        module_name = 'json'
        module = main.import_module(module_name)
        self.assertIsInstance(module, ModuleType, 'The imported module is not an instance of ModuleType')

    def test_call_function(self):
        module = main.import_module('math')
        function_name = 'sqrt'
        result = main.call_function(module, function_name, 16)
        self.assertEqual(result, 4, 'The call_function did not return the correct result for math.sqrt')

    @patch('importlib.reload')
    def test_reload_module(self, mock_reload):
        module = main.import_module('json')
        returned_module = main.reload_module(module)
        mock_reload.assert_called_with(module)
        self.assertEqual(returned_module, module, 'The reload_module function did not return the correct module')

    def test_create_module(self):
        module_name = 'dynamic_module'
        functions = {
            'greet': lambda name: f'Hello, {name}!'
        }
        module = main.create_module(module_name, functions)
        self.assertIsInstance(module, ModuleType, 'The create_module function did not return a ModuleType instance')
        self.assertTrue(hasattr(module, 'greet'), 'The dynamically created module does not have the greet function')
        self.assertEqual(module.greet(self.random_value), f'Hello, {self.random_value}!', 'The greet function did not return the correct result')


if __name__ == '__main__':
    unittest.main()
