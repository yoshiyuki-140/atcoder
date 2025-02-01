import itertools
from multiprocessing import Pool


def f(pair):
    a, b = pair
    return 1 if a <= b / 2 else 0


if __name__ == "__main__":
    # N = int(input())
    # A = list(map(lambda x: int(x), input().split()))
    N = 5 * (10**5)
    A = [i for i in range(N)]

    with Pool() as pool:
        results = sum(pool.map(f, itertools.permutations(A, r=2)))

    print(results)
