#### 二分查找

前提：
1. 有序，单调递增或者递减
2. 存在上下界, bounded
3. 能够通过索引访问(随机访问)

二分查找的难点在于: 明确搜索区间和搜索区间的变换

迭代也好，递归也罢，重复执行code片段的中的变量如何变换是关键问题，二分法类型题变换是**搜索区间**

```java 
//二分查找框架
public int binarySearch(int[] nums, int target) {
	int left = 0 , right = ...;

	while (...) {
		int mid = left + (right - left) >> 1;
		if (nums[mid] == target) {
			...
		} else if (nums[mid] < target) {
			left = ...
		} else if (nums[mid] > target) {
			right = ...
		}
	}
	return ...;

}
//寻找左侧边界二分搜索，解决目标元素重复问题(最左index) 坐闭右开
public int leftBound(int[] nums, int target) {
	if (nums.length == 0) return -1;
	int left = 0;
	int right = nums.length;//右开

	while (left < right) {
		int mid = left + (right - left) >> 1;
		if (num[mid] == target) {
			right == mid;
		} else if (nums[mid] < target) {
			left = mid + 1;
		} else if (nums[mid] > target) {
			right = mid;//
		}
	}
	//target > all element
	if (left == nums.length) return -1;

	return num[left] == target ? left : -1;
}
```