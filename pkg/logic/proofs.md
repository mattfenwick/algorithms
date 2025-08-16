
# Table of Contents

1. [P -> P](#proof-0)
2. [~ ( P ^ ~ P )](#proof-1)
3. [( P ^ Q ) -> ( Q ^ P )](#proof-2)
4. [( P v Q ) -> ( Q v P )](#proof-3)
5. [P v ~ P](#proof-4)
6. [( P -> Q ) -> ( ~ Q -> ~ P )](#proof-5)
7. [( ~ Q -> ~ P ) -> ( P -> Q )](#proof-6)
8. [Q -> ( P -> Q )](#proof-7)
9. [( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )](#proof-8)
10. [( P ^ P ) -> P](#proof-9)
11. [( P v P ) -> P](#proof-10)



# P -> P <a name="proof-0"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P</pre> | Assume |  |
| 2 | <pre>.   P</pre> | Repeat | 1 |
| 3 | <pre>P -> P</pre> | subproof implication | 1 - 2 |


# ~ ( P ^ ~ P ) <a name="proof-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ~ P</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P</pre> | E ^ (R) | 1 |
| 4 | <pre>~ ( P ^ ~ P )</pre> | subproof contradiction | 1 - 3 |


# ( P ^ Q ) -> ( Q ^ P ) <a name="proof-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ Q</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q ^ P</pre> | I ^ | 3, 2 |
| 5 | <pre>( P ^ Q ) -> ( Q ^ P )</pre> | subproof implication | 1 - 4 |


# ( P v Q ) -> ( Q v P ) <a name="proof-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v Q</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   Q v P</pre> | I v (R) | 2 |
| 4 | <pre>.   P -> ( Q v P )</pre> | subproof implication | 2 - 3 |
| 5 | <pre>.   .   Q</pre> | Assume |  |
| 6 | <pre>.   .   Q v P</pre> | I v (L) | 5 |
| 7 | <pre>.   Q -> ( Q v P )</pre> | subproof implication | 5 - 6 |
| 8 | <pre>.   Q v P</pre> | E v | 4, 7, 1 |
| 9 | <pre>( P v Q ) -> ( Q v P )</pre> | subproof implication | 1 - 8 |


# P v ~ P <a name="proof-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P v ~ P )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   P v ~ P</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ~ ( P v ~ P )</pre> | Reiterate | 1 |
| 5 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   P v ~ P</pre> | I v (R) | 5 |
| 7 | <pre>~ ~ ( P v ~ P )</pre> | subproof contradiction | 1 - 6 |
| 8 | <pre>P v ~ P</pre> | E ~ | 7 |


# ( P -> Q ) -> ( ~ Q -> ~ P ) <a name="proof-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> Q</pre> | Assume |  |
| 2 | <pre>.   .   ~ Q</pre> | Assume |  |
| 3 | <pre>.   .   .   P</pre> | Assume |  |
| 4 | <pre>.   .   .   P -> Q</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   Q</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   .   ~ Q</pre> | Reiterate | 2 |
| 7 | <pre>.   .   ~ P</pre> | subproof contradiction | 3 - 6 |
| 8 | <pre>.   ~ Q -> ~ P</pre> | subproof implication | 2 - 7 |
| 9 | <pre>( P -> Q ) -> ( ~ Q -> ~ P )</pre> | subproof implication | 1 - 8 |


# ( ~ Q -> ~ P ) -> ( P -> Q ) <a name="proof-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ Q -> ~ P</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   .   ~ Q</pre> | Assume |  |
| 4 | <pre>.   .   .   ~ Q -> ~ P</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   ~ P</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 7 | <pre>.   .   ~ ~ Q</pre> | subproof contradiction | 3 - 6 |
| 8 | <pre>.   .   Q</pre> | E ~ | 7 |
| 9 | <pre>.   P -> Q</pre> | subproof implication | 2 - 8 |
| 10 | <pre>( ~ Q -> ~ P ) -> ( P -> Q )</pre> | subproof implication | 1 - 9 |


# Q -> ( P -> Q ) <a name="proof-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   Q</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   Q</pre> | Reiterate | 1 |
| 4 | <pre>.   P -> Q</pre> | subproof implication | 2 - 3 |
| 5 | <pre>Q -> ( P -> Q )</pre> | subproof implication | 1 - 4 |


# ( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R ) <a name="proof-8"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( Q -> R )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   ( P -> Q ) ^ ( Q -> R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P -> Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   Q -> R</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E -> | 4, 2 |
| 7 | <pre>.   .   R</pre> | E -> | 5, 6 |
| 8 | <pre>.   P -> R</pre> | subproof implication | 2 - 7 |
| 9 | <pre>( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )</pre> | subproof implication | 1 - 8 |


# ( P ^ P ) -> P <a name="proof-9"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ P</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>( P ^ P ) -> P</pre> | subproof implication | 1 - 2 |


# ( P v P ) -> P <a name="proof-10"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v P</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   P</pre> | Repeat | 2 |
| 4 | <pre>.   P -> P</pre> | subproof implication | 2 - 3 |
| 5 | <pre>.   P</pre> | E v | 4, 4, 1 |
| 6 | <pre>( P v P ) -> P</pre> | subproof implication | 1 - 5 |
