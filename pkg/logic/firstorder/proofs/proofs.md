
# Table of Contents

1. [basics](#basics)
    1. [P v ~ P](#proof-1-1)
    2. [∀x.( Q(x) ) -> Q(a)](#proof-1-2)
    3. [Q(a) -> ∃x.( Q(x) )](#proof-1-3)
    4. [∃x.( P -> Q(x) ) -> ( P -> ∃y.( Q(y) ) )](#proof-1-4)
    5. [∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R](#proof-1-5)
    6. [( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R](#proof-1-6)
    7. [∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )](#proof-1-7)

# basics <a name="basics"></a>

## P v ~ P <a name="proof-1-1"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ~ ( P v ~ P )</pre> | Assume: contradiction |  |
| 2 | <pre>.   .   P</pre> | Assume: contradiction |  |
| 3 | <pre>.   .   P v ~ P</pre> | I v (L) | 2 |
| 4 | <pre>.   .   ~ ( P v ~ P )</pre> | Reiterate | 1 |
| 5 | <pre>.   ~ P</pre> | subproof contradiction | 2 - 4 |
| 6 | <pre>.   P v ~ P</pre> | I v (R) | 5 |
| 7 | <pre>P v ~ P</pre> | subproof contradiction | 1 - 6 |

## ∀x.( Q(x) ) -> Q(a) <a name="proof-1-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ∀x.( Q(x) )</pre> | Assume: implication |  |
| 2 | <pre>.   Q(a)</pre> | E ∀ | 1 |
| 3 | <pre>∀x.( Q(x) ) -> Q(a)</pre> | subproof implication | 1 - 2 |

## Q(a) -> ∃x.( Q(x) ) <a name="proof-1-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   Q(a)</pre> | Assume: implication |  |
| 2 | <pre>.   ∃x.( Q(x) )</pre> | I ∃ | 1 |
| 3 | <pre>Q(a) -> ∃x.( Q(x) )</pre> | subproof implication | 1 - 2 |

## ∃x.( P -> Q(x) ) -> ( P -> ∃y.( Q(y) ) ) <a name="proof-1-4"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ∃x.( P -> Q(x) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   P</pre> | Assume: implication |  |
| 3 | <pre>.   .   ∃x.( P -> Q(x) )</pre> | Reiterate | 1 |
| 4 | <pre>.   .   .   P -> Q(a)</pre> | Assume: ∃ elimination | 3 |
| 5 | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 6 | <pre>.   .   .   Q(a)</pre> | E -> | 4, 5 |
| 7 | <pre>.   .   .   ∃y.( Q(y) )</pre> | I ∃ | 6 |
| 8 | <pre>.   .   ∃y.( Q(y) )</pre> | subproof ∃ elimination | 4 - 7 |
| 9 | <pre>.   P -> ∃y.( Q(y) )</pre> | subproof implication | 2 - 8 |
| 10 | <pre>∃x.( P -> Q(x) ) -> ( P -> ∃y.( Q(y) ) )</pre> | subproof implication | 1 - 9 |

## ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R <a name="proof-1-5"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ∃x.( Q(x) ^ ( Q(x) -> R ) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   Q(a) ^ ( Q(a) -> R )</pre> | Assume: ∃ elimination | 1 |
| 3 | <pre>.   .   Q(a)</pre> | E ^ (L) | 2 |
| 4 | <pre>.   .   Q(a) -> R</pre> | E ^ (R) | 2 |
| 5 | <pre>.   .   R</pre> | E -> | 4, 3 |
| 6 | <pre>.   R</pre> | subproof ∃ elimination | 2 - 5 |
| 7 | <pre>∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R</pre> | subproof implication | 1 - 6 |

## ( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R <a name="proof-1-6"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R )</pre> | Assume: implication |  |
| 2 | <pre>.   ∀y.( Q(y) )</pre> | E ^ (L) | 1 |
| 3 | <pre>.   ∃x.( Q(x) -> R )</pre> | E ^ (R) | 1 |
| 4 | <pre>.   .   Q(a) -> R</pre> | Assume: ∃ elimination | 3 |
| 5 | <pre>.   .   ∀y.( Q(y) )</pre> | Reiterate | 2 |
| 6 | <pre>.   .   Q(a)</pre> | E ∀ | 5 |
| 7 | <pre>.   .   R</pre> | E -> | 4, 6 |
| 8 | <pre>.   R</pre> | subproof ∃ elimination | 4 - 7 |
| 9 | <pre>( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R</pre> | subproof implication | 1 - 8 |

## ∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) ) <a name="proof-1-7"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ∀x.( P(x) ^ Q(x) )</pre> | Assume: implication |  |
| 2 | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 3 | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ | 2 |
| 4 | <pre>.   .   P(a)</pre> | E ^ (L) | 3 |
| 5 | <pre>.   ∀y.( P(y) )</pre> | subproof ∃ elimination | 2 - 4 |
| 6 | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 7 | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ | 6 |
| 8 | <pre>.   .   Q(a)</pre> | E ^ (R) | 7 |
| 9 | <pre>.   ∀z.( Q(z) )</pre> | subproof ∃ elimination | 6 - 8 |
| 10 | <pre>.   ∀y.( P(y) ) ^ ∀z.( Q(z) )</pre> | I ^ | 5, 9 |
| 11 | <pre>∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )</pre> | subproof implication | 1 - 10 |

