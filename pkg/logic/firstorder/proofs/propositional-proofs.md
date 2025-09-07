
# Table of Contents

1. [theorems](#theorems)
    1. [P v ~ P](#proof-1-1)
    2. [( P -> Q ) <-> ( ~ Q -> ~ P )](#proof-1-2)
    3. [~ ( P -> Q ) <-> ( P ^ ~ Q )](#proof-1-3)
    4. [( P v Q ) <-> ( ~ P -> Q )](#proof-1-4)
    5. [~ ( P v Q ) <-> ( ~ P ^ ~ Q )](#proof-1-5)
    6. [~ ( P ^ Q ) <-> ( ~ P v ~ Q )](#proof-1-6)
    7. [( P <-> ~ Q ) <-> ~ ( P <-> Q )](#proof-1-7)
2. [reflexivity](#reflexivity)
    1. [P -> P](#proof-2-1)
    2. [( P ^ P ) <-> P](#proof-2-2)
    3. [( P v P ) <-> P](#proof-2-3)
    4. [P <-> P](#proof-2-4)
3. [basics](#basics)
    1. [~ ( P ^ ~ P )](#proof-3-1)
    2. [( P ^ Q ) -> ( P v Q )](#proof-3-2)
    3. [( P -> ~ P ) -> ~ P](#proof-3-3)
    4. [( ( P -> Q ) -> P ) -> P](#proof-3-4)
    5. [( P ^ ~ P ) -> Q](#proof-3-5)
    6. [P <-> ~ ~ P](#proof-3-6)
4. [arrows](#arrows)
    1. [Q -> ( P -> Q )](#proof-4-1)
    2. [~ P -> ( P -> Q )](#proof-4-2)
    3. [( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) )](#proof-4-3)
    4. [( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P](#proof-4-4)
    5. [( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q](#proof-4-5)
    6. [( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )](#proof-4-6)
5. [transitivity](#transitivity)
    1. [( P -> Q ) -> ( ( Q -> R ) -> ( P -> R ) )](#proof-5-1)
    2. [( P <-> Q ) -> ( ( Q <-> R ) -> ( P <-> R ) )](#proof-5-2)
6. [commutativity](#commutativity)
    1. [( P ^ Q ) -> ( Q ^ P )](#proof-6-1)
    2. [( P v Q ) -> ( Q v P )](#proof-6-2)
    3. [( P <-> Q ) -> ( Q <-> P )](#proof-6-3)
7. [associativity](#associativity)
    1. [( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) )](#proof-7-1)
    2. [( ( P v Q ) v R ) <-> ( P v ( Q v R ) )](#proof-7-2)
    3. [( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) )](#proof-7-3)
8. [distributivity](#distributivity)
    1. [( P -> ( Q -> R ) ) <-> ( ( P -> Q ) -> ( P -> R ) )](#proof-8-1)
    2. [( P -> ( Q v R ) ) <-> ( ( P -> Q ) v ( P -> R ) )](#proof-8-2)
    3. [( P -> ( Q ^ R ) ) <-> ( ( P -> Q ) ^ ( P -> R ) )](#proof-8-3)
    4. [( P ^ ( Q -> R ) ) -> ( ( P ^ Q ) -> ( P ^ R ) )](#proof-8-4)
    5. [( P ^ ( Q v R ) ) <-> ( ( P ^ Q ) v ( P ^ R ) )](#proof-8-5)
    6. [( P ^ ( Q ^ R ) ) <-> ( ( P ^ Q ) ^ ( P ^ R ) )](#proof-8-6)
    7. [( P v ( Q -> R ) ) <-> ( ( P v Q ) -> ( P v R ) )](#proof-8-7)
    8. [( P v ( Q v R ) ) <-> ( ( P v Q ) v ( P v R ) )](#proof-8-8)
    9. [( P v ( Q ^ R ) ) <-> ( ( P v Q ) ^ ( P v R ) )](#proof-8-9)
9. [anti-distributivity](#anti-distributivity)
    1. [( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) )](#proof-9-1)
    2. [( ( P ^ Q ) -> R ) <-> ( ( P -> R ) v ( Q -> R ) )](#proof-9-2)
10. [double-distribution](#double-distribution)
    1. [( ( P v Q ) ^ ( R v S ) ) <-> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )](#proof-10-1)
    2. [( ( P ^ Q ) v ( R ^ S ) ) <-> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )](#proof-10-2)
11. [disjunction](#disjunction)
    1. [( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )](#proof-11-1)
    2. [( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S )](#proof-11-2)
    3. [( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q )](#proof-11-3)
    4. [( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R](#proof-11-4)
    5. [( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q )](#proof-11-5)
12. [biconditional](#biconditional)
    1. [( ( P <-> Q ) ^ P ) -> Q](#proof-12-1)
    2. [( ( P <-> Q ) ^ ~ P ) -> ~ Q](#proof-12-2)
    3. [( P <-> ~ Q ) -> ( ~ P <-> Q )](#proof-12-3)
    4. [( P <-> Q ) <-> ( ~ P <-> ~ Q )](#proof-12-4)
    5. [( P <-> Q ) v ~ ( P <-> Q )](#proof-12-5)
    6. [( P <-> ~ Q ) <-> ( ( P v Q ) ^ ( ~ P v ~ Q ) )](#proof-12-6)
13. [miscellaneous](#miscellaneous)
    1. [( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )](#proof-13-1)
    2. [( P -> Q ) v ( Q -> R )](#proof-13-2)
    3. [( P -> Q ) v ( Q -> P )](#proof-13-3)

# theorems <a name="theorems"></a>

## P v ~ P <a name="proof-1-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( P v ~ P )</pre> | Assume: I/E ~ |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   P v ~ P</pre> | I v (L) | 2 |
| 4 |  | <pre>.   .   ~ ( P v ~ P )</pre> | Reiterate | 1 |
| 5 |  | <pre>.   ~ P</pre> | subproof I/E ~ | 2 - 4 |
| 6 |  | <pre>.   P v ~ P</pre> | I v (R) | 5 |
| 7 |  | <pre>P v ~ P</pre> | subproof I/E ~ | 1 - 6 |

## ( P -> Q ) <-> ( ~ Q -> ~ P ) <a name="proof-1-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   ~ Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: I/E ~ |  |
| 4 |  | <pre>.   .   .   P -> Q</pre> | Reiterate | 1 |
| 5 |  | <pre>.   .   .   Q</pre> | E -> | 4, 3 |
| 6 |  | <pre>.   .   .   ~ Q</pre> | Reiterate | 2 |
| 7 |  | <pre>.   .   ~ P</pre> | subproof I/E ~ | 3 - 6 |
| 8 |  | <pre>.   ~ Q -> ~ P</pre> | subproof I -> | 2 - 7 |
| 9 |  | <pre>( P -> Q ) -> ( ~ Q -> ~ P )</pre> | subproof I -> | 1 - 8 |
| 10 |  | <pre>.   ~ Q -> ~ P</pre> | Assume: I -> |  |
| 11 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 12 |  | <pre>.   .   .   ~ Q</pre> | Assume: I/E ~ |  |
| 13 |  | <pre>.   .   .   ~ Q -> ~ P</pre> | Reiterate | 10 |
| 14 |  | <pre>.   .   .   ~ P</pre> | E -> | 13, 12 |
| 15 |  | <pre>.   .   .   P</pre> | Reiterate | 11 |
| 16 |  | <pre>.   .   Q</pre> | subproof I/E ~ | 12 - 15 |
| 17 |  | <pre>.   P -> Q</pre> | subproof I -> | 11 - 16 |
| 18 |  | <pre>( ~ Q -> ~ P ) -> ( P -> Q )</pre> | subproof I -> | 10 - 17 |
| 19 |  | <pre>( P -> Q ) <-> ( ~ Q -> ~ P )</pre> | I <-> | 9, 18 |

## ~ ( P -> Q ) <-> ( P ^ ~ Q ) <a name="proof-1-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( P -> Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   ~ P</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   .   ~ Q</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 2 |
| 5 |  | <pre>.   .   ~ Q -> ~ P</pre> | subproof I -> | 3 - 4 |
| 6 |  | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 5 |
| 7 |  | <pre>.   .   ~ ( P -> Q )</pre> | Reiterate | 1 |
| 8 |  | <pre>.   P</pre> | subproof I/E ~ | 2 - 7 |
| 9 |  | <pre>.   .   Q</pre> | Assume: I/E ~ |  |
| 10 |  | <pre>.   .   P</pre> | Reiterate | 8 |
| 11 |  | <pre>.   .   P -> Q</pre> | I -> | 10, 9 |
| 12 |  | <pre>.   .   ~ ( P -> Q )</pre> | Reiterate | 1 |
| 13 |  | <pre>.   ~ Q</pre> | subproof I/E ~ | 9 - 12 |
| 14 |  | <pre>.   P ^ ~ Q</pre> | I ^ | 8, 13 |
| 15 |  | <pre>~ ( P -> Q ) -> ( P ^ ~ Q )</pre> | subproof I -> | 1 - 14 |
| 16 |  | <pre>.   P ^ ~ Q</pre> | Assume: I -> |  |
| 17 |  | <pre>.   .   P -> Q</pre> | Assume: I/E ~ |  |
| 18 |  | <pre>.   .   P ^ ~ Q</pre> | Reiterate | 16 |
| 19 |  | <pre>.   .   P</pre> | E ^ (L) | 18 |
| 20 |  | <pre>.   .   ~ Q</pre> | E ^ (R) | 18 |
| 21 |  | <pre>.   .   Q</pre> | E -> | 17, 19 |
| 22 |  | <pre>.   ~ ( P -> Q )</pre> | subproof I/E ~ | 17 - 21 |
| 23 |  | <pre>( P ^ ~ Q ) -> ~ ( P -> Q )</pre> | subproof I -> | 16 - 22 |
| 24 |  | <pre>~ ( P -> Q ) <-> ( P ^ ~ Q )</pre> | I <-> | 15, 23 |

## ( P v Q ) <-> ( ~ P -> Q ) <a name="proof-1-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   ~ P</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   Q</pre> | Reiterate | 2 |
| 5 |  | <pre>.   .   ~ P -> Q</pre> | subproof I -> | 3 - 4 |
| 6 |  | <pre>.   Q -> ( ~ P -> Q )</pre> | subproof I -> | 2 - 5 |
| 7 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 8 |  | <pre>.   .   .   ~ Q</pre> | Assume: I -> |  |
| 9 |  | <pre>.   .   .   P</pre> | Reiterate | 7 |
| 10 |  | <pre>.   .   ~ Q -> P</pre> | subproof I -> | 8 - 9 |
| 11 |  | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 10 |
| 12 |  | <pre>.   P -> ( ~ P -> Q )</pre> | subproof I -> | 7 - 11 |
| 13 |  | <pre>.   ~ P -> Q</pre> | E v | 12, 6, 1 |
| 14 |  | <pre>( P v Q ) -> ( ~ P -> Q )</pre> | subproof I -> | 1 - 13 |
| 15 |  | <pre>.   ~ P -> Q</pre> | Assume: I -> |  |
| 16 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 17 |  | <pre>.   .   P v Q</pre> | I v (L) | 16 |
| 18 |  | <pre>.   P -> ( P v Q )</pre> | subproof I -> | 16 - 17 |
| 19 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 20 |  | <pre>.   .   ~ P -> Q</pre> | Reiterate | 15 |
| 21 |  | <pre>.   .   Q</pre> | E -> | 20, 19 |
| 22 |  | <pre>.   .   P v Q</pre> | I v (R) | 21 |
| 23 |  | <pre>.   ~ P -> ( P v Q )</pre> | subproof I -> | 19 - 22 |
| 24 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 25 |  | <pre>.   P v Q</pre> | E v | 18, 23, 24 |
| 26 |  | <pre>( ~ P -> Q ) -> ( P v Q )</pre> | subproof I -> | 15 - 25 |
| 27 |  | <pre>( P v Q ) <-> ( ~ P -> Q )</pre> | I <-> | 14, 26 |

## ~ ( P v Q ) <-> ( ~ P ^ ~ Q ) <a name="proof-1-5"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( P v Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 |  | <pre>.   .   ~ ( P v Q )</pre> | Reiterate | 1 |
| 5 |  | <pre>.   ~ P</pre> | subproof I/E ~ | 2 - 4 |
| 6 |  | <pre>.   .   Q</pre> | Assume: I/E ~ |  |
| 7 |  | <pre>.   .   P v Q</pre> | I v (R) | 6 |
| 8 |  | <pre>.   .   ~ ( P v Q )</pre> | Reiterate | 1 |
| 9 |  | <pre>.   ~ Q</pre> | subproof I/E ~ | 6 - 8 |
| 10 |  | <pre>.   ~ P ^ ~ Q</pre> | I ^ | 5, 9 |
| 11 |  | <pre>~ ( P v Q ) -> ( ~ P ^ ~ Q )</pre> | subproof I -> | 1 - 10 |
| 12 |  | <pre>.   ~ P ^ ~ Q</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 14 |  | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume: I/E ~ |  |
| 15 |  | <pre>.   .   .   ~ P</pre> | E ^ (L) | 14 |
| 16 |  | <pre>.   .   .   P</pre> | Reiterate | 13 |
| 17 |  | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof I/E ~ | 14 - 16 |
| 18 |  | <pre>.   P -> ~ ( ~ P ^ ~ Q )</pre> | subproof I -> | 13 - 17 |
| 19 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 20 |  | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume: I/E ~ |  |
| 21 |  | <pre>.   .   .   ~ Q</pre> | E ^ (R) | 20 |
| 22 |  | <pre>.   .   .   Q</pre> | Reiterate | 19 |
| 23 |  | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof I/E ~ | 20 - 22 |
| 24 |  | <pre>.   Q -> ~ ( ~ P ^ ~ Q )</pre> | subproof I -> | 19 - 23 |
| 25 |  | <pre>.   .   P v Q</pre> | Assume: I/E ~ |  |
| 26 |  | <pre>.   .   P -> ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 18 |
| 27 |  | <pre>.   .   Q -> ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 24 |
| 28 |  | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | E v | 26, 27, 25 |
| 29 |  | <pre>.   .   ~ P ^ ~ Q</pre> | Reiterate | 12 |
| 30 |  | <pre>.   ~ ( P v Q )</pre> | subproof I/E ~ | 25 - 29 |
| 31 |  | <pre>( ~ P ^ ~ Q ) -> ~ ( P v Q )</pre> | subproof I -> | 12 - 30 |
| 32 |  | <pre>~ ( P v Q ) <-> ( ~ P ^ ~ Q )</pre> | I <-> | 11, 31 |

## ~ ( P ^ Q ) <-> ( ~ P v ~ Q ) <a name="proof-1-6"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( P ^ Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   ~ ( ~ P v ~ Q )</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   .   ~ P</pre> | Assume: I/E ~ |  |
| 4 |  | <pre>.   .   .   ~ P v ~ Q</pre> | I v (L) | 3 |
| 5 |  | <pre>.   .   .   ~ ( ~ P v ~ Q )</pre> | Reiterate | 2 |
| 6 |  | <pre>.   .   P</pre> | subproof I/E ~ | 3 - 5 |
| 7 |  | <pre>.   .   .   ~ Q</pre> | Assume: I/E ~ |  |
| 8 |  | <pre>.   .   .   ~ P v ~ Q</pre> | I v (R) | 7 |
| 9 |  | <pre>.   .   .   ~ ( ~ P v ~ Q )</pre> | Reiterate | 2 |
| 10 |  | <pre>.   .   Q</pre> | subproof I/E ~ | 7 - 9 |
| 11 |  | <pre>.   .   P ^ Q</pre> | I ^ | 6, 10 |
| 12 |  | <pre>.   .   ~ ( P ^ Q )</pre> | Reiterate | 1 |
| 13 |  | <pre>.   ~ P v ~ Q</pre> | subproof I/E ~ | 2 - 12 |
| 14 |  | <pre>~ ( P ^ Q ) -> ( ~ P v ~ Q )</pre> | subproof I -> | 1 - 13 |
| 15 |  | <pre>.   ~ P v ~ Q</pre> | Assume: I -> |  |
| 16 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 17 |  | <pre>.   .   .   P ^ Q</pre> | Assume: I/E ~ |  |
| 18 |  | <pre>.   .   .   P</pre> | E ^ (L) | 17 |
| 19 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 16 |
| 20 |  | <pre>.   .   ~ ( P ^ Q )</pre> | subproof I/E ~ | 17 - 19 |
| 21 |  | <pre>.   ~ P -> ~ ( P ^ Q )</pre> | subproof I -> | 16 - 20 |
| 22 |  | <pre>.   .   ~ Q</pre> | Assume: I -> |  |
| 23 |  | <pre>.   .   .   P ^ Q</pre> | Assume: I/E ~ |  |
| 24 |  | <pre>.   .   .   Q</pre> | E ^ (R) | 23 |
| 25 |  | <pre>.   .   .   ~ Q</pre> | Reiterate | 22 |
| 26 |  | <pre>.   .   ~ ( P ^ Q )</pre> | subproof I/E ~ | 23 - 25 |
| 27 |  | <pre>.   ~ Q -> ~ ( P ^ Q )</pre> | subproof I -> | 22 - 26 |
| 28 |  | <pre>.   ~ ( P ^ Q )</pre> | E v | 21, 27, 15 |
| 29 |  | <pre>( ~ P v ~ Q ) -> ~ ( P ^ Q )</pre> | subproof I -> | 15 - 28 |
| 30 |  | <pre>~ ( P ^ Q ) <-> ( ~ P v ~ Q )</pre> | I <-> | 14, 29 |

## ( P <-> ~ Q ) <-> ~ ( P <-> Q ) <a name="proof-1-7"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P <-> ~ Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 |  | <pre>.   ~ Q -> P</pre> | E <-> (R) | 1 |
| 4 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 5 |  | <pre>.   .   .   P <-> Q</pre> | Assume: I/E ~ |  |
| 6 |  | <pre>.   .   .   P -> ~ Q</pre> | Reiterate | 2 |
| 7 |  | <pre>.   .   .   P -> Q</pre> | E <-> (L) | 5 |
| 8 |  | <pre>.   .   .   P</pre> | Reiterate | 4 |
| 9 |  | <pre>.   .   .   Q</pre> | E -> | 7, 8 |
| 10 |  | <pre>.   .   .   ~ Q</pre> | E -> | 6, 8 |
| 11 |  | <pre>.   .   ~ ( P <-> Q )</pre> | subproof I/E ~ | 5 - 10 |
| 12 |  | <pre>.   P -> ~ ( P <-> Q )</pre> | subproof I -> | 4 - 11 |
| 13 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 14 |  | <pre>.   .   .   P <-> Q</pre> | Assume: I/E ~ |  |
| 15 |  | <pre>.   .   .   Q -> P</pre> | E <-> (R) | 14 |
| 16 |  | <pre>.   .   .   ~ P -> ~ Q</pre> | Theorem: contrapositive | 15 |
| 17 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 13 |
| 18 |  | <pre>.   .   .   ~ Q</pre> | E -> | 16, 17 |
| 19 |  | <pre>.   .   .   ~ Q -> P</pre> | Reiterate | 3 |
| 20 |  | <pre>.   .   .   P</pre> | E -> | 19, 18 |
| 21 |  | <pre>.   .   ~ ( P <-> Q )</pre> | subproof I/E ~ | 14 - 20 |
| 22 |  | <pre>.   ~ P -> ~ ( P <-> Q )</pre> | subproof I -> | 13 - 21 |
| 23 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 24 |  | <pre>.   ~ ( P <-> Q )</pre> | E v | 12, 22, 23 |
| 25 |  | <pre>( P <-> ~ Q ) -> ~ ( P <-> Q )</pre> | subproof I -> | 1 - 24 |
| 26 |  | <pre>.   ~ ( P <-> Q )</pre> | Assume: I -> |  |
| 27 |  | <pre>.   .   ~ ( P -> ~ Q )</pre> | Assume: I/E ~ |  |
| 28 |  | <pre>.   .   P ^ Q</pre> | Theorem: (~ ->) to ^ | 27 |
| 29 |  | <pre>.   .   P</pre> | E ^ (L) | 28 |
| 30 |  | <pre>.   .   Q</pre> | E ^ (R) | 28 |
| 31 |  | <pre>.   .   P -> Q</pre> | I -> | 29, 30 |
| 32 |  | <pre>.   .   Q -> P</pre> | I -> | 30, 29 |
| 33 |  | <pre>.   .   P <-> Q</pre> | I <-> | 31, 32 |
| 34 |  | <pre>.   .   ~ ( P <-> Q )</pre> | Reiterate | 26 |
| 35 |  | <pre>.   P -> ~ Q</pre> | subproof I/E ~ | 27 - 34 |
| 36 |  | <pre>.   .   ~ ( ~ Q -> P )</pre> | Assume: I/E ~ |  |
| 37 |  | <pre>.   .   ~ Q ^ ~ P</pre> | Theorem: (~ ->) to ^ | 36 |
| 38 |  | <pre>.   .   ~ Q</pre> | E ^ (L) | 37 |
| 39 |  | <pre>.   .   ~ P</pre> | E ^ (R) | 37 |
| 40 |  | <pre>.   .   ~ P -> ~ Q</pre> | I -> | 39, 38 |
| 41 |  | <pre>.   .   Q -> P</pre> | Theorem: contrapositive | 40 |
| 42 |  | <pre>.   .   ~ Q -> ~ P</pre> | I -> | 38, 39 |
| 43 |  | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 42 |
| 44 |  | <pre>.   .   P <-> Q</pre> | I <-> | 43, 41 |
| 45 |  | <pre>.   .   ~ ( P <-> Q )</pre> | Reiterate | 26 |
| 46 |  | <pre>.   ~ Q -> P</pre> | subproof I/E ~ | 36 - 45 |
| 47 |  | <pre>.   P <-> ~ Q</pre> | I <-> | 35, 46 |
| 48 |  | <pre>~ ( P <-> Q ) -> ( P <-> ~ Q )</pre> | subproof I -> | 26 - 47 |
| 49 |  | <pre>( P <-> ~ Q ) <-> ~ ( P <-> Q )</pre> | I <-> | 25, 48 |

# reflexivity <a name="reflexivity"></a>

## P -> P <a name="proof-2-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P</pre> | Assume: I -> |  |
| 2 |  | <pre>P -> P</pre> | subproof I -> | 1 - 1 |

## ( P ^ P ) <-> P <a name="proof-2-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>( P ^ P ) -> P</pre> | subproof I -> | 1 - 2 |
| 4 |  | <pre>.   P</pre> | Assume: I -> |  |
| 5 |  | <pre>.   P ^ P</pre> | I ^ | 4, 4 |
| 6 |  | <pre>P -> ( P ^ P )</pre> | subproof I -> | 4 - 5 |
| 7 |  | <pre>( P ^ P ) <-> P</pre> | I <-> | 3, 6 |

## ( P v P ) <-> P <a name="proof-2-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   P -> P</pre> | subproof I -> | 2 - 2 |
| 4 |  | <pre>.   P</pre> | E v | 3, 3, 1 |
| 5 |  | <pre>( P v P ) -> P</pre> | subproof I -> | 1 - 4 |
| 6 |  | <pre>.   P</pre> | Assume: I -> |  |
| 7 |  | <pre>.   P v P</pre> | I v (L) | 6 |
| 8 |  | <pre>P -> ( P v P )</pre> | subproof I -> | 6 - 7 |
| 9 |  | <pre>( P v P ) <-> P</pre> | I <-> | 5, 8 |

## P <-> P <a name="proof-2-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P</pre> | Assume: I -> |  |
| 2 |  | <pre>P -> P</pre> | subproof I -> | 1 - 1 |
| 3 |  | <pre>P <-> P</pre> | I <-> | 2, 2 |

# basics <a name="basics"></a>

## ~ ( P ^ ~ P ) <a name="proof-3-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ ~ P</pre> | Assume: I/E ~ |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ P</pre> | E ^ (R) | 1 |
| 4 |  | <pre>~ ( P ^ ~ P )</pre> | subproof I/E ~ | 1 - 3 |

## ( P ^ Q ) -> ( P v Q ) <a name="proof-3-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   P v Q</pre> | I v (L) | 2 |
| 4 |  | <pre>( P ^ Q ) -> ( P v Q )</pre> | subproof I -> | 1 - 3 |

## ( P -> ~ P ) -> ~ P <a name="proof-3-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> ~ P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   P -> ~ P</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   ~ P</pre> | E -> | 3, 2 |
| 5 |  | <pre>.   ~ P</pre> | subproof I/E ~ | 2 - 4 |
| 6 |  | <pre>( P -> ~ P ) -> ~ P</pre> | subproof I -> | 1 - 5 |

## ( ( P -> Q ) -> P ) -> P <a name="proof-3-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P -> Q ) -> P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   ~ ( P -> Q )</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P ^ ~ Q</pre> | Theorem: (~ ->) to ^ | 2 |
| 4 |  | <pre>.   .   P</pre> | E ^ (L) | 3 |
| 5 |  | <pre>.   ~ ( P -> Q ) -> P</pre> | subproof I -> | 2 - 4 |
| 6 |  | <pre>.   ( P -> Q ) v ~ ( P -> Q )</pre> | Theorem: excluded middle |  |
| 7 |  | <pre>.   P</pre> | E v | 1, 5, 6 |
| 8 |  | <pre>( ( P -> Q ) -> P ) -> P</pre> | subproof I -> | 1 - 7 |

## ( P ^ ~ P ) -> Q <a name="proof-3-5"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>~ ( P ^ ~ P )</pre> | Theorem: non-contradiction |  |
| 2 |  | <pre>.   ~ Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   ~ ( P ^ ~ P )</pre> | Reiterate | 1 |
| 4 |  | <pre>~ Q -> ~ ( P ^ ~ P )</pre> | subproof I -> | 2 - 3 |
| 5 |  | <pre>( P ^ ~ P ) -> Q</pre> | Theorem: contrapositive | 4 |

## P <-> ~ ~ P <a name="proof-3-6"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   ~ P</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   P</pre> | Reiterate | 1 |
| 4 |  | <pre>.   ~ ~ P</pre> | subproof I/E ~ | 2 - 3 |
| 5 |  | <pre>P -> ~ ~ P</pre> | subproof I -> | 1 - 4 |
| 6 |  | <pre>.   ~ ~ P</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   ~ P</pre> | Assume: I/E ~ |  |
| 8 |  | <pre>.   .   ~ ~ P</pre> | Reiterate | 6 |
| 9 |  | <pre>.   P</pre> | subproof I/E ~ | 7 - 8 |
| 10 |  | <pre>~ ~ P -> P</pre> | subproof I -> | 6 - 9 |
| 11 |  | <pre>P <-> ~ ~ P</pre> | I <-> | 5, 10 |

# arrows <a name="arrows"></a>

## Q -> ( P -> Q ) <a name="proof-4-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   Q</pre> | Reiterate | 1 |
| 4 |  | <pre>.   P -> Q</pre> | subproof I -> | 2 - 3 |
| 5 |  | <pre>Q -> ( P -> Q )</pre> | subproof I -> | 1 - 4 |

## ~ P -> ( P -> Q ) <a name="proof-4-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   ~ Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   ~ P</pre> | Reiterate | 1 |
| 4 |  | <pre>.   ~ Q -> ~ P</pre> | subproof I -> | 2 - 3 |
| 5 |  | <pre>.   P -> Q</pre> | Theorem: contrapositive | 4 |
| 6 |  | <pre>~ P -> ( P -> Q )</pre> | subproof I -> | 1 - 5 |

## ( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) ) <a name="proof-4-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P ^ Q ) -> R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 5 |  | <pre>.   .   .   P ^ Q</pre> | I ^ | 4, 3 |
| 6 |  | <pre>.   .   .   ( P ^ Q ) -> R</pre> | Reiterate | 1 |
| 7 |  | <pre>.   .   .   R</pre> | E -> | 6, 5 |
| 8 |  | <pre>.   .   Q -> R</pre> | subproof I -> | 3 - 7 |
| 9 |  | <pre>.   P -> ( Q -> R )</pre> | subproof I -> | 2 - 8 |
| 10 |  | <pre>( ( P ^ Q ) -> R ) -> ( P -> ( Q -> R ) )</pre> | subproof I -> | 1 - 9 |
| 11 |  | <pre>.   P -> ( Q -> R )</pre> | Assume: I -> |  |
| 12 |  | <pre>.   .   P ^ Q</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   P</pre> | E ^ (L) | 12 |
| 14 |  | <pre>.   .   Q</pre> | E ^ (R) | 12 |
| 15 |  | <pre>.   .   P -> ( Q -> R )</pre> | Reiterate | 11 |
| 16 |  | <pre>.   .   Q -> R</pre> | E -> | 15, 13 |
| 17 |  | <pre>.   .   R</pre> | E -> | 16, 14 |
| 18 |  | <pre>.   ( P ^ Q ) -> R</pre> | subproof I -> | 12 - 17 |
| 19 |  | <pre>( P -> ( Q -> R ) ) -> ( ( P ^ Q ) -> R )</pre> | subproof I -> | 11 - 18 |
| 20 |  | <pre>( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) )</pre> | I <-> | 10, 19 |

## ( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P <a name="proof-4-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P -> Q ) ^ ( P -> ~ Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   ( P -> Q ) ^ ( P -> ~ Q )</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   P -> Q</pre> | E ^ (L) | 3 |
| 5 |  | <pre>.   .   P -> ~ Q</pre> | E ^ (R) | 3 |
| 6 |  | <pre>.   .   Q</pre> | E -> | 4, 2 |
| 7 |  | <pre>.   .   ~ Q</pre> | E -> | 5, 2 |
| 8 |  | <pre>.   ~ P</pre> | subproof I/E ~ | 2 - 7 |
| 9 |  | <pre>( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P</pre> | subproof I -> | 1 - 8 |

## ( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q <a name="proof-4-5"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P -> Q ) ^ ( ~ P -> Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P -> Q</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ P -> Q</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 5 |  | <pre>.   Q</pre> | E v | 2, 3, 4 |
| 6 |  | <pre>( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q</pre> | subproof I -> | 1 - 5 |

## ( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) ) <a name="proof-4-6"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> ( Q -> R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P -> ( Q -> R )</pre> | Reiterate | 1 |
| 5 |  | <pre>.   .   .   Q -> R</pre> | E -> | 4, 3 |
| 6 |  | <pre>.   .   .   Q</pre> | Reiterate | 2 |
| 7 |  | <pre>.   .   .   R</pre> | E -> | 5, 6 |
| 8 |  | <pre>.   .   P -> R</pre> | subproof I -> | 3 - 7 |
| 9 |  | <pre>.   Q -> ( P -> R )</pre> | subproof I -> | 2 - 8 |
| 10 |  | <pre>( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )</pre> | subproof I -> | 1 - 9 |

# transitivity <a name="transitivity"></a>

## ( P -> Q ) -> ( ( Q -> R ) -> ( P -> R ) ) <a name="proof-5-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   Q -> R</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P -> Q</pre> | Reiterate | 1 |
| 5 |  | <pre>.   .   .   Q -> R</pre> | Reiterate | 2 |
| 6 |  | <pre>.   .   .   Q</pre> | E -> | 4, 3 |
| 7 |  | <pre>.   .   .   R</pre> | E -> | 5, 6 |
| 8 |  | <pre>.   .   P -> R</pre> | subproof I -> | 3 - 7 |
| 9 |  | <pre>.   ( Q -> R ) -> ( P -> R )</pre> | subproof I -> | 2 - 8 |
| 10 |  | <pre>( P -> Q ) -> ( ( Q -> R ) -> ( P -> R ) )</pre> | subproof I -> | 1 - 9 |

## ( P <-> Q ) -> ( ( Q <-> R ) -> ( P <-> R ) ) <a name="proof-5-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P <-> Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   Q <-> R</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P <-> Q</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   P -> Q</pre> | E <-> (L) | 3 |
| 5 |  | <pre>.   .   Q -> P</pre> | E <-> (R) | 3 |
| 6 |  | <pre>.   .   Q -> R</pre> | E <-> (L) | 2 |
| 7 |  | <pre>.   .   R -> Q</pre> | E <-> (R) | 2 |
| 8 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 9 |  | <pre>.   .   .   P -> Q</pre> | Reiterate | 4 |
| 10 |  | <pre>.   .   .   Q -> R</pre> | Reiterate | 6 |
| 11 |  | <pre>.   .   .   Q</pre> | E -> | 9, 8 |
| 12 |  | <pre>.   .   .   R</pre> | E -> | 10, 11 |
| 13 |  | <pre>.   .   P -> R</pre> | subproof I -> | 8 - 12 |
| 14 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 15 |  | <pre>.   .   .   R -> Q</pre> | Reiterate | 7 |
| 16 |  | <pre>.   .   .   Q -> P</pre> | Reiterate | 5 |
| 17 |  | <pre>.   .   .   Q</pre> | E -> | 15, 14 |
| 18 |  | <pre>.   .   .   P</pre> | E -> | 16, 17 |
| 19 |  | <pre>.   .   R -> P</pre> | subproof I -> | 14 - 18 |
| 20 |  | <pre>.   .   P <-> R</pre> | I <-> | 13, 19 |
| 21 |  | <pre>.   ( Q <-> R ) -> ( P <-> R )</pre> | subproof I -> | 2 - 20 |
| 22 |  | <pre>( P <-> Q ) -> ( ( Q <-> R ) -> ( P <-> R ) )</pre> | subproof I -> | 1 - 21 |

# commutativity <a name="commutativity"></a>

## ( P ^ Q ) -> ( Q ^ P ) <a name="proof-6-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   Q</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   Q ^ P</pre> | I ^ | 3, 2 |
| 5 |  | <pre>( P ^ Q ) -> ( Q ^ P )</pre> | subproof I -> | 1 - 4 |

## ( P v Q ) -> ( Q v P ) <a name="proof-6-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   Q v P</pre> | I v (R) | 2 |
| 4 |  | <pre>.   P -> ( Q v P )</pre> | subproof I -> | 2 - 3 |
| 5 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 6 |  | <pre>.   .   Q v P</pre> | I v (L) | 5 |
| 7 |  | <pre>.   Q -> ( Q v P )</pre> | subproof I -> | 5 - 6 |
| 8 |  | <pre>.   Q v P</pre> | E v | 4, 7, 1 |
| 9 |  | <pre>( P v Q ) -> ( Q v P )</pre> | subproof I -> | 1 - 8 |

## ( P <-> Q ) -> ( Q <-> P ) <a name="proof-6-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P <-> Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P -> Q</pre> | E <-> (L) | 1 |
| 3 |  | <pre>.   Q -> P</pre> | E <-> (R) | 1 |
| 4 |  | <pre>.   Q <-> P</pre> | I <-> | 3, 2 |
| 5 |  | <pre>( P <-> Q ) -> ( Q <-> P )</pre> | subproof I -> | 1 - 4 |

# associativity <a name="associativity"></a>

## ( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) ) <a name="proof-7-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P ^ Q ) ^ R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P ^ Q</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   R</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P</pre> | E ^ (L) | 2 |
| 5 |  | <pre>.   Q</pre> | E ^ (R) | 2 |
| 6 |  | <pre>.   Q ^ R</pre> | I ^ | 5, 3 |
| 7 |  | <pre>.   P ^ ( Q ^ R )</pre> | I ^ | 4, 6 |
| 8 |  | <pre>( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )</pre> | subproof I -> | 1 - 7 |
| 9 |  | <pre>.   P ^ ( Q ^ R )</pre> | Assume: I -> |  |
| 10 |  | <pre>.   P</pre> | E ^ (L) | 9 |
| 11 |  | <pre>.   Q ^ R</pre> | E ^ (R) | 9 |
| 12 |  | <pre>.   Q</pre> | E ^ (L) | 11 |
| 13 |  | <pre>.   R</pre> | E ^ (R) | 11 |
| 14 |  | <pre>.   P ^ Q</pre> | I ^ | 10, 12 |
| 15 |  | <pre>.   ( P ^ Q ) ^ R</pre> | I ^ | 14, 13 |
| 16 |  | <pre>( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )</pre> | subproof I -> | 9 - 15 |
| 17 |  | <pre>( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) )</pre> | I <-> | 8, 16 |

## ( ( P v Q ) v R ) <-> ( P v ( Q v R ) ) <a name="proof-7-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P v Q ) v R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P v Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P v ( Q v R )</pre> | I v (L) | 3 |
| 5 |  | <pre>.   .   P -> ( P v ( Q v R ) )</pre> | subproof I -> | 3 - 4 |
| 6 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   .   Q v R</pre> | I v (L) | 6 |
| 8 |  | <pre>.   .   .   P v ( Q v R )</pre> | I v (R) | 7 |
| 9 |  | <pre>.   .   Q -> ( P v ( Q v R ) )</pre> | subproof I -> | 6 - 8 |
| 10 |  | <pre>.   .   P v ( Q v R )</pre> | E v | 5, 9, 2 |
| 11 |  | <pre>.   ( P v Q ) -> ( P v ( Q v R ) )</pre> | subproof I -> | 2 - 10 |
| 12 |  | <pre>.   .   R</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   Q v R</pre> | I v (R) | 12 |
| 14 |  | <pre>.   .   P v ( Q v R )</pre> | I v (R) | 13 |
| 15 |  | <pre>.   R -> ( P v ( Q v R ) )</pre> | subproof I -> | 12 - 14 |
| 16 |  | <pre>.   P v ( Q v R )</pre> | E v | 11, 15, 1 |
| 17 |  | <pre>( ( P v Q ) v R ) -> ( P v ( Q v R ) )</pre> | subproof I -> | 1 - 16 |
| 18 |  | <pre>.   P v ( Q v R )</pre> | Assume: I -> |  |
| 19 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 20 |  | <pre>.   .   P v Q</pre> | I v (L) | 19 |
| 21 |  | <pre>.   .   ( P v Q ) v R</pre> | I v (L) | 20 |
| 22 |  | <pre>.   P -> ( ( P v Q ) v R )</pre> | subproof I -> | 19 - 21 |
| 23 |  | <pre>.   .   Q v R</pre> | Assume: I -> |  |
| 24 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 25 |  | <pre>.   .   .   P v Q</pre> | I v (R) | 24 |
| 26 |  | <pre>.   .   .   ( P v Q ) v R</pre> | I v (L) | 25 |
| 27 |  | <pre>.   .   Q -> ( ( P v Q ) v R )</pre> | subproof I -> | 24 - 26 |
| 28 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 29 |  | <pre>.   .   .   ( P v Q ) v R</pre> | I v (R) | 28 |
| 30 |  | <pre>.   .   R -> ( ( P v Q ) v R )</pre> | subproof I -> | 28 - 29 |
| 31 |  | <pre>.   .   ( P v Q ) v R</pre> | E v | 27, 30, 23 |
| 32 |  | <pre>.   ( Q v R ) -> ( ( P v Q ) v R )</pre> | subproof I -> | 23 - 31 |
| 33 |  | <pre>.   ( P v Q ) v R</pre> | E v | 22, 32, 18 |
| 34 |  | <pre>( P v ( Q v R ) ) -> ( ( P v Q ) v R )</pre> | subproof I -> | 18 - 33 |
| 35 |  | <pre>( ( P v Q ) v R ) <-> ( P v ( Q v R ) )</pre> | I <-> | 17, 34 |

## ( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) ) <a name="proof-7-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P <-> Q ) <-> R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   ( P <-> Q ) -> R</pre> | E <-> (L) | 1 |
| 3 |  | <pre>.   R -> ( P <-> Q )</pre> | E <-> (R) | 1 |
| 4 |  | <pre>.   .   ~ ( P -> ( Q <-> R ) )</pre> | Assume: I/E ~ |  |
| 5 |  | <pre>.   .   P ^ ~ ( Q <-> R )</pre> | Theorem: (~ ->) to ^ | 4 |
| 6 |  | <pre>.   .   P</pre> | E ^ (L) | 5 |
| 7 |  | <pre>.   .   ~ ( Q <-> R )</pre> | E ^ (R) | 5 |
| 8 |  | <pre>.   .   Q <-> ~ R</pre> | Theorem: <-> negation | 7 |
| 9 |  | <pre>.   .   .   R</pre> | Assume: I/E ~ |  |
| 10 |  | <pre>.   .   .   R -> ( P <-> Q )</pre> | Reiterate | 3 |
| 11 |  | <pre>.   .   .   P <-> Q</pre> | E -> | 10, 9 |
| 12 |  | <pre>.   .   .   P -> Q</pre> | E <-> (L) | 11 |
| 13 |  | <pre>.   .   .   P</pre> | Reiterate | 6 |
| 14 |  | <pre>.   .   .   Q</pre> | E -> | 12, 13 |
| 15 |  | <pre>.   .   .   Q <-> ~ R</pre> | Reiterate | 8 |
| 16 |  | <pre>.   .   .   Q -> ~ R</pre> | E <-> (L) | 15 |
| 17 |  | <pre>.   .   .   ~ R</pre> | E -> | 16, 14 |
| 18 |  | <pre>.   .   ~ R</pre> | subproof I/E ~ | 9 - 17 |
| 19 |  | <pre>.   .   ~ R -> Q</pre> | E <-> (R) | 8 |
| 20 |  | <pre>.   .   Q</pre> | E -> | 19, 18 |
| 21 |  | <pre>.   .   ( P <-> Q ) -> R</pre> | Reiterate | 2 |
| 22 |  | <pre>.   .   ~ R -> ~ ( P <-> Q )</pre> | Theorem: contrapositive | 21 |
| 23 |  | <pre>.   .   ~ ( P <-> Q )</pre> | E -> | 22, 18 |
| 24 |  | <pre>.   .   P <-> ~ Q</pre> | Theorem: <-> negation | 23 |
| 25 |  | <pre>.   .   P -> ~ Q</pre> | E <-> (L) | 24 |
| 26 |  | <pre>.   .   ~ Q</pre> | E -> | 25, 6 |
| 27 |  | <pre>.   P -> ( Q <-> R )</pre> | subproof I/E ~ | 4 - 26 |
| 28 |  | <pre>.   .   ~ ( ( Q <-> R ) -> P )</pre> | Assume: I/E ~ |  |
| 29 |  | <pre>.   .   ( Q <-> R ) ^ ~ P</pre> | Theorem: (~ ->) to ^ | 28 |
| 30 |  | <pre>.   .   Q <-> R</pre> | E ^ (L) | 29 |
| 31 |  | <pre>.   .   ~ P</pre> | E ^ (R) | 29 |
| 32 |  | <pre>.   .   .   R</pre> | Assume: I/E ~ |  |
| 33 |  | <pre>.   .   .   R -> ( P <-> Q )</pre> | Reiterate | 3 |
| 34 |  | <pre>.   .   .   Q <-> R</pre> | Reiterate | 30 |
| 35 |  | <pre>.   .   .   R -> Q</pre> | E <-> (R) | 34 |
| 36 |  | <pre>.   .   .   Q</pre> | E -> | 35, 32 |
| 37 |  | <pre>.   .   .   P <-> Q</pre> | E -> | 33, 32 |
| 38 |  | <pre>.   .   .   Q -> P</pre> | E <-> (R) | 37 |
| 39 |  | <pre>.   .   .   P</pre> | E -> | 38, 36 |
| 40 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 31 |
| 41 |  | <pre>.   .   ~ R</pre> | subproof I/E ~ | 32 - 40 |
| 42 |  | <pre>.   .   Q -> R</pre> | E <-> (L) | 30 |
| 43 |  | <pre>.   .   ~ R -> ~ Q</pre> | Theorem: contrapositive | 42 |
| 44 |  | <pre>.   .   ~ Q</pre> | E -> | 43, 41 |
| 45 |  | <pre>.   .   ( P <-> Q ) -> R</pre> | Reiterate | 2 |
| 46 |  | <pre>.   .   ~ R -> ~ ( P <-> Q )</pre> | Theorem: contrapositive | 45 |
| 47 |  | <pre>.   .   ~ ( P <-> Q )</pre> | E -> | 46, 41 |
| 48 |  | <pre>.   .   P <-> ~ Q</pre> | Theorem: <-> negation | 47 |
| 49 |  | <pre>.   .   ~ Q -> P</pre> | E <-> (R) | 48 |
| 50 |  | <pre>.   .   P</pre> | E -> | 49, 44 |
| 51 |  | <pre>.   ( Q <-> R ) -> P</pre> | subproof I/E ~ | 28 - 50 |
| 52 |  | <pre>.   P <-> ( Q <-> R )</pre> | I <-> | 27, 51 |
| 53 |  | <pre>( ( P <-> Q ) <-> R ) -> ( P <-> ( Q <-> R ) )</pre> | subproof I -> | 1 - 52 |
| 54 |  | <pre>.   P <-> ( Q <-> R )</pre> | Assume: I -> |  |
| 55 |  | <pre>.   P -> ( Q <-> R )</pre> | E <-> (L) | 54 |
| 56 |  | <pre>.   ( Q <-> R ) -> P</pre> | E <-> (R) | 54 |
| 57 |  | <pre>.   .   ~ ( ( P <-> Q ) -> R )</pre> | Assume: I/E ~ |  |
| 58 |  | <pre>.   .   ( P <-> Q ) ^ ~ R</pre> | Theorem: (~ ->) to ^ | 57 |
| 59 |  | <pre>.   .   P <-> Q</pre> | E ^ (L) | 58 |
| 60 |  | <pre>.   .   ~ R</pre> | E ^ (R) | 58 |
| 61 |  | <pre>.   .   P -> Q</pre> | E <-> (L) | 59 |
| 62 |  | <pre>.   .   .   P</pre> | Assume: I/E ~ |  |
| 63 |  | <pre>.   .   .   P -> Q</pre> | Reiterate | 61 |
| 64 |  | <pre>.   .   .   Q</pre> | E -> | 63, 62 |
| 65 |  | <pre>.   .   .   P -> ( Q <-> R )</pre> | Reiterate | 55 |
| 66 |  | <pre>.   .   .   Q <-> R</pre> | E -> | 65, 62 |
| 67 |  | <pre>.   .   .   Q -> R</pre> | E <-> (L) | 66 |
| 68 |  | <pre>.   .   .   R</pre> | E -> | 67, 64 |
| 69 |  | <pre>.   .   .   ~ R</pre> | Reiterate | 60 |
| 70 |  | <pre>.   .   ~ P</pre> | subproof I/E ~ | 62 - 69 |
| 71 |  | <pre>.   .   ( Q <-> R ) -> P</pre> | Reiterate | 56 |
| 72 |  | <pre>.   .   ~ P -> ~ ( Q <-> R )</pre> | Theorem: contrapositive | 71 |
| 73 |  | <pre>.   .   ~ ( Q <-> R )</pre> | E -> | 72, 70 |
| 74 |  | <pre>.   .   Q <-> ~ R</pre> | Theorem: <-> negation | 73 |
| 75 |  | <pre>.   .   ~ R -> Q</pre> | E <-> (R) | 74 |
| 76 |  | <pre>.   .   Q</pre> | E -> | 75, 60 |
| 77 |  | <pre>.   .   Q -> P</pre> | E <-> (R) | 59 |
| 78 |  | <pre>.   .   P</pre> | E -> | 77, 76 |
| 79 |  | <pre>.   ( P <-> Q ) -> R</pre> | subproof I/E ~ | 57 - 78 |
| 80 |  | <pre>.   .   ~ ( R -> ( P <-> Q ) )</pre> | Assume: I/E ~ |  |
| 81 |  | <pre>.   .   R ^ ~ ( P <-> Q )</pre> | Theorem: (~ ->) to ^ | 80 |
| 82 |  | <pre>.   .   R</pre> | E ^ (L) | 81 |
| 83 |  | <pre>.   .   ~ ( P <-> Q )</pre> | E ^ (R) | 81 |
| 84 |  | <pre>.   .   P <-> ~ Q</pre> | Theorem: <-> negation | 83 |
| 85 |  | <pre>.   .   P -> ~ Q</pre> | E <-> (L) | 84 |
| 86 |  | <pre>.   .   .   P</pre> | Assume: I/E ~ |  |
| 87 |  | <pre>.   .   .   P -> ~ Q</pre> | Reiterate | 85 |
| 88 |  | <pre>.   .   .   ~ Q</pre> | E -> | 87, 86 |
| 89 |  | <pre>.   .   .   P -> ( Q <-> R )</pre> | Reiterate | 55 |
| 90 |  | <pre>.   .   .   Q <-> R</pre> | E -> | 89, 86 |
| 91 |  | <pre>.   .   .   R -> Q</pre> | E <-> (R) | 90 |
| 92 |  | <pre>.   .   .   R</pre> | Reiterate | 82 |
| 93 |  | <pre>.   .   .   Q</pre> | E -> | 91, 92 |
| 94 |  | <pre>.   .   ~ P</pre> | subproof I/E ~ | 86 - 93 |
| 95 |  | <pre>.   .   ~ Q -> P</pre> | E <-> (R) | 84 |
| 96 |  | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 95 |
| 97 |  | <pre>.   .   Q</pre> | E -> | 96, 94 |
| 98 |  | <pre>.   .   ( Q <-> R ) -> P</pre> | Reiterate | 56 |
| 99 |  | <pre>.   .   ~ P -> ~ ( Q <-> R )</pre> | Theorem: contrapositive | 98 |
| 100 |  | <pre>.   .   ~ ( Q <-> R )</pre> | E -> | 99, 94 |
| 101 |  | <pre>.   .   Q <-> ~ R</pre> | Theorem: <-> negation | 100 |
| 102 |  | <pre>.   .   Q -> ~ R</pre> | E <-> (L) | 101 |
| 103 |  | <pre>.   .   ~ R</pre> | E -> | 102, 97 |
| 104 |  | <pre>.   R -> ( P <-> Q )</pre> | subproof I/E ~ | 80 - 103 |
| 105 |  | <pre>.   ( P <-> Q ) <-> R</pre> | I <-> | 79, 104 |
| 106 |  | <pre>( P <-> ( Q <-> R ) ) -> ( ( P <-> Q ) <-> R )</pre> | subproof I -> | 54 - 105 |
| 107 |  | <pre>( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) )</pre> | I <-> | 53, 106 |

# distributivity <a name="distributivity"></a>

## ( P -> ( Q -> R ) ) <-> ( ( P -> Q ) -> ( P -> R ) ) <a name="proof-8-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> ( Q -> R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P -> Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P -> ( Q -> R )</pre> | Reiterate | 1 |
| 5 |  | <pre>.   .   .   Q -> R</pre> | E -> | 4, 3 |
| 6 |  | <pre>.   .   .   P -> Q</pre> | Reiterate | 2 |
| 7 |  | <pre>.   .   .   Q</pre> | E -> | 6, 3 |
| 8 |  | <pre>.   .   .   R</pre> | E -> | 5, 7 |
| 9 |  | <pre>.   .   P -> R</pre> | subproof I -> | 3 - 8 |
| 10 |  | <pre>.   ( P -> Q ) -> ( P -> R )</pre> | subproof I -> | 2 - 9 |
| 11 |  | <pre>( P -> ( Q -> R ) ) -> ( ( P -> Q ) -> ( P -> R ) )</pre> | subproof I -> | 1 - 10 |
| 12 |  | <pre>.   ( P -> Q ) -> ( P -> R )</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 14 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 15 |  | <pre>.   .   .   ( P -> Q ) -> ( P -> R )</pre> | Reiterate | 12 |
| 16 |  | <pre>.   .   .   P</pre> | Reiterate | 13 |
| 17 |  | <pre>.   .   .   P -> Q</pre> | I -> | 16, 14 |
| 18 |  | <pre>.   .   .   P -> R</pre> | E -> | 15, 17 |
| 19 |  | <pre>.   .   .   R</pre> | E -> | 18, 16 |
| 20 |  | <pre>.   .   Q -> R</pre> | subproof I -> | 14 - 19 |
| 21 |  | <pre>.   P -> ( Q -> R )</pre> | subproof I -> | 13 - 20 |
| 22 |  | <pre>( ( P -> Q ) -> ( P -> R ) ) -> ( P -> ( Q -> R ) )</pre> | subproof I -> | 12 - 21 |
| 23 |  | <pre>( P -> ( Q -> R ) ) <-> ( ( P -> Q ) -> ( P -> R ) )</pre> | I <-> | 11, 22 |

## ( P -> ( Q v R ) ) <-> ( ( P -> Q ) v ( P -> R ) ) <a name="proof-8-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> ( Q v R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P -> ( Q v R )</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   Q v R</pre> | E -> | 3, 2 |
| 5 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 6 |  | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 7 |  | <pre>.   .   .   P -> Q</pre> | I -> | 6, 5 |
| 8 |  | <pre>.   .   .   ( P -> Q ) v ( P -> R )</pre> | I v (L) | 7 |
| 9 |  | <pre>.   .   Q -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof I -> | 5 - 8 |
| 10 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 11 |  | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 12 |  | <pre>.   .   .   P -> R</pre> | I -> | 11, 10 |
| 13 |  | <pre>.   .   .   ( P -> Q ) v ( P -> R )</pre> | I v (R) | 12 |
| 14 |  | <pre>.   .   R -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof I -> | 10 - 13 |
| 15 |  | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | E v | 9, 14, 4 |
| 16 |  | <pre>.   P -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof I -> | 2 - 15 |
| 17 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 18 |  | <pre>.   .   .   ~ Q</pre> | Assume: I -> |  |
| 19 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 17 |
| 20 |  | <pre>.   .   ~ Q -> ~ P</pre> | subproof I -> | 18 - 19 |
| 21 |  | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 20 |
| 22 |  | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | I v (L) | 21 |
| 23 |  | <pre>.   ~ P -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof I -> | 17 - 22 |
| 24 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 25 |  | <pre>.   ( P -> Q ) v ( P -> R )</pre> | E v | 16, 23, 24 |
| 26 |  | <pre>( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof I -> | 1 - 25 |
| 27 |  | <pre>.   ( P -> Q ) v ( P -> R )</pre> | Assume: I -> |  |
| 28 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 29 |  | <pre>.   .   .   P -> Q</pre> | Assume: I -> |  |
| 30 |  | <pre>.   .   .   P</pre> | Reiterate | 28 |
| 31 |  | <pre>.   .   .   Q</pre> | E -> | 29, 30 |
| 32 |  | <pre>.   .   .   Q v R</pre> | I v (L) | 31 |
| 33 |  | <pre>.   .   ( P -> Q ) -> ( Q v R )</pre> | subproof I -> | 29 - 32 |
| 34 |  | <pre>.   .   .   P -> R</pre> | Assume: I -> |  |
| 35 |  | <pre>.   .   .   P</pre> | Reiterate | 28 |
| 36 |  | <pre>.   .   .   R</pre> | E -> | 34, 35 |
| 37 |  | <pre>.   .   .   Q v R</pre> | I v (R) | 36 |
| 38 |  | <pre>.   .   ( P -> R ) -> ( Q v R )</pre> | subproof I -> | 34 - 37 |
| 39 |  | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | Reiterate | 27 |
| 40 |  | <pre>.   .   Q v R</pre> | E v | 33, 38, 39 |
| 41 |  | <pre>.   P -> ( Q v R )</pre> | subproof I -> | 28 - 40 |
| 42 |  | <pre>( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )</pre> | subproof I -> | 27 - 41 |
| 43 |  | <pre>( P -> ( Q v R ) ) <-> ( ( P -> Q ) v ( P -> R ) )</pre> | I <-> | 26, 42 |

## ( P -> ( Q ^ R ) ) <-> ( ( P -> Q ) ^ ( P -> R ) ) <a name="proof-8-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> ( Q ^ R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   Q ^ R</pre> | E -> | 3, 2 |
| 5 |  | <pre>.   .   Q</pre> | E ^ (L) | 4 |
| 6 |  | <pre>.   P -> Q</pre> | subproof I -> | 2 - 5 |
| 7 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 8 |  | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 9 |  | <pre>.   .   Q ^ R</pre> | E -> | 8, 7 |
| 10 |  | <pre>.   .   R</pre> | E ^ (R) | 9 |
| 11 |  | <pre>.   P -> R</pre> | subproof I -> | 7 - 10 |
| 12 |  | <pre>.   ( P -> Q ) ^ ( P -> R )</pre> | I ^ | 6, 11 |
| 13 |  | <pre>( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )</pre> | subproof I -> | 1 - 12 |
| 14 |  | <pre>.   ( P -> Q ) ^ ( P -> R )</pre> | Assume: I -> |  |
| 15 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 16 |  | <pre>.   .   ( P -> Q ) ^ ( P -> R )</pre> | Reiterate | 14 |
| 17 |  | <pre>.   .   P -> Q</pre> | E ^ (L) | 16 |
| 18 |  | <pre>.   .   P -> R</pre> | E ^ (R) | 16 |
| 19 |  | <pre>.   .   Q</pre> | E -> | 17, 15 |
| 20 |  | <pre>.   .   R</pre> | E -> | 18, 15 |
| 21 |  | <pre>.   .   Q ^ R</pre> | I ^ | 19, 20 |
| 22 |  | <pre>.   P -> ( Q ^ R )</pre> | subproof I -> | 15 - 21 |
| 23 |  | <pre>( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )</pre> | subproof I -> | 14 - 22 |
| 24 |  | <pre>( P -> ( Q ^ R ) ) <-> ( ( P -> Q ) ^ ( P -> R ) )</pre> | I <-> | 13, 23 |

## ( P ^ ( Q -> R ) ) -> ( ( P ^ Q ) -> ( P ^ R ) ) <a name="proof-8-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ ( Q -> R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P ^ Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P ^ ( Q -> R )</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 5 |  | <pre>.   .   Q</pre> | E ^ (R) | 2 |
| 6 |  | <pre>.   .   Q -> R</pre> | E ^ (R) | 3 |
| 7 |  | <pre>.   .   R</pre> | E -> | 6, 5 |
| 8 |  | <pre>.   .   P ^ R</pre> | I ^ | 4, 7 |
| 9 |  | <pre>.   ( P ^ Q ) -> ( P ^ R )</pre> | subproof I -> | 2 - 8 |
| 10 |  | <pre>( P ^ ( Q -> R ) ) -> ( ( P ^ Q ) -> ( P ^ R ) )</pre> | subproof I -> | 1 - 9 |

## ( P ^ ( Q v R ) ) <-> ( ( P ^ Q ) v ( P ^ R ) ) <a name="proof-8-5"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ ( Q v R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   Q v R</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 5 |  | <pre>.   .   P</pre> | Reiterate | 2 |
| 6 |  | <pre>.   .   P ^ Q</pre> | I ^ | 5, 4 |
| 7 |  | <pre>.   .   ( P ^ Q ) v ( P ^ R )</pre> | I v (L) | 6 |
| 8 |  | <pre>.   Q -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof I -> | 4 - 7 |
| 9 |  | <pre>.   .   R</pre> | Assume: I -> |  |
| 10 |  | <pre>.   .   P</pre> | Reiterate | 2 |
| 11 |  | <pre>.   .   P ^ R</pre> | I ^ | 10, 9 |
| 12 |  | <pre>.   .   ( P ^ Q ) v ( P ^ R )</pre> | I v (R) | 11 |
| 13 |  | <pre>.   R -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof I -> | 9 - 12 |
| 14 |  | <pre>.   ( P ^ Q ) v ( P ^ R )</pre> | E v | 8, 13, 3 |
| 15 |  | <pre>( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof I -> | 1 - 14 |
| 16 |  | <pre>.   ( P ^ Q ) v ( P ^ R )</pre> | Assume: I -> |  |
| 17 |  | <pre>.   .   P ^ Q</pre> | Assume: I -> |  |
| 18 |  | <pre>.   .   P</pre> | E ^ (L) | 17 |
| 19 |  | <pre>.   .   Q</pre> | E ^ (R) | 17 |
| 20 |  | <pre>.   .   Q v R</pre> | I v (L) | 19 |
| 21 |  | <pre>.   .   P ^ ( Q v R )</pre> | I ^ | 18, 20 |
| 22 |  | <pre>.   ( P ^ Q ) -> ( P ^ ( Q v R ) )</pre> | subproof I -> | 17 - 21 |
| 23 |  | <pre>.   .   P ^ R</pre> | Assume: I -> |  |
| 24 |  | <pre>.   .   P</pre> | E ^ (L) | 23 |
| 25 |  | <pre>.   .   R</pre> | E ^ (R) | 23 |
| 26 |  | <pre>.   .   Q v R</pre> | I v (R) | 25 |
| 27 |  | <pre>.   .   P ^ ( Q v R )</pre> | I ^ | 24, 26 |
| 28 |  | <pre>.   ( P ^ R ) -> ( P ^ ( Q v R ) )</pre> | subproof I -> | 23 - 27 |
| 29 |  | <pre>.   P ^ ( Q v R )</pre> | E v | 22, 28, 16 |
| 30 |  | <pre>( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )</pre> | subproof I -> | 16 - 29 |
| 31 |  | <pre>( P ^ ( Q v R ) ) <-> ( ( P ^ Q ) v ( P ^ R ) )</pre> | I <-> | 15, 30 |

## ( P ^ ( Q ^ R ) ) <-> ( ( P ^ Q ) ^ ( P ^ R ) ) <a name="proof-8-6"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ ( Q ^ R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   Q ^ R</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   Q</pre> | E ^ (L) | 3 |
| 5 |  | <pre>.   R</pre> | E ^ (R) | 3 |
| 6 |  | <pre>.   P ^ Q</pre> | I ^ | 2, 4 |
| 7 |  | <pre>.   P ^ R</pre> | I ^ | 2, 5 |
| 8 |  | <pre>.   ( P ^ Q ) ^ ( P ^ R )</pre> | I ^ | 6, 7 |
| 9 |  | <pre>( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ ( P ^ R ) )</pre> | subproof I -> | 1 - 8 |
| 10 |  | <pre>.   ( P ^ Q ) ^ ( P ^ R )</pre> | Assume: I -> |  |
| 11 |  | <pre>.   P ^ Q</pre> | E ^ (L) | 10 |
| 12 |  | <pre>.   P ^ R</pre> | E ^ (R) | 10 |
| 13 |  | <pre>.   P</pre> | E ^ (L) | 11 |
| 14 |  | <pre>.   Q</pre> | E ^ (R) | 11 |
| 15 |  | <pre>.   R</pre> | E ^ (R) | 12 |
| 16 |  | <pre>.   Q ^ R</pre> | I ^ | 14, 15 |
| 17 |  | <pre>.   P ^ ( Q ^ R )</pre> | I ^ | 13, 16 |
| 18 |  | <pre>( ( P ^ Q ) ^ ( P ^ R ) ) -> ( P ^ ( Q ^ R ) )</pre> | subproof I -> | 10 - 17 |
| 19 |  | <pre>( P ^ ( Q ^ R ) ) <-> ( ( P ^ Q ) ^ ( P ^ R ) )</pre> | I <-> | 9, 18 |

## ( P v ( Q -> R ) ) <-> ( ( P v Q ) -> ( P v R ) ) <a name="proof-8-7"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v ( Q -> R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P v Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P v R</pre> | I v (L) | 3 |
| 5 |  | <pre>.   .   P -> ( P v R )</pre> | subproof I -> | 3 - 4 |
| 6 |  | <pre>.   .   .   ~ P</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   .   P v ( Q -> R )</pre> | Reiterate | 1 |
| 8 |  | <pre>.   .   .   ~ P -> ( Q -> R )</pre> | Theorem: v to -> | 7 |
| 9 |  | <pre>.   .   .   Q -> R</pre> | E -> | 8, 6 |
| 10 |  | <pre>.   .   .   P v Q</pre> | Reiterate | 2 |
| 11 |  | <pre>.   .   .   ~ P -> Q</pre> | Theorem: v to -> | 10 |
| 12 |  | <pre>.   .   .   Q</pre> | E -> | 11, 6 |
| 13 |  | <pre>.   .   .   R</pre> | E -> | 9, 12 |
| 14 |  | <pre>.   .   .   P v R</pre> | I v (R) | 13 |
| 15 |  | <pre>.   .   ~ P -> ( P v R )</pre> | subproof I -> | 6 - 14 |
| 16 |  | <pre>.   .   P v ~ P</pre> | Theorem: excluded middle |  |
| 17 |  | <pre>.   .   P v R</pre> | E v | 5, 15, 16 |
| 18 |  | <pre>.   ( P v Q ) -> ( P v R )</pre> | subproof I -> | 2 - 17 |
| 19 |  | <pre>( P v ( Q -> R ) ) -> ( ( P v Q ) -> ( P v R ) )</pre> | subproof I -> | 1 - 18 |
| 20 |  | <pre>.   ( P v Q ) -> ( P v R )</pre> | Assume: I -> |  |
| 21 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 22 |  | <pre>.   .   P v ( Q -> R )</pre> | I v (L) | 21 |
| 23 |  | <pre>.   P -> ( P v ( Q -> R ) )</pre> | subproof I -> | 21 - 22 |
| 24 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 25 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 26 |  | <pre>.   .   .   ( P v Q ) -> ( P v R )</pre> | Reiterate | 20 |
| 27 |  | <pre>.   .   .   P v Q</pre> | I v (R) | 25 |
| 28 |  | <pre>.   .   .   P v R</pre> | E -> | 26, 27 |
| 29 |  | <pre>.   .   .   ~ P -> R</pre> | Theorem: v to -> | 28 |
| 30 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 24 |
| 31 |  | <pre>.   .   .   R</pre> | E -> | 29, 30 |
| 32 |  | <pre>.   .   .   Q -> R</pre> | I -> | 25, 31 |
| 33 |  | <pre>.   .   .   P v ( Q -> R )</pre> | I v (R) | 32 |
| 34 |  | <pre>.   .   Q -> ( P v ( Q -> R ) )</pre> | subproof I -> | 25 - 33 |
| 35 |  | <pre>.   .   .   ~ Q</pre> | Assume: I -> |  |
| 36 |  | <pre>.   .   .   ~ Q v R</pre> | I v (L) | 35 |
| 37 |  | <pre>.   .   .   Q -> R</pre> | Theorem: v to -> | 36 |
| 38 |  | <pre>.   .   .   P v ( Q -> R )</pre> | I v (R) | 37 |
| 39 |  | <pre>.   .   ~ Q -> ( P v ( Q -> R ) )</pre> | subproof I -> | 35 - 38 |
| 40 |  | <pre>.   .   Q v ~ Q</pre> | Theorem: excluded middle |  |
| 41 |  | <pre>.   .   P v ( Q -> R )</pre> | E v | 34, 39, 40 |
| 42 |  | <pre>.   ~ P -> ( P v ( Q -> R ) )</pre> | subproof I -> | 24 - 41 |
| 43 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 44 |  | <pre>.   P v ( Q -> R )</pre> | E v | 23, 42, 43 |
| 45 |  | <pre>( ( P v Q ) -> ( P v R ) ) -> ( P v ( Q -> R ) )</pre> | subproof I -> | 20 - 44 |
| 46 |  | <pre>( P v ( Q -> R ) ) <-> ( ( P v Q ) -> ( P v R ) )</pre> | I <-> | 19, 45 |

## ( P v ( Q v R ) ) <-> ( ( P v Q ) v ( P v R ) ) <a name="proof-8-8"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v ( Q v R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 |  | <pre>.   .   ( P v Q ) v ( P v R )</pre> | I v (L) | 3 |
| 5 |  | <pre>.   P -> ( ( P v Q ) v ( P v R ) )</pre> | subproof I -> | 2 - 4 |
| 6 |  | <pre>.   .   Q v R</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 8 |  | <pre>.   .   .   P v Q</pre> | I v (R) | 7 |
| 9 |  | <pre>.   .   .   ( P v Q ) v ( P v R )</pre> | I v (L) | 8 |
| 10 |  | <pre>.   .   Q -> ( ( P v Q ) v ( P v R ) )</pre> | subproof I -> | 7 - 9 |
| 11 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 12 |  | <pre>.   .   .   P v R</pre> | I v (R) | 11 |
| 13 |  | <pre>.   .   .   ( P v Q ) v ( P v R )</pre> | I v (R) | 12 |
| 14 |  | <pre>.   .   R -> ( ( P v Q ) v ( P v R ) )</pre> | subproof I -> | 11 - 13 |
| 15 |  | <pre>.   .   ( P v Q ) v ( P v R )</pre> | E v | 10, 14, 6 |
| 16 |  | <pre>.   ( Q v R ) -> ( ( P v Q ) v ( P v R ) )</pre> | subproof I -> | 6 - 15 |
| 17 |  | <pre>.   ( P v Q ) v ( P v R )</pre> | E v | 5, 16, 1 |
| 18 |  | <pre>( P v ( Q v R ) ) -> ( ( P v Q ) v ( P v R ) )</pre> | subproof I -> | 1 - 17 |
| 19 |  | <pre>.   ( P v Q ) v ( P v R )</pre> | Assume: I -> |  |
| 20 |  | <pre>.   .   P v Q</pre> | Assume: I -> |  |
| 21 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 22 |  | <pre>.   .   .   P v ( Q v R )</pre> | I v (L) | 21 |
| 23 |  | <pre>.   .   P -> ( P v ( Q v R ) )</pre> | subproof I -> | 21 - 22 |
| 24 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 25 |  | <pre>.   .   .   Q v R</pre> | I v (L) | 24 |
| 26 |  | <pre>.   .   .   P v ( Q v R )</pre> | I v (R) | 25 |
| 27 |  | <pre>.   .   Q -> ( P v ( Q v R ) )</pre> | subproof I -> | 24 - 26 |
| 28 |  | <pre>.   .   P v ( Q v R )</pre> | E v | 23, 27, 20 |
| 29 |  | <pre>.   ( P v Q ) -> ( P v ( Q v R ) )</pre> | subproof I -> | 20 - 28 |
| 30 |  | <pre>.   .   P v R</pre> | Assume: I -> |  |
| 31 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 32 |  | <pre>.   .   .   P v ( Q v R )</pre> | I v (L) | 31 |
| 33 |  | <pre>.   .   P -> ( P v ( Q v R ) )</pre> | subproof I -> | 31 - 32 |
| 34 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 35 |  | <pre>.   .   .   Q v R</pre> | I v (R) | 34 |
| 36 |  | <pre>.   .   .   P v ( Q v R )</pre> | I v (R) | 35 |
| 37 |  | <pre>.   .   R -> ( P v ( Q v R ) )</pre> | subproof I -> | 34 - 36 |
| 38 |  | <pre>.   .   P v ( Q v R )</pre> | E v | 33, 37, 30 |
| 39 |  | <pre>.   ( P v R ) -> ( P v ( Q v R ) )</pre> | subproof I -> | 30 - 38 |
| 40 |  | <pre>.   P v ( Q v R )</pre> | E v | 29, 39, 19 |
| 41 |  | <pre>( ( P v Q ) v ( P v R ) ) -> ( P v ( Q v R ) )</pre> | subproof I -> | 19 - 40 |
| 42 |  | <pre>( P v ( Q v R ) ) <-> ( ( P v Q ) v ( P v R ) )</pre> | I <-> | 18, 41 |

## ( P v ( Q ^ R ) ) <-> ( ( P v Q ) ^ ( P v R ) ) <a name="proof-8-9"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v ( Q ^ R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 |  | <pre>.   .   P v R</pre> | I v (L) | 2 |
| 5 |  | <pre>.   .   ( P v Q ) ^ ( P v R )</pre> | I ^ | 3, 4 |
| 6 |  | <pre>.   P -> ( ( P v Q ) ^ ( P v R ) )</pre> | subproof I -> | 2 - 5 |
| 7 |  | <pre>.   .   Q ^ R</pre> | Assume: I -> |  |
| 8 |  | <pre>.   .   Q</pre> | E ^ (L) | 7 |
| 9 |  | <pre>.   .   R</pre> | E ^ (R) | 7 |
| 10 |  | <pre>.   .   P v Q</pre> | I v (R) | 8 |
| 11 |  | <pre>.   .   P v R</pre> | I v (R) | 9 |
| 12 |  | <pre>.   .   ( P v Q ) ^ ( P v R )</pre> | I ^ | 10, 11 |
| 13 |  | <pre>.   ( Q ^ R ) -> ( ( P v Q ) ^ ( P v R ) )</pre> | subproof I -> | 7 - 12 |
| 14 |  | <pre>.   ( P v Q ) ^ ( P v R )</pre> | E v | 6, 13, 1 |
| 15 |  | <pre>( P v ( Q ^ R ) ) -> ( ( P v Q ) ^ ( P v R ) )</pre> | subproof I -> | 1 - 14 |
| 16 |  | <pre>.   ( P v Q ) ^ ( P v R )</pre> | Assume: I -> |  |
| 17 |  | <pre>.   P v Q</pre> | E ^ (L) | 16 |
| 18 |  | <pre>.   P v R</pre> | E ^ (R) | 16 |
| 19 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 20 |  | <pre>.   .   P v ( Q ^ R )</pre> | I v (L) | 19 |
| 21 |  | <pre>.   P -> ( P v ( Q ^ R ) )</pre> | subproof I -> | 19 - 20 |
| 22 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 23 |  | <pre>.   .   .   ~ Q</pre> | Assume: I -> |  |
| 24 |  | <pre>.   .   .   P</pre> | Reiterate | 22 |
| 25 |  | <pre>.   .   ~ Q -> P</pre> | subproof I -> | 23 - 24 |
| 26 |  | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 25 |
| 27 |  | <pre>.   P -> ( ~ P -> Q )</pre> | subproof I -> | 22 - 26 |
| 28 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 29 |  | <pre>.   .   .   ~ R</pre> | Assume: I -> |  |
| 30 |  | <pre>.   .   .   P</pre> | Reiterate | 28 |
| 31 |  | <pre>.   .   ~ R -> P</pre> | subproof I -> | 29 - 30 |
| 32 |  | <pre>.   .   ~ P -> R</pre> | Theorem: contrapositive | 31 |
| 33 |  | <pre>.   P -> ( ~ P -> R )</pre> | subproof I -> | 28 - 32 |
| 34 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 35 |  | <pre>.   .   P v Q</pre> | Reiterate | 17 |
| 36 |  | <pre>.   .   P v R</pre> | Reiterate | 18 |
| 37 |  | <pre>.   .   .   Q</pre> | Assume: I -> |  |
| 38 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 34 |
| 39 |  | <pre>.   .   .   ~ P -> Q</pre> | I -> | 38, 37 |
| 40 |  | <pre>.   .   Q -> ( ~ P -> Q )</pre> | subproof I -> | 37 - 39 |
| 41 |  | <pre>.   .   P -> ( ~ P -> Q )</pre> | Reiterate | 27 |
| 42 |  | <pre>.   .   ~ P -> Q</pre> | E v | 41, 40, 35 |
| 43 |  | <pre>.   .   Q</pre> | E -> | 42, 34 |
| 44 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 45 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 34 |
| 46 |  | <pre>.   .   .   ~ P -> R</pre> | I -> | 45, 44 |
| 47 |  | <pre>.   .   R -> ( ~ P -> R )</pre> | subproof I -> | 44 - 46 |
| 48 |  | <pre>.   .   P -> ( ~ P -> R )</pre> | Reiterate | 33 |
| 49 |  | <pre>.   .   ~ P -> R</pre> | E v | 48, 47, 36 |
| 50 |  | <pre>.   .   R</pre> | E -> | 49, 34 |
| 51 |  | <pre>.   .   Q ^ R</pre> | I ^ | 43, 50 |
| 52 |  | <pre>.   .   P v ( Q ^ R )</pre> | I v (R) | 51 |
| 53 |  | <pre>.   ~ P -> ( P v ( Q ^ R ) )</pre> | subproof I -> | 34 - 52 |
| 54 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 55 |  | <pre>.   P v ( Q ^ R )</pre> | E v | 21, 53, 54 |
| 56 |  | <pre>( ( P v Q ) ^ ( P v R ) ) -> ( P v ( Q ^ R ) )</pre> | subproof I -> | 16 - 55 |
| 57 |  | <pre>( P v ( Q ^ R ) ) <-> ( ( P v Q ) ^ ( P v R ) )</pre> | I <-> | 15, 56 |

# anti-distributivity <a name="anti-distributivity"></a>

## ( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) ) <a name="proof-9-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P v Q ) -> R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 |  | <pre>.   .   ( P v Q ) -> R</pre> | Reiterate | 1 |
| 5 |  | <pre>.   .   R</pre> | E -> | 4, 3 |
| 6 |  | <pre>.   P -> R</pre> | subproof I -> | 2 - 5 |
| 7 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 8 |  | <pre>.   .   P v Q</pre> | I v (R) | 7 |
| 9 |  | <pre>.   .   ( P v Q ) -> R</pre> | Reiterate | 1 |
| 10 |  | <pre>.   .   R</pre> | E -> | 9, 8 |
| 11 |  | <pre>.   Q -> R</pre> | subproof I -> | 7 - 10 |
| 12 |  | <pre>.   ( P -> R ) ^ ( Q -> R )</pre> | I ^ | 6, 11 |
| 13 |  | <pre>( ( P v Q ) -> R ) -> ( ( P -> R ) ^ ( Q -> R ) )</pre> | subproof I -> | 1 - 12 |
| 14 |  | <pre>.   ( P -> R ) ^ ( Q -> R )</pre> | Assume: I -> |  |
| 15 |  | <pre>.   .   P v Q</pre> | Assume: I -> |  |
| 16 |  | <pre>.   .   ( P -> R ) ^ ( Q -> R )</pre> | Reiterate | 14 |
| 17 |  | <pre>.   .   P -> R</pre> | E ^ (L) | 16 |
| 18 |  | <pre>.   .   Q -> R</pre> | E ^ (R) | 16 |
| 19 |  | <pre>.   .   R</pre> | E v | 17, 18, 15 |
| 20 |  | <pre>.   ( P v Q ) -> R</pre> | subproof I -> | 15 - 19 |
| 21 |  | <pre>( ( P -> R ) ^ ( Q -> R ) ) -> ( ( P v Q ) -> R )</pre> | subproof I -> | 14 - 20 |
| 22 |  | <pre>( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) )</pre> | I <-> | 13, 21 |

## ( ( P ^ Q ) -> R ) <-> ( ( P -> R ) v ( Q -> R ) ) <a name="proof-9-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P -> R ) v ( Q -> R )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P -> R</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   .   P ^ Q</pre> | Assume: I -> |  |
| 4 |  | <pre>.   .   .   P -> R</pre> | Reiterate | 2 |
| 5 |  | <pre>.   .   .   P</pre> | E ^ (L) | 3 |
| 6 |  | <pre>.   .   .   R</pre> | E -> | 4, 5 |
| 7 |  | <pre>.   .   ( P ^ Q ) -> R</pre> | subproof I -> | 3 - 6 |
| 8 |  | <pre>.   ( P -> R ) -> ( ( P ^ Q ) -> R )</pre> | subproof I -> | 2 - 7 |
| 9 |  | <pre>.   .   Q -> R</pre> | Assume: I -> |  |
| 10 |  | <pre>.   .   .   P ^ Q</pre> | Assume: I -> |  |
| 11 |  | <pre>.   .   .   Q -> R</pre> | Reiterate | 9 |
| 12 |  | <pre>.   .   .   Q</pre> | E ^ (R) | 10 |
| 13 |  | <pre>.   .   .   R</pre> | E -> | 11, 12 |
| 14 |  | <pre>.   .   ( P ^ Q ) -> R</pre> | subproof I -> | 10 - 13 |
| 15 |  | <pre>.   ( Q -> R ) -> ( ( P ^ Q ) -> R )</pre> | subproof I -> | 9 - 14 |
| 16 |  | <pre>.   ( P ^ Q ) -> R</pre> | E v | 8, 15, 1 |
| 17 |  | <pre>( ( P -> R ) v ( Q -> R ) ) -> ( ( P ^ Q ) -> R )</pre> | subproof I -> | 1 - 16 |
| 18 |  | <pre>.   ( P ^ Q ) -> R</pre> | Assume: I -> |  |
| 19 |  | <pre>.   .   ~ ( ( P -> R ) v ( Q -> R ) )</pre> | Assume: I/E ~ |  |
| 20 |  | <pre>.   .   ~ ( P -> R ) ^ ~ ( Q -> R )</pre> | DeMorgan's theorem: (~ v) to ^ | 19 |
| 21 |  | <pre>.   .   ~ ( P -> R )</pre> | E ^ (L) | 20 |
| 22 |  | <pre>.   .   ~ ( Q -> R )</pre> | E ^ (R) | 20 |
| 23 |  | <pre>.   .   P ^ ~ R</pre> | Theorem: (~ ->) to ^ | 21 |
| 24 |  | <pre>.   .   P</pre> | E ^ (L) | 23 |
| 25 |  | <pre>.   .   ~ R</pre> | E ^ (R) | 23 |
| 26 |  | <pre>.   .   Q ^ ~ R</pre> | Theorem: (~ ->) to ^ | 22 |
| 27 |  | <pre>.   .   Q</pre> | E ^ (L) | 26 |
| 28 |  | <pre>.   .   ( P ^ Q ) -> R</pre> | Reiterate | 18 |
| 29 |  | <pre>.   .   P ^ Q</pre> | I ^ | 24, 27 |
| 30 |  | <pre>.   .   R</pre> | E -> | 28, 29 |
| 31 |  | <pre>.   ( P -> R ) v ( Q -> R )</pre> | subproof I/E ~ | 19 - 30 |
| 32 |  | <pre>( ( P ^ Q ) -> R ) -> ( ( P -> R ) v ( Q -> R ) )</pre> | subproof I -> | 18 - 31 |
| 33 |  | <pre>( ( P ^ Q ) -> R ) <-> ( ( P -> R ) v ( Q -> R ) )</pre> | I <-> | 32, 17 |

# double-distribution <a name="double-distribution"></a>

## ( ( P v Q ) ^ ( R v S ) ) <-> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) ) <a name="proof-10-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P v Q ) ^ ( R v S )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P v Q</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   R v S</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 5 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 6 |  | <pre>.   .   .   P</pre> | Reiterate | 4 |
| 7 |  | <pre>.   .   .   P ^ R</pre> | I ^ | 6, 5 |
| 8 |  | <pre>.   .   .   ( P ^ R ) v ( P ^ S )</pre> | I v (L) | 7 |
| 9 |  | <pre>.   .   .   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | I v (L) | 8 |
| 10 |  | <pre>.   .   R -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 5 - 9 |
| 11 |  | <pre>.   .   .   S</pre> | Assume: I -> |  |
| 12 |  | <pre>.   .   .   P</pre> | Reiterate | 4 |
| 13 |  | <pre>.   .   .   P ^ S</pre> | I ^ | 12, 11 |
| 14 |  | <pre>.   .   .   ( P ^ R ) v ( P ^ S )</pre> | I v (R) | 13 |
| 15 |  | <pre>.   .   .   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | I v (L) | 14 |
| 16 |  | <pre>.   .   S -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 11 - 15 |
| 17 |  | <pre>.   .   R v S</pre> | Reiterate | 3 |
| 18 |  | <pre>.   .   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | E v | 10, 16, 17 |
| 19 |  | <pre>.   P -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 4 - 18 |
| 20 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 21 |  | <pre>.   .   .   R</pre> | Assume: I -> |  |
| 22 |  | <pre>.   .   .   Q</pre> | Reiterate | 20 |
| 23 |  | <pre>.   .   .   Q ^ R</pre> | I ^ | 22, 21 |
| 24 |  | <pre>.   .   .   ( Q ^ R ) v ( Q ^ S )</pre> | I v (L) | 23 |
| 25 |  | <pre>.   .   .   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | I v (R) | 24 |
| 26 |  | <pre>.   .   R -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 21 - 25 |
| 27 |  | <pre>.   .   .   S</pre> | Assume: I -> |  |
| 28 |  | <pre>.   .   .   Q</pre> | Reiterate | 20 |
| 29 |  | <pre>.   .   .   Q ^ S</pre> | I ^ | 28, 27 |
| 30 |  | <pre>.   .   .   ( Q ^ R ) v ( Q ^ S )</pre> | I v (R) | 29 |
| 31 |  | <pre>.   .   .   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | I v (R) | 30 |
| 32 |  | <pre>.   .   S -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 27 - 31 |
| 33 |  | <pre>.   .   R v S</pre> | Reiterate | 3 |
| 34 |  | <pre>.   .   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | E v | 26, 32, 33 |
| 35 |  | <pre>.   Q -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 20 - 34 |
| 36 |  | <pre>.   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | E v | 19, 35, 2 |
| 37 |  | <pre>( ( P v Q ) ^ ( R v S ) ) -> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | subproof I -> | 1 - 36 |
| 38 |  | <pre>.   ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) )</pre> | Assume: I -> |  |
| 39 |  | <pre>.   .   ( P ^ R ) v ( P ^ S )</pre> | Assume: I -> |  |
| 40 |  | <pre>.   .   .   P ^ R</pre> | Assume: I -> |  |
| 41 |  | <pre>.   .   .   P</pre> | E ^ (L) | 40 |
| 42 |  | <pre>.   .   .   R</pre> | E ^ (R) | 40 |
| 43 |  | <pre>.   .   .   P v Q</pre> | I v (L) | 41 |
| 44 |  | <pre>.   .   .   R v S</pre> | I v (L) | 42 |
| 45 |  | <pre>.   .   .   ( P v Q ) ^ ( R v S )</pre> | I ^ | 43, 44 |
| 46 |  | <pre>.   .   ( P ^ R ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 40 - 45 |
| 47 |  | <pre>.   .   .   P ^ S</pre> | Assume: I -> |  |
| 48 |  | <pre>.   .   .   P</pre> | E ^ (L) | 47 |
| 49 |  | <pre>.   .   .   S</pre> | E ^ (R) | 47 |
| 50 |  | <pre>.   .   .   P v Q</pre> | I v (L) | 48 |
| 51 |  | <pre>.   .   .   R v S</pre> | I v (R) | 49 |
| 52 |  | <pre>.   .   .   ( P v Q ) ^ ( R v S )</pre> | I ^ | 50, 51 |
| 53 |  | <pre>.   .   ( P ^ S ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 47 - 52 |
| 54 |  | <pre>.   .   ( P v Q ) ^ ( R v S )</pre> | E v | 46, 53, 39 |
| 55 |  | <pre>.   ( ( P ^ R ) v ( P ^ S ) ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 39 - 54 |
| 56 |  | <pre>.   .   ( Q ^ R ) v ( Q ^ S )</pre> | Assume: I -> |  |
| 57 |  | <pre>.   .   .   Q ^ R</pre> | Assume: I -> |  |
| 58 |  | <pre>.   .   .   Q</pre> | E ^ (L) | 57 |
| 59 |  | <pre>.   .   .   R</pre> | E ^ (R) | 57 |
| 60 |  | <pre>.   .   .   P v Q</pre> | I v (R) | 58 |
| 61 |  | <pre>.   .   .   R v S</pre> | I v (L) | 59 |
| 62 |  | <pre>.   .   .   ( P v Q ) ^ ( R v S )</pre> | I ^ | 60, 61 |
| 63 |  | <pre>.   .   ( Q ^ R ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 57 - 62 |
| 64 |  | <pre>.   .   .   Q ^ S</pre> | Assume: I -> |  |
| 65 |  | <pre>.   .   .   Q</pre> | E ^ (L) | 64 |
| 66 |  | <pre>.   .   .   S</pre> | E ^ (R) | 64 |
| 67 |  | <pre>.   .   .   P v Q</pre> | I v (R) | 65 |
| 68 |  | <pre>.   .   .   R v S</pre> | I v (R) | 66 |
| 69 |  | <pre>.   .   .   ( P v Q ) ^ ( R v S )</pre> | I ^ | 67, 68 |
| 70 |  | <pre>.   .   ( Q ^ S ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 64 - 69 |
| 71 |  | <pre>.   .   ( P v Q ) ^ ( R v S )</pre> | E v | 63, 70, 56 |
| 72 |  | <pre>.   ( ( Q ^ R ) v ( Q ^ S ) ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 56 - 71 |
| 73 |  | <pre>.   ( P v Q ) ^ ( R v S )</pre> | E v | 55, 72, 38 |
| 74 |  | <pre>( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) ) -> ( ( P v Q ) ^ ( R v S ) )</pre> | subproof I -> | 38 - 73 |
| 75 |  | <pre>( ( P v Q ) ^ ( R v S ) ) <-> ( ( ( P ^ R ) v ( P ^ S ) ) v ( ( Q ^ R ) v ( Q ^ S ) ) )</pre> | I <-> | 37, 74 |

## ( ( P ^ Q ) v ( R ^ S ) ) <-> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) ) <a name="proof-10-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P ^ Q ) v ( R ^ S )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P ^ Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 4 |  | <pre>.   .   Q</pre> | E ^ (R) | 2 |
| 5 |  | <pre>.   .   P v R</pre> | I v (L) | 3 |
| 6 |  | <pre>.   .   P v S</pre> | I v (L) | 3 |
| 7 |  | <pre>.   .   Q v R</pre> | I v (L) | 4 |
| 8 |  | <pre>.   .   Q v S</pre> | I v (L) | 4 |
| 9 |  | <pre>.   .   ( P v R ) ^ ( P v S )</pre> | I ^ | 5, 6 |
| 10 |  | <pre>.   .   ( Q v R ) ^ ( Q v S )</pre> | I ^ | 7, 8 |
| 11 |  | <pre>.   .   ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) )</pre> | I ^ | 9, 10 |
| 12 |  | <pre>.   ( P ^ Q ) -> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )</pre> | subproof I -> | 2 - 11 |
| 13 |  | <pre>.   .   R ^ S</pre> | Assume: I -> |  |
| 14 |  | <pre>.   .   R</pre> | E ^ (L) | 13 |
| 15 |  | <pre>.   .   S</pre> | E ^ (R) | 13 |
| 16 |  | <pre>.   .   P v R</pre> | I v (R) | 14 |
| 17 |  | <pre>.   .   P v S</pre> | I v (R) | 15 |
| 18 |  | <pre>.   .   Q v R</pre> | I v (R) | 14 |
| 19 |  | <pre>.   .   Q v S</pre> | I v (R) | 15 |
| 20 |  | <pre>.   .   ( P v R ) ^ ( P v S )</pre> | I ^ | 16, 17 |
| 21 |  | <pre>.   .   ( Q v R ) ^ ( Q v S )</pre> | I ^ | 18, 19 |
| 22 |  | <pre>.   .   ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) )</pre> | I ^ | 20, 21 |
| 23 |  | <pre>.   ( R ^ S ) -> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )</pre> | subproof I -> | 13 - 22 |
| 24 |  | <pre>.   ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) )</pre> | E v | 12, 23, 1 |
| 25 |  | <pre>( ( P ^ Q ) v ( R ^ S ) ) -> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )</pre> | subproof I -> | 1 - 24 |
| 26 |  | <pre>.   ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) )</pre> | Assume: I -> |  |
| 27 |  | <pre>.   ( P v R ) ^ ( P v S )</pre> | E ^ (L) | 26 |
| 28 |  | <pre>.   P v R</pre> | E ^ (L) | 27 |
| 29 |  | <pre>.   P v S</pre> | E ^ (R) | 27 |
| 30 |  | <pre>.   ( Q v R ) ^ ( Q v S )</pre> | E ^ (R) | 26 |
| 31 |  | <pre>.   Q v R</pre> | E ^ (L) | 30 |
| 32 |  | <pre>.   Q v S</pre> | E ^ (R) | 30 |
| 33 |  | <pre>.   .   P ^ Q</pre> | Assume: I -> |  |
| 34 |  | <pre>.   .   ( P ^ Q ) v ( R ^ S )</pre> | I v (L) | 33 |
| 35 |  | <pre>.   ( P ^ Q ) -> ( ( P ^ Q ) v ( R ^ S ) )</pre> | subproof I -> | 33 - 34 |
| 36 |  | <pre>.   .   ~ ( P ^ Q )</pre> | Assume: I -> |  |
| 37 |  | <pre>.   .   ~ P v ~ Q</pre> | DeMorgan's theorem: (~ ^) to v | 36 |
| 38 |  | <pre>.   .   .   ~ P</pre> | Assume: I -> |  |
| 39 |  | <pre>.   .   .   P v R</pre> | Reiterate | 28 |
| 40 |  | <pre>.   .   .   ~ P -> R</pre> | Theorem: v to -> | 39 |
| 41 |  | <pre>.   .   .   P v S</pre> | Reiterate | 29 |
| 42 |  | <pre>.   .   .   ~ P -> S</pre> | Theorem: v to -> | 41 |
| 43 |  | <pre>.   .   .   R</pre> | E -> | 40, 38 |
| 44 |  | <pre>.   .   .   S</pre> | E -> | 42, 38 |
| 45 |  | <pre>.   .   .   R ^ S</pre> | I ^ | 43, 44 |
| 46 |  | <pre>.   .   .   ( P ^ Q ) v ( R ^ S )</pre> | I v (R) | 45 |
| 47 |  | <pre>.   .   ~ P -> ( ( P ^ Q ) v ( R ^ S ) )</pre> | subproof I -> | 38 - 46 |
| 48 |  | <pre>.   .   .   ~ Q</pre> | Assume: I -> |  |
| 49 |  | <pre>.   .   .   Q v R</pre> | Reiterate | 31 |
| 50 |  | <pre>.   .   .   ~ Q -> R</pre> | Theorem: v to -> | 49 |
| 51 |  | <pre>.   .   .   Q v S</pre> | Reiterate | 32 |
| 52 |  | <pre>.   .   .   ~ Q -> S</pre> | Theorem: v to -> | 51 |
| 53 |  | <pre>.   .   .   R</pre> | E -> | 50, 48 |
| 54 |  | <pre>.   .   .   S</pre> | E -> | 52, 48 |
| 55 |  | <pre>.   .   .   R ^ S</pre> | I ^ | 53, 54 |
| 56 |  | <pre>.   .   .   ( P ^ Q ) v ( R ^ S )</pre> | I v (R) | 55 |
| 57 |  | <pre>.   .   ~ Q -> ( ( P ^ Q ) v ( R ^ S ) )</pre> | subproof I -> | 48 - 56 |
| 58 |  | <pre>.   .   ( P ^ Q ) v ( R ^ S )</pre> | E v | 47, 57, 37 |
| 59 |  | <pre>.   ~ ( P ^ Q ) -> ( ( P ^ Q ) v ( R ^ S ) )</pre> | subproof I -> | 36 - 58 |
| 60 |  | <pre>.   ( P ^ Q ) v ~ ( P ^ Q )</pre> | Theorem: excluded middle |  |
| 61 |  | <pre>.   ( P ^ Q ) v ( R ^ S )</pre> | E v | 35, 59, 60 |
| 62 |  | <pre>( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) ) -> ( ( P ^ Q ) v ( R ^ S ) )</pre> | subproof I -> | 26 - 61 |
| 63 |  | <pre>( ( P ^ Q ) v ( R ^ S ) ) <-> ( ( ( P v R ) ^ ( P v S ) ) ^ ( ( Q v R ) ^ ( Q v S ) ) )</pre> | I <-> | 25, 62 |

# disjunction <a name="disjunction"></a>

## ( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S ) <a name="proof-11-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   ( P -> R ) ^ ( Q -> S )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   P v Q</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P -> R</pre> | E ^ (L) | 2 |
| 5 |  | <pre>.   Q -> S</pre> | E ^ (R) | 2 |
| 6 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   P -> R</pre> | Reiterate | 4 |
| 8 |  | <pre>.   .   R</pre> | E -> | 7, 6 |
| 9 |  | <pre>.   .   R v S</pre> | I v (L) | 8 |
| 10 |  | <pre>.   P -> ( R v S )</pre> | subproof I -> | 6 - 9 |
| 11 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 12 |  | <pre>.   .   Q -> S</pre> | Reiterate | 5 |
| 13 |  | <pre>.   .   S</pre> | E -> | 12, 11 |
| 14 |  | <pre>.   .   R v S</pre> | I v (R) | 13 |
| 15 |  | <pre>.   Q -> ( R v S )</pre> | subproof I -> | 11 - 14 |
| 16 |  | <pre>.   R v S</pre> | E v | 10, 15, 3 |
| 17 |  | <pre>( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )</pre> | subproof I -> | 1 - 16 |

## ( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S ) <a name="proof-11-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   ( P -> R ) ^ ( Q -> S )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ R v Q</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P -> R</pre> | E ^ (L) | 2 |
| 5 |  | <pre>.   Q -> S</pre> | E ^ (R) | 2 |
| 6 |  | <pre>.   .   ~ R</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   P -> R</pre> | Reiterate | 4 |
| 8 |  | <pre>.   .   ~ R -> ~ P</pre> | Theorem: contrapositive | 7 |
| 9 |  | <pre>.   .   ~ P</pre> | E -> | 8, 6 |
| 10 |  | <pre>.   .   ~ P v S</pre> | I v (L) | 9 |
| 11 |  | <pre>.   ~ R -> ( ~ P v S )</pre> | subproof I -> | 6 - 10 |
| 12 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   Q -> S</pre> | Reiterate | 5 |
| 14 |  | <pre>.   .   S</pre> | E -> | 13, 12 |
| 15 |  | <pre>.   .   ~ P v S</pre> | I v (R) | 14 |
| 16 |  | <pre>.   Q -> ( ~ P v S )</pre> | subproof I -> | 12 - 15 |
| 17 |  | <pre>.   ~ P v S</pre> | E v | 11, 16, 3 |
| 18 |  | <pre>( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S )</pre> | subproof I -> | 1 - 17 |

## ( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q ) <a name="proof-11-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   ( P -> R ) ^ ( Q -> S )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ R v ~ S</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P -> R</pre> | E ^ (L) | 2 |
| 5 |  | <pre>.   Q -> S</pre> | E ^ (R) | 2 |
| 6 |  | <pre>.   .   ~ R</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   P -> R</pre> | Reiterate | 4 |
| 8 |  | <pre>.   .   ~ R -> ~ P</pre> | Theorem: contrapositive | 7 |
| 9 |  | <pre>.   .   ~ P</pre> | E -> | 8, 6 |
| 10 |  | <pre>.   .   ~ P v ~ Q</pre> | I v (L) | 9 |
| 11 |  | <pre>.   ~ R -> ( ~ P v ~ Q )</pre> | subproof I -> | 6 - 10 |
| 12 |  | <pre>.   .   ~ S</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   Q -> S</pre> | Reiterate | 5 |
| 14 |  | <pre>.   .   ~ S -> ~ Q</pre> | Theorem: contrapositive | 13 |
| 15 |  | <pre>.   .   ~ Q</pre> | E -> | 14, 12 |
| 16 |  | <pre>.   .   ~ P v ~ Q</pre> | I v (R) | 15 |
| 17 |  | <pre>.   ~ S -> ( ~ P v ~ Q )</pre> | subproof I -> | 12 - 16 |
| 18 |  | <pre>.   ~ P v ~ Q</pre> | E v | 11, 17, 3 |
| 19 |  | <pre>( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q )</pre> | subproof I -> | 1 - 18 |

## ( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R <a name="proof-11-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q )</pre> | Assume: I -> |  |
| 2 |  | <pre>.   ( P -> R ) v ( Q -> R )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   P ^ Q</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P</pre> | E ^ (L) | 3 |
| 5 |  | <pre>.   Q</pre> | E ^ (R) | 3 |
| 6 |  | <pre>.   .   P -> R</pre> | Assume: I -> |  |
| 7 |  | <pre>.   .   P</pre> | Reiterate | 4 |
| 8 |  | <pre>.   .   R</pre> | E -> | 6, 7 |
| 9 |  | <pre>.   ( P -> R ) -> R</pre> | subproof I -> | 6 - 8 |
| 10 |  | <pre>.   .   Q -> R</pre> | Assume: I -> |  |
| 11 |  | <pre>.   .   Q</pre> | Reiterate | 5 |
| 12 |  | <pre>.   .   R</pre> | E -> | 10, 11 |
| 13 |  | <pre>.   ( Q -> R ) -> R</pre> | subproof I -> | 10 - 12 |
| 14 |  | <pre>.   R</pre> | E v | 9, 13, 2 |
| 15 |  | <pre>( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R</pre> | subproof I -> | 1 - 14 |

## ( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q ) <a name="proof-11-5"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( ( P -> R ) v ( Q -> R ) ) ^ ~ R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   ( P -> R ) v ( Q -> R )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ R</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   .   P -> R</pre> | Assume: I -> |  |
| 5 |  | <pre>.   .   ~ R</pre> | Reiterate | 3 |
| 6 |  | <pre>.   .   ~ R -> ~ P</pre> | Theorem: contrapositive | 4 |
| 7 |  | <pre>.   .   ~ P</pre> | E -> | 6, 5 |
| 8 |  | <pre>.   .   ~ P v ~ Q</pre> | I v (L) | 7 |
| 9 |  | <pre>.   ( P -> R ) -> ( ~ P v ~ Q )</pre> | subproof I -> | 4 - 8 |
| 10 |  | <pre>.   .   Q -> R</pre> | Assume: I -> |  |
| 11 |  | <pre>.   .   ~ R</pre> | Reiterate | 3 |
| 12 |  | <pre>.   .   ~ R -> ~ Q</pre> | Theorem: contrapositive | 10 |
| 13 |  | <pre>.   .   ~ Q</pre> | E -> | 12, 11 |
| 14 |  | <pre>.   .   ~ P v ~ Q</pre> | I v (R) | 13 |
| 15 |  | <pre>.   ( Q -> R ) -> ( ~ P v ~ Q )</pre> | subproof I -> | 10 - 14 |
| 16 |  | <pre>.   ~ P v ~ Q</pre> | E v | 9, 15, 2 |
| 17 |  | <pre>( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q )</pre> | subproof I -> | 1 - 16 |

# biconditional <a name="biconditional"></a>

## ( ( P <-> Q ) ^ P ) -> Q <a name="proof-12-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P <-> Q ) ^ P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P <-> Q</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   P</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   P -> Q</pre> | E <-> (L) | 2 |
| 5 |  | <pre>.   Q</pre> | E -> | 4, 3 |
| 6 |  | <pre>( ( P <-> Q ) ^ P ) -> Q</pre> | subproof I -> | 1 - 5 |

## ( ( P <-> Q ) ^ ~ P ) -> ~ Q <a name="proof-12-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P <-> Q ) ^ ~ P</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P <-> Q</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ P</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   Q -> P</pre> | E <-> (R) | 2 |
| 5 |  | <pre>.   .   Q</pre> | Assume: I/E ~ |  |
| 6 |  | <pre>.   .   ~ P</pre> | Reiterate | 3 |
| 7 |  | <pre>.   .   Q -> P</pre> | Reiterate | 4 |
| 8 |  | <pre>.   .   P</pre> | E -> | 7, 5 |
| 9 |  | <pre>.   ~ Q</pre> | subproof I/E ~ | 5 - 8 |
| 10 |  | <pre>( ( P <-> Q ) ^ ~ P ) -> ~ Q</pre> | subproof I -> | 1 - 9 |

## ( P <-> ~ Q ) -> ( ~ P <-> Q ) <a name="proof-12-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P <-> ~ Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 |  | <pre>.   Q -> ~ P</pre> | Theorem: contrapositive | 2 |
| 4 |  | <pre>.   ~ Q -> P</pre> | E <-> (R) | 1 |
| 5 |  | <pre>.   ~ P -> Q</pre> | Theorem: contrapositive | 4 |
| 6 |  | <pre>.   ~ P <-> Q</pre> | I <-> | 5, 3 |
| 7 |  | <pre>( P <-> ~ Q ) -> ( ~ P <-> Q )</pre> | subproof I -> | 1 - 6 |

## ( P <-> Q ) <-> ( ~ P <-> ~ Q ) <a name="proof-12-4"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P <-> Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P -> Q</pre> | E <-> (L) | 1 |
| 3 |  | <pre>.   ~ Q -> ~ P</pre> | Theorem: contrapositive | 2 |
| 4 |  | <pre>.   Q -> P</pre> | E <-> (R) | 1 |
| 5 |  | <pre>.   ~ P -> ~ Q</pre> | Theorem: contrapositive | 4 |
| 6 |  | <pre>.   ~ P <-> ~ Q</pre> | I <-> | 5, 3 |
| 7 |  | <pre>( P <-> Q ) -> ( ~ P <-> ~ Q )</pre> | subproof I -> | 1 - 6 |
| 8 |  | <pre>.   ~ P <-> ~ Q</pre> | Assume: I -> |  |
| 9 |  | <pre>.   ~ P -> ~ Q</pre> | E <-> (L) | 8 |
| 10 |  | <pre>.   Q -> P</pre> | Theorem: contrapositive | 9 |
| 11 |  | <pre>.   ~ Q -> ~ P</pre> | E <-> (R) | 8 |
| 12 |  | <pre>.   P -> Q</pre> | Theorem: contrapositive | 11 |
| 13 |  | <pre>.   P <-> Q</pre> | I <-> | 12, 10 |
| 14 |  | <pre>( ~ P <-> ~ Q ) -> ( P <-> Q )</pre> | subproof I -> | 8 - 13 |
| 15 |  | <pre>( P <-> Q ) <-> ( ~ P <-> ~ Q )</pre> | I <-> | 7, 14 |

## ( P <-> Q ) v ~ ( P <-> Q ) <a name="proof-12-5"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | Assume: I/E ~ |  |
| 2 |  | <pre>.   .   P <-> Q</pre> | Assume: I/E ~ |  |
| 3 |  | <pre>.   .   ( P <-> Q ) v ~ ( P <-> Q )</pre> | I v (L) | 2 |
| 4 |  | <pre>.   .   ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | Reiterate | 1 |
| 5 |  | <pre>.   ~ ( P <-> Q )</pre> | subproof I/E ~ | 2 - 4 |
| 6 |  | <pre>.   ( P <-> Q ) v ~ ( P <-> Q )</pre> | I v (R) | 5 |
| 7 |  | <pre>( P <-> Q ) v ~ ( P <-> Q )</pre> | subproof I/E ~ | 1 - 6 |

## ( P <-> ~ Q ) <-> ( ( P v Q ) ^ ( ~ P v ~ Q ) ) <a name="proof-12-6"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P <-> ~ Q</pre> | Assume: I -> |  |
| 2 |  | <pre>.   P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 |  | <pre>.   ~ Q -> P</pre> | E <-> (R) | 1 |
| 4 |  | <pre>.   ~ P -> Q</pre> | Theorem: contrapositive | 3 |
| 5 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 6 |  | <pre>.   .   P v Q</pre> | I v (L) | 5 |
| 7 |  | <pre>.   .   P -> ~ Q</pre> | Reiterate | 2 |
| 8 |  | <pre>.   .   ~ Q</pre> | E -> | 7, 5 |
| 9 |  | <pre>.   .   ~ P v ~ Q</pre> | I v (R) | 8 |
| 10 |  | <pre>.   .   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | I ^ | 6, 9 |
| 11 |  | <pre>.   P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | subproof I -> | 5 - 10 |
| 12 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 13 |  | <pre>.   .   ~ P v ~ Q</pre> | I v (L) | 12 |
| 14 |  | <pre>.   .   ~ P -> Q</pre> | Reiterate | 4 |
| 15 |  | <pre>.   .   Q</pre> | E -> | 14, 12 |
| 16 |  | <pre>.   .   P v Q</pre> | I v (R) | 15 |
| 17 |  | <pre>.   .   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | I ^ | 16, 13 |
| 18 |  | <pre>.   ~ P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | subproof I -> | 12 - 17 |
| 19 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 20 |  | <pre>.   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | E v | 11, 18, 19 |
| 21 |  | <pre>( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | subproof I -> | 1 - 20 |
| 22 |  | <pre>.   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | Assume: I -> |  |
| 23 |  | <pre>.   P v Q</pre> | E ^ (L) | 22 |
| 24 |  | <pre>.   ~ P v ~ Q</pre> | E ^ (R) | 22 |
| 25 |  | <pre>.   .   ~ P</pre> | Assume: I -> |  |
| 26 |  | <pre>.   .   .   P ^ Q</pre> | Assume: I/E ~ |  |
| 27 |  | <pre>.   .   .   P</pre> | E ^ (L) | 26 |
| 28 |  | <pre>.   .   .   ~ P</pre> | Reiterate | 25 |
| 29 |  | <pre>.   .   ~ ( P ^ Q )</pre> | subproof I/E ~ | 26 - 28 |
| 30 |  | <pre>.   ~ P -> ~ ( P ^ Q )</pre> | subproof I -> | 25 - 29 |
| 31 |  | <pre>.   .   ~ Q</pre> | Assume: I -> |  |
| 32 |  | <pre>.   .   .   P ^ Q</pre> | Assume: I/E ~ |  |
| 33 |  | <pre>.   .   .   Q</pre> | E ^ (R) | 32 |
| 34 |  | <pre>.   .   .   ~ Q</pre> | Reiterate | 31 |
| 35 |  | <pre>.   .   ~ ( P ^ Q )</pre> | subproof I/E ~ | 32 - 34 |
| 36 |  | <pre>.   ~ Q -> ~ ( P ^ Q )</pre> | subproof I -> | 31 - 35 |
| 37 |  | <pre>.   ~ ( P ^ Q )</pre> | E v | 30, 36, 24 |
| 38 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 39 |  | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume: I/E ~ |  |
| 40 |  | <pre>.   .   .   ~ P</pre> | E ^ (L) | 39 |
| 41 |  | <pre>.   .   .   P</pre> | Reiterate | 38 |
| 42 |  | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof I/E ~ | 39 - 41 |
| 43 |  | <pre>.   P -> ~ ( ~ P ^ ~ Q )</pre> | subproof I -> | 38 - 42 |
| 44 |  | <pre>.   .   Q</pre> | Assume: I -> |  |
| 45 |  | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume: I/E ~ |  |
| 46 |  | <pre>.   .   .   ~ Q</pre> | E ^ (R) | 45 |
| 47 |  | <pre>.   .   .   Q</pre> | Reiterate | 44 |
| 48 |  | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof I/E ~ | 45 - 47 |
| 49 |  | <pre>.   Q -> ~ ( ~ P ^ ~ Q )</pre> | subproof I -> | 44 - 48 |
| 50 |  | <pre>.   ~ ( ~ P ^ ~ Q )</pre> | E v | 43, 49, 23 |
| 51 |  | <pre>.   .   P</pre> | Assume: I -> |  |
| 52 |  | <pre>.   .   .   Q</pre> | Assume: I/E ~ |  |
| 53 |  | <pre>.   .   .   P</pre> | Reiterate | 51 |
| 54 |  | <pre>.   .   .   P ^ Q</pre> | I ^ | 53, 52 |
| 55 |  | <pre>.   .   .   ~ ( P ^ Q )</pre> | Reiterate | 37 |
| 56 |  | <pre>.   .   ~ Q</pre> | subproof I/E ~ | 52 - 55 |
| 57 |  | <pre>.   P -> ~ Q</pre> | subproof I -> | 51 - 56 |
| 58 |  | <pre>.   .   ~ Q</pre> | Assume: I -> |  |
| 59 |  | <pre>.   .   .   ~ P</pre> | Assume: I/E ~ |  |
| 60 |  | <pre>.   .   .   ~ Q</pre> | Reiterate | 58 |
| 61 |  | <pre>.   .   .   ~ P ^ ~ Q</pre> | I ^ | 59, 60 |
| 62 |  | <pre>.   .   .   ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 50 |
| 63 |  | <pre>.   .   P</pre> | subproof I/E ~ | 59 - 62 |
| 64 |  | <pre>.   ~ Q -> P</pre> | subproof I -> | 58 - 63 |
| 65 |  | <pre>.   P <-> ~ Q</pre> | I <-> | 57, 64 |
| 66 |  | <pre>( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )</pre> | subproof I -> | 22 - 65 |
| 67 |  | <pre>( P <-> ~ Q ) <-> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | I <-> | 21, 66 |

# miscellaneous <a name="miscellaneous"></a>

## ( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) ) <a name="proof-13-1"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ( P -> Q ) -> R</pre> | Assume: I -> |  |
| 2 |  | <pre>.   .   P -> Q</pre> | Assume: I -> |  |
| 3 |  | <pre>.   .   ( P -> Q ) -> R</pre> | Reiterate | 1 |
| 4 |  | <pre>.   .   R</pre> | E -> | 3, 2 |
| 5 |  | <pre>.   .   .   P</pre> | Assume: I -> |  |
| 6 |  | <pre>.   .   .   R</pre> | Reiterate | 4 |
| 7 |  | <pre>.   .   P -> R</pre> | subproof I -> | 5 - 6 |
| 8 |  | <pre>.   ( P -> Q ) -> ( P -> R )</pre> | subproof I -> | 2 - 7 |
| 9 |  | <pre>( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )</pre> | subproof I -> | 1 - 8 |

## ( P -> Q ) v ( Q -> R ) <a name="proof-13-2"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( ( P -> Q ) v ( Q -> R ) )</pre> | Assume: I/E ~ |  |
| 2 |  | <pre>.   ~ ( P -> Q ) ^ ~ ( Q -> R )</pre> | DeMorgan's theorem: (~ v) to ^ | 1 |
| 3 |  | <pre>.   ~ ( P -> Q )</pre> | E ^ (L) | 2 |
| 4 |  | <pre>.   ~ ( Q -> R )</pre> | E ^ (R) | 2 |
| 5 |  | <pre>.   P ^ ~ Q</pre> | Theorem: (~ ->) to ^ | 3 |
| 6 |  | <pre>.   ~ Q</pre> | E ^ (R) | 5 |
| 7 |  | <pre>.   Q ^ ~ R</pre> | Theorem: (~ ->) to ^ | 4 |
| 8 |  | <pre>.   Q</pre> | E ^ (L) | 7 |
| 9 |  | <pre>( P -> Q ) v ( Q -> R )</pre> | subproof I/E ~ | 1 - 8 |

## ( P -> Q ) v ( Q -> P ) <a name="proof-13-3"></a>

| Line | Term vars | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ( ( P -> Q ) v ( Q -> P ) )</pre> | Assume: I/E ~ |  |
| 2 |  | <pre>.   ~ ( P -> Q ) ^ ~ ( Q -> P )</pre> | DeMorgan's theorem: (~ v) to ^ | 1 |
| 3 |  | <pre>.   ~ ( P -> Q )</pre> | E ^ (L) | 2 |
| 4 |  | <pre>.   ~ ( Q -> P )</pre> | E ^ (R) | 2 |
| 5 |  | <pre>.   P ^ ~ Q</pre> | Theorem: (~ ->) to ^ | 3 |
| 6 |  | <pre>.   ~ Q</pre> | E ^ (R) | 5 |
| 7 |  | <pre>.   Q ^ ~ P</pre> | Theorem: (~ ->) to ^ | 4 |
| 8 |  | <pre>.   Q</pre> | E ^ (L) | 7 |
| 9 |  | <pre>( P -> Q ) v ( Q -> P )</pre> | subproof I/E ~ | 1 - 8 |

