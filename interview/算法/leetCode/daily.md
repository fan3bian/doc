#### 1.两数之和
#### 2.两数相加
```java
//预先指针
ListNode pre = new ListNode(0);
ListNode cur = pre;
int carry = 0;//进位

while(l1 !=null || l2 !=null){
	int x = l1==null?0:l1.varl;
	int y = l2==null?0:l2.varl;
	int sum = x + y + carry;
	 carry = sum /10;
	 sum =sum % 10;
	 cur.next = new ListNode(sum);
	 cur = cur.next;
	 if(l1!=null){
	 	l1 = l1.next;
	 }
	 if(l2!=null){
	 	l2 = l2.next;
	 }
}
if(carry ==1){
	cur.next = new ListNode(carry)
}
return pre.next;
//返回结果为头结点
```