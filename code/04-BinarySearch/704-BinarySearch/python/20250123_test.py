import pytest
from main import Solution


def test_search():
    test_cases = [
        {"nums": [-1, 0, 3, 5, 9, 12], "target": 9, "want": 4},
        {"nums": [-1, 0, 3, 5, 9, 12], "target": 2, "want": -1},
    ]

    solution = Solution()
    for case in test_cases:
        assert solution.search(case["nums"], case["target"]) == case["want"]
