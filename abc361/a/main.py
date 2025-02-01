N,K,X = input().split()
A = [int(a) for a in input().split()]
A.insert(int(K),int(X))
for a in A:
    print(a,end=" ")
