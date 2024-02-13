## ccwc

This is an attempt to solve [this](https://codingchallenges.fyi/challenges/challenge-wc)

A custom implementation of wc tool of linux in golang.

Options:

1. `-c` count the number of bytes
2. `-l` count the number of lines
3. `-w` count number of words
4. `-m` count number of utf-8 bytes. (runes)

---

Usage

```sh
ccwc -w test.txt
```

or it takes input from the stdin stream as well

```sh
cat test.txt | ccwc -c
```

