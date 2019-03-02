Math
---

## Catalan Number

Catalan numbers is a sequence of natural numbers that occur in various counting problems.

![catalan number](/images/catalan.svg)

which is a solution of this recurrence formula.

![catalan number](/images/catalan2.svg)

The first catalan numbers for n = 0,1,2,3,... are

```
1, 1, 2, 5, 14, 42, 132, 429 ...
```

the sequence form such as:

- The number of a convex polygon with n+2 sides, which can be cut into triangles by connecting vertices with non-crossing line segments.
- The number of ways to tile a stair step of height n with n rectangles.
- The number of tournament games by n+1 people in total.

## Euler's totient function

Euler's totient function is the number of integers from 0 to N that are relatively prime to n.

[wiki: Euler's totient function](https://en.wikipedia.org/wiki/Euler%27s_totient_function)


## Extended Euclidean algorithm

```
ax + by = gcd(a, b)
```

Integers x and y are always exist and calculable. Extended Euclidean algorithm can calcurate them.

e.g
```
41x + 28y = 1 then x=-12,y=17
```

[wiki: https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm](https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm)


## Congruence relation (Congruence)

[wiki: Congruence relation](https://en.wikipedia.org/wiki/Congruence_relation)

Congruence relation is an equivalence relation, which has a same remainder of N.

```
  a ≡ b (mod b)
```

e.g.

```
  37 ≡ 57 (mod 10)
```

because `37 % 10 = 7` and `57 % 10 = 7`.

```
  9 / 4 ≡ 12 (mod 13)
```

because `4 * 12 = 48 ≡ 9 (mod 13)`.

## Inverse element

[wiki: Inverse element](https://en.wikipedia.org/wiki/Inverse_element)

```
a * x = 1
```

When we have this equation, `x` is __inverse number__ of `a`.

```
a % y = 1
```

Now, when we have this equation, `y` is __inverse element__ of `a`.


e.g What is X in this equation? X is inverse element of `5`.

```
  5 * X = 40 ≡ 1 (mod 13)
```


## Felmat's little therem

[wiki: Felmat's little therem](https://en.wikipedia.org/wiki/Fermat%27s_little_theorem)

Inverse element can be solved by Felmat's little therem, which is:

```
  a^(p-1) ≡ 1 (mod p)
```

it is same with:

```
  a^(-1) ≡ a^(p-2) (mod p)
```

it is also same with:

```
  b / a ≡ b * a^(p-2)
```

## Combination (mod 10^9+7) used by Felmat's little therem

TODO

Example: https://atcoder.jp/contests/abc110/tasks/abc110_d
