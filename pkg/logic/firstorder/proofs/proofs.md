
# Table of Contents

1. [basics](#basics)
    1. [∀x.( P(x) v ~ P(x) )](#proof-1-1)
    2. [∀x.( ~ ( P(x) ^ ~ P(x) ) )](#proof-1-2)
    3. [( P ^ ~ P ) -> Q](#proof-1-3)
2. [quantifiers](#quantifiers)
    1. [( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )](#proof-2-1)
    2. [∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R](#proof-2-2)
    3. [( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R](#proof-2-3)
    4. [∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )](#proof-2-4)
    5. [∃x.( R ) <-> R](#proof-2-5)
    6. [( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) )](#proof-2-6)
3. [DeMorgan's](#DeMorgan's)
    1. [∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) )](#proof-3-1)
    2. [~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) )](#proof-3-2)
4. [distributive](#distributive)
    1. [∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )](#proof-4-1)
    2. [∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )](#proof-4-2)

# basics <a name="basics"></a>

## ∀x.( P(x) v ~ P(x) ) <a name="proof-1-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 | a | <pre>.   </pre> | term var: ∀ introduction |  |
| 2 | a | <pre>.   .   ~ ( P(a) v ~ P(a) )</pre> | Assume: contradiction |  |
| 3 | a | <pre>.   .   .   P(a)</pre> | Assume: contradiction |  |
| 4 | a | <pre>.   .   .   P(a) v ~ P(a)</pre> | I v (L) | 3 |
| 5 | a | <pre>.   .   .   ~ ( P(a) v ~ P(a) )</pre> | Reiterate | 2 |
| 6 | a | <pre>.   .   ~ P(a)</pre> | subproof contradiction | 3 - 5 |
| 7 | a | <pre>.   .   P(a) v ~ P(a)</pre> | I v (R) | 6 |
| 8 | a | <pre>.   P(a) v ~ P(a)</pre> | subproof contradiction | 2 - 7 |
| 9 |  | <pre>∀x.( P(x) v ~ P(x) )</pre> | subproof ∀ introduction | 1 - 8 |

## ∀x.( ~ ( P(x) ^ ~ P(x) ) ) <a name="proof-1-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 | a | <pre>.   </pre> | term var: ∀ introduction |  |
| 2 | a | <pre>.   .   P(a) ^ ~ P(a)</pre> | Assume: contradiction |  |
| 3 | a | <pre>.   .   P(a)</pre> | E ^ (L) | 2 |
| 4 | a | <pre>.   .   ~ P(a)</pre> | E ^ (R) | 2 |
| 5 | a | <pre>.   ~ ( P(a) ^ ~ P(a) )</pre> | subproof contradiction | 2 - 4 |
| 6 |  | <pre>∀x.( ~ ( P(x) ^ ~ P(x) ) )</pre> | subproof ∀ introduction | 1 - 5 |

## ( P ^ ~ P ) -> Q <a name="proof-1-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>~ ( P ^ ~ P )</pre> | Theorem: non-contradiction |  |
| 2 |  | <pre>.   ~ Q</pre> | Assume: implication |  |
| 3 |  | <pre>.   ~ ( P ^ ~ P )</pre> | Reiterate | 1 |
| 4 |  | <pre>~ Q -> ~ ( P ^ ~ P )</pre> | subproof implication | 2 - 3 |
| 5 |  | <pre>( P ^ ~ P ) -> Q</pre> | Theorem: contrapositive | 4 |

# quantifiers <a name="quantifiers"></a>

## ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) ) <a name="proof-2-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) )</pre> | Assume: implication |  |
| 2 |  | <pre>.   ∃x.( T )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   P -> ∃x.( Q(x) )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 5 | a | <pre>.   .   T</pre> | Assume: ∃ elimination [x -> a] | 2 |
| 6 | a | <pre>.   .   .   P</pre> | Assume: implication |  |
| 7 | a | <pre>.   .   .   P -> ∃x.( Q(x) )</pre> | Reiterate | 3 |
| 8 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | E -> | 7, 6 |
| 9 | a, b | <pre>.   .   .   .   </pre> | term var: ∃ elimination |  |
| 10 | a, b | <pre>.   .   .   .   Q(b)</pre> | Assume: ∃ elimination [x -> b] | 8 |
| 11 | a, b | <pre>.   .   .   .   P</pre> | Reiterate | 6 |
| 12 | a, b | <pre>.   .   .   .   P -> Q(b)</pre> | I -> | 11, 10 |
| 13 | a, b | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> | I ∃ [b -> x] | 12 |
| 14 | a | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> | subproof ∃ elimination | 9 - 13 |
| 15 | a | <pre>.   .   P -> ∃x.( P -> Q(x) )</pre> | subproof implication | 6 - 14 |
| 16 | a | <pre>.   .   .   ~ P</pre> | Assume: implication |  |
| 17 | a | <pre>.   .   .   .   ~ Q(a)</pre> | Assume: implication |  |
| 18 | a | <pre>.   .   .   .   ~ P</pre> | Reiterate | 16 |
| 19 | a | <pre>.   .   .   ~ Q(a) -> ~ P</pre> | subproof implication | 17 - 18 |
| 20 | a | <pre>.   .   .   P -> Q(a)</pre> | Theorem: contrapositive | 19 |
| 21 | a | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> | I ∃ [a -> x] | 20 |
| 22 | a | <pre>.   .   ~ P -> ∃x.( P -> Q(x) )</pre> | subproof implication | 16 - 21 |
| 23 | a | <pre>.   .   P v ~ P</pre> | Theorem: excluded middle |  |
| 24 | a | <pre>.   .   ∃x.( P -> Q(x) )</pre> | E v | 15, 22, 23 |
| 25 |  | <pre>.   ∃x.( P -> Q(x) )</pre> | subproof ∃ elimination | 4 - 24 |
| 26 |  | <pre>( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) -> ∃x.( P -> Q(x) )</pre> | subproof implication | 1 - 25 |
| 27 |  | <pre>.   ∃x.( P -> Q(x) )</pre> | Assume: implication |  |
| 28 |  | <pre>.   .   P</pre> | Assume: implication |  |
| 29 |  | <pre>.   .   ∃x.( P -> Q(x) )</pre> | Reiterate | 27 |
| 30 | a | <pre>.   .   .   </pre> | term var: ∃ elimination |  |
| 31 | a | <pre>.   .   .   P -> Q(a)</pre> | Assume: ∃ elimination [x -> a] | 29 |
| 32 | a | <pre>.   .   .   P</pre> | Reiterate | 28 |
| 33 | a | <pre>.   .   .   Q(a)</pre> | E -> | 31, 32 |
| 34 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 33 |
| 35 |  | <pre>.   .   ∃x.( Q(x) )</pre> | subproof ∃ elimination | 30 - 34 |
| 36 |  | <pre>.   P -> ∃x.( Q(x) )</pre> | subproof implication | 28 - 35 |
| 37 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 38 | a | <pre>.   .   P -> Q(a)</pre> | Assume: ∃ elimination [x -> a] | 27 |
| 39 | a | <pre>.   .   T</pre> | Reiterate | 0 |
| 40 | a | <pre>.   .   ∃x.( T )</pre> | I ∃ [a -> x] | 39 |
| 41 |  | <pre>.   ∃x.( T )</pre> | subproof ∃ elimination | 37 - 40 |
| 42 |  | <pre>.   ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) )</pre> | I ^ | 41, 36 |
| 43 |  | <pre>∃x.( P -> Q(x) ) -> ( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) )</pre> | subproof implication | 27 - 42 |
| 44 |  | <pre>( ∃x.( T ) ^ ( P -> ∃x.( Q(x) ) ) ) <-> ∃x.( P -> Q(x) )</pre> | I <-> | 26, 43 |

## ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R <a name="proof-2-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( Q(x) ^ ( Q(x) -> R ) )</pre> | Assume: implication |  |
| 2 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 3 | a | <pre>.   .   Q(a) ^ ( Q(a) -> R )</pre> | Assume: ∃ elimination [x -> a] | 1 |
| 4 | a | <pre>.   .   Q(a)</pre> | E ^ (L) | 3 |
| 5 | a | <pre>.   .   Q(a) -> R</pre> | E ^ (R) | 3 |
| 6 | a | <pre>.   .   R</pre> | E -> | 5, 4 |
| 7 |  | <pre>.   R</pre> | subproof ∃ elimination | 2 - 6 |
| 8 |  | <pre>∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R</pre> | subproof implication | 1 - 7 |

## ( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R <a name="proof-2-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R )</pre> | Assume: implication |  |
| 2 |  | <pre>.   ∀y.( Q(y) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ∃x.( Q(x) -> R )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 5 | a | <pre>.   .   Q(a) -> R</pre> | Assume: ∃ elimination [x -> a] | 3 |
| 6 | a | <pre>.   .   ∀y.( Q(y) )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   Q(a)</pre> | E ∀ [y -> a] | 6, 4 |
| 8 | a | <pre>.   .   R</pre> | E -> | 5, 7 |
| 9 |  | <pre>.   R</pre> | subproof ∃ elimination | 4 - 8 |
| 10 |  | <pre>( ∀y.( Q(y) ) ^ ∃x.( Q(x) -> R ) ) -> R</pre> | subproof implication | 1 - 9 |

## ∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) ) <a name="proof-2-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( P(x) ^ Q(x) )</pre> | Assume: implication |  |
| 2 | a | <pre>.   .   </pre> | term var: ∀ introduction |  |
| 3 | a | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 4 | a | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ [x -> a] | 3, 2 |
| 5 | a | <pre>.   .   P(a)</pre> | E ^ (L) | 4 |
| 6 |  | <pre>.   ∀y.( P(y) )</pre> | subproof ∀ introduction | 2 - 5 |
| 7 | a | <pre>.   .   </pre> | term var: ∀ introduction |  |
| 8 | a | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 9 | a | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ [x -> a] | 8, 7 |
| 10 | a | <pre>.   .   Q(a)</pre> | E ^ (R) | 9 |
| 11 |  | <pre>.   ∀z.( Q(z) )</pre> | subproof ∀ introduction | 7 - 10 |
| 12 |  | <pre>.   ∀y.( P(y) ) ^ ∀z.( Q(z) )</pre> | I ^ | 6, 11 |
| 13 |  | <pre>∀x.( P(x) ^ Q(x) ) -> ( ∀y.( P(y) ) ^ ∀z.( Q(z) ) )</pre> | subproof implication | 1 - 12 |

## ∃x.( R ) <-> R <a name="proof-2-5"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( R )</pre> | Assume: implication |  |
| 2 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 3 | a | <pre>.   .   R</pre> | Assume: ∃ elimination [x -> a] | 1 |
| 4 |  | <pre>.   R</pre> | subproof ∃ elimination | 2 - 3 |
| 5 |  | <pre>∃x.( R ) -> R</pre> | subproof implication | 1 - 4 |
| 6 |  | <pre>.   R</pre> | Assume: implication |  |
| 7 |  | <pre>.   ∃x.( R )</pre> | I ∃ [a -> x] | 6 |
| 8 |  | <pre>R -> ∃x.( R )</pre> | subproof implication | 6 - 7 |
| 9 |  | <pre>∃x.( R ) <-> R</pre> | I <-> | 5, 8 |

## ( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) ) <a name="proof-2-6"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( Q(x) ) ^ ∃x.( T )</pre> | Assume: implication |  |
| 2 |  | <pre>.   ∀x.( Q(x) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ∃x.( T )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 5 | a | <pre>.   .   T</pre> | Assume: ∃ elimination [x -> a] | 3 |
| 6 | a | <pre>.   .   ∀x.( Q(x) )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   Q(a)</pre> | E ∀ [x -> a] | 6, 4 |
| 8 | a | <pre>.   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 7 |
| 9 |  | <pre>.   ∃x.( Q(x) )</pre> | subproof ∃ elimination | 4 - 8 |
| 10 |  | <pre>( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) )</pre> | subproof implication | 1 - 9 |

# DeMorgan's <a name="DeMorgan's"></a>

## ∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) ) <a name="proof-3-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( ~ Q(x) )</pre> | Assume: implication |  |
| 2 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 3 | a | <pre>.   .   ∃x.( Q(x) )</pre> | Assume: ∃ contradiction |  |
| 4 | a | <pre>.   .   Q(a)</pre> | Assume: ∃ contradiction [x -> a] |  |
| 5 | a | <pre>.   .   ∀x.( ~ Q(x) )</pre> | Reiterate | 1 |
| 6 | a | <pre>.   .   ~ Q(a)</pre> | E ∀ [x -> a] | 5, 2 |
| 7 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | subproof ∃ elimination | 2 - 6 |
| 8 |  | <pre>∀x.( ~ Q(x) ) -> ~ ∃x.( Q(x) )</pre> | subproof implication | 1 - 7 |
| 9 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | Assume: implication |  |
| 10 | a | <pre>.   .   </pre> | term var: ∀ introduction |  |
| 11 | a | <pre>.   .   .   Q(a)</pre> | Assume: contradiction |  |
| 12 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 11 |
| 13 | a | <pre>.   .   .   ~ ∃x.( Q(x) )</pre> | Reiterate | 9 |
| 14 | a | <pre>.   .   ~ Q(a)</pre> | subproof contradiction | 11 - 13 |
| 15 |  | <pre>.   ∀x.( ~ Q(x) )</pre> | subproof ∀ introduction | 10 - 14 |
| 16 |  | <pre>~ ∃x.( Q(x) ) -> ∀x.( ~ Q(x) )</pre> | subproof implication | 9 - 15 |
| 17 |  | <pre>∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) )</pre> | I <-> | 8, 16 |

## ~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) ) <a name="proof-3-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ∀x.( Q(x) )</pre> | Assume: implication |  |
| 2 |  | <pre>.   .   ~ ∃x.( ~ Q(x) )</pre> | Assume: contradiction |  |
| 3 | a | <pre>.   .   .   </pre> | term var: ∀ introduction |  |
| 4 | a | <pre>.   .   .   .   ~ Q(a)</pre> | Assume: contradiction |  |
| 5 | a | <pre>.   .   .   .   ∃x.( ~ Q(x) )</pre> | I ∃ [a -> x] | 4 |
| 6 | a | <pre>.   .   .   .   ~ ∃x.( ~ Q(x) )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   .   Q(a)</pre> | subproof contradiction | 4 - 6 |
| 8 |  | <pre>.   .   ∀x.( Q(x) )</pre> | subproof ∀ introduction | 3 - 7 |
| 9 |  | <pre>.   .   ~ ∀x.( Q(x) )</pre> | Reiterate | 1 |
| 10 |  | <pre>.   ∃x.( ~ Q(x) )</pre> | subproof contradiction | 2 - 9 |
| 11 |  | <pre>~ ∀x.( Q(x) ) -> ∃x.( ~ Q(x) )</pre> | subproof implication | 1 - 10 |
| 12 |  | <pre>.   ∃x.( ~ Q(x) )</pre> | Assume: implication |  |
| 13 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 14 | a | <pre>.   .   ~ Q(a)</pre> | Assume: ∃ elimination [x -> a] | 12 |
| 15 | a | <pre>.   .   .   ∀x.( Q(x) )</pre> | Assume: contradiction |  |
| 16 | a | <pre>.   .   .   Q(a)</pre> | E ∀ [x -> a] | 15, 13 |
| 17 | a | <pre>.   .   .   ~ Q(a)</pre> | Reiterate | 14 |
| 18 | a | <pre>.   .   ~ ∀x.( Q(x) )</pre> | subproof contradiction | 15 - 17 |
| 19 |  | <pre>.   ~ ∀x.( Q(x) )</pre> | subproof ∃ elimination | 13 - 18 |
| 20 |  | <pre>∃x.( ~ Q(x) ) -> ~ ∀x.( Q(x) )</pre> | subproof implication | 12 - 19 |
| 21 |  | <pre>~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) )</pre> | I <-> | 11, 20 |

# distributive <a name="distributive"></a>

## ∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) ) <a name="proof-4-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( P(x) ^ Q(x) )</pre> | Assume: implication |  |
| 2 | a | <pre>.   .   </pre> | term var: ∀ introduction |  |
| 3 | a | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 4 | a | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ [x -> a] | 3, 2 |
| 5 | a | <pre>.   .   P(a)</pre> | E ^ (L) | 4 |
| 6 |  | <pre>.   ∀x.( P(x) )</pre> | subproof ∀ introduction | 2 - 5 |
| 7 | a | <pre>.   .   </pre> | term var: ∀ introduction |  |
| 8 | a | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 9 | a | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ [x -> a] | 8, 7 |
| 10 | a | <pre>.   .   Q(a)</pre> | E ^ (R) | 9 |
| 11 |  | <pre>.   ∀x.( Q(x) )</pre> | subproof ∀ introduction | 7 - 10 |
| 12 |  | <pre>.   ∀x.( P(x) ) ^ ∀x.( Q(x) )</pre> | I ^ | 6, 11 |
| 13 |  | <pre>∀x.( P(x) ^ Q(x) ) -> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )</pre> | subproof implication | 1 - 12 |
| 14 |  | <pre>.   ∀x.( P(x) ) ^ ∀x.( Q(x) )</pre> | Assume: implication |  |
| 15 | a | <pre>.   .   </pre> | term var: ∀ introduction |  |
| 16 | a | <pre>.   .   ∀x.( P(x) ) ^ ∀x.( Q(x) )</pre> | Reiterate | 14 |
| 17 | a | <pre>.   .   ∀x.( P(x) )</pre> | E ^ (L) | 16 |
| 18 | a | <pre>.   .   ∀x.( Q(x) )</pre> | E ^ (R) | 16 |
| 19 | a | <pre>.   .   P(a)</pre> | E ∀ [x -> a] | 17, 15 |
| 20 | a | <pre>.   .   Q(a)</pre> | E ∀ [x -> a] | 18, 15 |
| 21 | a | <pre>.   .   P(a) ^ Q(a)</pre> | I ^ | 19, 20 |
| 22 |  | <pre>.   ∀x.( P(x) ^ Q(x) )</pre> | subproof ∀ introduction | 15 - 21 |
| 23 |  | <pre>( ∀x.( P(x) ) ^ ∀x.( Q(x) ) ) -> ∀x.( P(x) ^ Q(x) )</pre> | subproof implication | 14 - 22 |
| 24 |  | <pre>∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )</pre> | I <-> | 13, 23 |

## ∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) ) <a name="proof-4-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( P(x) v Q(x) )</pre> | Assume: implication |  |
| 2 | a | <pre>.   .   </pre> | term var: ∃ elimination |  |
| 3 | a | <pre>.   .   P(a) v Q(a)</pre> | Assume: ∃ elimination [x -> a] | 1 |
| 4 | a | <pre>.   .   .   P(a)</pre> | Assume: implication |  |
| 5 | a | <pre>.   .   .   ∃x.( P(x) )</pre> | I ∃ [a -> x] | 4 |
| 6 | a | <pre>.   .   .   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | I v (L) | 5 |
| 7 | a | <pre>.   .   P(a) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | subproof implication | 4 - 6 |
| 8 | a | <pre>.   .   .   Q(a)</pre> | Assume: implication |  |
| 9 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 8 |
| 10 | a | <pre>.   .   .   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | I v (R) | 9 |
| 11 | a | <pre>.   .   Q(a) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | subproof implication | 8 - 10 |
| 12 | a | <pre>.   .   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | E v | 7, 11, 3 |
| 13 |  | <pre>.   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | subproof ∃ elimination | 2 - 12 |
| 14 |  | <pre>∃x.( P(x) v Q(x) ) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | subproof implication | 1 - 13 |
| 15 |  | <pre>.   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | Assume: implication |  |
| 16 |  | <pre>.   .   ∃x.( P(x) )</pre> | Assume: implication |  |
| 17 | a | <pre>.   .   .   </pre> | term var: ∃ elimination |  |
| 18 | a | <pre>.   .   .   P(a)</pre> | Assume: ∃ elimination [x -> a] | 16 |
| 19 | a | <pre>.   .   .   P(a) v Q(a)</pre> | I v (L) | 18 |
| 20 | a | <pre>.   .   .   ∃x.( P(x) v Q(x) )</pre> | I ∃ [a -> x] | 19 |
| 21 |  | <pre>.   .   ∃x.( P(x) v Q(x) )</pre> | subproof ∃ elimination | 17 - 20 |
| 22 |  | <pre>.   ∃x.( P(x) ) -> ∃x.( P(x) v Q(x) )</pre> | subproof implication | 16 - 21 |
| 23 |  | <pre>.   .   ∃x.( Q(x) )</pre> | Assume: implication |  |
| 24 | a | <pre>.   .   .   </pre> | term var: ∃ elimination |  |
| 25 | a | <pre>.   .   .   Q(a)</pre> | Assume: ∃ elimination [x -> a] | 23 |
| 26 | a | <pre>.   .   .   P(a) v Q(a)</pre> | I v (R) | 25 |
| 27 | a | <pre>.   .   .   ∃x.( P(x) v Q(x) )</pre> | I ∃ [a -> x] | 26 |
| 28 |  | <pre>.   .   ∃x.( P(x) v Q(x) )</pre> | subproof ∃ elimination | 24 - 27 |
| 29 |  | <pre>.   ∃x.( Q(x) ) -> ∃x.( P(x) v Q(x) )</pre> | subproof implication | 23 - 28 |
| 30 |  | <pre>.   ∃x.( P(x) v Q(x) )</pre> | E v | 22, 29, 15 |
| 31 |  | <pre>( ∃x.( P(x) ) v ∃x.( Q(x) ) ) -> ∃x.( P(x) v Q(x) )</pre> | subproof implication | 15 - 30 |
| 32 |  | <pre>∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | I <-> | 14, 31 |

