
# Table of Contents

1. [basics](#basics)
    1. [P -> P](#proof-1-1)
    2. [~ ( P ^ ~ P )](#proof-1-2)
    3. [P v ~ P](#proof-1-3)
    4. [( P ^ P ) -> P](#proof-1-4)
    5. [( P v P ) -> P](#proof-1-5)
2. [commutativity](#commutativity)
    1. [( P ^ Q ) -> ( Q ^ P )](#proof-2-1)
    2. [( P v Q ) -> ( Q v P )](#proof-2-2)
3. [associativity](#associativity)
    1. [( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) )](#proof-3-1)
    2. [( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R )](#proof-3-2)
    3. [( ( P v Q ) v R ) -> ( P v ( Q v R ) )](#proof-3-3)
    4. [( P v ( Q v R ) ) -> ( ( P v Q ) v R )](#proof-3-4)
4. [distributivity](#distributivity)
5. [arrows](#arrows)
    1. [( P -> Q ) -> ( ~ Q -> ~ P )](#proof-5-1)
    2. [( ~ Q -> ~ P ) -> ( P -> Q )](#proof-5-2)
    3. [Q -> ( P -> Q )](#proof-5-3)
    4. [( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R )](#proof-5-4)

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

# commutativity <a name="commutativity"></a>

## ( P ^ Q ) -> ( Q ^ P ) <a name="proof-2-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   P ^ Q</pre> | Assume |  |
| 2 | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 | <pre>.   Q</pre> | E ^ (R) | 1 |
| 4 | <pre>.   Q ^ P</pre> | I ^ | 3, 2 |
| 5 | <pre>( P ^ Q ) -> ( Q ^ P )</pre> | subproof implication | 1 - 4 |

## ( P v Q ) -> ( Q v P ) <a name="proof-2-2"></a>

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

## ( ( P ^ Q ) ^ R ) -> ( P ^ ( Q ^ R ) ) <a name="proof-3-1"></a>

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

## ( P ^ ( Q ^ R ) ) -> ( ( P ^ Q ) ^ R ) <a name="proof-3-2"></a>

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

## ( ( P v Q ) v R ) -> ( P v ( Q v R ) ) <a name="proof-3-3"></a>

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

## ( P v ( Q v R ) ) -> ( ( P v Q ) v R ) <a name="proof-3-4"></a>

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

# arrows <a name="arrows"></a>

## ( P -> Q ) -> ( ~ Q -> ~ P ) <a name="proof-5-1"></a>

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

## ( ~ Q -> ~ P ) -> ( P -> Q ) <a name="proof-5-2"></a>

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

## Q -> ( P -> Q ) <a name="proof-5-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   Q</pre> | Assume |  |
| 2 | <pre>.   .   P</pre> | Assume |  |
| 3 | <pre>.   .   Q</pre> | Reiterate | 1 |
| 4 | <pre>.   P -> Q</pre> | subproof implication | 2 - 3 |
| 5 | <pre>Q -> ( P -> Q )</pre> | subproof implication | 1 - 4 |

## ( ( P -> Q ) ^ ( Q -> R ) ) -> ( P -> R ) <a name="proof-5-4"></a>

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

