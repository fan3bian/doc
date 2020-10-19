## Tree
https://mp.weixin.qq.com/s/_ymfWYvTNd2GvWvC5HOE4A
## Binary Tree
二叉树：每个节点至多有两个子节点

### 深度优先遍历(DFS:Depth First Search)

- 前序遍历：中左右
- 中序遍历：左中右
- 后序遍历：左右中

注：前中后描述的是**根节点**的相对位置。

按照访问左子树——根节点——右子树的方式遍历这棵树，而在访问左子树或者右子树的时候我们按照同样的方式遍历，直到遍历完整棵树。因

```java
//递归
public List<Integer> inorderTraversal(TreeNode root) {
	List<Integer> res = new ArrayList<Integer>();
	inorder(root, res);
	return res;
}
public void inorder(TreeNode root, List<Integer> res) {
	if (root == null) {
		return ;
	}
	inorder(root.left);
	res.add(root.val);
	inorder(root.right);
}
//迭代
public List<Integer> inorderTraversal(TreeNode root) {
	List<Integer> res = new ArrayList<>();
	Deque<TreeNode> deque = new LinkedList<>();
	while(root != null || !deque.isEmpty()) {
		while (root != null) {
			deque.push(root);
			root = root.left;
		}
		root = deque.pop();
		res.add(root.val);
		root = root.right;
	}
	return res;
}


```
## Binary Search Tree
有序的二叉树


