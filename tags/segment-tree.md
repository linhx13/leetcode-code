# Segment Tree

## 原理

- <https://en.wikipedia.org/wiki/Segment_tree>
- <https://cp-algorithms.com/data_structures/segment_tree.html>
- <https://en.wikipedia.org/wiki/Interval_tree>

The structure of Segment Tree is a binary tree which each node has two attributes start and end denote an segment / interval.

start and end are both integers, they should be assigned in following rules:

- The root's start and end is given by build method.
- The left child of node A has start=A.start, end=(A.start + A.end) / 2.
- The right child of node A has start=(A.start + A.end) / 2 + 1, end=A.end.
- if start equals to end, there will be no children for this node.
