
# Table of Contents

1. [theorems](#theorems)
    1. [P v ~ P](#proof-1-1)
    2. [( P -> Q ) -> ( ~ Q -> ~ P )](#proof-1-2)
    3. [( ~ Q -> ~ P ) -> ( P -> Q )](#proof-1-3)
    4. [~ ( P -> Q ) -> ( P ^ ~ Q )](#proof-1-4)
    5. [( P ^ ~ Q ) -> ~ ( P -> Q )](#proof-1-5)
    6. [( P v Q ) -> ( ~ P -> Q )](#proof-1-6)
    7. [( ~ P -> Q ) -> ( P v Q )](#proof-1-7)
2. [basics](#basics)
    1. [P -> P](#proof-2-1)
    2. [~ ( P ^ ~ P )](#proof-2-2)
    3. [( P ^ P ) -> P](#proof-2-3)
    4. [P -> ( P ^ P )](#proof-2-4)
    5. [( P v P ) -> P](#proof-2-5)
    6. [P -> ( P v P )](#proof-2-6)
    7. [( P ^ Q ) -> ( P v Q )](#proof-2-7)
    8. [( P -> ~ P ) -> ~ P](#proof-2-8)
3. [arrows](#arrows)
    1. [Q -> ( P -> Q )](#proof-3-1)
    2. [~ P -> ( P -> Q )](#proof-3-2)
    3. [( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )](#proof-3-3)
    4. [( ( P ^ Q ) -> R ) -> ( P -> ( Q -> R ) )](#proof-3-4)
    5. [( P -> ( Q -> R ) ) -> ( ( P ^ Q ) -> R )](#proof-3-5)
    6. [( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P](#proof-3-6)
    7. [( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q](#proof-3-7)
    8. [( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) )](#proof-3-8)
    9. [( P -> Q ) -> ( P -> ( Q v R ) )](#proof-3-9)
    10. [( ( P v Q ) -> R ) -> ( P -> R )](#proof-3-10)
    11. [( P -> R ) -> ( ( P ^ Q ) -> R )](#proof-3-11)
    12. [( P -> ( Q ^ R ) ) -> ( P -> Q )](#proof-3-12)
4. [commutativity](#commutativity)
    1. [( P ^ Q ) -> ( Q ^ P )](#proof-4-1)
    2. [( P v Q ) -> ( Q v P )](#proof-4-2)
    3. [( P <-> Q ) -> ( Q <-> P )](#proof-4-3)
5. [associativity](#associativity)
    1. [( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )](#proof-5-1)
    2. [( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )](#proof-5-2)
    3. [( ( P v Q ) v R ) -> ( P v ( Q v R ) )](#proof-5-3)
    4. [( P v ( Q v R ) ) -> ( ( P v Q ) v R )](#proof-5-4)
    5. [( ( P <-> Q ) <-> R ) -> ( P <-> ( Q <-> R ) )](#proof-5-5)
    6. [( P <-> ( Q <-> R ) ) -> ( ( P <-> Q ) <-> R )](#proof-5-6)
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
    11. [~ ( P v Q ) -> ( ~ P ^ ~ Q )](#proof-6-11)
    12. [( ~ P ^ ~ Q ) -> ~ ( P v Q )](#proof-6-12)
    13. [~ ( P ^ Q ) -> ( ~ P v ~ Q )](#proof-6-13)
    14. [( ~ P v ~ Q ) -> ~ ( P ^ Q )](#proof-6-14)
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
    4. [( P <-> ~ Q ) -> ~ ( P <-> Q )](#proof-8-4)
    5. [~ ( P <-> Q ) -> ( P <-> ~ Q )](#proof-8-5)
    6. [( P <-> Q ) -> ( ~ P <-> ~ Q )](#proof-8-6)
    7. [( ~ P <-> ~ Q ) -> ( P <-> Q )](#proof-8-7)
    8. [( P <-> Q ) v ~ ( P <-> Q )](#proof-8-8)
    9. [( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) )](#proof-8-9)
    10. [( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )](#proof-8-10)
9. [miscellaneous](#miscellaneous)
    1. [( ( P -> Q ) -> R ) -> ( ( P -> Q ) -> ( P -> R ) )](#proof-9-1)

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
| 7 | <pre>~ ~ ( P v ~ P )</pre> | subproof contradiction | 1 - 6 |
| 8 | <pre>P v ~ P</pre> | E ~ | 7 |

## ( P -> Q ) -> ( ~ Q -> ~ P ) <a name="proof-1-2"></a>

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

## ( ~ Q -> ~ P ) -> ( P -> Q ) <a name="proof-1-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ Q -> ~ P</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   ~ Q</pre> | Assume Contradiction |  |
| 4 | <pre>.   .   .   ~ Q -> ~ P</pre> | Reiterate | 1 |
| 5 | <pre>.   .   .   ~ P</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 7 | <pre>.   .   ~ ~ Q</pre> | subproof contradiction | 3 - 6 |
| 8 | <pre>.   .   Q</pre> | E ~ | 7 |
| 9 | <pre>.   P -> Q</pre> | subproof implication | 2 - 8 |
| 10 | <pre>( ~ Q -> ~ P ) -> ( P -> Q )</pre> | subproof implication | 1 - 9 |

## ~ ( P -> Q ) -> ( P ^ ~ Q ) <a name="proof-1-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P -> Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   .   ~ Q</pre> | Assume Implication |  |
| 4 | <pre>.   .   .   ~ P</pre> | Reiterate | 2 |
| 5 | <pre>.   .   ~ Q -> ~ P</pre> | subproof implication | 3 - 4 |
| 6 | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 5 |
| 7 | <pre>.   .   ~ ( P -> Q )</pre> | Reiterate | 1 |
| 8 | <pre>.   ~ ~ P</pre> | subproof contradiction | 2 - 7 |
| 9 | <pre>.   P</pre> | E ~ | 8 |
| 10 | <pre>.   .   Q</pre> | Assume Contradiction |  |
| 11 | <pre>.   .   P</pre> | Reiterate | 9 |
| 12 | <pre>.   .   P -> Q</pre> | I -> | 11, 10 |
| 13 | <pre>.   .   ~ ( P -> Q )</pre> | Reiterate | 1 |
| 14 | <pre>.   ~ Q</pre> | subproof contradiction | 10 - 13 |
| 15 | <pre>.   P ^ ~ Q</pre> | I ^ | 9, 14 |
| 16 | <pre>~ ( P -> Q ) -> ( P ^ ~ Q )</pre> | subproof implication | 1 - 15 |

## ( P ^ ~ Q ) -> ~ ( P -> Q ) <a name="proof-1-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   P -> Q</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   P ^ ~ Q</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   ~ Q</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E -> | 2, 4 |
| 7 | <pre>.   ~ ( P -> Q )</pre> | subproof contradiction | 2 - 6 |
| 8 | <pre>( P ^ ~ Q ) -> ~ ( P -> Q )</pre> | subproof implication | 1 - 7 |

## ( P v Q ) -> ( ~ P -> Q ) <a name="proof-1-6"></a>

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

## ( ~ P -> Q ) -> ( P v Q ) <a name="proof-1-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ P -> Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   P -> ( P v Q )</pre> | subproof implication | 2 - 3 |
| 5 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 6 | <pre>.   .   ~ P -> Q</pre> | Reiterate | 1 |
| 7 | <pre>.   .   Q</pre> | E -> | 6, 5 |
| 8 | <pre>.   .   P v Q</pre> | I v (R) | 7 |
| 9 | <pre>.   ~ P -> ( P v Q )</pre> | subproof implication | 5 - 8 |
| 10 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 11 | <pre>.   P v Q</pre> | E v | 4, 9, 10 |
| 12 | <pre>( ~ P -> Q ) -> ( P v Q )</pre> | subproof implication | 1 - 11 |

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

## ( P ^ P ) -> P <a name="proof-2-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ P</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>( P ^ P ) -> P</pre> | subproof implication | 1 - 2 |

## P -> ( P ^ P ) <a name="proof-2-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P</pre> | Assume Implication |  |
| 2 | <pre>.   P ^ P</pre> | I ^ | 1, 1 |
| 3 | <pre>P -> ( P ^ P )</pre> | subproof implication | 1 - 2 |

## ( P v P ) -> P <a name="proof-2-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v P</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   P -> P</pre> | subproof implication | 2 - 2 |
| 4 | <pre>.   P</pre> | E v | 3, 3, 1 |
| 5 | <pre>( P v P ) -> P</pre> | subproof implication | 1 - 4 |

## P -> ( P v P ) <a name="proof-2-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P</pre> | Assume Implication |  |
| 2 | <pre>.   P v P</pre> | I v (L) | 1 |
| 3 | <pre>P -> ( P v P )</pre> | subproof implication | 1 - 2 |

## ( P ^ Q ) -> ( P v Q ) <a name="proof-2-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ Q</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>( P ^ Q ) -> ( P v Q )</pre> | subproof implication | 1 - 3 |

## ( P -> ~ P ) -> ~ P <a name="proof-2-8"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ~ P</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   P -> ~ P</pre> | Reiterate | 1 |
| 4 | <pre>.   .   ~ P</pre> | E -> | 3, 2 |
| 5 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>( P -> ~ P ) -> ~ P</pre> | subproof implication | 1 - 5 |

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

## ( ( P ^ Q ) -> R ) -> ( P -> ( Q -> R ) ) <a name="proof-3-4"></a>

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

## ( P -> ( Q -> R ) ) -> ( ( P ^ Q ) -> R ) <a name="proof-3-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q -> R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P ^ Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 4 | <pre>.   .   Q</pre> | E ^ (R) | 2 |
| 5 | <pre>.   .   P -> ( Q -> R )</pre> | Reiterate | 1 |
| 6 | <pre>.   .   Q -> R</pre> | E -> | 5, 3 |
| 7 | <pre>.   .   R</pre> | E -> | 6, 4 |
| 8 | <pre>.   ( P ^ Q ) -> R</pre> | subproof implication | 2 - 7 |
| 9 | <pre>( P -> ( Q -> R ) ) -> ( ( P ^ Q ) -> R )</pre> | subproof implication | 1 - 8 |

## ( ( P -> Q ) ^ ( P -> ~ Q ) ) -> ~ P <a name="proof-3-6"></a>

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

## ( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q <a name="proof-3-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( ~ P -> Q )</pre> | Assume Implication |  |
| 2 | <pre>.   P -> Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P -> Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 5 | <pre>.   Q</pre> | E v | 2, 3, 4 |
| 6 | <pre>( ( P -> Q ) ^ ( ~ P -> Q ) ) -> Q</pre> | subproof implication | 1 - 5 |

## ( P -> ( Q -> R ) ) -> ( Q -> ( P -> R ) ) <a name="proof-3-8"></a>

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

## ( P -> Q ) -> ( P -> ( Q v R ) ) <a name="proof-3-9"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P -> Q</pre> | Reiterate | 1 |
| 4 | <pre>.   .   Q</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   Q v R</pre> | I v (L) | 4 |
| 6 | <pre>.   P -> ( Q v R )</pre> | subproof implication | 2 - 5 |
| 7 | <pre>( P -> Q ) -> ( P -> ( Q v R ) )</pre> | subproof implication | 1 - 6 |

## ( ( P v Q ) -> R ) -> ( P -> R ) <a name="proof-3-10"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P v Q ) -> R</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ( P v Q ) -> R</pre> | Reiterate | 1 |
| 5 | <pre>.   .   R</pre> | E -> | 4, 3 |
| 6 | <pre>.   P -> R</pre> | subproof implication | 2 - 5 |
| 7 | <pre>( ( P v Q ) -> R ) -> ( P -> R )</pre> | subproof implication | 1 - 6 |

## ( P -> R ) -> ( ( P ^ Q ) -> R ) <a name="proof-3-11"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> R</pre> | Assume Implication |  |
| 2 | <pre>.   .   P ^ Q</pre> | Assume Implication |  |
| 3 | <pre>.   .   P -> R</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 5 | <pre>.   .   R</pre> | E -> | 3, 4 |
| 6 | <pre>.   ( P ^ Q ) -> R</pre> | subproof implication | 2 - 5 |
| 7 | <pre>( P -> R ) -> ( ( P ^ Q ) -> R )</pre> | subproof implication | 1 - 6 |

## ( P -> ( Q ^ R ) ) -> ( P -> Q ) <a name="proof-3-12"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q ^ R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   Q ^ R</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   Q</pre> | E ^ (L) | 4 |
| 6 | <pre>.   P -> Q</pre> | subproof implication | 2 - 5 |
| 7 | <pre>( P -> ( Q ^ R ) ) -> ( P -> Q )</pre> | subproof implication | 1 - 6 |

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

## ( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) ) <a name="proof-5-1"></a>

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

## ( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R ) <a name="proof-5-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ( Q ^ R )</pre> | Assume Implication |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q ^ R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   R</pre> | E ^ (R) | 3 |
| 6 | <pre>.   P ^ Q</pre> | I ^ | 2, 4 |
| 7 | <pre>.   ( P ^ Q ) ^ R</pre> | I ^ | 6, 5 |
| 8 | <pre>( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )</pre> | subproof implication | 1 - 7 |

## ( ( P v Q ) v R ) -> ( P v ( Q v R ) ) <a name="proof-5-3"></a>

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

## ( P v ( Q v R ) ) -> ( ( P v Q ) v R ) <a name="proof-5-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v ( Q v R )</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ( P v Q ) v R</pre> | I v (L) | 3 |
| 5 | <pre>.   P -> ( ( P v Q ) v R )</pre> | subproof implication | 2 - 4 |
| 6 | <pre>.   .   Q v R</pre> | Assume Implication |  |
| 7 | <pre>.   .   .   Q</pre> | Assume Implication |  |
| 8 | <pre>.   .   .   P v Q</pre> | I v (R) | 7 |
| 9 | <pre>.   .   .   ( P v Q ) v R</pre> | I v (L) | 8 |
| 10 | <pre>.   .   Q -> ( ( P v Q ) v R )</pre> | subproof implication | 7 - 9 |
| 11 | <pre>.   .   .   R</pre> | Assume Implication |  |
| 12 | <pre>.   .   .   ( P v Q ) v R</pre> | I v (R) | 11 |
| 13 | <pre>.   .   R -> ( ( P v Q ) v R )</pre> | subproof implication | 11 - 12 |
| 14 | <pre>.   .   ( P v Q ) v R</pre> | E v | 10, 13, 6 |
| 15 | <pre>.   ( Q v R ) -> ( ( P v Q ) v R )</pre> | subproof implication | 6 - 14 |
| 16 | <pre>.   ( P v Q ) v R</pre> | E v | 5, 15, 1 |
| 17 | <pre>( P v ( Q v R ) ) -> ( ( P v Q ) v R )</pre> | subproof implication | 1 - 16 |

## ( ( P <-> Q ) <-> R ) -> ( P <-> ( Q <-> R ) ) <a name="proof-5-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P <-> Q ) <-> R</pre> | Assume Implication |  |
| 2 | <pre>.   ( P <-> Q ) -> R</pre> | E <-> (L) | 1 |
| 3 | <pre>.   R -> ( P <-> Q )</pre> | E <-> (R) | 1 |
| 4 | <pre>.   .   ~ ( P -> ( Q <-> R ) )</pre> | Assume Contradiction |  |
| 5 | <pre>.   .   P ^ ~ ( Q <-> R )</pre> | Theorem: arrow negation | 4 |
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
| 27 | <pre>.   ~ ~ ( P -> ( Q <-> R ) )</pre> | subproof contradiction | 4 - 26 |
| 28 | <pre>.   P -> ( Q <-> R )</pre> | E ~ | 27 |
| 29 | <pre>.   .   ~ ( ( Q <-> R ) -> P )</pre> | Assume Contradiction |  |
| 30 | <pre>.   .   ( Q <-> R ) ^ ~ P</pre> | Theorem: arrow negation | 29 |
| 31 | <pre>.   .   Q <-> R</pre> | E ^ (L) | 30 |
| 32 | <pre>.   .   ~ P</pre> | E ^ (R) | 30 |
| 33 | <pre>.   .   .   R</pre> | Assume Contradiction |  |
| 34 | <pre>.   .   .   R -> ( P <-> Q )</pre> | Reiterate | 3 |
| 35 | <pre>.   .   .   Q <-> R</pre> | Reiterate | 31 |
| 36 | <pre>.   .   .   R -> Q</pre> | E <-> (R) | 35 |
| 37 | <pre>.   .   .   Q</pre> | E -> | 36, 33 |
| 38 | <pre>.   .   .   P <-> Q</pre> | E -> | 34, 33 |
| 39 | <pre>.   .   .   Q -> P</pre> | E <-> (R) | 38 |
| 40 | <pre>.   .   .   P</pre> | E -> | 39, 37 |
| 41 | <pre>.   .   .   ~ P</pre> | Reiterate | 32 |
| 42 | <pre>.   .   ~ R</pre> | subproof contradiction | 33 - 41 |
| 43 | <pre>.   .   Q -> R</pre> | E <-> (L) | 31 |
| 44 | <pre>.   .   ~ R -> ~ Q</pre> | Theorem: contrapositive | 43 |
| 45 | <pre>.   .   ~ Q</pre> | E -> | 44, 42 |
| 46 | <pre>.   .   ( P <-> Q ) -> R</pre> | Reiterate | 2 |
| 47 | <pre>.   .   ~ R -> ~ ( P <-> Q )</pre> | Theorem: contrapositive | 46 |
| 48 | <pre>.   .   ~ ( P <-> Q )</pre> | E -> | 47, 42 |
| 49 | <pre>.   .   P <-> ~ Q</pre> | Theorem: biconditional negation | 48 |
| 50 | <pre>.   .   ~ Q -> P</pre> | E <-> (R) | 49 |
| 51 | <pre>.   .   P</pre> | E -> | 50, 45 |
| 52 | <pre>.   ~ ~ ( ( Q <-> R ) -> P )</pre> | subproof contradiction | 29 - 51 |
| 53 | <pre>.   ( Q <-> R ) -> P</pre> | E ~ | 52 |
| 54 | <pre>.   P <-> ( Q <-> R )</pre> | I <-> | 28, 53 |
| 55 | <pre>( ( P <-> Q ) <-> R ) -> ( P <-> ( Q <-> R ) )</pre> | subproof implication | 1 - 54 |

## ( P <-> ( Q <-> R ) ) -> ( ( P <-> Q ) <-> R ) <a name="proof-5-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> ( Q <-> R )</pre> | Assume Implication |  |
| 2 | <pre>.   P -> ( Q <-> R )</pre> | E <-> (L) | 1 |
| 3 | <pre>.   ( Q <-> R ) -> P</pre> | E <-> (R) | 1 |
| 4 | <pre>.   .   ~ ( ( P <-> Q ) -> R )</pre> | Assume Contradiction |  |
| 5 | <pre>.   .   ( P <-> Q ) ^ ~ R</pre> | Theorem: arrow negation | 4 |
| 6 | <pre>.   .   P <-> Q</pre> | E ^ (L) | 5 |
| 7 | <pre>.   .   ~ R</pre> | E ^ (R) | 5 |
| 8 | <pre>.   .   P -> Q</pre> | E <-> (L) | 6 |
| 9 | <pre>.   .   .   P</pre> | Assume Contradiction |  |
| 10 | <pre>.   .   .   P -> Q</pre> | Reiterate | 8 |
| 11 | <pre>.   .   .   Q</pre> | E -> | 10, 9 |
| 12 | <pre>.   .   .   P -> ( Q <-> R )</pre> | Reiterate | 2 |
| 13 | <pre>.   .   .   Q <-> R</pre> | E -> | 12, 9 |
| 14 | <pre>.   .   .   Q -> R</pre> | E <-> (L) | 13 |
| 15 | <pre>.   .   .   R</pre> | E -> | 14, 11 |
| 16 | <pre>.   .   .   ~ R</pre> | Reiterate | 7 |
| 17 | <pre>.   .   ~ P</pre> | subproof contradiction | 9 - 16 |
| 18 | <pre>.   .   ( Q <-> R ) -> P</pre> | Reiterate | 3 |
| 19 | <pre>.   .   ~ P -> ~ ( Q <-> R )</pre> | Theorem: contrapositive | 18 |
| 20 | <pre>.   .   ~ ( Q <-> R )</pre> | E -> | 19, 17 |
| 21 | <pre>.   .   Q <-> ~ R</pre> | Theorem: biconditional negation | 20 |
| 22 | <pre>.   .   ~ R -> Q</pre> | E <-> (R) | 21 |
| 23 | <pre>.   .   Q</pre> | E -> | 22, 7 |
| 24 | <pre>.   .   Q -> P</pre> | E <-> (R) | 6 |
| 25 | <pre>.   .   P</pre> | E -> | 24, 23 |
| 26 | <pre>.   ~ ~ ( ( P <-> Q ) -> R )</pre> | subproof contradiction | 4 - 25 |
| 27 | <pre>.   ( P <-> Q ) -> R</pre> | E ~ | 26 |
| 28 | <pre>.   .   ~ ( R -> ( P <-> Q ) )</pre> | Assume Contradiction |  |
| 29 | <pre>.   .   R ^ ~ ( P <-> Q )</pre> | Theorem: arrow negation | 28 |
| 30 | <pre>.   .   R</pre> | E ^ (L) | 29 |
| 31 | <pre>.   .   ~ ( P <-> Q )</pre> | E ^ (R) | 29 |
| 32 | <pre>.   .   P <-> ~ Q</pre> | Theorem: biconditional negation | 31 |
| 33 | <pre>.   .   P -> ~ Q</pre> | E <-> (L) | 32 |
| 34 | <pre>.   .   .   P</pre> | Assume Contradiction |  |
| 35 | <pre>.   .   .   P -> ~ Q</pre> | Reiterate | 33 |
| 36 | <pre>.   .   .   ~ Q</pre> | E -> | 35, 34 |
| 37 | <pre>.   .   .   P -> ( Q <-> R )</pre> | Reiterate | 2 |
| 38 | <pre>.   .   .   Q <-> R</pre> | E -> | 37, 34 |
| 39 | <pre>.   .   .   R -> Q</pre> | E <-> (R) | 38 |
| 40 | <pre>.   .   .   R</pre> | Reiterate | 30 |
| 41 | <pre>.   .   .   Q</pre> | E -> | 39, 40 |
| 42 | <pre>.   .   ~ P</pre> | subproof contradiction | 34 - 41 |
| 43 | <pre>.   .   ~ Q -> P</pre> | E <-> (R) | 32 |
| 44 | <pre>.   .   ~ P -> Q</pre> | Theorem: contrapositive | 43 |
| 45 | <pre>.   .   Q</pre> | E -> | 44, 42 |
| 46 | <pre>.   .   ( Q <-> R ) -> P</pre> | Reiterate | 3 |
| 47 | <pre>.   .   ~ P -> ~ ( Q <-> R )</pre> | Theorem: contrapositive | 46 |
| 48 | <pre>.   .   ~ ( Q <-> R )</pre> | E -> | 47, 42 |
| 49 | <pre>.   .   Q <-> ~ R</pre> | Theorem: biconditional negation | 48 |
| 50 | <pre>.   .   Q -> ~ R</pre> | E <-> (L) | 49 |
| 51 | <pre>.   .   ~ R</pre> | E -> | 50, 45 |
| 52 | <pre>.   ~ ~ ( R -> ( P <-> Q ) )</pre> | subproof contradiction | 28 - 51 |
| 53 | <pre>.   R -> ( P <-> Q )</pre> | E ~ | 52 |
| 54 | <pre>.   ( P <-> Q ) <-> R</pre> | I <-> | 27, 53 |
| 55 | <pre>( P <-> ( Q <-> R ) ) -> ( ( P <-> Q ) <-> R )</pre> | subproof implication | 1 - 54 |

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

## ~ ( P v Q ) -> ( ~ P ^ ~ Q ) <a name="proof-6-11"></a>

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

## ( ~ P ^ ~ Q ) -> ~ ( P v Q ) <a name="proof-6-12"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ P ^ ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   P</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume Contradiction |  |
| 4 | <pre>.   .   .   ~ P</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 6 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof contradiction | 3 - 5 |
| 7 | <pre>.   P -> ~ ( ~ P ^ ~ Q )</pre> | subproof implication | 2 - 6 |
| 8 | <pre>.   .   Q</pre> | Assume Implication |  |
| 9 | <pre>.   .   .   ~ P ^ ~ Q</pre> | Assume Contradiction |  |
| 10 | <pre>.   .   .   ~ Q</pre> | E ^ (R) | 9 |
| 11 | <pre>.   .   .   Q</pre> | Reiterate | 8 |
| 12 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | subproof contradiction | 9 - 11 |
| 13 | <pre>.   Q -> ~ ( ~ P ^ ~ Q )</pre> | subproof implication | 8 - 12 |
| 14 | <pre>.   .   P v Q</pre> | Assume Contradiction |  |
| 15 | <pre>.   .   P -> ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 7 |
| 16 | <pre>.   .   Q -> ~ ( ~ P ^ ~ Q )</pre> | Reiterate | 13 |
| 17 | <pre>.   .   ~ ( ~ P ^ ~ Q )</pre> | E v | 15, 16, 14 |
| 18 | <pre>.   .   ~ P ^ ~ Q</pre> | Reiterate | 1 |
| 19 | <pre>.   ~ ( P v Q )</pre> | subproof contradiction | 14 - 18 |
| 20 | <pre>( ~ P ^ ~ Q ) -> ~ ( P v Q )</pre> | subproof implication | 1 - 19 |

## ~ ( P ^ Q ) -> ( ~ P v ~ Q ) <a name="proof-6-13"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P ^ Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ ( ~ P v ~ Q )</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   .   ~ P</pre> | Assume Contradiction |  |
| 4 | <pre>.   .   .   ~ P v ~ Q</pre> | I v (L) | 3 |
| 5 | <pre>.   .   .   ~ ( ~ P v ~ Q )</pre> | Reiterate | 2 |
| 6 | <pre>.   .   ~ ~ P</pre> | subproof contradiction | 3 - 5 |
| 7 | <pre>.   .   P</pre> | E ~ | 6 |
| 8 | <pre>.   .   .   ~ Q</pre> | Assume Contradiction |  |
| 9 | <pre>.   .   .   ~ P v ~ Q</pre> | I v (R) | 8 |
| 10 | <pre>.   .   .   ~ ( ~ P v ~ Q )</pre> | Reiterate | 2 |
| 11 | <pre>.   .   ~ ~ Q</pre> | subproof contradiction | 8 - 10 |
| 12 | <pre>.   .   Q</pre> | E ~ | 11 |
| 13 | <pre>.   .   P ^ Q</pre> | I ^ | 7, 12 |
| 14 | <pre>.   .   ~ ( P ^ Q )</pre> | Reiterate | 1 |
| 15 | <pre>.   ~ ~ ( ~ P v ~ Q )</pre> | subproof contradiction | 2 - 14 |
| 16 | <pre>.   ~ P v ~ Q</pre> | E ~ | 15 |
| 17 | <pre>~ ( P ^ Q ) -> ( ~ P v ~ Q )</pre> | subproof implication | 1 - 16 |

## ( ~ P v ~ Q ) -> ~ ( P ^ Q ) <a name="proof-6-14"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ P v ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ P</pre> | Assume Implication |  |
| 3 | <pre>.   .   .   P ^ Q</pre> | Assume Contradiction |  |
| 4 | <pre>.   .   .   P</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   .   ~ P</pre> | Reiterate | 2 |
| 6 | <pre>.   .   ~ ( P ^ Q )</pre> | subproof contradiction | 3 - 5 |
| 7 | <pre>.   ~ P -> ~ ( P ^ Q )</pre> | subproof implication | 2 - 6 |
| 8 | <pre>.   .   ~ Q</pre> | Assume Implication |  |
| 9 | <pre>.   .   .   P ^ Q</pre> | Assume Contradiction |  |
| 10 | <pre>.   .   .   Q</pre> | E ^ (R) | 9 |
| 11 | <pre>.   .   .   ~ Q</pre> | Reiterate | 8 |
| 12 | <pre>.   .   ~ ( P ^ Q )</pre> | subproof contradiction | 9 - 11 |
| 13 | <pre>.   ~ Q -> ~ ( P ^ Q )</pre> | subproof implication | 8 - 12 |
| 14 | <pre>.   ~ ( P ^ Q )</pre> | E v | 7, 13, 1 |
| 15 | <pre>( ~ P v ~ Q ) -> ~ ( P ^ Q )</pre> | subproof implication | 1 - 14 |

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

## ( P <-> ~ Q ) -> ~ ( P <-> Q ) <a name="proof-8-4"></a>

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

## ~ ( P <-> Q ) -> ( P <-> ~ Q ) <a name="proof-8-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P <-> Q )</pre> | Assume Implication |  |
| 2 | <pre>.   .   ~ ( P -> ~ Q )</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   P ^ ~ ~ Q</pre> | Theorem: arrow negation | 2 |
| 4 | <pre>.   .   P</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   ~ ~ Q</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E ~ | 5 |
| 7 | <pre>.   .   P -> Q</pre> | I -> | 4, 6 |
| 8 | <pre>.   .   Q -> P</pre> | I -> | 6, 4 |
| 9 | <pre>.   .   P <-> Q</pre> | I <-> | 7, 8 |
| 10 | <pre>.   .   ~ ( P <-> Q )</pre> | Reiterate | 1 |
| 11 | <pre>.   ~ ~ ( P -> ~ Q )</pre> | subproof contradiction | 2 - 10 |
| 12 | <pre>.   P -> ~ Q</pre> | E ~ | 11 |
| 13 | <pre>.   .   ~ ( ~ Q -> P )</pre> | Assume Contradiction |  |
| 14 | <pre>.   .   ~ Q ^ ~ P</pre> | Theorem: arrow negation | 13 |
| 15 | <pre>.   .   ~ Q</pre> | E ^ (L) | 14 |
| 16 | <pre>.   .   ~ P</pre> | E ^ (R) | 14 |
| 17 | <pre>.   .   ~ P -> ~ Q</pre> | I -> | 16, 15 |
| 18 | <pre>.   .   Q -> P</pre> | Theorem: contrapositive | 17 |
| 19 | <pre>.   .   ~ Q -> ~ P</pre> | I -> | 15, 16 |
| 20 | <pre>.   .   P -> Q</pre> | Theorem: contrapositive | 19 |
| 21 | <pre>.   .   P <-> Q</pre> | I <-> | 20, 18 |
| 22 | <pre>.   .   ~ ( P <-> Q )</pre> | Reiterate | 1 |
| 23 | <pre>.   ~ ~ ( ~ Q -> P )</pre> | subproof contradiction | 13 - 22 |
| 24 | <pre>.   ~ Q -> P</pre> | E ~ | 23 |
| 25 | <pre>.   P <-> ~ Q</pre> | I <-> | 12, 24 |
| 26 | <pre>~ ( P <-> Q ) -> ( P <-> ~ Q )</pre> | subproof implication | 1 - 25 |

## ( P <-> Q ) -> ( ~ P <-> ~ Q ) <a name="proof-8-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P <-> Q</pre> | Assume Implication |  |
| 2 | <pre>.   P -> Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   ~ Q -> ~ P</pre> | Theorem: contrapositive | 2 |
| 4 | <pre>.   Q -> P</pre> | E <-> (R) | 1 |
| 5 | <pre>.   ~ P -> ~ Q</pre> | Theorem: contrapositive | 4 |
| 6 | <pre>.   ~ P <-> ~ Q</pre> | I <-> | 5, 3 |
| 7 | <pre>( P <-> Q ) -> ( ~ P <-> ~ Q )</pre> | subproof implication | 1 - 6 |

## ( ~ P <-> ~ Q ) -> ( P <-> Q ) <a name="proof-8-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ P <-> ~ Q</pre> | Assume Implication |  |
| 2 | <pre>.   ~ P -> ~ Q</pre> | E <-> (L) | 1 |
| 3 | <pre>.   Q -> P</pre> | Theorem: contrapositive | 2 |
| 4 | <pre>.   ~ Q -> ~ P</pre> | E <-> (R) | 1 |
| 5 | <pre>.   P -> Q</pre> | Theorem: contrapositive | 4 |
| 6 | <pre>.   P <-> Q</pre> | I <-> | 5, 3 |
| 7 | <pre>( ~ P <-> ~ Q ) -> ( P <-> Q )</pre> | subproof implication | 1 - 6 |

## ( P <-> Q ) v ~ ( P <-> Q ) <a name="proof-8-8"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | Assume Contradiction |  |
| 2 | <pre>.   .   P <-> Q</pre> | Assume Contradiction |  |
| 3 | <pre>.   .   ( P <-> Q ) v ~ ( P <-> Q )</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | Reiterate | 1 |
| 5 | <pre>.   ~ ( P <-> Q )</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   ( P <-> Q ) v ~ ( P <-> Q )</pre> | I v (R) | 5 |
| 7 | <pre>~ ~ ( ( P <-> Q ) v ~ ( P <-> Q ) )</pre> | subproof contradiction | 1 - 6 |
| 8 | <pre>( P <-> Q ) v ~ ( P <-> Q )</pre> | E ~ | 7 |

## ( P <-> ~ Q ) -> ( ( P v Q ) ^ ( ~ P v ~ Q ) ) <a name="proof-8-9"></a>

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

## ( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q ) <a name="proof-8-10"></a>

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
| 42 | <pre>.   .   ~ ~ P</pre> | subproof contradiction | 38 - 41 |
| 43 | <pre>.   .   P</pre> | E ~ | 42 |
| 44 | <pre>.   ~ Q -> P</pre> | subproof implication | 37 - 43 |
| 45 | <pre>.   P <-> ~ Q</pre> | I <-> | 36, 44 |
| 46 | <pre>( ( P v Q ) ^ ( ~ P v ~ Q ) ) -> ( P <-> ~ Q )</pre> | subproof implication | 1 - 45 |

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

