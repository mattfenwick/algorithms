
# Table of Contents

1. [basics](#basics)
    1. [P v ~ P](#proof-1-1)
    2. [∀x.( Q(x) ) -> Q(a)](#proof-1-2)
    3. [Q(a) -> ∃x.( Q(x) )](#proof-1-3)
    4. [( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )](#proof-1-4)
    5. [∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R](#proof-1-5)
    6. [( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R](#proof-1-6)
    7. [∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )](#proof-1-7)

# basics <a name="basics"></a>

## P v ~ P <a name="proof-1-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   ~ ( P v ~ P )</pre> | Assume: contradiction |  |
| 2 | <pre>.   .   .   P</pre> | Assume: contradiction |  |
| 3 | <pre>.   .   .   P v ~ P</pre> | I v (L) | 2 |
| 4 | <pre>.   .   .   ~ ( P v ~ P )</pre> | Reiterate | 1 |
| 5 | <pre>.   .   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   .   P v ~ P</pre> | I v (R) | 5 |
| 7 | <pre>.   P v ~ P</pre> | subproof contradiction | 1 - 6 |

## ∀x.( Q(x) ) -> Q(a) <a name="proof-1-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   ∀x.( Q(x) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   Q(a)</pre> | E ∀ | 1 |
| 3 | <pre>.   ∀x.( Q(x) ) -> Q(a)</pre> | subproof implication | 1 - 2 |

## Q(a) -> ∃x.( Q(x) ) <a name="proof-1-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   Q(a)</pre> | Assume: implication |  |
| 2 | <pre>.   .   ∃x.( Q(x) )</pre> | I ∃ | 1 |
| 3 | <pre>.   Q(a) -> ∃x.( Q(x) )</pre> | subproof implication | 1 - 2 |

## ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) ) <a name="proof-1-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   ∃x.( T )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   .   P -> ∃x.( Q(x) )</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   .   T</pre> | Assume: ∃ elimination | 2 |
| 5 | <pre>.   .   .   .   P</pre> | Assume: implication |  |
| 6 | <pre>.   .   .   .   P -> ∃x.( Q(x) )</pre> | Reiterate | 3 |
| 7 | <pre>.   .   .   .   ∃x.( Q(x) )</pre> | E -> | 6, 5 |
| 8 | <pre>.   .   .   .   .   Q(b)</pre> | Assume: ∃ elimination | 7 |
| 9 | <pre>.   .   .   .   .   P</pre> | Reiterate | 5 |
| 10 | <pre>.   .   .   .   .   P -> Q(b)</pre> | I -> | 9, 8 |
| 11 | <pre>.   .   .   .   .   ∃x.( P -> Q(x) )</pre> | I ∃ | 10 |
| 12 | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> | subproof ∃ elimination | 8 - 11 |
| 13 | <pre>.   .   .   P -> ∃x.( P -> Q(x) )</pre> | subproof implication | 5 - 12 |
| 14 | <pre>.   .   .   .   ~ P</pre> | Assume: implication |  |
| 15 | <pre>.   .   .   .   .   ~ Q(a)</pre> | Assume: implication |  |
| 16 | <pre>.   .   .   .   .   ~ P</pre> | Reiterate | 14 |
| 17 | <pre>.   .   .   .   ~ Q(a) -> ~ P</pre> | subproof implication | 15 - 16 |
| 18 | <pre>.   .   .   .   P -> Q(a)</pre> | Theorem: contrapositive | 17 |
| 19 | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> | I ∃ | 18 |
| 20 | <pre>.   .   .   ~ P -> ∃x.( P -> Q(x) )</pre> | subproof implication | 14 - 19 |
| 21 | <pre>.   .   .   P v ~ P</pre> | Theorem: excluded middle |  |
| 22 | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> | E v | 13, 20, 21 |
| 23 | <pre>.   .   ∃x.( P -> Q(x) )</pre> | subproof ∃ elimination | 4 - 22 |
| 24 | <pre>.   ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) -> ∃x.( P -> Q(x) )</pre> | subproof implication | 1 - 23 |
| 25 | <pre>.   .   ∃x.( P -> Q(x) )</pre> | Assume: implication |  |
| 26 | <pre>.   .   .   P</pre> | Assume: implication |  |
| 27 | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> | Reiterate | 25 |
| 28 | <pre>.   .   .   .   P -> Q(a)</pre> | Assume: ∃ elimination | 27 |
| 29 | <pre>.   .   .   .   P</pre> | Reiterate | 26 |
| 30 | <pre>.   .   .   .   Q(a)</pre> | E -> | 28, 29 |
| 31 | <pre>.   .   .   .   ∃x.( Q(x) )</pre> | I ∃ | 30 |
| 32 | <pre>.   .   .   ∃x.( Q(x) )</pre> | subproof ∃ elimination | 28 - 31 |
| 33 | <pre>.   .   P -> ∃x.( Q(x) )</pre> | subproof implication | 26 - 32 |
| 34 | <pre>.   .   .   P -> Q(a)</pre> | Assume: ∃ elimination | 25 |
| 35 | <pre>.   .   .   T</pre> | Reiterate | 0 |
| 36 | <pre>.   .   .   ∃x.( T )</pre> | I ∃ | 35 |
| 37 | <pre>.   .   ∃x.( T )</pre> | subproof ∃ elimination | 34 - 36 |
| 38 | <pre>.   .   ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) )</pre> | I ^ | 37, 33 |
| 39 | <pre>.   ∃x.( P -> Q(x) ) -> ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) )</pre> | subproof implication | 25 - 38 |
| 40 | <pre>.   ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )</pre> | I <-> | 24, 39 |

## ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R <a name="proof-1-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   ∃x.( Q(x) ^ ( Q(x) -> R ) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   .   Q(a) ^ ( Q(a) -> R )</pre> | Assume: ∃ elimination | 1 |
| 3 | <pre>.   .   .   Q(a)</pre> | E ^ (L) | 2 |
| 4 | <pre>.   .   .   Q(a) -> R</pre> | E ^ (R) | 2 |
| 5 | <pre>.   .   .   R</pre> | E -> | 4, 3 |
| 6 | <pre>.   .   R</pre> | subproof ∃ elimination | 2 - 5 |
| 7 | <pre>.   ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R</pre> | subproof implication | 1 - 6 |

## ( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R <a name="proof-1-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R )</pre> | Assume: implication |  |
| 2 | <pre>.   .   ∀y.( Q(y) )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   .   ∃x.( Q(x) -> R )</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   .   Q(a) -> R</pre> | Assume: ∃ elimination | 3 |
| 5 | <pre>.   .   .   ∀y.( Q(y) )</pre> | Reiterate | 2 |
| 6 | <pre>.   .   .   Q(a)</pre> | E ∀ | 5 |
| 7 | <pre>.   .   .   R</pre> | E -> | 4, 6 |
| 8 | <pre>.   .   R</pre> | subproof ∃ elimination | 4 - 7 |
| 9 | <pre>.   ( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R</pre> | subproof implication | 1 - 8 |

## ∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) ) <a name="proof-1-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 3 | <pre>.   .   .   P(a) ^ Q(a)</pre> | E ∀ | 2 |
| 4 | <pre>.   .   .   P(a)</pre> | E ^ (L) | 3 |
| 5 | <pre>.   .   ∀y.( P(y) )</pre> | subproof ∀ introduction | 2 - 4 |
| 6 | <pre>.   .   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 7 | <pre>.   .   .   P(a) ^ Q(a)</pre> | E ∀ | 6 |
| 8 | <pre>.   .   .   Q(a)</pre> | E ^ (R) | 7 |
| 9 | <pre>.   .   ∀z.( Q(z) )</pre> | subproof ∀ introduction | 6 - 8 |
| 10 | <pre>.   .   ∀y.( P(y) ) ^ ∀z.( Q(z) )</pre> | I ^ | 5, 9 |
| 11 | <pre>.   ∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )</pre> | subproof implication | 1 - 10 |

