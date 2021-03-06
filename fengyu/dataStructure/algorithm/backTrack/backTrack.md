## 回溯法
1. 回溯法是一种搜索方式，也是递归，常用于解决**组合、排列**问题。
2. 回溯法性能较差，回溯的本质是穷举。
3. 解决一个回溯问题，其实就是**遍历决策树**

- 路径: 做过的选择
- 选择列表：当前的选择
- 结束条件：到达决策树底层，无法做出选择的条件

```java
//全排列
List<List<Integer>> res = new ArrayList();
public List<List<Integer>> permute (int[] nums) {
	Deque<Integer> track = new LinkedList<>();
	dfs(track, nums);
	return res;
}
private void dfs (Deque<Integer> track, int[] nums) {
	if (track.size() == nums.length) {
		res.add(new ArrayList(track));
		return;
	}
	for (int num : nums) {
		//选择
		if (track.contains(num)) {
			continue;
		}
		track.add(num);
		dfs(track, nums);
		track.removeLast();
	}
}
/**
 全排列：每个返回值中的元素都与原数组元素一致，这就是结束的条件，路径长度==数组长度。
 数组的长度决定了决策树的高度，也同时决定路径的长度。
**/


//子集
List<List<Integer>> res = new ArrayList();
public List<List<Integer>> subSets(int[] nums) {
	Deque<Integer> track = new LinkedList<>();
	dfs(track, 0, nums);
	return res;
}

private void dfs (Deque<Integer> track, int start, int[] nums) {
	res.add(new ArrayList(track));
	for (int i = start; i < nums.length; i++) {
		track.add(nums[i]);
		dfs(track, i + 1, nums);
		track.removeLast();
	}
}

//组合
List<List<Integer>> res = new ArrayList<>();
public List<List<Integer>> combine (int n, int k) {
	Deque<Integer> track = new LinkedList<>();
}
private void dfs (Deque<Integer> track, int start, int n ,int k) {
	if (track.size() == k) {
		res.add(new ArrayList(track));
	}
	for (int i = start; i <= n; i++) {
		track.add(i);
		dfs(track, i + 1, n , k);
		track.removeLast();
	}
}
```

#### 剪枝

可以剪枝的地⽅就在递归中每⼀层的for循环所选择的起始位置。
如果for循环选择的起始位置之后的元素个数 已经不⾜ 我们需要的元素个数了，那么就没有必要搜索
了。
处理过程和回溯过程是一一对应

回溯法抽象为决策树结构后，其遍历过程就是：for循环横向遍历，递归纵向遍历，回溯不断调整结果集。

排列问题树的深度是固定的，组合问题树的深度是可变的。

836080779