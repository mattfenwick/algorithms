
# Table of Contents

1. [basics](#basics)
    1. [P v ~ P](#proof-1-1)
    2. [∀x.( Q(x) ) -> Q(a)](#proof-1-2)
    3. [Q(a) -> ∃x.( Q(x) )](#proof-1-3)

# basics <a name="basics"></a>

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

## ∀x.( Q(x) ) -> Q(a) <a name="proof-1-2"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   ∀x.( Q(x) )</pre> | Assume Implication |  |
| 2 | <pre>.   Q(a)</pre> | E ∀ | 1 |
| 3 | <pre>∀x.( Q(x) ) -> Q(a)</pre> | subproof implication | 1 - 2 |

## Q(a) -> ∃x.( Q(x) ) <a name="proof-1-3"></a>

| Line | Formula | Justification | Lines used |
| - | - | - | - |
| 1 | <pre>.   Q(a)</pre> | Assume Implication |  |
| 2 | <pre>.   ∃x.( Q(x) )</pre> | I ∃ | 1 |
| 3 | <pre>Q(a) -> ∃x.( Q(x) )</pre> | subproof implication | 1 - 2 |

