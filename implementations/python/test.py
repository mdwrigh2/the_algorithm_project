#!/usr/bin/env python
import sys
import unittest

from algorithms import test_sorting, test_array
from data_structures import test_heaps

loader = unittest.TestLoader()

data_structures = {}
data_structures["heaps"] = loader.loadTestsFromModule(test_heaps)

algorithms = {}
algorithms["sorting"] = loader.loadTestsFromModule(test_sorting)
algorithms["array"] = loader.loadTestsFromModule(test_array)

data_structures_suite = unittest.TestSuite()
for suite in data_structures.values():
    data_structures_suite.addTests(suite)

algorithms_suite = unittest.TestSuite()
for suite in algorithms.values():
    algorithms_suite.addTests(suite)

all_suite = unittest.TestSuite()
all_suite.addTests(algorithms_suite)
all_suite.addTests(data_structures_suite)

suites = {"all": all_suite, "algorithms": algorithms_suite, "data_structures": data_structures_suite}
suites.update(algorithms)
suites.update(data_structures)

if __name__ == "__main__":
    if len(sys.argv) > 1:
        suite = suites[sys.argv[1]]
        if not suite:
            print "Error: No such test suite"
            sys.exit(1)
    else:
        suite = all_suite
    unittest.TextTestRunner(verbosity=2).run(suite)
