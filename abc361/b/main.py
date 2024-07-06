a, b, c, d, e, f = [int(i) for i in input().split()]
g, h, i, j, k, l = [int(i) for i in input().split()]
A_xmin, A_xmax, A_ymin, A_ymax, A_zmin, A_zmax = a, d, b, e, c, f
B_xmin, B_xmax, B_ymin, B_ymax, B_zmin, B_zmax = g, j, h, k, j, l


def intersecting_volume(A, B):
    # 各軸の重なり範囲を計算
    xmin_common = max(A["xmin"], B["xmin"])
    xmax_common = min(A["xmax"], B["xmax"])
    ymin_common = max(A["ymin"], B["ymin"])
    ymax_common = min(A["ymax"], B["ymax"])
    zmin_common = max(A["zmin"], B["zmin"])
    zmax_common = min(A["zmax"], B["zmax"])

    # 重なり範囲が存在するか確認
    if (
        xmin_common < xmax_common
        and ymin_common < ymax_common
        and zmin_common < zmax_common
    ):
        # 共通部分の体積を計算
        volume = (
            (xmax_common - xmin_common)
            * (ymax_common - ymin_common)
            * (zmax_common - zmin_common)
        )
    else:
        volume = 0

    return volume


# 例: 直方体AとBの定義
A = {
    "xmin": A_xmin,
    "xmax": A_xmax,
    "ymin": A_ymin,
    "ymax": A_ymax,
    "zmin": A_zmin,
    "zmax": A_zmax,
}
B = {
    "xmin": B_xmin,
    "xmax": B_xmax,
    "ymin": B_ymin,
    "ymax": B_ymax,
    "zmin": B_zmin,
    "zmax": B_zmax,
}


# 共通部分の体積を計算
volume = intersecting_volume(A, B)
if volume <= 0:
    print("No")
else:
    print("Yes")
