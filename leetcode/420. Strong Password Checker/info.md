[420. Strong Password Checker](https://leetcode.com/problems/strong-password-checker/)

Hard

Topics

Companies

A password is considered strong if the below conditions are all met:

- It has at least `6` characters and at most `20` characters.
- It contains at least **one lowercase** letter, at least **one uppercase** letter, and at least **one digit** .
- It does not contain three repeating characters in a row (i.e., `"B<u><strong>aaa</strong></u>bb0"` is weak, but `"B<strong><u>aa</u></strong>b<u><strong>a</strong></u>0"` is strong).

Given a string `password`, return _the minimum number of steps required to make `password` strong. if `password` is already strong, return `0`._

In one step, you can:

- Insert one character to `password`,
- Delete one character from `password`, or
- Replace one character of `password` with another character.

**Example 1:**

<pre><strong>Input:</strong> password = "a"
<strong>Output:</strong> 5
</pre>

**Example 2:**

<pre><strong>Input:</strong> password = "aA1"
<strong>Output:</strong> 3
</pre>

**Example 3:**

<pre><strong>Input:</strong> password = "1337C0d3"
<strong>Output:</strong> 0
</pre>

**Constraints:**

- `1 <= password.length <= 50`
- `password` consists of letters, digits, dot `'.'` or exclamation mark `'!'`.
