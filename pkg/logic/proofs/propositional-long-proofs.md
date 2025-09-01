
# Table of Contents

1. [theorems](#theorems)
    1. [P v ~ P](#proof-1-1)
    2. [( P -> Q ) <-> ( ~ Q -> ~ P )](#proof-1-2)
    3. [~ ( P -> Q ) <-> ( P ^ ~ Q )](#proof-1-3)
    4. [( P v Q ) <-> ( ~ P -> Q )](#proof-1-4)
    5. [~ ( P v Q ) <-> ( ~ P ^ ~ Q )](#proof-1-5)
    6. [~ ( P ^ Q ) <-> ( ~ P v ~ Q )](#proof-1-6)
    7. [( P <-> ~ Q ) <-> ~ ( P <-> Q )](#proof-1-7)
2. [basics](#basics)
    1. [P -> P](#proof-2-1)
    2. [~ ( P ^ ~ P )](#proof-2-2)
    3. [( P ^ P ) <-> P](#proof-2-3)
    4. [( P v P ) <-> P](#proof-2-4)
    5. [( P ^ Q ) -> ( P v Q )](#proof-2-5)
    6. [( P -> ~ P ) -> ~ P](#proof-2-6)
    7. [( ( P -> Q ) -> P ) -> P](#proof-2-7)
3. [arrows](#arrows)
    1. [Q -> ( P -> Q )](#proof-3-1)
    2. [~ P -> ( P -> Q )](#proof-3-2)
    3. [( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )](#proof-3-3)
    4. [( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) )](#proof-3-4)
    5. [( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P](#proof-3-5)
    6. [( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q](#proof-3-6)
    7. [( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )](#proof-3-7)
    8. [( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) )](#proof-3-8)
    9. [( ( P -> R ) v ( Q -> R ) ) <-> ( ( P ^ Q ) -> R )](#proof-3-9)
4. [commutativity](#commutativity)
    1. [( P ^ Q ) -> ( Q ^ P )](#proof-4-1)
    2. [( P v Q ) -> ( Q v P )](#proof-4-2)
    3. [( P <-> Q ) -> ( Q <-> P )](#proof-4-3)
5. [associativity](#associativity)
    1. [( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) )](#proof-5-1)
    2. [( ( P v Q ) v R ) <-> ( P v ( Q v R ) )](#proof-5-2)
    3. [( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) )](#proof-5-3)
6. [distributivity](#distributivity)
    1. [( P -> ( Q -> R ) ) -> ( ( P -> Q ) -> ( P -> R ) )](#proof-6-1)
    2. [( ( P -> Q ) -> ( P -> R ) ) -> ( P -> ( Q -> R ) )](#proof-6-2)
    3. [( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )](#proof-6-3)
    4. [( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )](#proof-6-4)
    5. [( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )](#proof-6-5)
    6. [( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )](#proof-6-6)
    7. [( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )](#proof-6-7)
    8. [( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )](#proof-6-8)
    9. [( P v ( Q ^ R ) ) -> ( ( P v Q ) ^ ( P v R ) )](#proof-6-9)
    10. [( ( P v Q ) ^ ( P v R ) ) -> ( P v ( Q ^ R ) )](#proof-6-10)
7. [disjunction](#disjunction)
    1. [( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )](#proof-7-1)
    2. [( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S )](#proof-7-2)
    3. [( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q )](#proof-7-3)
    4. [( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R](#proof-7-4)
    5. [( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q )](#proof-7-5)
8. [biconditional](#biconditional)
    1. [( ( P <-> Q ) ^ P ) -> Q](#proof-8-1)
    2. [( ( P <-> Q ) ^ ~ P ) -> ~ Q](#proof-8-2)
    3. [( P <-> ~ Q ) -> ( ~ P <-> Q )](#proof-8-3)
    4. [( P <-> Q ) -> ( ~ P <-> ~ Q )](#proof-8-4)
    5. [( ~ P <-> ~ Q ) -> ( P <-> Q )](#proof-8-5)
    6. [( P <-> Q ) v ~ ( P <-> Q )](#proof-8-6)
    7. [( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )](#proof-8-7)
    8. [( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )](#proof-8-8)
9. [miscellaneous](#miscellaneous)
    1. [( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )](#proof-9-1)
    2. [( P -> Q ) v ( Q -> R )](#proof-9-2)
    3. [( P -> Q ) v ( Q -> P )](#proof-9-3)

# theorems <a name="theorems"></a>

## P v ~ P <a name="proof-1-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P v ~ P )</pre> | Assume Contradiction |  |
| 2 | <pre>.   .   P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   P v ~ P</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ~ ( P v ~ P )</pre> | Reiterate | 1 |
| 5 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   P v ~ P</pre> | I v (R) | 5 |
| 7 | <pre>P v ~ P</pre> | subproof contradiction | 1 - 6 |

## ( P -> Q ) <-> ( ~ Q -> ~ P ) <a name="proof-1-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P</pre> | Assume Contradiction |  |
| 4 | <pre>.   .   .   P -> Q</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   Q</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   .   ~ Q</pre> | Reiterate | 2 |
| 7 | <pre>.   .   ~ P</pre> | subproof contradiction | 3 - 6 |
| 8 | <pre>.   ~ Q -> ~ P</pre> | subproof implication | 2 - 7 |
| 9 | <pre>( P -> Q ) -> ( ~ Q -> ~ P )</pre> | subproof implication | 1 - 8 |
| 10 | <pre>.   ~ Q -> ~ P</pre> | Assume Implication |  |
| 11 | <pre>.   .   P</pre> | Assume Implication |  |
| 12 | <pre>.   .   .   ~ Q</pre> | Assume Contradiction |  |
| 13 | <pre>.   .   .   ~ Q -> ~ P</pre> | Reiterate | 10 |
| 14 | <pre>.   .   .   ~ P</pre> | E -> | 13, 12 |
| 15 | <pre>.   .   .   P</pre> | Reiterate | 11 |
| 16 | <pre>.   .   Q</pre> | subproof contradiction | 12 - 15 |
| 17 | <pre>.   P -> Q</pre> | subproof implication | 11 - 16 |
| 18 | <pre>( ~ Q -> ~ P ) -> ( P -> Q )</pre> | subproof implication | 10 - 17 |
| 19 | <pre>( P -> Q ) <-> ( ~ Q -> ~ P )</pre> | I <-> | 9, 18 |

## ~ ( P -> Q ) <-> ( P ^ ~ Q ) <a name="proof-1-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P -> Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   .   ~ Q</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   ~ P</pre> | Reiterate | 2 |
| 5 | <pre>.   .   ~ Q -> ~ P</pre> | subproof implication | 3 - 4 |
| 6 | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 5 |
| 7 | <pre>.   .   ~ ( P -> Q )</pre> | Reiterate | 1 |
| 8 | <pre>.   P</pre> | subproof contradiction | 2 - 7 |
| 9 | <pre>.   .   Q</pre> | Assume Contradiction |  |
| 10 | <pre>.   .   P</pre> | Reiterate | 8 |
| 11 | <pre>.   .   P -> Q</pre> | I -> | 10, 9 |
| 12 | <pre>.   .   ~ ( P -> Q )</pre> | Reiterate | 1 |
| 13 | <pre>.   ~ Q</pre> | subproof contradiction | 9 - 12 |
| 14 | <pre>.   P ^ ~ Q</pre> | I ^ | 8, 13 |
| 15 | <pre>~ ( P -> Q ) -> ( P ^ ~ Q )</pre> | subproof implication | 1 - 14 |
| 16 | <pre>.   P ^ ~ Q</pre> | Assume Implication |  |
| 17 | <pre>.   .   P -> Q</pre> | Assume Contradiction |  |
| 18 | <pre>.   .   P ^ ~ Q</pre> | Reiterate | 16 |
| 19 | <pre>.   .   P</pre> | E ^ (L) | 18 |
| 20 | <pre>.   .   ~ Q</pre> | E ^ (R) | 18 |
| 21 | <pre>.   .   Q</pre> | E -> | 17, 19 |
| 22 | <pre>.   ~ ( P -> Q )</pre> | subproof contradiction | 17 - 21 |
| 23 | <pre>( P ^ ~ Q ) -> ~ ( P -> Q )</pre> | subproof implication | 16 - 22 |
| 24 | <pre>~ ( P -> Q ) <-> ( P ^ ~ Q )</pre> | I <-> | 15, 23 |

## ( P v Q ) <-> ( ~ P -> Q ) <a name="proof-1-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   ~ P</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   Q</pre> | Reiterate | 2 |
| 5 | <pre>.   .   ~ P -> Q</pre> | subproof implication | 3 - 4 |
| 6 | <pre>.   Q -> ( ~ P -> Q )</pre> | subproof implication | 2 - 5 |
| 7 | <pre>.   .   P</pre> | Assume Implication |  |
| 8 | <pre>.   .   .   ~ Q</pre> | Assume Implication |  |
| 9 | <pre>.   .   .   P</pre> | Reiterate | 7 |
| 10 | <pre>.   .   ~ Q -> P</pre> | subproof implication | 8 - 9 |
| 11 | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 10 |
| 12 | <pre>.   P -> ( ~ P -> Q )</pre> | subproof implication | 7 - 11 |
| 13 | <pre>.   ~ P -> Q</pre> | E v | 12, 6, 1 |
| 14 | <pre>( P v Q ) -> ( ~ P -> Q )</pre> | subproof implication | 1 - 13 |
| 15 | <pre>.   ~ P -> Q</pre> | Assume Implication |  |
| 16 | <pre>.   .   P</pre> | Assume Implication |  |
| 17 | <pre>.   .   P v Q</pre> | I v (L) | 16 |
| 18 | <pre>.   P -> ( P v Q )</pre> | subproof implication | 16 - 17 |
| 19 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 20 | <pre>.   .   ~ P -> Q</pre> | Reiterate | 15 |
| 21 | <pre>.   .   Q</pre> | E -> | 20, 19 |
| 22 | <pre>.   .   P v Q</pre> | I v (R) | 21 |
| 23 | <pre>.   ~ P -> ( P v Q )</pre> | subproof implication | 19 - 22 |
| 24 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 25 | <pre>.   P v Q</pre> | E v | 18, 23, 24 |
| 26 | <pre>( ~ P -> Q ) -> ( P v Q )</pre> | subproof implication | 15 - 25 |
| 27 | <pre>( P v Q ) <-> ( ~ P -> Q )</pre> | I <-> | 14, 26 |

## ~ ( P v Q ) <-> ( ~ P ^ ~ Q ) <a name="proof-1-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P v Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ~ ( P v Q )</pre> | Reiterate | 1 |
| 5 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   .   Q</pre> | Assume Contradiction |  |
| 7 | <pre>.   .   P v Q</pre> | I v (R) | 6 |
| 8 | <pre>.   .   ~ ( P v Q )</pre> | Reiterate | 1 |
| 9 | <pre>.   ~ Q</pre> | subproof contradiction | 6 - 8 |
| 10 | <pre>.   ~ P ^ ~ Q</pre> | I ^ | 5, 9 |
| 11 | <pre>~ ( P v Q ) -> ( ~ P ^ ~ Q )</pre> | subproof implication | 1 - 10 |
| 12 | <pre>.   ~ P ^ ~ Q</pre> | Assume Implication |  |
| 13 | <pre>.   .   P</pre> | Assume Implication |  |
| 14 | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume Contradiction |  |
| 15 | <pre>.   .   .   ~ P</pre> | E ^ (L) | 14 |
| 16 | <pre>.   .   .   P</pre> | Reiterate | 13 |
| 17 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof contradiction | 14 - 16 |
| 18 | <pre>.   P -> ~ ( ~ P ^ ~ Q )</pre> | subproof implication | 13 - 17 |
| 19 | <pre>.   .   Q</pre> | Assume Implication |  |
| 20 | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume Contradiction |  |
| 21 | <pre>.   .   .   ~ Q</pre> | E ^ (R) | 20 |
| 22 | <pre>.   .   .   Q</pre> | Reiterate | 19 |
| 23 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof contradiction | 20 - 22 |
| 24 | <pre>.   Q -> ~ ( ~ P ^ ~ Q )</pre> | subproof implication | 19 - 23 |
| 25 | <pre>.   .   P v Q</pre> | Assume Contradiction |  |
| 26 | <pre>.   .   P -> ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 18 |
| 27 | <pre>.   .   Q -> ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 24 |
| 28 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | E v | 26, 27, 25 |
| 29 | <pre>.   .   ~ P ^ ~ Q</pre> | Reiterate | 12 |
| 30 | <pre>.   ~ ( P v Q )</pre> | subproof contradiction | 25 - 29 |
| 31 | <pre>( ~ P ^ ~ Q ) -> ~ ( P v Q )</pre> | subproof implication | 12 - 30 |
| 32 | <pre>~ ( P v Q ) <-> ( ~ P ^ ~ Q )</pre> | I <-> | 11, 31 |

## ~ ( P ^ Q ) <-> ( ~ P v ~ Q ) <a name="proof-1-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P ^ Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ ( ~ P v ~ Q )</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   .   ~ P</pre> | Assume Contradiction |  |
| 4 | <pre>.   .   .   ~ P v ~ Q</pre> | I v (L) | 3 |
| 5 | <pre>.   .   .   ~ ( ~ P v ~ Q )</pre> | Reiterate | 2 |
| 6 | <pre>.   .   P</pre> | subproof contradiction | 3 - 5 |
| 7 | <pre>.   .   .   ~ Q</pre> | Assume Contradiction |  |
| 8 | <pre>.   .   .   ~ P v ~ Q</pre> | I v (R) | 7 |
| 9 | <pre>.   .   .   ~ ( ~ P v ~ Q )</pre> | Reiterate | 2 |
| 10 | <pre>.   .   Q</pre> | subproof contradiction | 7 - 9 |
| 11 | <pre>.   .   P ^ Q</pre> | I ^ | 6, 10 |
| 12 | <pre>.   .   ~ ( P ^ Q )</pre> | Reiterate | 1 |
| 13 | <pre>.   ~ P v ~ Q</pre> | subproof contradiction | 2 - 12 |
| 14 | <pre>~ ( P ^ Q ) -> ( ~ P v ~ Q )</pre> | subproof implication | 1 - 13 |
| 15 | <pre>.   ~ P v ~ Q</pre> | Assume Implication |  |
| 16 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 17 | <pre>.   .   .   P ^ Q</pre> | Assume Contradiction |  |
| 18 | <pre>.   .   .   P</pre> | E ^ (L) | 17 |
| 19 | <pre>.   .   .   ~ P</pre> | Reiterate | 16 |
| 20 | <pre>.   .   ~ ( P ^ Q )</pre> | subproof contradiction | 17 - 19 |
| 21 | <pre>.   ~ P -> ~ ( P ^ Q )</pre> | subproof implication | 16 - 20 |
| 22 | <pre>.   .   ~ Q</pre> | Assume Implication |  |
| 23 | <pre>.   .   .   P ^ Q</pre> | Assume Contradiction |  |
| 24 | <pre>.   .   .   Q</pre> | E ^ (R) | 23 |
| 25 | <pre>.   .   .   ~ Q</pre> | Reiterate | 22 |
| 26 | <pre>.   .   ~ ( P ^ Q )</pre> | subproof contradiction | 23 - 25 |
| 27 | <pre>.   ~ Q -> ~ ( P ^ Q )</pre> | subproof implication | 22 - 26 |
| 28 | <pre>.   ~ ( P ^ Q )</pre> | E v | 21, 27, 15 |
| 29 | <pre>( ~ P v ~ Q ) -> ~ ( P ^ Q )</pre> | subproof implication | 15 - 28 |
| 30 | <pre>~ ( P ^ Q ) <-> ( ~ P v ~ Q )</pre> | I <-> | 14, 29 |

## ( P <-> ~ Q ) <-> ~ ( P <-> Q ) <a name="proof-1-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   ~ Q -> P</pre> | E <-> (R) | 1 |
| 4 | <pre>.   .   P</pre> | Assume Implication |  |
| 5 | <pre>.   .   .   P <-> Q</pre> | Assume Contradiction |  |
| 6 | <pre>.   .   .   P -> ~ Q</pre> | Reiterate | 2 |
| 7 | <pre>.   .   .   P -> Q</pre> | E <-> (L) | 5 |
| 8 | <pre>.   .   .   P</pre> | Reiterate | 4 |
| 9 | <pre>.   .   .   Q</pre> | E -> | 7, 8 |
| 10 | <pre>.   .   .   ~ Q</pre> | E -> | 6, 8 |
| 11 | <pre>.   .   ~ ( P <-> Q )</pre> | subproof contradiction | 5 - 10 |
| 12 | <pre>.   P -> ~ ( P <-> Q )</pre> | subproof implication | 4 - 11 |
| 13 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 14 | <pre>.   .   .   P <-> Q</pre> | Assume Contradiction |  |
| 15 | <pre>.   .   .   Q -> P</pre> | E <-> (R) | 14 |
| 16 | <pre>.   .   .   ~ P -> ~ Q</pre> | Theorem: contrapositive | 15 |
| 17 | <pre>.   .   .   ~ P</pre> | Reiterate | 13 |
| 18 | <pre>.   .   .   ~ Q</pre> | E -> | 16, 17 |
| 19 | <pre>.   .   .   ~ Q -> P</pre> | Reiterate | 3 |
| 20 | <pre>.   .   .   P</pre> | E -> | 19, 18 |
| 21 | <pre>.   .   ~ ( P <-> Q )</pre> | subproof contradiction | 14 - 20 |
| 22 | <pre>.   ~ P -> ~ ( P <-> Q )</pre> | subproof implication | 13 - 21 |
| 23 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 24 | <pre>.   ~ ( P <-> Q )</pre> | E v | 12, 22, 23 |
| 25 | <pre>( P <-> ~ Q ) -> ~ ( P <-> Q )</pre> | subproof implication | 1 - 24 |
| 26 | <pre>.   ~ ( P <-> Q )</pre> | Assume Implication |  |
| 27 | <pre>.   .   ~ ( P -> ~ Q )</pre> | Assume Contradiction |  |
| 28 | <pre>.   .   P ^ Q</pre> | Theorem: (~ ->) to ^ | 27 |
| 29 | <pre>.   .   P</pre> | E ^ (L) | 28 |
| 30 | <pre>.   .   Q</pre> | E ^ (R) | 28 |
| 31 | <pre>.   .   P -> Q</pre> | I -> | 29, 30 |
| 32 | <pre>.   .   Q -> P</pre> | I -> | 30, 29 |
| 33 | <pre>.   .   P <-> Q</pre> | I <-> | 31, 32 |
| 34 | <pre>.   .   ~ ( P <-> Q )</pre> | Reiterate | 26 |
| 35 | <pre>.   P -> ~ Q</pre> | subproof contradiction | 27 - 34 |
| 36 | <pre>.   .   ~ ( ~ Q -> P )</pre> | Assume Contradiction |  |
| 37 | <pre>.   .   ~ Q ^ ~ P</pre> | Theorem: (~ ->) to ^ | 36 |
| 38 | <pre>.   .   ~ Q</pre> | E ^ (L) | 37 |
| 39 | <pre>.   .   ~ P</pre> | E ^ (R) | 37 |
| 40 | <pre>.   .   ~ P -> ~ Q</pre> | I -> | 39, 38 |
| 41 | <pre>.   .   Q -> P</pre> | Theorem: contrapositive | 40 |
| 42 | <pre>.   .   ~ Q -> ~ P</pre> | I -> | 38, 39 |
| 43 | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 42 |
| 44 | <pre>.   .   P <-> Q</pre> | I <-> | 43, 41 |
| 45 | <pre>.   .   ~ ( P <-> Q )</pre> | Reiterate | 26 |
| 46 | <pre>.   ~ Q -> P</pre> | subproof contradiction | 36 - 45 |
| 47 | <pre>.   P <-> ~ Q</pre> | I <-> | 35, 46 |
| 48 | <pre>~ ( P <-> Q ) -> ( P <-> ~ Q )</pre> | subproof implication | 26 - 47 |
| 49 | <pre>( P <-> ~ Q ) <-> ~ ( P <-> Q )</pre> | I <-> | 25, 48 |

# basics <a name="basics"></a>

## P -> P <a name="proof-2-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P</pre> | Assume Implication |  |
| 2 | <pre>P -> P</pre> | subproof implication | 1 - 1 |

## ~ ( P ^ ~ P ) <a name="proof-2-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ~ P</pre> | Assume Contradiction |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P</pre> | E ^ (R) | 1 |
| 4 | <pre>~ ( P ^ ~ P )</pre> | subproof contradiction | 1 - 3 |

## ( P ^ P ) <-> P <a name="proof-2-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ P</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>( P ^ P ) -> P</pre> | subproof implication | 1 - 2 |
| 4 | <pre>.   P</pre> | Assume Implication |  |
| 5 | <pre>.   P ^ P</pre> | I ^ | 4, 4 |
| 6 | <pre>P -> ( P ^ P )</pre> | subproof implication | 4 - 5 |
| 7 | <pre>( P ^ P ) <-> P</pre> | I <-> | 3, 6 |

## ( P v P ) <-> P <a name="proof-2-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v P</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   P -> P</pre> | subproof implication | 2 - 2 |
| 4 | <pre>.   P</pre> | E v | 3, 3, 1 |
| 5 | <pre>( P v P ) -> P</pre> | subproof implication | 1 - 4 |
| 6 | <pre>.   P</pre> | Assume Implication |  |
| 7 | <pre>.   P v P</pre> | I v (L) | 6 |
| 8 | <pre>P -> ( P v P )</pre> | subproof implication | 6 - 7 |
| 9 | <pre>( P v P ) <-> P</pre> | I <-> | 5, 8 |

## ( P ^ Q ) -> ( P v Q ) <a name="proof-2-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ Q</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>( P ^ Q ) -> ( P v Q )</pre> | subproof implication | 1 - 3 |

## ( P -> ~ P ) -> ~ P <a name="proof-2-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ~ P</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   P -> ~ P</pre> | Reiterate | 1 |
| 4 | <pre>.   .   ~ P</pre> | E -> | 3, 2 |
| 5 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>( P -> ~ P ) -> ~ P</pre> | subproof implication | 1 - 5 |

## ( ( P -> Q ) -> P ) -> P <a name="proof-2-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) -> P</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ ( P -> Q )</pre> | Assume Implication |  |
| 3 | <pre>.   .   P ^ ~ Q</pre> | Theorem: (~ ->) to ^ | 2 |
| 4 | <pre>.   .   P</pre> | E ^ (L) | 3 |
| 5 | <pre>.   ~ ( P -> Q ) -> P</pre> | subproof implication | 2 - 4 |
| 6 | <pre>.   ( P -> Q ) v ~ ( P -> Q )</pre> | Theorem: excluded middle |  |
| 7 | <pre>.   P</pre> | E v | 1, 5, 6 |
| 8 | <pre>( ( P -> Q ) -> P ) -> P</pre> | subproof implication | 1 - 7 |

# arrows <a name="arrows"></a>

## Q -> ( P -> Q ) <a name="proof-3-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   Q</pre> | Reiterate | 1 |
| 4 | <pre>.   P -> Q</pre> | subproof implication | 2 - 3 |
| 5 | <pre>Q -> ( P -> Q )</pre> | subproof implication | 1 - 4 |

## ~ P -> ( P -> Q ) <a name="proof-3-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ P</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   ~ P</pre> | Reiterate | 1 |
| 4 | <pre>.   ~ Q -> ~ P</pre> | subproof implication | 2 - 3 |
| 5 | <pre>.   P -> Q</pre> | Theorem: contrapositive | 4 |
| 6 | <pre>~ P -> ( P -> Q )</pre> | subproof implication | 1 - 5 |

## ( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R ) <a name="proof-3-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( Q -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   ( P -> Q ) ^ ( Q -> R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P -> Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   Q -> R</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E -> | 4, 2 |
| 7 | <pre>.   .   R</pre> | E -> | 5, 6 |
| 8 | <pre>.   P -> R</pre> | subproof implication | 2 - 7 |
| 9 | <pre>( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )</pre> | subproof implication | 1 - 8 |

## ( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) ) <a name="proof-3-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P ^ Q ) -> R</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 5 | <pre>.   .   .   P ^ Q</pre> | I ^ | 4, 3 |
| 6 | <pre>.   .   .   ( P ^ Q ) -> R</pre> | Reiterate | 1 |
| 7 | <pre>.   .   .   R</pre> | E -> | 6, 5 |
| 8 | <pre>.   .   Q -> R</pre> | subproof implication | 3 - 7 |
| 9 | <pre>.   P -> ( Q -> R )</pre> | subproof implication | 2 - 8 |
| 10 | <pre>( ( P ^ Q ) -> R ) -> ( P -> ( Q -> R ) )</pre> | subproof implication | 1 - 9 |
| 11 | <pre>.   P -> ( Q -> R )</pre> | Assume Implication |  |
| 12 | <pre>.   .   P ^ Q</pre> | Assume Implication |  |
| 13 | <pre>.   .   P</pre> | E ^ (L) | 12 |
| 14 | <pre>.   .   Q</pre> | E ^ (R) | 12 |
| 15 | <pre>.   .   P -> ( Q -> R )</pre> | Reiterate | 11 |
| 16 | <pre>.   .   Q -> R</pre> | E -> | 15, 13 |
| 17 | <pre>.   .   R</pre> | E -> | 16, 14 |
| 18 | <pre>.   ( P ^ Q ) -> R</pre> | subproof implication | 12 - 17 |
| 19 | <pre>( P -> ( Q -> R ) ) -> ( ( P ^ Q ) -> R )</pre> | subproof implication | 11 - 18 |
| 20 | <pre>( ( P ^ Q ) -> R ) <-> ( P -> ( Q -> R ) )</pre> | I <-> | 10, 19 |

## ( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P <a name="proof-3-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( P -> ~ Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   ( P -> Q ) ^ ( P -> ~ Q )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P -> Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   P -> ~ Q</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E -> | 4, 2 |
| 7 | <pre>.   .   ~ Q</pre> | E -> | 5, 2 |
| 8 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 7 |
| 9 | <pre>( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P</pre> | subproof implication | 1 - 8 |

## ( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q <a name="proof-3-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( ~ P -> Q )</pre> | Assume Implication |  |
| 2 | <pre>.   P -> Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P -> Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 5 | <pre>.   Q</pre> | E v | 2, 3, 4 |
| 6 | <pre>( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q</pre> | subproof implication | 1 - 5 |

## ( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) ) <a name="proof-3-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   P -> ( Q -> R )</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   Q -> R</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   .   Q</pre> | Reiterate | 2 |
| 7 | <pre>.   .   .   R</pre> | E -> | 5, 6 |
| 8 | <pre>.   .   P -> R</pre> | subproof implication | 3 - 7 |
| 9 | <pre>.   Q -> ( P -> R )</pre> | subproof implication | 2 - 8 |
| 10 | <pre>( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )</pre> | subproof implication | 1 - 9 |

## ( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) ) <a name="proof-3-8"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P v Q ) -> R</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ( P v Q ) -> R</pre> | Reiterate | 1 |
| 5 | <pre>.   .   R</pre> | E -> | 4, 3 |
| 6 | <pre>.   P -> R</pre> | subproof implication | 2 - 5 |
| 7 | <pre>.   .   Q</pre> | Assume Implication |  |
| 8 | <pre>.   .   P v Q</pre> | I v (R) | 7 |
| 9 | <pre>.   .   ( P v Q ) -> R</pre> | Reiterate | 1 |
| 10 | <pre>.   .   R</pre> | E -> | 9, 8 |
| 11 | <pre>.   Q -> R</pre> | subproof implication | 7 - 10 |
| 12 | <pre>.   ( P -> R ) ^ ( Q -> R )</pre> | I ^ | 6, 11 |
| 13 | <pre>( ( P v Q ) -> R ) -> ( ( P -> R ) ^ ( Q -> R ) )</pre> | subproof implication | 1 - 12 |
| 14 | <pre>.   ( P -> R ) ^ ( Q -> R )</pre> | Assume Implication |  |
| 15 | <pre>.   .   P v Q</pre> | Assume Implication |  |
| 16 | <pre>.   .   ( P -> R ) ^ ( Q -> R )</pre> | Reiterate | 14 |
| 17 | <pre>.   .   P -> R</pre> | E ^ (L) | 16 |
| 18 | <pre>.   .   Q -> R</pre> | E ^ (R) | 16 |
| 19 | <pre>.   .   R</pre> | E v | 17, 18, 15 |
| 20 | <pre>.   ( P v Q ) -> R</pre> | subproof implication | 15 - 19 |
| 21 | <pre>( ( P -> R ) ^ ( Q -> R ) ) -> ( ( P v Q ) -> R )</pre> | subproof implication | 14 - 20 |
| 22 | <pre>( ( P v Q ) -> R ) <-> ( ( P -> R ) ^ ( Q -> R ) )</pre> | I <-> | 13, 21 |

## ( ( P -> R ) v ( Q -> R ) ) <-> ( ( P ^ Q ) -> R ) <a name="proof-3-9"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> R ) v ( Q -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P -> R</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P ^ Q</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   P -> R</pre> | Reiterate | 2 |
| 5 | <pre>.   .   .   P</pre> | E ^ (L) | 3 |
| 6 | <pre>.   .   .   R</pre> | E -> | 4, 5 |
| 7 | <pre>.   .   ( P ^ Q ) -> R</pre> | subproof implication | 3 - 6 |
| 8 | <pre>.   ( P -> R ) -> ( ( P ^ Q ) -> R )</pre> | subproof implication | 2 - 7 |
| 9 | <pre>.   .   Q -> R</pre> | Assume Implication |  |
| 10 | <pre>.   .   .   P ^ Q</pre> | Assume Implication |  |
| 11 | <pre>.   .   .   Q -> R</pre> | Reiterate | 9 |
| 12 | <pre>.   .   .   Q</pre> | E ^ (R) | 10 |
| 13 | <pre>.   .   .   R</pre> | E -> | 11, 12 |
| 14 | <pre>.   .   ( P ^ Q ) -> R</pre> | subproof implication | 10 - 13 |
| 15 | <pre>.   ( Q -> R ) -> ( ( P ^ Q ) -> R )</pre> | subproof implication | 9 - 14 |
| 16 | <pre>.   ( P ^ Q ) -> R</pre> | E v | 8, 15, 1 |
| 17 | <pre>( ( P -> R ) v ( Q -> R ) ) -> ( ( P ^ Q ) -> R )</pre> | subproof implication | 1 - 16 |
| 18 | <pre>.   ( P ^ Q ) -> R</pre> | Assume Implication |  |
| 19 | <pre>.   .   ~ ( ( P -> R ) v ( Q -> R ) )</pre> | Assume Contradiction |  |
| 20 | <pre>.   .   ~ ( P -> R ) ^ ~ ( Q -> R )</pre> | DeMorgan's theorem: (~ v) to ^ | 19 |
| 21 | <pre>.   .   ~ ( P -> R )</pre> | E ^ (L) | 20 |
| 22 | <pre>.   .   ~ ( Q -> R )</pre> | E ^ (R) | 20 |
| 23 | <pre>.   .   P ^ ~ R</pre> | Theorem: (~ ->) to ^ | 21 |
| 24 | <pre>.   .   P</pre> | E ^ (L) | 23 |
| 25 | <pre>.   .   ~ R</pre> | E ^ (R) | 23 |
| 26 | <pre>.   .   Q ^ ~ R</pre> | Theorem: (~ ->) to ^ | 22 |
| 27 | <pre>.   .   Q</pre> | E ^ (L) | 26 |
| 28 | <pre>.   .   ( P ^ Q ) -> R</pre> | Reiterate | 18 |
| 29 | <pre>.   .   P ^ Q</pre> | I ^ | 24, 27 |
| 30 | <pre>.   .   R</pre> | E -> | 28, 29 |
| 31 | <pre>.   ( P -> R ) v ( Q -> R )</pre> | subproof contradiction | 19 - 30 |
| 32 | <pre>( ( P ^ Q ) -> R ) -> ( ( P -> R ) v ( Q -> R ) )</pre> | subproof implication | 18 - 31 |
| 33 | <pre>( ( P -> R ) v ( Q -> R ) ) <-> ( ( P ^ Q ) -> R )</pre> | I <-> | 17, 32 |

# commutativity <a name="commutativity"></a>

## ( P ^ Q ) -> ( Q ^ P ) <a name="proof-4-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ Q</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q ^ P</pre> | I ^ | 3, 2 |
| 5 | <pre>( P ^ Q ) -> ( Q ^ P )</pre> | subproof implication | 1 - 4 |

## ( P v Q ) -> ( Q v P ) <a name="proof-4-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   Q v P</pre> | I v (R) | 2 |
| 4 | <pre>.   P -> ( Q v P )</pre> | subproof implication | 2 - 3 |
| 5 | <pre>.   .   Q</pre> | Assume Implication |  |
| 6 | <pre>.   .   Q v P</pre> | I v (L) | 5 |
| 7 | <pre>.   Q -> ( Q v P )</pre> | subproof implication | 5 - 6 |
| 8 | <pre>.   Q v P</pre> | E v | 4, 7, 1 |
| 9 | <pre>( P v Q ) -> ( Q v P )</pre> | subproof implication | 1 - 8 |

## ( P <-> Q ) -> ( Q <-> P ) <a name="proof-4-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> Q</pre> | Assume Implication |  |
| 2 | <pre>.   P -> Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   Q -> P</pre> | E <-> (R) | 1 |
| 4 | <pre>.   Q <-> P</pre> | I <-> | 3, 2 |
| 5 | <pre>( P <-> Q ) -> ( Q <-> P )</pre> | subproof implication | 1 - 4 |

# associativity <a name="associativity"></a>

## ( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) ) <a name="proof-5-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P ^ Q ) ^ R</pre> | Assume Implication |  |
| 2 | <pre>.   P ^ Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P</pre> | E ^ (L) | 2 |
| 5 | <pre>.   Q</pre> | E ^ (R) | 2 |
| 6 | <pre>.   Q ^ R</pre> | I ^ | 5, 3 |
| 7 | <pre>.   P ^ ( Q ^ R )</pre> | I ^ | 4, 6 |
| 8 | <pre>( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )</pre> | subproof implication | 1 - 7 |
| 9 | <pre>.   P ^ ( Q ^ R )</pre> | Assume Implication |  |
| 10 | <pre>.   P</pre> | E ^ (L) | 9 |
| 11 | <pre>.   Q ^ R</pre> | E ^ (R) | 9 |
| 12 | <pre>.   Q</pre> | E ^ (L) | 11 |
| 13 | <pre>.   R</pre> | E ^ (R) | 11 |
| 14 | <pre>.   P ^ Q</pre> | I ^ | 10, 12 |
| 15 | <pre>.   ( P ^ Q ) ^ R</pre> | I ^ | 14, 13 |
| 16 | <pre>( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )</pre> | subproof implication | 9 - 15 |
| 17 | <pre>( ( P ^ Q ) ^ R ) <-> ( P ^ ( Q ^ R ) )</pre> | I <-> | 8, 16 |

## ( ( P v Q ) v R ) <-> ( P v ( Q v R ) ) <a name="proof-5-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P v Q ) v R</pre> | Assume Implication |  |
| 2 | <pre>.   .   P v Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   P v ( Q v R )</pre> | I v (L) | 3 |
| 5 | <pre>.   .   P -> ( P v ( Q v R ) )</pre> | subproof implication | 3 - 4 |
| 6 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 7 | <pre>.   .   .   Q v R</pre> | I v (L) | 6 |
| 8 | <pre>.   .   .   P v ( Q v R )</pre> | I v (R) | 7 |
| 9 | <pre>.   .   Q -> ( P v ( Q v R ) )</pre> | subproof implication | 6 - 8 |
| 10 | <pre>.   .   P v ( Q v R )</pre> | E v | 5, 9, 2 |
| 11 | <pre>.   ( P v Q ) -> ( P v ( Q v R ) )</pre> | subproof implication | 2 - 10 |
| 12 | <pre>.   .   R</pre> | Assume Implication |  |
| 13 | <pre>.   .   Q v R</pre> | I v (R) | 12 |
| 14 | <pre>.   .   P v ( Q v R )</pre> | I v (R) | 13 |
| 15 | <pre>.   R -> ( P v ( Q v R ) )</pre> | subproof implication | 12 - 14 |
| 16 | <pre>.   P v ( Q v R )</pre> | E v | 11, 15, 1 |
| 17 | <pre>( ( P v Q ) v R ) -> ( P v ( Q v R ) )</pre> | subproof implication | 1 - 16 |
| 18 | <pre>.   P v ( Q v R )</pre> | Assume Implication |  |
| 19 | <pre>.   .   P</pre> | Assume Implication |  |
| 20 | <pre>.   .   P v Q</pre> | I v (L) | 19 |
| 21 | <pre>.   .   ( P v Q ) v R</pre> | I v (L) | 20 |
| 22 | <pre>.   P -> ( ( P v Q ) v R )</pre> | subproof implication | 19 - 21 |
| 23 | <pre>.   .   Q v R</pre> | Assume Implication |  |
| 24 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 25 | <pre>.   .   .   P v Q</pre> | I v (R) | 24 |
| 26 | <pre>.   .   .   ( P v Q ) v R</pre> | I v (L) | 25 |
| 27 | <pre>.   .   Q -> ( ( P v Q ) v R )</pre> | subproof implication | 24 - 26 |
| 28 | <pre>.   .   .   R</pre> | Assume Implication |  |
| 29 | <pre>.   .   .   ( P v Q ) v R</pre> | I v (R) | 28 |
| 30 | <pre>.   .   R -> ( ( P v Q ) v R )</pre> | subproof implication | 28 - 29 |
| 31 | <pre>.   .   ( P v Q ) v R</pre> | E v | 27, 30, 23 |
| 32 | <pre>.   ( Q v R ) -> ( ( P v Q ) v R )</pre> | subproof implication | 23 - 31 |
| 33 | <pre>.   ( P v Q ) v R</pre> | E v | 22, 32, 18 |
| 34 | <pre>( P v ( Q v R ) ) -> ( ( P v Q ) v R )</pre> | subproof implication | 18 - 33 |
| 35 | <pre>( ( P v Q ) v R ) <-> ( P v ( Q v R ) )</pre> | I <-> | 17, 34 |

## ( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) ) <a name="proof-5-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P <-> Q ) <-> R</pre> | Assume Implication |  |
| 2 | <pre>.   ( P <-> Q ) -> R</pre> | E <-> (L) | 1 |
| 3 | <pre>.   R -> ( P <-> Q )</pre> | E <-> (R) | 1 |
| 4 | <pre>.   .   ~ ( P -> ( Q <-> R ) )</pre> | Assume Contradiction |  |
| 5 | <pre>.   .   P ^ ~ ( Q <-> R )</pre> | Theorem: (~ ->) to ^ | 4 |
| 6 | <pre>.   .   P</pre> | E ^ (L) | 5 |
| 7 | <pre>.   .   ~ ( Q <-> R )</pre> | E ^ (R) | 5 |
| 8 | <pre>.   .   Q <-> ~ R</pre> | Theorem: biconditional negation | 7 |
| 9 | <pre>.   .   .   R</pre> | Assume Contradiction |  |
| 10 | <pre>.   .   .   R -> ( P <-> Q )</pre> | Reiterate | 3 |
| 11 | <pre>.   .   .   P <-> Q</pre> | E -> | 10, 9 |
| 12 | <pre>.   .   .   P -> Q</pre> | E <-> (L) | 11 |
| 13 | <pre>.   .   .   P</pre> | Reiterate | 6 |
| 14 | <pre>.   .   .   Q</pre> | E -> | 12, 13 |
| 15 | <pre>.   .   .   Q <-> ~ R</pre> | Reiterate | 8 |
| 16 | <pre>.   .   .   Q -> ~ R</pre> | E <-> (L) | 15 |
| 17 | <pre>.   .   .   ~ R</pre> | E -> | 16, 14 |
| 18 | <pre>.   .   ~ R</pre> | subproof contradiction | 9 - 17 |
| 19 | <pre>.   .   ~ R -> Q</pre> | E <-> (R) | 8 |
| 20 | <pre>.   .   Q</pre> | E -> | 19, 18 |
| 21 | <pre>.   .   ( P <-> Q ) -> R</pre> | Reiterate | 2 |
| 22 | <pre>.   .   ~ R -> ~ ( P <-> Q )</pre> | Theorem: contrapositive | 21 |
| 23 | <pre>.   .   ~ ( P <-> Q )</pre> | E -> | 22, 18 |
| 24 | <pre>.   .   P <-> ~ Q</pre> | Theorem: biconditional negation | 23 |
| 25 | <pre>.   .   P -> ~ Q</pre> | E <-> (L) | 24 |
| 26 | <pre>.   .   ~ Q</pre> | E -> | 25, 6 |
| 27 | <pre>.   P -> ( Q <-> R )</pre> | subproof contradiction | 4 - 26 |
| 28 | <pre>.   .   ~ ( ( Q <-> R ) -> P )</pre> | Assume Contradiction |  |
| 29 | <pre>.   .   ( Q <-> R ) ^ ~ P</pre> | Theorem: (~ ->) to ^ | 28 |
| 30 | <pre>.   .   Q <-> R</pre> | E ^ (L) | 29 |
| 31 | <pre>.   .   ~ P</pre> | E ^ (R) | 29 |
| 32 | <pre>.   .   .   R</pre> | Assume Contradiction |  |
| 33 | <pre>.   .   .   R -> ( P <-> Q )</pre> | Reiterate | 3 |
| 34 | <pre>.   .   .   Q <-> R</pre> | Reiterate | 30 |
| 35 | <pre>.   .   .   R -> Q</pre> | E <-> (R) | 34 |
| 36 | <pre>.   .   .   Q</pre> | E -> | 35, 32 |
| 37 | <pre>.   .   .   P <-> Q</pre> | E -> | 33, 32 |
| 38 | <pre>.   .   .   Q -> P</pre> | E <-> (R) | 37 |
| 39 | <pre>.   .   .   P</pre> | E -> | 38, 36 |
| 40 | <pre>.   .   .   ~ P</pre> | Reiterate | 31 |
| 41 | <pre>.   .   ~ R</pre> | subproof contradiction | 32 - 40 |
| 42 | <pre>.   .   Q -> R</pre> | E <-> (L) | 30 |
| 43 | <pre>.   .   ~ R -> ~ Q</pre> | Theorem: contrapositive | 42 |
| 44 | <pre>.   .   ~ Q</pre> | E -> | 43, 41 |
| 45 | <pre>.   .   ( P <-> Q ) -> R</pre> | Reiterate | 2 |
| 46 | <pre>.   .   ~ R -> ~ ( P <-> Q )</pre> | Theorem: contrapositive | 45 |
| 47 | <pre>.   .   ~ ( P <-> Q )</pre> | E -> | 46, 41 |
| 48 | <pre>.   .   P <-> ~ Q</pre> | Theorem: biconditional negation | 47 |
| 49 | <pre>.   .   ~ Q -> P</pre> | E <-> (R) | 48 |
| 50 | <pre>.   .   P</pre> | E -> | 49, 44 |
| 51 | <pre>.   ( Q <-> R ) -> P</pre> | subproof contradiction | 28 - 50 |
| 52 | <pre>.   P <-> ( Q <-> R )</pre> | I <-> | 27, 51 |
| 53 | <pre>( ( P <-> Q ) <-> R ) -> ( P <-> ( Q <-> R ) )</pre> | subproof implication | 1 - 52 |
| 54 | <pre>.   P <-> ( Q <-> R )</pre> | Assume Implication |  |
| 55 | <pre>.   P -> ( Q <-> R )</pre> | E <-> (L) | 54 |
| 56 | <pre>.   ( Q <-> R ) -> P</pre> | E <-> (R) | 54 |
| 57 | <pre>.   .   ~ ( ( P <-> Q ) -> R )</pre> | Assume Contradiction |  |
| 58 | <pre>.   .   ( P <-> Q ) ^ ~ R</pre> | Theorem: (~ ->) to ^ | 57 |
| 59 | <pre>.   .   P <-> Q</pre> | E ^ (L) | 58 |
| 60 | <pre>.   .   ~ R</pre> | E ^ (R) | 58 |
| 61 | <pre>.   .   P -> Q</pre> | E <-> (L) | 59 |
| 62 | <pre>.   .   .   P</pre> | Assume Contradiction |  |
| 63 | <pre>.   .   .   P -> Q</pre> | Reiterate | 61 |
| 64 | <pre>.   .   .   Q</pre> | E -> | 63, 62 |
| 65 | <pre>.   .   .   P -> ( Q <-> R )</pre> | Reiterate | 55 |
| 66 | <pre>.   .   .   Q <-> R</pre> | E -> | 65, 62 |
| 67 | <pre>.   .   .   Q -> R</pre> | E <-> (L) | 66 |
| 68 | <pre>.   .   .   R</pre> | E -> | 67, 64 |
| 69 | <pre>.   .   .   ~ R</pre> | Reiterate | 60 |
| 70 | <pre>.   .   ~ P</pre> | subproof contradiction | 62 - 69 |
| 71 | <pre>.   .   ( Q <-> R ) -> P</pre> | Reiterate | 56 |
| 72 | <pre>.   .   ~ P -> ~ ( Q <-> R )</pre> | Theorem: contrapositive | 71 |
| 73 | <pre>.   .   ~ ( Q <-> R )</pre> | E -> | 72, 70 |
| 74 | <pre>.   .   Q <-> ~ R</pre> | Theorem: biconditional negation | 73 |
| 75 | <pre>.   .   ~ R -> Q</pre> | E <-> (R) | 74 |
| 76 | <pre>.   .   Q</pre> | E -> | 75, 60 |
| 77 | <pre>.   .   Q -> P</pre> | E <-> (R) | 59 |
| 78 | <pre>.   .   P</pre> | E -> | 77, 76 |
| 79 | <pre>.   ( P <-> Q ) -> R</pre> | subproof contradiction | 57 - 78 |
| 80 | <pre>.   .   ~ ( R -> ( P <-> Q ) )</pre> | Assume Contradiction |  |
| 81 | <pre>.   .   R ^ ~ ( P <-> Q )</pre> | Theorem: (~ ->) to ^ | 80 |
| 82 | <pre>.   .   R</pre> | E ^ (L) | 81 |
| 83 | <pre>.   .   ~ ( P <-> Q )</pre> | E ^ (R) | 81 |
| 84 | <pre>.   .   P <-> ~ Q</pre> | Theorem: biconditional negation | 83 |
| 85 | <pre>.   .   P -> ~ Q</pre> | E <-> (L) | 84 |
| 86 | <pre>.   .   .   P</pre> | Assume Contradiction |  |
| 87 | <pre>.   .   .   P -> ~ Q</pre> | Reiterate | 85 |
| 88 | <pre>.   .   .   ~ Q</pre> | E -> | 87, 86 |
| 89 | <pre>.   .   .   P -> ( Q <-> R )</pre> | Reiterate | 55 |
| 90 | <pre>.   .   .   Q <-> R</pre> | E -> | 89, 86 |
| 91 | <pre>.   .   .   R -> Q</pre> | E <-> (R) | 90 |
| 92 | <pre>.   .   .   R</pre> | Reiterate | 82 |
| 93 | <pre>.   .   .   Q</pre> | E -> | 91, 92 |
| 94 | <pre>.   .   ~ P</pre> | subproof contradiction | 86 - 93 |
| 95 | <pre>.   .   ~ Q -> P</pre> | E <-> (R) | 84 |
| 96 | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 95 |
| 97 | <pre>.   .   Q</pre> | E -> | 96, 94 |
| 98 | <pre>.   .   ( Q <-> R ) -> P</pre> | Reiterate | 56 |
| 99 | <pre>.   .   ~ P -> ~ ( Q <-> R )</pre> | Theorem: contrapositive | 98 |
| 100 | <pre>.   .   ~ ( Q <-> R )</pre> | E -> | 99, 94 |
| 101 | <pre>.   .   Q <-> ~ R</pre> | Theorem: biconditional negation | 100 |
| 102 | <pre>.   .   Q -> ~ R</pre> | E <-> (L) | 101 |
| 103 | <pre>.   .   ~ R</pre> | E -> | 102, 97 |
| 104 | <pre>.   R -> ( P <-> Q )</pre> | subproof contradiction | 80 - 103 |
| 105 | <pre>.   ( P <-> Q ) <-> R</pre> | I <-> | 79, 104 |
| 106 | <pre>( P <-> ( Q <-> R ) ) -> ( ( P <-> Q ) <-> R )</pre> | subproof implication | 54 - 105 |
| 107 | <pre>( ( P <-> Q ) <-> R ) <-> ( P <-> ( Q <-> R ) )</pre> | I <-> | 53, 106 |

# distributivity <a name="distributivity"></a>

## ( P -> ( Q -> R ) ) -> ( ( P -> Q ) -> ( P -> R ) ) <a name="proof-6-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P -> Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   P -> ( Q -> R )</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   Q -> R</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   .   P -> Q</pre> | Reiterate | 2 |
| 7 | <pre>.   .   .   Q</pre> | E -> | 6, 3 |
| 8 | <pre>.   .   .   R</pre> | E -> | 5, 7 |
| 9 | <pre>.   .   P -> R</pre> | subproof implication | 3 - 8 |
| 10 | <pre>.   ( P -> Q ) -> ( P -> R )</pre> | subproof implication | 2 - 9 |
| 11 | <pre>( P -> ( Q -> R ) ) -> ( ( P -> Q ) -> ( P -> R ) )</pre> | subproof implication | 1 - 10 |

## ( ( P -> Q ) -> ( P -> R ) ) -> ( P -> ( Q -> R ) ) <a name="proof-6-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) -> ( P -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   ( P -> Q ) -> ( P -> R )</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 6 | <pre>.   .   .   P -> Q</pre> | I -> | 5, 3 |
| 7 | <pre>.   .   .   P -> R</pre> | E -> | 4, 6 |
| 8 | <pre>.   .   .   R</pre> | E -> | 7, 5 |
| 9 | <pre>.   .   Q -> R</pre> | subproof implication | 3 - 8 |
| 10 | <pre>.   P -> ( Q -> R )</pre> | subproof implication | 2 - 9 |
| 11 | <pre>( ( P -> Q ) -> ( P -> R ) ) -> ( P -> ( Q -> R ) )</pre> | subproof implication | 1 - 10 |

## ( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) ) <a name="proof-6-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q v R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P -> ( Q v R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   Q v R</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 6 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 7 | <pre>.   .   .   P -> Q</pre> | I -> | 6, 5 |
| 8 | <pre>.   .   .   ( P -> Q ) v ( P -> R )</pre> | I v (L) | 7 |
| 9 | <pre>.   .   Q -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 5 - 8 |
| 10 | <pre>.   .   .   R</pre> | Assume Implication |  |
| 11 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 12 | <pre>.   .   .   P -> R</pre> | I -> | 11, 10 |
| 13 | <pre>.   .   .   ( P -> Q ) v ( P -> R )</pre> | I v (R) | 12 |
| 14 | <pre>.   .   R -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 10 - 13 |
| 15 | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | E v | 9, 14, 4 |
| 16 | <pre>.   P -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 2 - 15 |
| 17 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 18 | <pre>.   .   .   ~ Q</pre> | Assume Implication |  |
| 19 | <pre>.   .   .   ~ P</pre> | Reiterate | 17 |
| 20 | <pre>.   .   ~ Q -> ~ P</pre> | subproof implication | 18 - 19 |
| 21 | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 20 |
| 22 | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | I v (L) | 21 |
| 23 | <pre>.   ~ P -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 17 - 22 |
| 24 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 25 | <pre>.   ( P -> Q ) v ( P -> R )</pre> | E v | 16, 23, 24 |
| 26 | <pre>( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 1 - 25 |

## ( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) ) <a name="proof-6-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) v ( P -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P -> Q</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 5 | <pre>.   .   .   Q</pre> | E -> | 3, 4 |
| 6 | <pre>.   .   .   Q v R</pre> | I v (L) | 5 |
| 7 | <pre>.   .   ( P -> Q ) -> ( Q v R )</pre> | subproof implication | 3 - 6 |
| 8 | <pre>.   .   .   P -> R</pre> | Assume Implication |  |
| 9 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 10 | <pre>.   .   .   R</pre> | E -> | 8, 9 |
| 11 | <pre>.   .   .   Q v R</pre> | I v (R) | 10 |
| 12 | <pre>.   .   ( P -> R ) -> ( Q v R )</pre> | subproof implication | 8 - 11 |
| 13 | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | Reiterate | 1 |
| 14 | <pre>.   .   Q v R</pre> | E v | 7, 12, 13 |
| 15 | <pre>.   P -> ( Q v R )</pre> | subproof implication | 2 - 14 |
| 16 | <pre>( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )</pre> | subproof implication | 1 - 15 |

## ( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) ) <a name="proof-6-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q ^ R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   Q ^ R</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   Q</pre> | E ^ (L) | 4 |
| 6 | <pre>.   P -> Q</pre> | subproof implication | 2 - 5 |
| 7 | <pre>.   .   P</pre> | Assume Implication |  |
| 8 | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 9 | <pre>.   .   Q ^ R</pre> | E -> | 8, 7 |
| 10 | <pre>.   .   R</pre> | E ^ (R) | 9 |
| 11 | <pre>.   P -> R</pre> | subproof implication | 7 - 10 |
| 12 | <pre>.   ( P -> Q ) ^ ( P -> R )</pre> | I ^ | 6, 11 |
| 13 | <pre>( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )</pre> | subproof implication | 1 - 12 |

## ( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) ) <a name="proof-6-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( P -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   ( P -> Q ) ^ ( P -> R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P -> Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   P -> R</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E -> | 4, 2 |
| 7 | <pre>.   .   R</pre> | E -> | 5, 2 |
| 8 | <pre>.   .   Q ^ R</pre> | I ^ | 6, 7 |
| 9 | <pre>.   P -> ( Q ^ R )</pre> | subproof implication | 2 - 8 |
| 10 | <pre>( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )</pre> | subproof implication | 1 - 9 |

## ( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) ) <a name="proof-6-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ( Q v R )</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q v R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   Q</pre> | Assume Implication |  |
| 5 | <pre>.   .   P</pre> | Reiterate | 2 |
| 6 | <pre>.   .   P ^ Q</pre> | I ^ | 5, 4 |
| 7 | <pre>.   .   ( P ^ Q ) v ( P ^ R )</pre> | I v (L) | 6 |
| 8 | <pre>.   Q -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof implication | 4 - 7 |
| 9 | <pre>.   .   R</pre> | Assume Implication |  |
| 10 | <pre>.   .   P</pre> | Reiterate | 2 |
| 11 | <pre>.   .   P ^ R</pre> | I ^ | 10, 9 |
| 12 | <pre>.   .   ( P ^ Q ) v ( P ^ R )</pre> | I v (R) | 11 |
| 13 | <pre>.   R -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof implication | 9 - 12 |
| 14 | <pre>.   ( P ^ Q ) v ( P ^ R )</pre> | E v | 8, 13, 3 |
| 15 | <pre>( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof implication | 1 - 14 |

## ( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) ) <a name="proof-6-8"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P ^ Q ) v ( P ^ R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P ^ Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 4 | <pre>.   .   Q</pre> | E ^ (R) | 2 |
| 5 | <pre>.   .   Q v R</pre> | I v (L) | 4 |
| 6 | <pre>.   .   P ^ ( Q v R )</pre> | I ^ | 3, 5 |
| 7 | <pre>.   ( P ^ Q ) -> ( P ^ ( Q v R ) )</pre> | subproof implication | 2 - 6 |
| 8 | <pre>.   .   P ^ R</pre> | Assume Implication |  |
| 9 | <pre>.   .   P</pre> | E ^ (L) | 8 |
| 10 | <pre>.   .   R</pre> | E ^ (R) | 8 |
| 11 | <pre>.   .   Q v R</pre> | I v (R) | 10 |
| 12 | <pre>.   .   P ^ ( Q v R )</pre> | I ^ | 9, 11 |
| 13 | <pre>.   ( P ^ R ) -> ( P ^ ( Q v R ) )</pre> | subproof implication | 8 - 12 |
| 14 | <pre>.   P ^ ( Q v R )</pre> | E v | 7, 13, 1 |
| 15 | <pre>( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )</pre> | subproof implication | 1 - 14 |

## ( P v ( Q ^ R ) ) -> ( ( P v Q ) ^ ( P v R ) ) <a name="proof-6-9"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v ( Q ^ R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   .   P v R</pre> | I v (L) | 2 |
| 5 | <pre>.   .   ( P v Q ) ^ ( P v R )</pre> | I ^ | 3, 4 |
| 6 | <pre>.   P -> ( ( P v Q ) ^ ( P v R ) )</pre> | subproof implication | 2 - 5 |
| 7 | <pre>.   .   Q ^ R</pre> | Assume Implication |  |
| 8 | <pre>.   .   Q</pre> | E ^ (L) | 7 |
| 9 | <pre>.   .   R</pre> | E ^ (R) | 7 |
| 10 | <pre>.   .   P v Q</pre> | I v (R) | 8 |
| 11 | <pre>.   .   P v R</pre> | I v (R) | 9 |
| 12 | <pre>.   .   ( P v Q ) ^ ( P v R )</pre> | I ^ | 10, 11 |
| 13 | <pre>.   ( Q ^ R ) -> ( ( P v Q ) ^ ( P v R ) )</pre> | subproof implication | 7 - 12 |
| 14 | <pre>.   ( P v Q ) ^ ( P v R )</pre> | E v | 6, 13, 1 |
| 15 | <pre>( P v ( Q ^ R ) ) -> ( ( P v Q ) ^ ( P v R ) )</pre> | subproof implication | 1 - 14 |

## ( ( P v Q ) ^ ( P v R ) ) -> ( P v ( Q ^ R ) ) <a name="proof-6-10"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P v Q ) ^ ( P v R )</pre> | Assume Implication |  |
| 2 | <pre>.   P v Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   P v R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   P</pre> | Assume Implication |  |
| 5 | <pre>.   .   P v ( Q ^ R )</pre> | I v (L) | 4 |
| 6 | <pre>.   P -> ( P v ( Q ^ R ) )</pre> | subproof implication | 4 - 5 |
| 7 | <pre>.   .   P</pre> | Assume Implication |  |
| 8 | <pre>.   .   .   ~ Q</pre> | Assume Implication |  |
| 9 | <pre>.   .   .   P</pre> | Reiterate | 7 |
| 10 | <pre>.   .   ~ Q -> P</pre> | subproof implication | 8 - 9 |
| 11 | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 10 |
| 12 | <pre>.   P -> ( ~ P -> Q )</pre> | subproof implication | 7 - 11 |
| 13 | <pre>.   .   P</pre> | Assume Implication |  |
| 14 | <pre>.   .   .   ~ R</pre> | Assume Implication |  |
| 15 | <pre>.   .   .   P</pre> | Reiterate | 13 |
| 16 | <pre>.   .   ~ R -> P</pre> | subproof implication | 14 - 15 |
| 17 | <pre>.   .   ~ P -> R</pre> | Theorem: contrapositive | 16 |
| 18 | <pre>.   P -> ( ~ P -> R )</pre> | subproof implication | 13 - 17 |
| 19 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 20 | <pre>.   .   P v Q</pre> | Reiterate | 2 |
| 21 | <pre>.   .   P v R</pre> | Reiterate | 3 |
| 22 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 23 | <pre>.   .   .   ~ P</pre> | Reiterate | 19 |
| 24 | <pre>.   .   .   ~ P -> Q</pre> | I -> | 23, 22 |
| 25 | <pre>.   .   Q -> ( ~ P -> Q )</pre> | subproof implication | 22 - 24 |
| 26 | <pre>.   .   P -> ( ~ P -> Q )</pre> | Reiterate | 12 |
| 27 | <pre>.   .   ~ P -> Q</pre> | E v | 26, 25, 20 |
| 28 | <pre>.   .   Q</pre> | E -> | 27, 19 |
| 29 | <pre>.   .   .   R</pre> | Assume Implication |  |
| 30 | <pre>.   .   .   ~ P</pre> | Reiterate | 19 |
| 31 | <pre>.   .   .   ~ P -> R</pre> | I -> | 30, 29 |
| 32 | <pre>.   .   R -> ( ~ P -> R )</pre> | subproof implication | 29 - 31 |
| 33 | <pre>.   .   P -> ( ~ P -> R )</pre> | Reiterate | 18 |
| 34 | <pre>.   .   ~ P -> R</pre> | E v | 33, 32, 21 |
| 35 | <pre>.   .   R</pre> | E -> | 34, 19 |
| 36 | <pre>.   .   Q ^ R</pre> | I ^ | 28, 35 |
| 37 | <pre>.   .   P v ( Q ^ R )</pre> | I v (R) | 36 |
| 38 | <pre>.   ~ P -> ( P v ( Q ^ R ) )</pre> | subproof implication | 19 - 37 |
| 39 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 40 | <pre>.   P v ( Q ^ R )</pre> | E v | 6, 38, 39 |
| 41 | <pre>( ( P v Q ) ^ ( P v R ) ) -> ( P v ( Q ^ R ) )</pre> | subproof implication | 1 - 40 |

# disjunction <a name="disjunction"></a>

## ( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S ) <a name="proof-7-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q )</pre> | Assume Implication |  |
| 2 | <pre>.   ( P -> R ) ^ ( Q -> S )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   P v Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P -> R</pre> | E ^ (L) | 2 |
| 5 | <pre>.   Q -> S</pre> | E ^ (R) | 2 |
| 6 | <pre>.   .   P</pre> | Assume Implication |  |
| 7 | <pre>.   .   P -> R</pre> | Reiterate | 4 |
| 8 | <pre>.   .   R</pre> | E -> | 7, 6 |
| 9 | <pre>.   .   R v S</pre> | I v (L) | 8 |
| 10 | <pre>.   P -> ( R v S )</pre> | subproof implication | 6 - 9 |
| 11 | <pre>.   .   Q</pre> | Assume Implication |  |
| 12 | <pre>.   .   Q -> S</pre> | Reiterate | 5 |
| 13 | <pre>.   .   S</pre> | E -> | 12, 11 |
| 14 | <pre>.   .   R v S</pre> | I v (R) | 13 |
| 15 | <pre>.   Q -> ( R v S )</pre> | subproof implication | 11 - 14 |
| 16 | <pre>.   R v S</pre> | E v | 10, 15, 3 |
| 17 | <pre>( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( P v Q ) ) -> ( R v S )</pre> | subproof implication | 1 - 16 |

## ( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S ) <a name="proof-7-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q )</pre> | Assume Implication |  |
| 2 | <pre>.   ( P -> R ) ^ ( Q -> S )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ R v Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P -> R</pre> | E ^ (L) | 2 |
| 5 | <pre>.   Q -> S</pre> | E ^ (R) | 2 |
| 6 | <pre>.   .   ~ R</pre> | Assume Implication |  |
| 7 | <pre>.   .   P -> R</pre> | Reiterate | 4 |
| 8 | <pre>.   .   ~ R -> ~ P</pre> | Theorem: contrapositive | 7 |
| 9 | <pre>.   .   ~ P</pre> | E -> | 8, 6 |
| 10 | <pre>.   .   ~ P v S</pre> | I v (L) | 9 |
| 11 | <pre>.   ~ R -> ( ~ P v S )</pre> | subproof implication | 6 - 10 |
| 12 | <pre>.   .   Q</pre> | Assume Implication |  |
| 13 | <pre>.   .   Q -> S</pre> | Reiterate | 5 |
| 14 | <pre>.   .   S</pre> | E -> | 13, 12 |
| 15 | <pre>.   .   ~ P v S</pre> | I v (R) | 14 |
| 16 | <pre>.   Q -> ( ~ P v S )</pre> | subproof implication | 12 - 15 |
| 17 | <pre>.   ~ P v S</pre> | E v | 11, 16, 3 |
| 18 | <pre>( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v Q ) ) -> ( ~ P v S )</pre> | subproof implication | 1 - 17 |

## ( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q ) <a name="proof-7-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S )</pre> | Assume Implication |  |
| 2 | <pre>.   ( P -> R ) ^ ( Q -> S )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ R v ~ S</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P -> R</pre> | E ^ (L) | 2 |
| 5 | <pre>.   Q -> S</pre> | E ^ (R) | 2 |
| 6 | <pre>.   .   ~ R</pre> | Assume Implication |  |
| 7 | <pre>.   .   P -> R</pre> | Reiterate | 4 |
| 8 | <pre>.   .   ~ R -> ~ P</pre> | Theorem: contrapositive | 7 |
| 9 | <pre>.   .   ~ P</pre> | E -> | 8, 6 |
| 10 | <pre>.   .   ~ P v ~ Q</pre> | I v (L) | 9 |
| 11 | <pre>.   ~ R -> ( ~ P v ~ Q )</pre> | subproof implication | 6 - 10 |
| 12 | <pre>.   .   ~ S</pre> | Assume Implication |  |
| 13 | <pre>.   .   Q -> S</pre> | Reiterate | 5 |
| 14 | <pre>.   .   ~ S -> ~ Q</pre> | Theorem: contrapositive | 13 |
| 15 | <pre>.   .   ~ Q</pre> | E -> | 14, 12 |
| 16 | <pre>.   .   ~ P v ~ Q</pre> | I v (R) | 15 |
| 17 | <pre>.   ~ S -> ( ~ P v ~ Q )</pre> | subproof implication | 12 - 16 |
| 18 | <pre>.   ~ P v ~ Q</pre> | E v | 11, 17, 3 |
| 19 | <pre>( ( ( P -> R ) ^ ( Q -> S ) ) ^ ( ~ R v ~ S ) ) -> ( ~ P v ~ Q )</pre> | subproof implication | 1 - 18 |

## ( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R <a name="proof-7-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q )</pre> | Assume Implication |  |
| 2 | <pre>.   ( P -> R ) v ( Q -> R )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   P ^ Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P</pre> | E ^ (L) | 3 |
| 5 | <pre>.   Q</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   P -> R</pre> | Assume Implication |  |
| 7 | <pre>.   .   P</pre> | Reiterate | 4 |
| 8 | <pre>.   .   R</pre> | E -> | 6, 7 |
| 9 | <pre>.   ( P -> R ) -> R</pre> | subproof implication | 6 - 8 |
| 10 | <pre>.   .   Q -> R</pre> | Assume Implication |  |
| 11 | <pre>.   .   Q</pre> | Reiterate | 5 |
| 12 | <pre>.   .   R</pre> | E -> | 10, 11 |
| 13 | <pre>.   ( Q -> R ) -> R</pre> | subproof implication | 10 - 12 |
| 14 | <pre>.   R</pre> | E v | 9, 13, 2 |
| 15 | <pre>( ( ( P -> R ) v ( Q -> R ) ) ^ ( P ^ Q ) ) -> R</pre> | subproof implication | 1 - 14 |

## ( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q ) <a name="proof-7-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( ( P -> R ) v ( Q -> R ) ) ^ ~ R</pre> | Assume Implication |  |
| 2 | <pre>.   ( P -> R ) v ( Q -> R )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   P -> R</pre> | Assume Implication |  |
| 5 | <pre>.   .   ~ R</pre> | Reiterate | 3 |
| 6 | <pre>.   .   ~ R -> ~ P</pre> | Theorem: contrapositive | 4 |
| 7 | <pre>.   .   ~ P</pre> | E -> | 6, 5 |
| 8 | <pre>.   .   ~ P v ~ Q</pre> | I v (L) | 7 |
| 9 | <pre>.   ( P -> R ) -> ( ~ P v ~ Q )</pre> | subproof implication | 4 - 8 |
| 10 | <pre>.   .   Q -> R</pre> | Assume Implication |  |
| 11 | <pre>.   .   ~ R</pre> | Reiterate | 3 |
| 12 | <pre>.   .   ~ R -> ~ Q</pre> | Theorem: contrapositive | 10 |
| 13 | <pre>.   .   ~ Q</pre> | E -> | 12, 11 |
| 14 | <pre>.   .   ~ P v ~ Q</pre> | I v (R) | 13 |
| 15 | <pre>.   ( Q -> R ) -> ( ~ P v ~ Q )</pre> | subproof implication | 10 - 14 |
| 16 | <pre>.   ~ P v ~ Q</pre> | E v | 9, 15, 2 |
| 17 | <pre>( ( ( P -> R ) v ( Q -> R ) ) ^ ~ R ) -> ( ~ P v ~ Q )</pre> | subproof implication | 1 - 16 |

# biconditional <a name="biconditional"></a>

## ( ( P <-> Q ) ^ P ) -> Q <a name="proof-8-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P <-> Q ) ^ P</pre> | Assume Implication |  |
| 2 | <pre>.   P <-> Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   P</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P -> Q</pre> | E <-> (L) | 2 |
| 5 | <pre>.   Q</pre> | E -> | 4, 3 |
| 6 | <pre>( ( P <-> Q ) ^ P ) -> Q</pre> | subproof implication | 1 - 5 |

## ( ( P <-> Q ) ^ ~ P ) -> ~ Q <a name="proof-8-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P <-> Q ) ^ ~ P</pre> | Assume Implication |  |
| 2 | <pre>.   P <-> Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q -> P</pre> | E <-> (R) | 2 |
| 5 | <pre>.   .   Q</pre> | Assume Contradiction |  |
| 6 | <pre>.   .   ~ P</pre> | Reiterate | 3 |
| 7 | <pre>.   .   Q -> P</pre> | Reiterate | 4 |
| 8 | <pre>.   .   P</pre> | E -> | 7, 5 |
| 9 | <pre>.   ~ Q</pre> | subproof contradiction | 5 - 8 |
| 10 | <pre>( ( P <-> Q ) ^ ~ P ) -> ~ Q</pre> | subproof implication | 1 - 9 |

## ( P <-> ~ Q ) -> ( ~ P <-> Q ) <a name="proof-8-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   Q -> ~ P</pre> | Theorem: contrapositive | 2 |
| 4 | <pre>.   ~ Q -> P</pre> | E <-> (R) | 1 |
| 5 | <pre>.   ~ P -> Q</pre> | Theorem: contrapositive | 4 |
| 6 | <pre>.   ~ P <-> Q</pre> | I <-> | 5, 3 |
| 7 | <pre>( P <-> ~ Q ) -> ( ~ P <-> Q )</pre> | subproof implication | 1 - 6 |

## ( P <-> Q ) -> ( ~ P <-> ~ Q ) <a name="proof-8-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> Q</pre> | Assume Implication |  |
| 2 | <pre>.   P -> Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   ~ Q -> ~ P</pre> | Theorem: contrapositive | 2 |
| 4 | <pre>.   Q -> P</pre> | E <-> (R) | 1 |
| 5 | <pre>.   ~ P -> ~ Q</pre> | Theorem: contrapositive | 4 |
| 6 | <pre>.   ~ P <-> ~ Q</pre> | I <-> | 5, 3 |
| 7 | <pre>( P <-> Q ) -> ( ~ P <-> ~ Q )</pre> | subproof implication | 1 - 6 |

## ( ~ P <-> ~ Q ) -> ( P <-> Q ) <a name="proof-8-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ P <-> ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   ~ P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   Q -> P</pre> | Theorem: contrapositive | 2 |
| 4 | <pre>.   ~ Q -> ~ P</pre> | E <-> (R) | 1 |
| 5 | <pre>.   P -> Q</pre> | Theorem: contrapositive | 4 |
| 6 | <pre>.   P <-> Q</pre> | I <-> | 5, 3 |
| 7 | <pre>( ~ P <-> ~ Q ) -> ( P <-> Q )</pre> | subproof implication | 1 - 6 |

## ( P <-> Q ) v ~ ( P <-> Q ) <a name="proof-8-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | Assume Contradiction |  |
| 2 | <pre>.   .   P <-> Q</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   ( P <-> Q ) v ~ ( P <-> Q )</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | Reiterate | 1 |
| 5 | <pre>.   ~ ( P <-> Q )</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   ( P <-> Q ) v ~ ( P <-> Q )</pre> | I v (R) | 5 |
| 7 | <pre>( P <-> Q ) v ~ ( P <-> Q )</pre> | subproof contradiction | 1 - 6 |

## ( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) ) <a name="proof-8-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   ~ Q -> P</pre> | E <-> (R) | 1 |
| 4 | <pre>.   ~ P -> Q</pre> | Theorem: contrapositive | 3 |
| 5 | <pre>.   .   P</pre> | Assume Implication |  |
| 6 | <pre>.   .   P v Q</pre> | I v (L) | 5 |
| 7 | <pre>.   .   P -> ~ Q</pre> | Reiterate | 2 |
| 8 | <pre>.   .   ~ Q</pre> | E -> | 7, 5 |
| 9 | <pre>.   .   ~ P v ~ Q</pre> | I v (R) | 8 |
| 10 | <pre>.   .   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | I ^ | 6, 9 |
| 11 | <pre>.   P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | subproof implication | 5 - 10 |
| 12 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 13 | <pre>.   .   ~ P v ~ Q</pre> | I v (L) | 12 |
| 14 | <pre>.   .   ~ P -> Q</pre> | Reiterate | 4 |
| 15 | <pre>.   .   Q</pre> | E -> | 14, 12 |
| 16 | <pre>.   .   P v Q</pre> | I v (R) | 15 |
| 17 | <pre>.   .   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | I ^ | 16, 13 |
| 18 | <pre>.   ~ P -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | subproof implication | 12 - 17 |
| 19 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 20 | <pre>.   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | E v | 11, 18, 19 |
| 21 | <pre>( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )</pre> | subproof implication | 1 - 20 |

## ( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q ) <a name="proof-8-8"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P v Q ) ^ ( ~ P v ~ Q )</pre> | Assume Implication |  |
| 2 | <pre>.   P v Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P v ~ Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 5 | <pre>.   .   .   P ^ Q</pre> | Assume Contradiction |  |
| 6 | <pre>.   .   .   P</pre> | E ^ (L) | 5 |
| 7 | <pre>.   .   .   ~ P</pre> | Reiterate | 4 |
| 8 | <pre>.   .   ~ ( P ^ Q )</pre> | subproof contradiction | 5 - 7 |
| 9 | <pre>.   ~ P -> ~ ( P ^ Q )</pre> | subproof implication | 4 - 8 |
| 10 | <pre>.   .   ~ Q</pre> | Assume Implication |  |
| 11 | <pre>.   .   .   P ^ Q</pre> | Assume Contradiction |  |
| 12 | <pre>.   .   .   Q</pre> | E ^ (R) | 11 |
| 13 | <pre>.   .   .   ~ Q</pre> | Reiterate | 10 |
| 14 | <pre>.   .   ~ ( P ^ Q )</pre> | subproof contradiction | 11 - 13 |
| 15 | <pre>.   ~ Q -> ~ ( P ^ Q )</pre> | subproof implication | 10 - 14 |
| 16 | <pre>.   ~ ( P ^ Q )</pre> | E v | 9, 15, 3 |
| 17 | <pre>.   .   P</pre> | Assume Implication |  |
| 18 | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume Contradiction |  |
| 19 | <pre>.   .   .   ~ P</pre> | E ^ (L) | 18 |
| 20 | <pre>.   .   .   P</pre> | Reiterate | 17 |
| 21 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof contradiction | 18 - 20 |
| 22 | <pre>.   P -> ~ ( ~ P ^ ~ Q )</pre> | subproof implication | 17 - 21 |
| 23 | <pre>.   .   Q</pre> | Assume Implication |  |
| 24 | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume Contradiction |  |
| 25 | <pre>.   .   .   ~ Q</pre> | E ^ (R) | 24 |
| 26 | <pre>.   .   .   Q</pre> | Reiterate | 23 |
| 27 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof contradiction | 24 - 26 |
| 28 | <pre>.   Q -> ~ ( ~ P ^ ~ Q )</pre> | subproof implication | 23 - 27 |
| 29 | <pre>.   ~ ( ~ P ^ ~ Q )</pre> | E v | 22, 28, 2 |
| 30 | <pre>.   .   P</pre> | Assume Implication |  |
| 31 | <pre>.   .   .   Q</pre> | Assume Contradiction |  |
| 32 | <pre>.   .   .   P</pre> | Reiterate | 30 |
| 33 | <pre>.   .   .   P ^ Q</pre> | I ^ | 32, 31 |
| 34 | <pre>.   .   .   ~ ( P ^ Q )</pre> | Reiterate | 16 |
| 35 | <pre>.   .   ~ Q</pre> | subproof contradiction | 31 - 34 |
| 36 | <pre>.   P -> ~ Q</pre> | subproof implication | 30 - 35 |
| 37 | <pre>.   .   ~ Q</pre> | Assume Implication |  |
| 38 | <pre>.   .   .   ~ P</pre> | Assume Contradiction |  |
| 39 | <pre>.   .   .   ~ Q</pre> | Reiterate | 37 |
| 40 | <pre>.   .   .   ~ P ^ ~ Q</pre> | I ^ | 38, 39 |
| 41 | <pre>.   .   .   ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 29 |
| 42 | <pre>.   .   P</pre> | subproof contradiction | 38 - 41 |
| 43 | <pre>.   ~ Q -> P</pre> | subproof implication | 37 - 42 |
| 44 | <pre>.   P <-> ~ Q</pre> | I <-> | 36, 43 |
| 45 | <pre>( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )</pre> | subproof implication | 1 - 44 |

# miscellaneous <a name="miscellaneous"></a>

## ( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) ) <a name="proof-9-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) -> R</pre> | Assume Implication |  |
| 2 | <pre>.   .   P -> Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   ( P -> Q ) -> R</pre> | Reiterate | 1 |
| 4 | <pre>.   .   R</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   .   P</pre> | Assume Implication |  |
| 6 | <pre>.   .   .   R</pre> | Reiterate | 4 |
| 7 | <pre>.   .   P -> R</pre> | subproof implication | 5 - 6 |
| 8 | <pre>.   ( P -> Q ) -> ( P -> R )</pre> | subproof implication | 2 - 7 |
| 9 | <pre>( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )</pre> | subproof implication | 1 - 8 |

## ( P -> Q ) v ( Q -> R ) <a name="proof-9-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( ( P -> Q ) v ( Q -> R ) )</pre> | Assume Contradiction |  |
| 2 | <pre>.   ~ ( P -> Q ) ^ ~ ( Q -> R )</pre> | DeMorgan's theorem: (~ v) to ^ | 1 |
| 3 | <pre>.   ~ ( P -> Q )</pre> | E ^ (L) | 2 |
| 4 | <pre>.   ~ ( Q -> R )</pre> | E ^ (R) | 2 |
| 5 | <pre>.   P ^ ~ Q</pre> | Theorem: (~ ->) to ^ | 3 |
| 6 | <pre>.   ~ Q</pre> | E ^ (R) | 5 |
| 7 | <pre>.   Q ^ ~ R</pre> | Theorem: (~ ->) to ^ | 4 |
| 8 | <pre>.   Q</pre> | E ^ (L) | 7 |
| 9 | <pre>( P -> Q ) v ( Q -> R )</pre> | subproof contradiction | 1 - 8 |

## ( P -> Q ) v ( Q -> P ) <a name="proof-9-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( ( P -> Q ) v ( Q -> P ) )</pre> | Assume Contradiction |  |
| 2 | <pre>.   ~ ( P -> Q ) ^ ~ ( Q -> P )</pre> | DeMorgan's theorem: (~ v) to ^ | 1 |
| 3 | <pre>.   ~ ( P -> Q )</pre> | E ^ (L) | 2 |
| 4 | <pre>.   ~ ( Q -> P )</pre> | E ^ (R) | 2 |
| 5 | <pre>.   P ^ ~ Q</pre> | Theorem: (~ ->) to ^ | 3 |
| 6 | <pre>.   ~ Q</pre> | E ^ (R) | 5 |
| 7 | <pre>.   Q ^ ~ P</pre> | Theorem: (~ ->) to ^ | 4 |
| 8 | <pre>.   Q</pre> | E ^ (L) | 7 |
| 9 | <pre>( P -> Q ) v ( Q -> P )</pre> | subproof contradiction | 1 - 8 |

