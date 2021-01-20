## Recursion(递归)

自己调用自己

例子
1. 阶乘
2. 斐波那契数列

#### 递归函数模板
1. recursion terminator: 递归终止条件
2. porcess logic current level: 处理当前层逻辑
3. drill down: 下探到下一层
4. reverse the current level status：清理当前层状态

**寻找最近重复子问题**

#### 递归代码模板
```java
// Java
public void recur(int level, int param) { 


  // terminator 
  if (level > MAX_LEVEL) { 
    // process result 
    return; 
  }


  // process current logic 
  process(level, param); 


  // drill down 
  recur( level: level + 1, newParam); 


  // restore current status 
```

对于递归而言，最重要的是明确递归函数的意义。

递归灵魂三问:
1. 这个函数是干什么的? 
2. 这个函数的变量是什么? 通过参数变量值变化，进行状态转移  a -> a.next; a > a - 1; a -> a.left;
3. 得到递归的结果，怎么做?
