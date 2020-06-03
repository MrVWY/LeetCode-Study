## LeetCode Question collection and classification!!! (记录自己Leetcode的解题之路)

### 传送门
#### 串烧类题目

+ [买卖股票的最佳时机(Stock)](https://github.com/MrVWY/LeetCode-Study/tree/master/Dynamic_programming/Stock)  
(DP)
  + 121 买卖股票的最佳时机
  + 122 买卖股票的最佳时机 II
  + 123 买卖股票的最佳时机 III 
  + 188 买卖股票的最佳时机 IV (未解决)(difficult)
  + 309 最佳买卖股票时机含冷冻期
  + 714 买卖股票的最佳时机含手续费
+ [打家劫舍(House Robber)](https://github.com/MrVWY/LeetCode-Study/tree/master/House_Robber)  
(DP)(dfs)
  + 198 打家劫舍
  + 213 打家劫舍 II
  + 337 打家劫舍 III
+ [跳跃游戏(Jump Game)](https://github.com/MrVWY/LeetCode-Study/tree/master/Jump_Game)
  + 55 跳跃游戏
  + 45 跳跃游戏 II
  + 1306 跳跃游戏 III
  + 1345 跳跃游戏 IV (未解决)(difficult)
  + 1340 跳跃游戏 V (未解决)(difficult)
+ [螺旋矩阵(Spiral_matrix)](https://github.com/MrVWY/LeetCode-Study/blob/master/Spiral_matrix)
(数学)
  + 54 螺旋矩阵
  + 59 螺旋矩阵 II
  + 885 螺旋矩阵 III (未解决)(difficult)
+ [从上到下打印二叉树](https://github.com/MrVWY/LeetCode-Study/blob/master/Breadth_First_Search/Breadth_first_search_2.go)  
(BFS)
  + 从上到下打印二叉树
  + 从上到下打印二叉树 II
  + 从上到下打印二叉树 III
+ [接雨水](https://github.com/MrVWY/LeetCode-Study/tree/master/Catch_rain)
  + 接雨水 I
  + 接雨水 II (未解决)(difficult)
+ [课程表]()  
(拓扑排序)
  + 207 课程表
  + 210 课程表 II
  + 630 课程表 III (未解决)(difficult)
+ [单词接龙](https://github.com/MrVWY/LeetCode-Study/blob/master/Breadth_First_Search/Word_Solitaire/Word_Solitaire.go)  
(BFS)(双向BFS`未完成`)
  + 127 单词接龙
  + 126 单词接龙 II (未解决)(difficult)
+ [只出现一次的数字](https://github.com/MrVWY/LeetCode-Study/blob/master/bitwise_operation/bitwise_operation_2.go)  
(位运算)
  + 136 只出现一次的数字
  + 137 只出现一次的数字 II
  + 260 只出现一次的数字 III
+ [石子游戏(stone_Game)](https://github.com/MrVWY/LeetCode-Study/tree/master/Dynamic_programming/Stone_Game)  
(博弈动态规划)
  + 877 石子游戏
  + 1140 石子游戏 II (未解决)(medium)
  + 1406 石子游戏 III (未解决)(difficult)
+ [子集](https://github.com/MrVWY/LeetCode-Study/blob/master/bitwise_operation/bitwise_operation_3.go)  
(位运算)(回溯算法`未完成`)
  + 78 子集 I
  + 90 子集 II
+ [全排列](https://github.com/MrVWY/LeetCode-Study/blob/master/Braktrace/Braktrace_1.go)  
(回溯算法+剪枝)
  + 46 全排列
  + 47 全排列 II
+ [组合总和](https://github.com/MrVWY/LeetCode-Study/blob/master/Braktrace/Braktrace_2.go)  
(回溯算法+剪枝)
  + 39 组合总和
  + 40 组合总和 II 
  + 216 组合总和 III 
  + 377 组合总和 Ⅳ (未解决)(medium)
+ [环形链表](https://github.com/MrVWY/LeetCode-Study/blob/master/List/Circular_linked_list.go)  
  + 141 环形链表
  + 142 环形链表 II
+ [反转链表](https://github.com/MrVWY/LeetCode-Study/blob/master/List/Reverse_List/reverseList1.go)  
  + 206 反转链表
  + 25 K 个一组翻转链表

#### 经典
+ [七大排序](https://github.com/MrVWY/LeetCode-Study/blob/master/Sort/Sort.go)
+ [位运算](https://github.com/MrVWY/LeetCode-Study/tree/master/bitwise_operation/bitwise_operation_1.go)
+ [约瑟夫环问题]()
+ [堆](https://github.com/MrVWY/LeetCode-Study/tree/master/heap)  

#### 周赛新世界
+ (夜喵场-第27场) 5410 课程安排 IV  ---打表法

#### 算法总结

+ 动态规划  
1、明确状态定义  
2、写出状态转移方程  
3、思考初始化,每个动态规划问题都要思考初始值是什么  
4、思考输出  
例如：在串烧类题目中的买卖股票中,0表示不持股,1表示持股,接着写出状态方程来表示如果今天持股或者不持股是由前一天对股票做了什么操作得来的,然后取最优解。
最后思考输出,在买卖股票中,要想赚的多那么必然最后一天不持股的时候才赚的多。  

最后这类买卖股票的题目可以算是经典的动态规划问题,我们要明确动态规划的4个过程来对题目进行思考。

+ BFS 和 DFS  
  + BFS伪代码  
    `  
    queue = make(...) //定义一个队列(也就是数组)queue  
    for len(queuq) > 0 {  
        size := len(queue) //当前层数有多少个节点  
        for ... {  
            //开始遍历当前层的每一个节点  
            queue = queue[1:] //出队  
            new_node = 下一层节点  
            if 判断是否符合条件 {
                return  
            } 
             
        }  
        queue = append(queue, new_node)//将下一层节点添加到队列中
    }  
    
    `
    