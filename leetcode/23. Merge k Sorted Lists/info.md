[23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

Hard

Topics

Companies

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

_Merge all the linked-lists into one sorted linked-list and return it._

**Example 1:**

<pre><strong>Input:</strong> lists = [[1,4,5],[1,3,4],[2,6]]
<strong>Output:</strong> [1,1,2,3,4,4,5,6]
<strong>Explanation:</strong> The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
</pre>

**Example 2:**

<pre><strong>Input:</strong> lists = []
<strong>Output:</strong> []
</pre>

**Example 3:**

<pre><strong>Input:</strong> lists = [[]]
<strong>Output:</strong> []
</pre>

**Constraints:**

- `k == lists.length`
- `0 <= k <= 10<sup>4</sup>`
- `0 <= lists[i].length <= 500`
- `-10<sup>4</sup> <= lists[i][j] <= 10<sup>4</sup>`
- `lists[i]` is sorted in **ascending order** .
- The sum of `lists[i].length` will not exceed `10<sup>4</sup>`.
