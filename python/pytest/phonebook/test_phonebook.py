# python -m unittest

import unittest

from phonebook import Phonebook

#declare a class containing tests
class PhonebookTest(unittest.TestCase):

    #test fixtures
    def setUp(self) -> None: #called first
        self.phonebook = Phonebook()

    def tearDown(self) -> None: #often unecessary unless external connections
        pass

    #tests: arrange, act, assert
    def test_lookup_by_name(self):
        self.phonebook.add("Bob", "12345")
        number = self.phonebook.lookup("Bob")
        self.assertEqual(number, "12345")

    def test_missing_name(self):
        with self.assertRaises(KeyError): #asserting and error
            self.phonebook.lookup("missing")

    def test_empty_phonebook_is_consistent(self):
        is_consistent = self.phonebook.is_consistent()
        self.assertTrue(is_consistent)

    def test_is_consistent_with_different_entries(self):
        self.phonebook.add("Bob", "12345")
        self.phonebook.add("Anna", "012345")
        self.assertTrue(self.phonebook.is_consistent())

    def test_inconsistent_with_duplicate_entries(self):
        self.phonebook.add("Bob", "12345")
        self.phonebook.add("Sue", "12345")
        self.assertFalse(self.phonebook.is_consistent())

    def test_inconsistent_with_duplicate_prefix(self):
        self.phonebook.add("Bob", "12345")
        self.phonebook.add("Sue", "123")
        self.assertFalse(self.phonebook.is_consistent())
