## 数组

数组算法题的几种常见解法

1. 二分法：解决**有序**数组的指定元素检索  
2. 双指针法：删除、替换数组中的元素的方法   No.35 

3. 滑动窗口：查找数组中**子序列(连续元素)**问题，滑动窗口属于双指针中的特例
滑动窗口：就是不断的调节子序列的起始位置和终止位置，从而得出我们要想的结果
No.209

```java
//数组中两个元素交换位置
void swap(int[] arr, int l, int r){
	//i,r < arr.length
	int temp = arr[l];
	arr[l] = arr[r];
	arr[r] = temp; 
}
```