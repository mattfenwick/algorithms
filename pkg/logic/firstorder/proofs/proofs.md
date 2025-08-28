
# Table of Contents

1. [theorems](#theorems)
    1. [P v ~ P](#proof-1-1)

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

