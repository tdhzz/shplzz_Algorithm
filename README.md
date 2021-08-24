# shplzz_Algorithm
# 算法训练营规则

## 基本规则
1. 每人每周3道题，周日22:00:00 CST +08:00结束统计
2. 轮流值班，按照拼音排序
3. 执行sql：
```
SELECT question.frontend_id, question.title, question_tag.tag 
FROM question 
LEFT OUTER JOIN question_tag on question.id = question_tag.question_id 
WHERE 
     question.status = 0 
AND
    question.difficulty = “Easy”
ORDER BY random() LIMIT 5
```
## 值班人员职责
1. tool tom来题目抽取
2. 维护每周题目完成情况在线文档

## 题目管理流程
1. 每周上传代码到github仓库

## 本周题目
{medium：[537, 1513],
easy: [1408, 788, 58, 929]}

##题目类型
Sort, Bit Manipulation, Array, Math, Greedy，String, Binary Search, Hash Table, Stack, Geometry, Linked List, Graph, Recursion, Tree, Two Pointers, Breadth-first Search, Brainteaser, Heap, Queue, Dynamic Programming, Depth-first Search, Design, Trie, Minimax, Backtracking, Divide and Conquer, Sliding Window, Segment Tree, Union Find, Binary Search Tree, Line Sweep, Random, Ordered Map, Rejection Sampling, Reservoir Sampling, Binary Indexed Tree, Topological Sort, Suffix Array, Rolling Hash, Memoization

## 奖惩
1. 每周提交题目数n,不足3道题，需要交(3-n)x3元团建费;
## 
owner每周统计：

仓库地址：
https://github.com/tdhzz/shplzz_Algorithm
Git仓库地址：
https://github.com/tdhzz/shplzz_Algorithm
