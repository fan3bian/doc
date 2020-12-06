## 反转
1. leetcode344

```java
//中位对称交换
public void reverseString(char[] s) {
    int len = s.length;
    int mid = len >> 1;
    for (int i = 0; i < mid; i++) {
        char temp = s[i];
        s[i] = s[len - i - 1];
        s[len - 1 -i]  = temp;
    }
}
//双指针 l和r对换
public void reverseString(char[] s) {
    int len = s.length;
    for (int l = 0, r = len -1; i < l; ++i, --r) {
    	char temp = s[l];
    	s[l] = s[r];
   		s[r] = temp;
    }
}	
//双指针
void reverseString(char[] s) {
    int l = 0, r = s.length - 1;
    while (l < r) {
        char temp = s[l];
        s[l++] = s[r];
        s[r--] = temp; 
    }
}
```
字符串反转 #344 #541