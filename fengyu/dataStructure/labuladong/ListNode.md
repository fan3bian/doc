1. 反转链表
```java
public ListNode reverse(ListNode head) {
	if (head == null || head.next == null) {
		return head;
	}
	ListNode p = reverse(head.next);
	head.next.next = head;
	head.next = null;
	return p;
}
//反转前N
ListNode suffix;
public ListNode reverseN(ListNode head, int n) {
	if (n == 1) {//最后一个需要反转的节点
		suffix = head.next;
		return head;
	}
	ListNode p = reverseN(head.next, n - 1);
	head.next.next = head;
	head.next = suffix;//连接
	return p;
	//1.反转一个元素就是他本身，要记录	 后驱节点
}

public ListNode reverseBetween(ListNode head, int m, int n) {
	if (m == 1) {
		return reverse(head,n);
	} 
	head.next = reverseBetween(head.next, m - 1, n - 1);
	return head;
}

public ListNode reverseKGroup(ListNode head, int k){
	if (head == null) return head;
	ListNode left = head;
	ListNode right = head;
	for (int i = 0; i < k; i++) {
		if (right != null) return head;//不足K，不反转
		right = right.next;
	}
	ListNode newHead = reverse(left, right);//反转子列表 K个元素
	left.next = reverseKGroup(right, k);
	return newHead;
}
//反转 l到r之间的节点 [l,r)
private ListNode reverse(ListNode l, ListNode r) {
	ListNode prev = null;
	ListNode cur = l;
	while (l != r) {
		ListNode next = l.next;
		l.next = cur;
		cur = prev;
		prev = next;
	}
	return prev;
}
```
//K个一组反转链表 
子问题：如何表示一个链表？
数组表示一个子数组 `subArr = arr, leftIndex, rightIndex`，链表head, leftNode, rightNode。
如何动态表示表示一个子链表
数组可以通过下标作为变量，链表通过节点指针P。 K个一组，每一组是一个子链表，多组则是P指针的移动完成遍历