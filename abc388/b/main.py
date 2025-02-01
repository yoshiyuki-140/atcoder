N, D = list(map(lambda x: int(x), input().split()))
T, L = [], []
for i in range(N):
    tmp = input().split()
    T.append(tmp[0])
    L.append(tmp[1])

# 整数化
T, L = list(map(lambda x: int(x), T)), list(map(lambda x: int(x), L))

# 結果を格納する変数
results = []

# 重さを導き出している
for i in range(D):
    # 長さを1増やす
    L = list(map(lambda x: x + 1, L))
    # 重さの最大値を求める
    weight = 0
    for t, l in zip(T, L):
        result = t * l
        if weight <= result:
            weight = result
    results.append(weight)


# print処理
for r in results:
    print(r)
