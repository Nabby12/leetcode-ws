import pytest
from main import Solution


def test_threeSum():
    test_cases = [
        {"nums": [-1, 0, 1, 2, -1, -4], "want": [[-1, -1, 2], [-1, 0, 1]]},
        {"nums": [0, 1, 1], "want": []},
        {"nums": [0, 0, 0], "want": [[0, 0, 0]]},
    ]

    solution = Solution()
    for case in test_cases:
        assert solution.threeSum(case["nums"]) == case["want"]
