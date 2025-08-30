
# Table of Contents

1. [basics](#basics)
    1. [P v ~ P](#proof-1-1)
    2. [Q(a) -> ∃x.( Q(x) )](#proof-1-2)
    3. [( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )](#proof-1-3)
    4. [∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R](#proof-1-4)
    5. [( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R](#proof-1-5)
    6. [∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )](#proof-1-6)

# basics <a name="basics"></a>

## P v ~ P <a name="proof-1-1"></a>

| Line | Formula | Term var | Justification | Lines used |
| - | - | - | - | - |
| 1 | <pre>.   .   ~ ( P v ~ P )</pre> |  | Assume: contradiction |  |
| 2 | <pre>.   .   .   P</pre> |  | Assume: contradiction |  |
| 3 | <pre>.   .   .   P v ~ P</pre> |  | I v (L) | 2 |
| 4 | <pre>.   .   .   ~ ( P v ~ P )</pre> |  | Reiterate | 1 |
| 5 | <pre>.   .   ~ P</pre> |  | subproof contradiction | 2 - 4 |
| 6 | <pre>.   .   P v ~ P</pre> |  | I v (R) | 5 |
| 7 | <pre>.   P v ~ P</pre> |  | subproof contradiction | 1 - 6 |

## Q(a) -> ∃x.( Q(x) ) <a name="proof-1-2"></a>

| Line | Formula | Term var | Justification | Lines used |
| - | - | - | - | - |
| 1 | <pre>.   .   Q(a)</pre> |  | Assume: implication |  |
| 2 | <pre>.   .   ∃x.( Q(x) )</pre> |  | I ∃ | 1 |
| 3 | <pre>.   Q(a) -> ∃x.( Q(x) )</pre> |  | subproof implication | 1 - 2 |

## ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) ) <a name="proof-1-3"></a>

| Line | Formula | Term var | Justification | Lines used |
| - | - | - | - | - |
| 1 | <pre>.   .   ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) )</pre> |  | Assume: implication |  |
| 2 | <pre>.   .   ∃x.( T )</pre> |  | E ^ (L) | 1 |
| 3 | <pre>.   .   P -> ∃x.( Q(x) )</pre> |  | E ^ (R) | 1 |
| 4 | <pre>.   .   .   </pre> | a | Term var |  |
| 5 | <pre>.   .   .   T</pre> |  | Assume: ∃ elimination | 2 |
| 6 | <pre>.   .   .   .   P</pre> |  | Assume: implication |  |
| 7 | <pre>.   .   .   .   P -> ∃x.( Q(x) )</pre> |  | Reiterate | 3 |
| 8 | <pre>.   .   .   .   ∃x.( Q(x) )</pre> |  | E -> | 7, 6 |
| 9 | <pre>.   .   .   .   .   </pre> | b | Term var |  |
| 10 | <pre>.   .   .   .   .   Q(b)</pre> |  | Assume: ∃ elimination | 8 |
| 11 | <pre>.   .   .   .   .   P</pre> |  | Reiterate | 6 |
| 12 | <pre>.   .   .   .   .   P -> Q(b)</pre> |  | I -> | 11, 10 |
| 13 | <pre>.   .   .   .   .   ∃x.( P -> Q(x) )</pre> |  | I ∃ | 12 |
| 14 | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> |  | subproof ∃ elimination | 9 - 13 |
| 15 | <pre>.   .   .   P -> ∃x.( P -> Q(x) )</pre> |  | subproof implication | 6 - 14 |
| 16 | <pre>.   .   .   .   ~ P</pre> |  | Assume: implication |  |
| 17 | <pre>.   .   .   .   .   ~ Q(a)</pre> |  | Assume: implication |  |
| 18 | <pre>.   .   .   .   .   ~ P</pre> |  | Reiterate | 16 |
| 19 | <pre>.   .   .   .   ~ Q(a) -> ~ P</pre> |  | subproof implication | 17 - 18 |
| 20 | <pre>.   .   .   .   P -> Q(a)</pre> |  | Theorem: contrapositive | 19 |
| 21 | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> |  | I ∃ | 20 |
| 22 | <pre>.   .   .   ~ P -> ∃x.( P -> Q(x) )</pre> |  | subproof implication | 16 - 21 |
| 23 | <pre>.   .   .   P v ~ P</pre> |  | Theorem: excluded middle |  |
| 24 | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> |  | E v | 15, 22, 23 |
| 25 | <pre>.   .   ∃x.( P -> Q(x) )</pre> |  | subproof ∃ elimination | 4 - 24 |
| 26 | <pre>.   ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) -> ∃x.( P -> Q(x) )</pre> |  | subproof implication | 1 - 25 |
| 27 | <pre>.   .   ∃x.( P -> Q(x) )</pre> |  | Assume: implication |  |
| 28 | <pre>.   .   .   P</pre> |  | Assume: implication |  |
| 29 | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> |  | Reiterate | 27 |
| 30 | <pre>.   .   .   .   </pre> | a | Term var |  |
| 31 | <pre>.   .   .   .   P -> Q(a)</pre> |  | Assume: ∃ elimination | 29 |
| 32 | <pre>.   .   .   .   P</pre> |  | Reiterate | 28 |
| 33 | <pre>.   .   .   .   Q(a)</pre> |  | E -> | 31, 32 |
| 34 | <pre>.   .   .   .   ∃x.( Q(x) )</pre> |  | I ∃ | 33 |
| 35 | <pre>.   .   .   ∃x.( Q(x) )</pre> |  | subproof ∃ elimination | 30 - 34 |
| 36 | <pre>.   .   P -> ∃x.( Q(x) )</pre> |  | subproof implication | 28 - 35 |
| 37 | <pre>.   .   .   </pre> | a | Term var |  |
| 38 | <pre>.   .   .   P -> Q(a)</pre> |  | Assume: ∃ elimination | 27 |
| 39 | <pre>.   .   .   T</pre> |  | Reiterate | 0 |
| 40 | <pre>.   .   .   ∃x.( T )</pre> |  | I ∃ | 39 |
| 41 | <pre>.   .   ∃x.( T )</pre> |  | subproof ∃ elimination | 37 - 40 |
| 42 | <pre>.   .   ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) )</pre> |  | I ^ | 41, 36 |
| 43 | <pre>.   ∃x.( P -> Q(x) ) -> ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) )</pre> |  | subproof implication | 27 - 42 |
| 44 | <pre>.   ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )</pre> |  | I <-> | 26, 43 |

## ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R <a name="proof-1-4"></a>

| Line | Formula | Term var | Justification | Lines used |
| - | - | - | - | - |
| 1 | <pre>.   .   ∃x.( Q(x) ^ ( Q(x) -> R ) )</pre> |  | Assume: implication |  |
| 2 | <pre>.   .   .   </pre> | a | Term var |  |
| 3 | <pre>.   .   .   Q(a) ^ ( Q(a) -> R )</pre> |  | Assume: ∃ elimination | 1 |
| 4 | <pre>.   .   .   Q(a)</pre> |  | E ^ (L) | 3 |
| 5 | <pre>.   .   .   Q(a) -> R</pre> |  | E ^ (R) | 3 |
| 6 | <pre>.   .   .   R</pre> |  | E -> | 5, 4 |
| 7 | <pre>.   .   R</pre> |  | subproof ∃ elimination | 2 - 6 |
| 8 | <pre>.   ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R</pre> |  | subproof implication | 1 - 7 |

## ( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R <a name="proof-1-5"></a>

| Line | Formula | Term var | Justification | Lines used |
| - | - | - | - | - |
| 1 | <pre>.   .   ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R )</pre> |  | Assume: implication |  |
| 2 | <pre>.   .   ∀y.( Q(y) )</pre> |  | E ^ (L) | 1 |
| 3 | <pre>.   .   ∃x.( Q(x) -> R )</pre> |  | E ^ (R) | 1 |
| 4 | <pre>.   .   .   </pre> | a | Term var |  |
| 5 | <pre>.   .   .   Q(a) -> R</pre> |  | Assume: ∃ elimination | 3 |
| 6 | <pre>.   .   .   ∀y.( Q(y) )</pre> |  | Reiterate | 2 |
| 7 | <pre>.   .   .   Q(a)</pre> |  | E ∀ | 6, 4 |
| 8 | <pre>.   .   .   R</pre> |  | E -> | 5, 7 |
| 9 | <pre>.   .   R</pre> |  | subproof ∃ elimination | 4 - 8 |
| 10 | <pre>.   ( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R</pre> |  | subproof implication | 1 - 9 |

## ∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) ) <a name="proof-1-6"></a>

| Line | Formula | Term var | Justification | Lines used |
| - | - | - | - | - |
| 1 | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> |  | Assume: implication |  |
| 2 | <pre>.   .   .   </pre> | a | Term var |  |
| 3 | <pre>.   .   .   ∀x.( P(x) ^ Q(x) )</pre> |  | Reiterate | 1 |
| 4 | <pre>.   .   .   P(a) ^ Q(a)</pre> |  | E ∀ | 3, 2 |
| 5 | <pre>.   .   .   P(a)</pre> |  | E ^ (L) | 4 |
| 6 | <pre>.   .   ∀y.( P(y) )</pre> |  | subproof ∀ introduction | 2 - 5 |
| 7 | <pre>.   .   .   </pre> | a | Term var |  |
| 8 | <pre>.   .   .   ∀x.( P(x) ^ Q(x) )</pre> |  | Reiterate | 1 |
| 9 | <pre>.   .   .   P(a) ^ Q(a)</pre> |  | E ∀ | 8, 7 |
| 10 | <pre>.   .   .   Q(a)</pre> |  | E ^ (R) | 9 |
| 11 | <pre>.   .   ∀z.( Q(z) )</pre> |  | subproof ∀ introduction | 7 - 10 |
| 12 | <pre>.   .   ∀y.( P(y) ) ^ ∀z.( Q(z) )</pre> |  | I ^ | 6, 11 |
| 13 | <pre>.   ∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )</pre> |  | subproof implication | 1 - 12 |

