树的算法：把题目的要求细化，搞清楚**根节点**应该做什么，然后剩下的事情抛给前/中/后序的遍历框架就行了

#### 遍历的选择
具体使用前序、中序、后序遍历取决于处理当前根节点时需依赖信息(无、左子树、子树)

```java
public String collect(TreeNode node) {
    if (node == null) return "#";
    String serial = node.val + "," + collect(node.left) + "," + collect(node.right);
    count.put(serial, count.getOrDefault(serial, 0) + 1);
    if (count.get(serial) == 2)
        ans.add(node);
    return serial;
}
```
什么时候该选择树的后序遍历？
处理当前节点需要依靠推导左右子树的信息。

 
```java
//计算二叉树节点个数
public int countNode(TreeNode root) {
	if (root == null) return 0;

	int leftCount = countNode(root.left);
	int rightCount = countNode(root.right);
	return 1 + leftCount + rightCount;
}
//计算满二叉树节点个数
public int countNode(TreeNode root) {
	//2^h - 1
	int h = 0;
	while (root != null) {
		root = root.left;
		h++;
	}
	return (int)Math.pow(2,h) - 1
}
//计算完全二叉树节点个数
public void countNode(TreeNode root) {
	//完全二叉树左右节点，必有一个满二叉树
	int hl = 0, hr = 0;
	TreeNode l = root, r = root;
	while (l != null) {
		l = l.left;
		hl++;
	}
	while (r != null) {
		r = r.right;
		hr++;
	}
	if (hl == hr) {
		return (int)Math.pow(2,hl) - 1;
	}
	return 1 + countNode(root.left, root.right);
}
```