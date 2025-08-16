
# Table of Contents

1. [basics](#basics)
    1. [P -> P](#proof-1-1)
    2. [~ ( P ^ ~ P )](#proof-1-2)
    3. [P v ~ P](#proof-1-3)
    4. [( P ^ P ) -> P](#proof-1-4)
    5. [( P v P ) -> P](#proof-1-5)
2. [arrows](#arrows)
    1. [( P -> Q ) -> ( ~ Q -> ~ P )](#proof-2-1)
    2. [( ~ Q -> ~ P ) -> ( P -> Q )](#proof-2-2)
    3. [Q -> ( P -> Q )](#proof-2-3)
    4. [( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )](#proof-2-4)
3. [commutativity](#commutativity)
    1. [( P ^ Q ) -> ( Q ^ P )](#proof-3-1)
    2. [( P v Q ) -> ( Q v P )](#proof-3-2)
4. [associativity](#associativity)
    1. [( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )](#proof-4-1)
    2. [( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )](#proof-4-2)
    3. [( ( P v Q ) v R ) -> ( P v ( Q v R ) )](#proof-4-3)
    4. [( P v ( Q v R ) ) -> ( ( P v Q ) v R )](#proof-4-4)
5. [distributivity](#distributivity)
    1. [( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )](#proof-5-1)
    2. [( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )](#proof-5-2)
    3. [( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )](#proof-5-3)
    4. [( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )](#proof-5-4)
    5. [( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )](#proof-5-5)
    6. [( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )](#proof-5-6)

# basics <a name="basics"></a>

## P -> P <a name="proof-1-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P</pre> | Assume |  |
| 2 | <pre>.   P</pre> | Repeat | 1 |
| 3 | <pre>P -> P</pre> | subproof implication | 1 - 2 |

## ~ ( P ^ ~ P ) <a name="proof-1-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ~ P</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ~ P</pre> | E ^ (R) | 1 |
| 4 | <pre>~ ( P ^ ~ P )</pre> | subproof contradiction | 1 - 3 |

## P v ~ P <a name="proof-1-3"></a>

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

## ( P ^ P ) -> P <a name="proof-1-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ P</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>( P ^ P ) -> P</pre> | subproof implication | 1 - 2 |

## ( P v P ) -> P <a name="proof-1-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v P</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   P</pre> | Repeat | 2 |
| 4 | <pre>.   P -> P</pre> | subproof implication | 2 - 3 |
| 5 | <pre>.   P</pre> | E v | 4, 4, 1 |
| 6 | <pre>( P v P ) -> P</pre> | subproof implication | 1 - 5 |

# arrows <a name="arrows"></a>

## ( P -> Q ) -> ( ~ Q -> ~ P ) <a name="proof-2-1"></a>

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

## ( ~ Q -> ~ P ) -> ( P -> Q ) <a name="proof-2-2"></a>

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

## Q -> ( P -> Q ) <a name="proof-2-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   Q</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   Q</pre> | Reiterate | 1 |
| 4 | <pre>.   P -> Q</pre> | subproof implication | 2 - 3 |
| 5 | <pre>Q -> ( P -> Q )</pre> | subproof implication | 1 - 4 |

## ( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R ) <a name="proof-2-4"></a>

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

# commutativity <a name="commutativity"></a>

## ( P ^ Q ) -> ( Q ^ P ) <a name="proof-3-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ Q</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q ^ P</pre> | I ^ | 3, 2 |
| 5 | <pre>( P ^ Q ) -> ( Q ^ P )</pre> | subproof implication | 1 - 4 |

## ( P v Q ) -> ( Q v P ) <a name="proof-3-2"></a>

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

# associativity <a name="associativity"></a>

## ( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) ) <a name="proof-4-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P ^ Q ) ^ R</pre> | Assume |  |
| 2 | <pre>.   P ^ Q</pre> | E ^ (L) | 1 |
| 3 | <pre>.   R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   P</pre> | E ^ (L) | 2 |
| 5 | <pre>.   Q</pre> | E ^ (R) | 2 |
| 6 | <pre>.   Q ^ R</pre> | I ^ | 5, 3 |
| 7 | <pre>.   P ^ ( Q ^ R )</pre> | I ^ | 4, 6 |
| 8 | <pre>( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )</pre> | subproof implication | 1 - 7 |

## ( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R ) <a name="proof-4-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ( Q ^ R )</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q ^ R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   R</pre> | E ^ (R) | 3 |
| 6 | <pre>.   P ^ Q</pre> | I ^ | 2, 4 |
| 7 | <pre>.   ( P ^ Q ) ^ R</pre> | I ^ | 6, 5 |
| 8 | <pre>( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )</pre> | subproof implication | 1 - 7 |

## ( ( P v Q ) v R ) -> ( P v ( Q v R ) ) <a name="proof-4-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P v Q ) v R</pre> | Assume |  |
| 2 | <pre>.   .   P v Q</pre> | Assume |  |
| 3 | <pre>.   .   .   P</pre> | Assume |  |
| 4 | <pre>.   .   .   P v ( Q v R )</pre> | I v (L) | 3 |
| 5 | <pre>.   .   P -> ( P v ( Q v R ) )</pre> | subproof implication | 3 - 4 |
| 6 | <pre>.   .   .   Q</pre> | Assume |  |
| 7 | <pre>.   .   .   Q v R</pre> | I v (L) | 6 |
| 8 | <pre>.   .   .   P v ( Q v R )</pre> | I v (R) | 7 |
| 9 | <pre>.   .   Q -> ( P v ( Q v R ) )</pre> | subproof implication | 6 - 8 |
| 10 | <pre>.   .   P v ( Q v R )</pre> | E v | 5, 9, 2 |
| 11 | <pre>.   ( P v Q ) -> ( P v ( Q v R ) )</pre> | subproof implication | 2 - 10 |
| 12 | <pre>.   .   R</pre> | Assume |  |
| 13 | <pre>.   .   Q v R</pre> | I v (R) | 12 |
| 14 | <pre>.   .   P v ( Q v R )</pre> | I v (R) | 13 |
| 15 | <pre>.   R -> ( P v ( Q v R ) )</pre> | subproof implication | 12 - 14 |
| 16 | <pre>.   P v ( Q v R )</pre> | E v | 11, 15, 1 |
| 17 | <pre>( ( P v Q ) v R ) -> ( P v ( Q v R ) )</pre> | subproof implication | 1 - 16 |

## ( P v ( Q v R ) ) -> ( ( P v Q ) v R ) <a name="proof-4-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P v ( Q v R )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   P v Q</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ( P v Q ) v R</pre> | I v (L) | 3 |
| 5 | <pre>.   P -> ( ( P v Q ) v R )</pre> | subproof implication | 2 - 4 |
| 6 | <pre>.   .   Q v R</pre> | Assume |  |
| 7 | <pre>.   .   .   Q</pre> | Assume |  |
| 8 | <pre>.   .   .   P v Q</pre> | I v (R) | 7 |
| 9 | <pre>.   .   .   ( P v Q ) v R</pre> | I v (L) | 8 |
| 10 | <pre>.   .   Q -> ( ( P v Q ) v R )</pre> | subproof implication | 7 - 9 |
| 11 | <pre>.   .   .   R</pre> | Assume |  |
| 12 | <pre>.   .   .   ( P v Q ) v R</pre> | I v (R) | 11 |
| 13 | <pre>.   .   R -> ( ( P v Q ) v R )</pre> | subproof implication | 11 - 12 |
| 14 | <pre>.   .   ( P v Q ) v R</pre> | E v | 10, 13, 6 |
| 15 | <pre>.   ( Q v R ) -> ( ( P v Q ) v R )</pre> | subproof implication | 6 - 14 |
| 16 | <pre>.   ( P v Q ) v R</pre> | E v | 5, 15, 1 |
| 17 | <pre>( P v ( Q v R ) ) -> ( ( P v Q ) v R )</pre> | subproof implication | 1 - 16 |

# distributivity <a name="distributivity"></a>

## ( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) ) <a name="proof-5-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q v R )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   P -> ( Q v R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   Q v R</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   .   Q</pre> | Assume |  |
| 6 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 7 | <pre>.   .   .   P -> Q</pre> | I -> | 6, 5 |
| 8 | <pre>.   .   .   ( P -> Q ) v ( P -> R )</pre> | I v (L) | 7 |
| 9 | <pre>.   .   Q -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 5 - 8 |
| 10 | <pre>.   .   .   R</pre> | Assume |  |
| 11 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 12 | <pre>.   .   .   P -> R</pre> | I -> | 11, 10 |
| 13 | <pre>.   .   .   ( P -> Q ) v ( P -> R )</pre> | I v (R) | 12 |
| 14 | <pre>.   .   R -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 10 - 13 |
| 15 | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | E v | 9, 14, 4 |
| 16 | <pre>.   P -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 2 - 15 |
| 17 | <pre>.   .   ~ P</pre> | Assume |  |
| 18 | <pre>.   .   .   ~ Q</pre> | Assume |  |
| 19 | <pre>.   .   .   ~ P</pre> | Reiterate | 17 |
| 20 | <pre>.   .   ~ Q -> ~ P</pre> | subproof implication | 18 - 19 |
| 21 | <pre>.   .   .   P</pre> | Assume |  |
| 22 | <pre>.   .   .   .   ~ Q</pre> | Assume |  |
| 23 | <pre>.   .   .   .   ~ Q -> ~ P</pre> | Reiterate | 20 |
| 24 | <pre>.   .   .   .   ~ P</pre> | E -> | 23, 22 |
| 25 | <pre>.   .   .   .   P</pre> | Reiterate | 21 |
| 26 | <pre>.   .   .   ~ ~ Q</pre> | subproof contradiction | 22 - 25 |
| 27 | <pre>.   .   .   Q</pre> | E ~ | 26 |
| 28 | <pre>.   .   P -> Q</pre> | subproof implication | 21 - 27 |
| 29 | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | I v (L) | 28 |
| 30 | <pre>.   ~ P -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 17 - 29 |
| 31 | <pre>.   .   ~ ( P v ~ P )</pre> | Assume |  |
| 32 | <pre>.   .   .   P</pre> | Assume |  |
| 33 | <pre>.   .   .   P v ~ P</pre> | I v (L) | 32 |
| 34 | <pre>.   .   .   ~ ( P v ~ P )</pre> | Reiterate | 31 |
| 35 | <pre>.   .   ~ P</pre> | subproof contradiction | 32 - 34 |
| 36 | <pre>.   .   P v ~ P</pre> | I v (R) | 35 |
| 37 | <pre>.   ~ ~ ( P v ~ P )</pre> | subproof contradiction | 31 - 36 |
| 38 | <pre>.   P v ~ P</pre> | E ~ | 37 |
| 39 | <pre>.   ( P -> Q ) v ( P -> R )</pre> | E v | 16, 30, 38 |
| 40 | <pre>( P -> ( Q v R ) ) -> ( ( P -> Q ) v ( P -> R ) )</pre> | subproof implication | 1 - 39 |

## ( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) ) <a name="proof-5-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) v ( P -> R )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   .   P -> Q</pre> | Assume |  |
| 4 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 5 | <pre>.   .   .   Q</pre> | E -> | 3, 4 |
| 6 | <pre>.   .   .   Q v R</pre> | I v (L) | 5 |
| 7 | <pre>.   .   ( P -> Q ) -> ( Q v R )</pre> | subproof implication | 3 - 6 |
| 8 | <pre>.   .   .   P -> R</pre> | Assume |  |
| 9 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 10 | <pre>.   .   .   R</pre> | E -> | 8, 9 |
| 11 | <pre>.   .   .   Q v R</pre> | I v (R) | 10 |
| 12 | <pre>.   .   ( P -> R ) -> ( Q v R )</pre> | subproof implication | 8 - 11 |
| 13 | <pre>.   .   ( P -> Q ) v ( P -> R )</pre> | Reiterate | 1 |
| 14 | <pre>.   .   Q v R</pre> | E v | 7, 12, 13 |
| 15 | <pre>.   P -> ( Q v R )</pre> | subproof implication | 2 - 14 |
| 16 | <pre>( ( P -> Q ) v ( P -> R ) ) -> ( P -> ( Q v R ) )</pre> | subproof implication | 1 - 15 |

## ( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) ) <a name="proof-5-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P -> ( Q ^ R )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   Q ^ R</pre> | E -> | 3, 2 |
| 5 | <pre>.   .   Q</pre> | E ^ (L) | 4 |
| 6 | <pre>.   P -> Q</pre> | subproof implication | 2 - 5 |
| 7 | <pre>.   .   P</pre> | Assume |  |
| 8 | <pre>.   .   P -> ( Q ^ R )</pre> | Reiterate | 1 |
| 9 | <pre>.   .   Q ^ R</pre> | E -> | 8, 7 |
| 10 | <pre>.   .   R</pre> | E ^ (R) | 9 |
| 11 | <pre>.   P -> R</pre> | subproof implication | 7 - 10 |
| 12 | <pre>.   ( P -> Q ) ^ ( P -> R )</pre> | I ^ | 6, 11 |
| 13 | <pre>( P -> ( Q ^ R ) ) -> ( ( P -> Q ) ^ ( P -> R ) )</pre> | subproof implication | 1 - 12 |

## ( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) ) <a name="proof-5-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P -> Q ) ^ ( P -> R )</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   ( P -> Q ) ^ ( P -> R )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   P -> Q</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   P -> R</pre> | E ^ (R) | 3 |
| 6 | <pre>.   .   Q</pre> | E -> | 4, 2 |
| 7 | <pre>.   .   R</pre> | E -> | 5, 2 |
| 8 | <pre>.   .   Q ^ R</pre> | I ^ | 6, 7 |
| 9 | <pre>.   P -> ( Q ^ R )</pre> | subproof implication | 2 - 8 |
| 10 | <pre>( ( P -> Q ) ^ ( P -> R ) ) -> ( P -> ( Q ^ R ) )</pre> | subproof implication | 1 - 9 |

## ( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) ) <a name="proof-5-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ ( Q v R )</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q v R</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   Q</pre> | Assume |  |
| 5 | <pre>.   .   P</pre> | Reiterate | 2 |
| 6 | <pre>.   .   P ^ Q</pre> | I ^ | 5, 4 |
| 7 | <pre>.   .   ( P ^ Q ) v ( P ^ R )</pre> | I v (L) | 6 |
| 8 | <pre>.   Q -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof implication | 4 - 7 |
| 9 | <pre>.   .   R</pre> | Assume |  |
| 10 | <pre>.   .   P</pre> | Reiterate | 2 |
| 11 | <pre>.   .   P ^ R</pre> | I ^ | 10, 9 |
| 12 | <pre>.   .   ( P ^ Q ) v ( P ^ R )</pre> | I v (R) | 11 |
| 13 | <pre>.   R -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof implication | 9 - 12 |
| 14 | <pre>.   ( P ^ Q ) v ( P ^ R )</pre> | E v | 8, 13, 3 |
| 15 | <pre>( P ^ ( Q v R ) ) -> ( ( P ^ Q ) v ( P ^ R ) )</pre> | subproof implication | 1 - 14 |

## ( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) ) <a name="proof-5-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ( P ^ Q ) v ( P ^ R )</pre> | Assume |  |
| 2 | <pre>.   .   P ^ Q</pre> | Assume |  |
| 3 | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 4 | <pre>.   .   Q</pre> | E ^ (R) | 2 |
| 5 | <pre>.   .   Q v R</pre> | I v (L) | 4 |
| 6 | <pre>.   .   P ^ ( Q v R )</pre> | I ^ | 3, 5 |
| 7 | <pre>.   ( P ^ Q ) -> ( P ^ ( Q v R ) )</pre> | subproof implication | 2 - 6 |
| 8 | <pre>.   .   P ^ R</pre> | Assume |  |
| 9 | <pre>.   .   P</pre> | E ^ (L) | 8 |
| 10 | <pre>.   .   R</pre> | E ^ (R) | 8 |
| 11 | <pre>.   .   Q v R</pre> | I v (R) | 10 |
| 12 | <pre>.   .   P ^ ( Q v R )</pre> | I ^ | 9, 11 |
| 13 | <pre>.   ( P ^ R ) -> ( P ^ ( Q v R ) )</pre> | subproof implication | 8 - 12 |
| 14 | <pre>.   P ^ ( Q v R )</pre> | E v | 7, 13, 1 |
| 15 | <pre>( ( P ^ Q ) v ( P ^ R ) ) -> ( P ^ ( Q v R ) )</pre> | subproof implication | 1 - 14 |

