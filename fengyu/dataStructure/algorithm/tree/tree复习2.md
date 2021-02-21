## 二叉树专题
写递归算法的关键是要明确函数的「定义」是什么，然后相信这个定义，利用这个定义推导最终结果，绝不要跳入递归的细节。
写树相关的算法，简单说就是，先搞清楚当前 root 节点该做什么，然后根据函数定义递归调用子节点
二叉树题目的一个难点就是，如何把题目的要求细化成每个节点需要做的事情。

#### 关键点
1. 确定每个节点应该做的事情，剩下的交给前中后序遍历。
2. 确定使用哪种前中后序遍历，明确递归终止条件，明确方法返回值。

## 例题
1. 计算二叉树结点数
```java
public void count(TreeNode root) {
	if (root == null) retur 0;
	//后序遍历
	// int leftCount = count(root.left);
	// int rightCount = count(root.right);
	// return 1 + leftCount + rightCount;
	//简化版
	return 1 + count(root.left) + count(root.right);
}
```
2. 翻转二叉树 leetcode 226
```java
//通过观察，我们发现只要把二叉树上的每一个节点的左右子节点进行交换，最后的结果就是完全翻转之后的二叉树。
public TreeNode invertTree(TreeNode root) {
	if (root == null) return root;
	//左右结点互换
	TreeNode tempNode = root.left;
	root.left = root.right;
	root.right = tempNode;
	invertTree(root.left);
	invertTree(root.right);
	return root;
}
```
3. 填充二叉树节点的右侧指针 leetcode 116 redo
```java
public void connect(Node node) {
	if (root == null) return root;
	connect(root.left, root.right);
	return root;
}
public void connect(Node node1, Node node2) {//同一层的两个node
	if (node1 == null || node2 == null) {
		return ;
	}
	node1.next = node2;
	//同一父节点
	connect(node1.left, node1.right);
	connect(node2.left, node2.right);
	//不同父节点
	connect(node1.right, node2.left);
}
```
3. 二叉树展开为链表 leetcode 114 redo
```java
public void flatten(TreeNode root) {
	if (root == null) return ;
	flatten(root.left);
	flatten(root.right);
	//后序遍历
	TreeNode temp = root.right;//右节点暂存
	root.right = root.left;//右节点替换为左节点
	root.left = null;//
	TreeNode p = root;
	while (p.right != null) {
		p = p.right;
	}
	p.right = temp;
}
```
4. 构造最大二叉树 leetcode 654
```java 
public TreeNode constructMaximumBinaryTree(int[] nums) {
	return construct(nums, 0, nums.length);
}	
private TreeNode construct(int[] nums, int l, int r) {
	if (l == r) return null;
	//先序遍历
	int index = getIndex(nums, l, r);
	TreeNode root = new TreeNode(nums[index]);
	root.left = construct(nums, l, index);
	root.right = construct(nums, index + 1, r);
	return root;
}
private int getIndex(int[] nums, int l, int r) {
	int index = l;
	for (int i = l + 1; i < r; i++) {
		if (nums[i] > nums[index]) {
			index = i;
		}
	}
	return index;
}
```
5. 通过先序和中序遍历构造二叉树 leetcode 105
```java
public TreeNode buildTree(int[] preorder, int[] inorder) {
	return buildTree(preorder, inorder, 0, preorder.length, 0 ,inorder.length);
}
private TreeNode buildTree(int[] preorder, int[] inorder, int pl, int pr, int il, int ir) {
	if (pl == pr) return null;
	int val = preorder[pl];
	TreeNode root = new TreeNode(val);
	int iIndex = getIndex(inorder, il, ir, val);
	int pIndex = pl + iIndex - il;
	root.left = buildTree(preorder, inorder, pl + 1, pIndex + 1 , il, iIndex);
	root.right = buildTree(preorder, inorder, pIndex + 1, pr, iIndex + 1, ir);
	return root;
}
private int getIndex(int[] nums, int l, int r, int val) {
	for (int i = l; i < r; i++) {
		if (nums[i] == val) {
			return i;
		}
	}
	return -1;
}
```
6.通过中序和后序遍历构造二叉树 leetcode 106
```java
public TreeNode buildTree(int[] inorder, int[] postorder) {
	return buildTree(postorder, inorder, 0, postorder.length - 1, 0 ,inorder.length - 1);
}
private TreeNode buildTree(int[] postorder, int[] inorder, int pl, int pr, int il, int ir) {
	if (pl > pr) return null;
	int val = postorder[pr];
	TreeNode root = new TreeNode(val);
	int iIndex = getIndex(inorder, il, ir, val);
	int pIndex = pr + iIndex - ir;
	root.left = buildTree(postorder, inorder, pl, pIndex - 1 , il, iIndex - 1);
	root.right = buildTree(postorder, inorder, pIndex, pr - 1, iIndex + 1, ir);
	return root;
}
private int getIndex(int[] nums, int l, int r, int val) {
	for (int i = l; i <= r; i++) {
		if (nums[i] == val) {
			return i;
		}
	}
	return -1;
}
```
7. 寻找重复子树 leetcode 652
```java
//如何判断树重复 序列化结合Map
List<TreeNode> res = new ArrayList<>();
Map<String,Integer> map = new HashMap<>();
public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
	dfs(root);	
	return res;
}
private String dfs(TreeNode root) {
	if (root == null) return "#";
	// String subTree = dfs(root.left) + "," + root.val + "," + dfs(root.right);
	String subTree = dfs(root.left) + "," + dfs(root.right) + "," + root.val;
	int freq = map.getOrDefault(subTree, 0);
	if (freq == 1) {
		res.add(subTree);
	}
	map.put(subTree,freq + 1);
	return subTree;
}
```
