---

date: 2025-04-01T15:23:58+0800
title: 学习坐标变换
layout: post

---

```python
import matplotlib.pyplot as plt

# 创建图形和坐标轴
fig, ax = plt.subplots()

# 设置坐标轴范围
ax.set_xlim(-6, 6)
ax.set_ylim(-6, 6)

# 将坐标轴移动到中心
ax.spines["left"].set_position("center")
ax.spines["bottom"].set_position("center")
ax.spines["right"].set_visible(False)
ax.spines["top"].set_visible(False)

# 设置刻度标记（显示所有整数）
ax.set_xticks(range(-6, 6))
ax.set_yticks(range(-6, 6))

# 绘制半径为5的圆
circle = plt.Circle((0, 0), 5, color="blue", fill=False)
ax.add_patch(circle)

# 绘制圆上的点(3,4)
ax.plot(3, 4, "ro")  # 红色圆点
# 添加坐标值标注
ax.annotate(
    "(3,4)",
    xy=(3, 4),
    xytext=(3.5, 4.5),
    # arrowprops=dict(facecolor="black", shrink=0.05),
    bbox=dict(boxstyle="round,pad=0.5", fc="yellow", alpha=0.5),
)

# 绘制到坐标轴的虚线
ax.plot([3, 3], [4, 0], "--", color="gray")  # 垂直虚线到x轴
ax.plot([3, 0], [4, 4], "--", color="gray")  # 水平虚线到y轴

# 绘制圆上的点(-3,-4)
ax.plot(-3, -4, "ro")  # 红色圆点

# 绘制到坐标轴的虚线
ax.plot([-3, -3], [-4, 0], "--", color="gray")  # 垂直虚线到x轴
ax.plot([-3, 0], [-4, -4], "--", color="gray")  # 垂直虚线到x轴

ax.plot([-3, 3], [-4, 4], "--", color="gray")  # 垂直虚线到x轴


# 设置等比例显示
ax.set_aspect("equal")

# 显示图形
plt.show()
```

![](https://ohsaisai.oss-cn-shanghai.aliyuncs.com/2025/04/Figure_1.png)

想给赛赛讲一下坐标变换的，比如向左旋转90度，坐标值是怎么变换的。
