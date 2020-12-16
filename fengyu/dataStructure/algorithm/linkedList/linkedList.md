## LinkedList
```java
public class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}
```
节点、值、引用

#### 移除节点

1. 移除结点

被移除的结点可能是：头结点、普通结点、尾结点
单链表仅持有对下一节点的引用(next)，所以在移除时，我们无法对本结点操作。只能通过操作前一个结点来移除元素。

```java 
//移除投结点
if (head != null) {
	head = head.next;
}
//删除普通结点
cur.next = cur.next.next;
//删除尾结点
cur.next = null;
```
技巧：设置一个dummy node,避免对头结点特殊判断

```java
void remove(ListNode head, int val) {
	ListNode dummy = new ListNode(0);
	dummy.next = head;//

}
```






