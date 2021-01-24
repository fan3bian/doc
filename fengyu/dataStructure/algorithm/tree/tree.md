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

1、对于 BST 的每一个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大。
2、对于 BST 的每一个节点node，它的左侧子树和右侧子树都是 BST。
特性：BST 的中序遍历结果是有序的（升序）。

#### BST遍历
```java
//搜索
public void searchBST(TreeNode root, int target) {
	if (root.val == target) {
		//do something
	}
	if (root.val < target) {
		searchBST(root.right, target);
	}
	if (root.val > target) {
		searchBST(root.left, target);
	}
}
//插入
TreeNode insertBST(TreeNode root, int val) {
	if (root == null) {
		return new TreeNode(val);
	}
	if (root.val == val) {
		throws new IllegalArgumentException("illegal val");
	}
	if (root.val < val) {
		root.right = insertBST(root.right, val);
	}
	if (root.val > val) {
		root.left = insertBST(root.left, val);
	}
	return root;
}
//删除
TreeNode deleteBST(TreeNode root, int val) {
	if (root == null) {
		throws new IllegalArgumentException("illegal val");
	}
	if (root.val == val) {
		//do lelete
		if (root.left == null) {
			return root.right;
		}
		if (root.right == null) {
			return root.left;
		}
		TreeNode minNode = getMin(root.right);
		root.val = minNode.val;
		root.right = deleteBST(root.right, minNode.val);
	}	
	if (root.val < val) {
		root.right = deleteBST(root.right, val);
	}
	if (root.val > val) {
		root.left = deleteBST(root.left, val);
	}
	return root;
}
TreeNode getMin(TreeNode node) {
	//BST最左边的树
	while (node.left ！= null) {
		node = node.left;
	}
	return node;
}
```

前序遍历的代码在进入某一个节点之前的那个时间点执行，后序遍历代码在离开某个节点之后的那个时间点执行