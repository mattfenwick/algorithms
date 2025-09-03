
# Table of Contents

1. [basics](#basics)
    1. [∀x.( P(x) v ~ P(x) )](#proof-1-1)
    2. [~ ∃x.( ~ ( P(x) v ~ P(x) ) )](#proof-1-2)
    3. [∀x.( ~ ( P(x) ^ ~ P(x) ) )](#proof-1-3)
    4. [~ ∃x.( P(x) ^ ~ P(x) )](#proof-1-4)
2. [DeMorgan's](#DeMorgan's)
    1. [∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) )](#proof-2-1)
    2. [~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) )](#proof-2-2)
3. [exist-or-empty](#exist-or-empty)
    1. [( ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) ) ) <-> ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) )](#proof-3-1)
    2. [( ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) ) ) <-> ( ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) ) )](#proof-3-2)
    3. [( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) ) <-> ∃x.( T )](#proof-3-3)
    4. [( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) <-> ~ ∃x.( T )](#proof-3-4)
4. [distributive-and](#distributive-and)
    1. [( P ^ ∃x.( Q(x) ) ) <-> ∃x.( P ^ Q(x) )](#proof-4-1)
    2. [∃x.( T ) -> ( ( P ^ ∀x.( Q(x) ) ) <-> ∀x.( P ^ Q(x) ) )](#proof-4-2)
    3. [∃x.( P(x) ^ Q(x) ) -> ( ∃x.( P(x) ) ^ ∃x.( Q(x) ) )](#proof-4-3)
    4. [∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )](#proof-4-4)
5. [distributive-or](#distributive-or)
    1. [∃x.( T ) -> ( ( P v ∃x.( Q(x) ) ) <-> ∃x.( P v Q(x) ) )](#proof-5-1)
    2. [( P v ∀x.( Q(x) ) ) <-> ∀x.( P v Q(x) )](#proof-5-2)
    3. [∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )](#proof-5-3)
    4. [( ∀x.( P(x) ) v ∀x.( Q(x) ) ) -> ∀x.( P(x) v Q(x) )](#proof-5-4)
6. [distributive-arrow](#distributive-arrow)
    1. [∃x.( T ) -> ( ( P -> ∃x.( Q(x) ) ) <-> ∃x.( P -> Q(x) ) )](#proof-6-1)
    2. [( P -> ∀x.( Q(x) ) ) <-> ∀x.( P -> Q(x) )](#proof-6-2)
    3. [( ∃x.( P(x) ) -> ∃x.( Q(x) ) ) -> ( ∃x.( T ) -> ∃x.( P(x) -> Q(x) ) )](#proof-6-3)
    4. [∀x.( P(x) -> Q(x) ) -> ( ∀x.( P(x) ) -> ∀x.( Q(x) ) )](#proof-6-4)
7. [commutativity](#commutativity)
    1. [∃x.( ∃y.( P(x,y) ) ) -> ∃y.( ∃x.( P(x,y) ) )](#proof-7-1)
    2. [∀x.( ∀y.( P(x,y) ) ) -> ∀y.( ∀x.( P(x,y) ) )](#proof-7-2)
    3. [∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) )](#proof-7-3)
    4. [∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) )](#proof-7-4)
8. [random](#random)
    1. [P -> ( ∀x.( ~ P ) -> ~ ∃x.( T ) )](#proof-8-1)
    2. [P -> ( ∃x.( T ) -> ~ ∀x.( ~ P ) )](#proof-8-2)
    3. [∃x.( T ) -> ( ∀x.( ~ P ) -> ~ P )](#proof-8-3)
    4. [∃x.( R ) -> R](#proof-8-4)
    5. [R -> ∀x.( R )](#proof-8-5)
    6. [∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R](#proof-8-6)
    7. [( ∀x.( Q(x) ) ^ ∃x.( Q(x) -> R ) ) -> R](#proof-8-7)
    8. [( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) )](#proof-8-8)

# basics <a name="basics"></a>

## ∀x.( P(x) v ~ P(x) ) <a name="proof-1-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 | a | <pre>.   </pre> | term var: I ∀ |  |
| 2 | a | <pre>.   .   ~ ( P(a) v ~ P(a) )</pre> | Assume: ~ |  |
| 3 | a | <pre>.   .   .   P(a)</pre> | Assume: ~ |  |
| 4 | a | <pre>.   .   .   P(a) v ~ P(a)</pre> | I v (L) | 3 |
| 5 | a | <pre>.   .   .   ~ ( P(a) v ~ P(a) )</pre> | Reiterate | 2 |
| 6 | a | <pre>.   .   ~ P(a)</pre> | subproof ~ | 3 - 5 |
| 7 | a | <pre>.   .   P(a) v ~ P(a)</pre> | I v (R) | 6 |
| 8 | a | <pre>.   P(a) v ~ P(a)</pre> | subproof ~ | 2 - 7 |
| 9 |  | <pre>∀x.( P(x) v ~ P(x) )</pre> | subproof I ∀ [a -> x] | 1 - 8 |

## ~ ∃x.( ~ ( P(x) v ~ P(x) ) ) <a name="proof-1-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 | a | <pre>.   </pre> | term var: ~ ∃ |  |
| 2 | a | <pre>.   ∃x.( ~ ( P(x) v ~ P(x) ) )</pre> | Assume: ~ ∃ |  |
| 3 | a | <pre>.   ~ ( P(a) v ~ P(a) )</pre> | Assume: ~ ∃ [x -> a] |  |
| 4 | a | <pre>.   .   P(a)</pre> | Assume: ~ |  |
| 5 | a | <pre>.   .   P(a) v ~ P(a)</pre> | I v (L) | 4 |
| 6 | a | <pre>.   .   ~ ( P(a) v ~ P(a) )</pre> | Reiterate | 3 |
| 7 | a | <pre>.   ~ P(a)</pre> | subproof ~ | 4 - 6 |
| 8 | a | <pre>.   P(a) v ~ P(a)</pre> | I v (R) | 7 |
| 9 |  | <pre>~ ∃x.( ~ ( P(x) v ~ P(x) ) )</pre> | subproof ~ ∃ | 1 - 8 |

## ∀x.( ~ ( P(x) ^ ~ P(x) ) ) <a name="proof-1-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 | a | <pre>.   </pre> | term var: I ∀ |  |
| 2 | a | <pre>.   .   P(a) ^ ~ P(a)</pre> | Assume: ~ |  |
| 3 | a | <pre>.   .   P(a)</pre> | E ^ (L) | 2 |
| 4 | a | <pre>.   .   ~ P(a)</pre> | E ^ (R) | 2 |
| 5 | a | <pre>.   ~ ( P(a) ^ ~ P(a) )</pre> | subproof ~ | 2 - 4 |
| 6 |  | <pre>∀x.( ~ ( P(x) ^ ~ P(x) ) )</pre> | subproof I ∀ [a -> x] | 1 - 5 |

## ~ ∃x.( P(x) ^ ~ P(x) ) <a name="proof-1-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 | a | <pre>.   </pre> | term var: ~ ∃ |  |
| 2 | a | <pre>.   ∃x.( P(x) ^ ~ P(x) )</pre> | Assume: ~ ∃ |  |
| 3 | a | <pre>.   P(a) ^ ~ P(a)</pre> | Assume: ~ ∃ [x -> a] |  |
| 4 | a | <pre>.   P(a)</pre> | E ^ (L) | 3 |
| 5 | a | <pre>.   ~ P(a)</pre> | E ^ (R) | 3 |
| 6 |  | <pre>~ ∃x.( P(x) ^ ~ P(x) )</pre> | subproof ~ ∃ | 1 - 5 |

# DeMorgan's <a name="DeMorgan's"></a>

## ∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) ) <a name="proof-2-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( ~ Q(x) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: ~ ∃ |  |
| 3 | a | <pre>.   .   ∃x.( Q(x) )</pre> | Assume: ~ ∃ |  |
| 4 | a | <pre>.   .   Q(a)</pre> | Assume: ~ ∃ [x -> a] |  |
| 5 | a | <pre>.   .   ∀x.( ~ Q(x) )</pre> | Reiterate | 1 |
| 6 | a | <pre>.   .   ~ Q(a)</pre> | E ∀ [x -> a] | 5, 2 |
| 7 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | subproof ~ ∃ | 2 - 6 |
| 8 |  | <pre>∀x.( ~ Q(x) ) -> ~ ∃x.( Q(x) )</pre> | subproof -> | 1 - 7 |
| 9 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | Assume: -> |  |
| 10 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 11 | a | <pre>.   .   .   Q(a)</pre> | Assume: ~ |  |
| 12 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 11, 10 |
| 13 | a | <pre>.   .   .   ~ ∃x.( Q(x) )</pre> | Reiterate | 9 |
| 14 | a | <pre>.   .   ~ Q(a)</pre> | subproof ~ | 11 - 13 |
| 15 |  | <pre>.   ∀x.( ~ Q(x) )</pre> | subproof I ∀ [a -> x] | 10 - 14 |
| 16 |  | <pre>~ ∃x.( Q(x) ) -> ∀x.( ~ Q(x) )</pre> | subproof -> | 9 - 15 |
| 17 |  | <pre>∀x.( ~ Q(x) ) <-> ~ ∃x.( Q(x) )</pre> | I <-> | 8, 16 |

## ~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) ) <a name="proof-2-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ∀x.( Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ~ ∃x.( ~ Q(x) )</pre> | Assume: ~ |  |
| 3 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 4 | a | <pre>.   .   .   .   ~ Q(a)</pre> | Assume: ~ |  |
| 5 | a | <pre>.   .   .   .   ∃x.( ~ Q(x) )</pre> | I ∃ [a -> x] | 4, 3 |
| 6 | a | <pre>.   .   .   .   ~ ∃x.( ~ Q(x) )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   .   Q(a)</pre> | subproof ~ | 4 - 6 |
| 8 |  | <pre>.   .   ∀x.( Q(x) )</pre> | subproof I ∀ [a -> x] | 3 - 7 |
| 9 |  | <pre>.   .   ~ ∀x.( Q(x) )</pre> | Reiterate | 1 |
| 10 |  | <pre>.   ∃x.( ~ Q(x) )</pre> | subproof ~ | 2 - 9 |
| 11 |  | <pre>~ ∀x.( Q(x) ) -> ∃x.( ~ Q(x) )</pre> | subproof -> | 1 - 10 |
| 12 |  | <pre>.   ∃x.( ~ Q(x) )</pre> | Assume: -> |  |
| 13 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 14 | a | <pre>.   .   ~ Q(a)</pre> | Assume: E ∃ [x -> a] | 12 |
| 15 | a | <pre>.   .   .   ∀x.( Q(x) )</pre> | Assume: ~ |  |
| 16 | a | <pre>.   .   .   Q(a)</pre> | E ∀ [x -> a] | 15, 13 |
| 17 | a | <pre>.   .   .   ~ Q(a)</pre> | Reiterate | 14 |
| 18 | a | <pre>.   .   ~ ∀x.( Q(x) )</pre> | subproof ~ | 15 - 17 |
| 19 |  | <pre>.   ~ ∀x.( Q(x) )</pre> | subproof E ∃ | 13 - 18 |
| 20 |  | <pre>∃x.( ~ Q(x) ) -> ~ ∀x.( Q(x) )</pre> | subproof -> | 12 - 19 |
| 21 |  | <pre>~ ∀x.( Q(x) ) <-> ∃x.( ~ Q(x) )</pre> | I <-> | 11, 20 |

# exist-or-empty <a name="exist-or-empty"></a>

## ( ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) ) ) <-> ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) <a name="proof-3-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   ∀x.( Q(x) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ∀x.( ~ Q(x) )</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   ~ ∃x.( ~ Q(x) )</pre> | DeMorgan's theorem: ∀ to (~ ∃) | 2 |
| 5 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | DeMorgan's theorem: ∀ to (~ ∃) | 3 |
| 6 |  | <pre>.   ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) )</pre> | I ^ | 5, 4 |
| 7 |  | <pre>( ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) ) ) -> ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) )</pre> | subproof -> | 1 - 6 |
| 8 |  | <pre>.   ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) )</pre> | Assume: -> |  |
| 9 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | E ^ (L) | 8 |
| 10 |  | <pre>.   ~ ∃x.( ~ Q(x) )</pre> | E ^ (R) | 8 |
| 11 |  | <pre>.   ∀x.( ~ Q(x) )</pre> | DeMorgan's theorem: (~ ∃) to ∀ | 9 |
| 12 |  | <pre>.   ∀x.( Q(x) )</pre> | DeMorgan's theorem: (~ ∃) to ∀ | 10 |
| 13 |  | <pre>.   ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) )</pre> | I ^ | 12, 11 |
| 14 |  | <pre>( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) -> ( ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) ) )</pre> | subproof -> | 8 - 13 |
| 15 |  | <pre>( ∀x.( Q(x) ) ^ ∀x.( ~ Q(x) ) ) <-> ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) )</pre> | I <-> | 7, 14 |

## ( ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) ) ) <-> ( ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) ) ) <a name="proof-3-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   ~ ∀x.( Q(x) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ ∀x.( ~ Q(x) )</pre> | E ^ (R) | 1 |
| 4 |  | <pre>.   ∃x.( ~ Q(x) )</pre> | DeMorgan's theorem: (~ ∀) to ∃ | 2 |
| 5 |  | <pre>.   ∃x.( Q(x) )</pre> | DeMorgan's theorem: (~ ∀) to ∃ | 3 |
| 6 |  | <pre>.   ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) )</pre> | I ^ | 5, 4 |
| 7 |  | <pre>( ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) ) ) -> ( ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) ) )</pre> | subproof -> | 1 - 6 |
| 8 |  | <pre>.   ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) )</pre> | Assume: -> |  |
| 9 |  | <pre>.   ∃x.( Q(x) )</pre> | E ^ (L) | 8 |
| 10 |  | <pre>.   ∃x.( ~ Q(x) )</pre> | E ^ (R) | 8 |
| 11 |  | <pre>.   ~ ∀x.( ~ Q(x) )</pre> | DeMorgan's theorem: ∃ to (~ ∀) | 9 |
| 12 |  | <pre>.   ~ ∀x.( Q(x) )</pre> | DeMorgan's theorem: ∃ to (~ ∀) | 10 |
| 13 |  | <pre>.   ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) )</pre> | I ^ | 12, 11 |
| 14 |  | <pre>( ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) ) ) -> ( ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) ) )</pre> | subproof -> | 8 - 13 |
| 15 |  | <pre>( ∃x.( Q(x) ) ^ ∃x.( ~ Q(x) ) ) <-> ( ~ ∀x.( Q(x) ) ^ ~ ∀x.( ~ Q(x) ) )</pre> | I <-> | 14, 7 |

## ( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) ) <-> ∃x.( T ) <a name="proof-3-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( Q(x) ) v ∃x.( ~ Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ∃x.( Q(x) )</pre> | Assume: -> |  |
| 3 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 4 | a | <pre>.   .   .   Q(a)</pre> | Assume: E ∃ [x -> a] | 2 |
| 5 | a | <pre>.   .   .   T</pre> | Reiterate | 0 |
| 6 | a | <pre>.   .   .   ∃x.( T )</pre> | I ∃ [a -> x] | 5, 3 |
| 7 |  | <pre>.   .   ∃x.( T )</pre> | subproof E ∃ | 3 - 6 |
| 8 |  | <pre>.   ∃x.( Q(x) ) -> ∃x.( T )</pre> | subproof -> | 2 - 7 |
| 9 |  | <pre>.   .   ∃x.( ~ Q(x) )</pre> | Assume: -> |  |
| 10 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 11 | a | <pre>.   .   .   ~ Q(a)</pre> | Assume: E ∃ [x -> a] | 9 |
| 12 | a | <pre>.   .   .   T</pre> | Reiterate | 0 |
| 13 | a | <pre>.   .   .   ∃x.( T )</pre> | I ∃ [a -> x] | 12, 10 |
| 14 |  | <pre>.   .   ∃x.( T )</pre> | subproof E ∃ | 10 - 13 |
| 15 |  | <pre>.   ∃x.( ~ Q(x) ) -> ∃x.( T )</pre> | subproof -> | 9 - 14 |
| 16 |  | <pre>.   ∃x.( T )</pre> | E v | 8, 15, 1 |
| 17 |  | <pre>( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) ) -> ∃x.( T )</pre> | subproof -> | 1 - 16 |
| 18 |  | <pre>.   ∃x.( T )</pre> | Assume: -> |  |
| 19 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 20 | a | <pre>.   .   T</pre> | Assume: E ∃ [x -> a] | 18 |
| 21 | a | <pre>.   .   Q(a) v ~ Q(a)</pre> | Theorem: excluded middle |  |
| 22 | a | <pre>.   .   .   Q(a)</pre> | Assume: -> |  |
| 23 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 22, 19 |
| 24 | a | <pre>.   .   .   ∃x.( Q(x) ) v ∃x.( ~ Q(x) )</pre> | I v (L) | 23 |
| 25 | a | <pre>.   .   Q(a) -> ( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) )</pre> | subproof -> | 22 - 24 |
| 26 | a | <pre>.   .   .   ~ Q(a)</pre> | Assume: -> |  |
| 27 | a | <pre>.   .   .   ∃x.( ~ Q(x) )</pre> | I ∃ [a -> x] | 26, 19 |
| 28 | a | <pre>.   .   .   ∃x.( Q(x) ) v ∃x.( ~ Q(x) )</pre> | I v (R) | 27 |
| 29 | a | <pre>.   .   ~ Q(a) -> ( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) )</pre> | subproof -> | 26 - 28 |
| 30 | a | <pre>.   .   ∃x.( Q(x) ) v ∃x.( ~ Q(x) )</pre> | E v | 25, 29, 21 |
| 31 |  | <pre>.   ∃x.( Q(x) ) v ∃x.( ~ Q(x) )</pre> | subproof E ∃ | 19 - 30 |
| 32 |  | <pre>∃x.( T ) -> ( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) )</pre> | subproof -> | 18 - 31 |
| 33 |  | <pre>( ∃x.( Q(x) ) v ∃x.( ~ Q(x) ) ) <-> ∃x.( T )</pre> | I <-> | 17, 32 |

## ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) <-> ~ ∃x.( T ) <a name="proof-3-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ~ ∃x.( ~ Q(x) )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: ~ ∃ |  |
| 5 | a | <pre>.   .   ∃x.( T )</pre> | Assume: ~ ∃ |  |
| 6 | a | <pre>.   .   T</pre> | Assume: ~ ∃ [x -> a] |  |
| 7 | a | <pre>.   .   .   Q(a)</pre> | Assume: ~ |  |
| 8 | a | <pre>.   .   .   ~ ∃x.( Q(x) )</pre> | Reiterate | 2 |
| 9 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 7, 4 |
| 10 | a | <pre>.   .   ~ Q(a)</pre> | subproof ~ | 7 - 9 |
| 11 | a | <pre>.   .   ~ ∃x.( ~ Q(x) )</pre> | Reiterate | 3 |
| 12 | a | <pre>.   .   ∃x.( ~ Q(x) )</pre> | I ∃ [a -> x] | 10, 4 |
| 13 |  | <pre>.   ~ ∃x.( T )</pre> | subproof ~ ∃ | 4 - 12 |
| 14 |  | <pre>( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) -> ~ ∃x.( T )</pre> | subproof -> | 1 - 13 |
| 15 |  | <pre>.   ~ ∃x.( T )</pre> | Assume: -> |  |
| 16 | a | <pre>.   .   </pre> | term var: ~ ∃ |  |
| 17 | a | <pre>.   .   ∃x.( Q(x) )</pre> | Assume: ~ ∃ |  |
| 18 | a | <pre>.   .   Q(a)</pre> | Assume: ~ ∃ [x -> a] |  |
| 19 | a | <pre>.   .   T</pre> | Reiterate | 0 |
| 20 | a | <pre>.   .   ∃x.( T )</pre> | I ∃ [a -> x] | 19, 16 |
| 21 | a | <pre>.   .   ~ ∃x.( T )</pre> | Reiterate | 15 |
| 22 |  | <pre>.   ~ ∃x.( Q(x) )</pre> | subproof ~ ∃ | 16 - 21 |
| 23 | a | <pre>.   .   </pre> | term var: ~ ∃ |  |
| 24 | a | <pre>.   .   ∃x.( ~ Q(x) )</pre> | Assume: ~ ∃ |  |
| 25 | a | <pre>.   .   ~ Q(a)</pre> | Assume: ~ ∃ [x -> a] |  |
| 26 | a | <pre>.   .   T</pre> | Reiterate | 0 |
| 27 | a | <pre>.   .   ∃x.( T )</pre> | I ∃ [a -> x] | 26, 23 |
| 28 | a | <pre>.   .   ~ ∃x.( T )</pre> | Reiterate | 15 |
| 29 |  | <pre>.   ~ ∃x.( ~ Q(x) )</pre> | subproof ~ ∃ | 23 - 28 |
| 30 |  | <pre>.   ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) )</pre> | I ^ | 22, 29 |
| 31 |  | <pre>~ ∃x.( T ) -> ( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) )</pre> | subproof -> | 15 - 30 |
| 32 |  | <pre>( ~ ∃x.( Q(x) ) ^ ~ ∃x.( ~ Q(x) ) ) <-> ~ ∃x.( T )</pre> | I <-> | 14, 31 |

# distributive-and <a name="distributive-and"></a>

## ( P ^ ∃x.( Q(x) ) ) <-> ∃x.( P ^ Q(x) ) <a name="proof-4-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P ^ ∃x.( Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   P</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ∃x.( Q(x) )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 5 | a | <pre>.   .   Q(a)</pre> | Assume: E ∃ [x -> a] | 3 |
| 6 | a | <pre>.   .   P</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   P ^ Q(a)</pre> | I ^ | 6, 5 |
| 8 | a | <pre>.   .   ∃x.( P ^ Q(x) )</pre> | I ∃ [a -> x] | 7, 4 |
| 9 |  | <pre>.   ∃x.( P ^ Q(x) )</pre> | subproof E ∃ | 4 - 8 |
| 10 |  | <pre>( P ^ ∃x.( Q(x) ) ) -> ∃x.( P ^ Q(x) )</pre> | subproof -> | 1 - 9 |
| 11 |  | <pre>.   ∃x.( P ^ Q(x) )</pre> | Assume: -> |  |
| 12 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 13 | a | <pre>.   .   P ^ Q(a)</pre> | Assume: E ∃ [x -> a] | 11 |
| 14 | a | <pre>.   .   P</pre> | E ^ (L) | 13 |
| 15 | a | <pre>.   .   Q(a)</pre> | E ^ (R) | 13 |
| 16 | a | <pre>.   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 15, 12 |
| 17 | a | <pre>.   .   P ^ ∃x.( Q(x) )</pre> | I ^ | 14, 16 |
| 18 |  | <pre>.   P ^ ∃x.( Q(x) )</pre> | subproof E ∃ | 12 - 17 |
| 19 |  | <pre>∃x.( P ^ Q(x) ) -> ( P ^ ∃x.( Q(x) ) )</pre> | subproof -> | 11 - 18 |
| 20 |  | <pre>( P ^ ∃x.( Q(x) ) ) <-> ∃x.( P ^ Q(x) )</pre> | I <-> | 10, 19 |

## ∃x.( T ) -> ( ( P ^ ∀x.( Q(x) ) ) <-> ∀x.( P ^ Q(x) ) ) <a name="proof-4-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( T )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   P ^ ∀x.( Q(x) )</pre> | Assume: -> |  |
| 3 |  | <pre>.   .   P</pre> | E ^ (L) | 2 |
| 4 |  | <pre>.   .   ∀x.( Q(x) )</pre> | E ^ (R) | 2 |
| 5 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 6 | a | <pre>.   .   .   ∀x.( Q(x) )</pre> | Reiterate | 4 |
| 7 | a | <pre>.   .   .   Q(a)</pre> | E ∀ [x -> a] | 6, 5 |
| 8 | a | <pre>.   .   .   P</pre> | Reiterate | 3 |
| 9 | a | <pre>.   .   .   P ^ Q(a)</pre> | I ^ | 8, 7 |
| 10 |  | <pre>.   .   ∀x.( P ^ Q(x) )</pre> | subproof I ∀ [a -> x] | 5 - 9 |
| 11 |  | <pre>.   ( P ^ ∀x.( Q(x) ) ) -> ∀x.( P ^ Q(x) )</pre> | subproof -> | 2 - 10 |
| 12 |  | <pre>.   .   ∀x.( P ^ Q(x) )</pre> | Assume: -> |  |
| 13 |  | <pre>.   .   ∃x.( T )</pre> | Reiterate | 1 |
| 14 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 15 | a | <pre>.   .   .   T</pre> | Assume: E ∃ [x -> a] | 13 |
| 16 | a | <pre>.   .   .   ∀x.( P ^ Q(x) )</pre> | Reiterate | 12 |
| 17 | a | <pre>.   .   .   P ^ Q(a)</pre> | E ∀ [x -> a] | 16, 14 |
| 18 | a | <pre>.   .   .   P</pre> | E ^ (L) | 17 |
| 19 |  | <pre>.   .   P</pre> | subproof E ∃ | 14 - 18 |
| 20 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 21 | a | <pre>.   .   .   ∀x.( P ^ Q(x) )</pre> | Reiterate | 12 |
| 22 | a | <pre>.   .   .   P ^ Q(a)</pre> | E ∀ [x -> a] | 21, 20 |
| 23 | a | <pre>.   .   .   Q(a)</pre> | E ^ (R) | 22 |
| 24 |  | <pre>.   .   ∀x.( Q(x) )</pre> | subproof I ∀ [a -> x] | 20 - 23 |
| 25 |  | <pre>.   .   P ^ ∀x.( Q(x) )</pre> | I ^ | 19, 24 |
| 26 |  | <pre>.   ∀x.( P ^ Q(x) ) -> ( P ^ ∀x.( Q(x) ) )</pre> | subproof -> | 12 - 25 |
| 27 |  | <pre>.   ( P ^ ∀x.( Q(x) ) ) <-> ∀x.( P ^ Q(x) )</pre> | I <-> | 11, 26 |
| 28 |  | <pre>∃x.( T ) -> ( ( P ^ ∀x.( Q(x) ) ) <-> ∀x.( P ^ Q(x) ) )</pre> | subproof -> | 1 - 27 |

## ∃x.( P(x) ^ Q(x) ) -> ( ∃x.( P(x) ) ^ ∃x.( Q(x) ) ) <a name="proof-4-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( P(x) ^ Q(x) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 3 | a | <pre>.   .   P(a) ^ Q(a)</pre> | Assume: E ∃ [x -> a] | 1 |
| 4 | a | <pre>.   .   P(a)</pre> | E ^ (L) | 3 |
| 5 | a | <pre>.   .   Q(a)</pre> | E ^ (R) | 3 |
| 6 | a | <pre>.   .   ∃x.( P(x) )</pre> | I ∃ [a -> x] | 4, 2 |
| 7 | a | <pre>.   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 5, 2 |
| 8 | a | <pre>.   .   ∃x.( P(x) ) ^ ∃x.( Q(x) )</pre> | I ^ | 6, 7 |
| 9 |  | <pre>.   ∃x.( P(x) ) ^ ∃x.( Q(x) )</pre> | subproof E ∃ | 2 - 8 |
| 10 |  | <pre>∃x.( P(x) ^ Q(x) ) -> ( ∃x.( P(x) ) ^ ∃x.( Q(x) ) )</pre> | subproof -> | 1 - 9 |

## ∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) ) <a name="proof-4-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( P(x) ^ Q(x) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 3 | a | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 4 | a | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ [x -> a] | 3, 2 |
| 5 | a | <pre>.   .   P(a)</pre> | E ^ (L) | 4 |
| 6 |  | <pre>.   ∀x.( P(x) )</pre> | subproof I ∀ [a -> x] | 2 - 5 |
| 7 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 8 | a | <pre>.   .   ∀x.( P(x) ^ Q(x) )</pre> | Reiterate | 1 |
| 9 | a | <pre>.   .   P(a) ^ Q(a)</pre> | E ∀ [x -> a] | 8, 7 |
| 10 | a | <pre>.   .   Q(a)</pre> | E ^ (R) | 9 |
| 11 |  | <pre>.   ∀x.( Q(x) )</pre> | subproof I ∀ [a -> x] | 7 - 10 |
| 12 |  | <pre>.   ∀x.( P(x) ) ^ ∀x.( Q(x) )</pre> | I ^ | 6, 11 |
| 13 |  | <pre>∀x.( P(x) ^ Q(x) ) -> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )</pre> | subproof -> | 1 - 12 |
| 14 |  | <pre>.   ∀x.( P(x) ) ^ ∀x.( Q(x) )</pre> | Assume: -> |  |
| 15 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 16 | a | <pre>.   .   ∀x.( P(x) ) ^ ∀x.( Q(x) )</pre> | Reiterate | 14 |
| 17 | a | <pre>.   .   ∀x.( P(x) )</pre> | E ^ (L) | 16 |
| 18 | a | <pre>.   .   ∀x.( Q(x) )</pre> | E ^ (R) | 16 |
| 19 | a | <pre>.   .   P(a)</pre> | E ∀ [x -> a] | 17, 15 |
| 20 | a | <pre>.   .   Q(a)</pre> | E ∀ [x -> a] | 18, 15 |
| 21 | a | <pre>.   .   P(a) ^ Q(a)</pre> | I ^ | 19, 20 |
| 22 |  | <pre>.   ∀x.( P(x) ^ Q(x) )</pre> | subproof I ∀ [a -> x] | 15 - 21 |
| 23 |  | <pre>( ∀x.( P(x) ) ^ ∀x.( Q(x) ) ) -> ∀x.( P(x) ^ Q(x) )</pre> | subproof -> | 14 - 22 |
| 24 |  | <pre>∀x.( P(x) ^ Q(x) ) <-> ( ∀x.( P(x) ) ^ ∀x.( Q(x) ) )</pre> | I <-> | 13, 23 |

# distributive-or <a name="distributive-or"></a>

## ∃x.( T ) -> ( ( P v ∃x.( Q(x) ) ) <-> ∃x.( P v Q(x) ) ) <a name="proof-5-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( T )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   P v ∃x.( Q(x) )</pre> | Assume: -> |  |
| 3 |  | <pre>.   .   .   P</pre> | Assume: -> |  |
| 4 |  | <pre>.   .   .   ∃x.( T )</pre> | Reiterate | 1 |
| 5 | a | <pre>.   .   .   .   </pre> | term var: E ∃ |  |
| 6 | a | <pre>.   .   .   .   T</pre> | Assume: E ∃ [x -> a] | 4 |
| 7 | a | <pre>.   .   .   .   P</pre> | Reiterate | 3 |
| 8 | a | <pre>.   .   .   .   P v Q(a)</pre> | I v (L) | 7 |
| 9 | a | <pre>.   .   .   .   ∃x.( P v Q(x) )</pre> | I ∃ [a -> x] | 8, 5 |
| 10 |  | <pre>.   .   .   ∃x.( P v Q(x) )</pre> | subproof E ∃ | 5 - 9 |
| 11 |  | <pre>.   .   P -> ∃x.( P v Q(x) )</pre> | subproof -> | 3 - 10 |
| 12 |  | <pre>.   .   .   ∃x.( Q(x) )</pre> | Assume: -> |  |
| 13 | a | <pre>.   .   .   .   </pre> | term var: E ∃ |  |
| 14 | a | <pre>.   .   .   .   Q(a)</pre> | Assume: E ∃ [x -> a] | 12 |
| 15 | a | <pre>.   .   .   .   P v Q(a)</pre> | I v (R) | 14 |
| 16 | a | <pre>.   .   .   .   ∃x.( P v Q(x) )</pre> | I ∃ [a -> x] | 15, 13 |
| 17 |  | <pre>.   .   .   ∃x.( P v Q(x) )</pre> | subproof E ∃ | 13 - 16 |
| 18 |  | <pre>.   .   ∃x.( Q(x) ) -> ∃x.( P v Q(x) )</pre> | subproof -> | 12 - 17 |
| 19 |  | <pre>.   .   ∃x.( P v Q(x) )</pre> | E v | 11, 18, 2 |
| 20 |  | <pre>.   ( P v ∃x.( Q(x) ) ) -> ∃x.( P v Q(x) )</pre> | subproof -> | 2 - 19 |
| 21 |  | <pre>.   .   ∃x.( P v Q(x) )</pre> | Assume: -> |  |
| 22 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 23 | a | <pre>.   .   .   P v Q(a)</pre> | Assume: E ∃ [x -> a] | 21 |
| 24 | a | <pre>.   .   .   .   P</pre> | Assume: -> |  |
| 25 | a | <pre>.   .   .   .   P v ∃x.( Q(x) )</pre> | I v (L) | 24 |
| 26 | a | <pre>.   .   .   P -> ( P v ∃x.( Q(x) ) )</pre> | subproof -> | 24 - 25 |
| 27 | a | <pre>.   .   .   .   Q(a)</pre> | Assume: -> |  |
| 28 | a | <pre>.   .   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 27, 22 |
| 29 | a | <pre>.   .   .   .   P v ∃x.( Q(x) )</pre> | I v (R) | 28 |
| 30 | a | <pre>.   .   .   Q(a) -> ( P v ∃x.( Q(x) ) )</pre> | subproof -> | 27 - 29 |
| 31 | a | <pre>.   .   .   P v ∃x.( Q(x) )</pre> | E v | 26, 30, 23 |
| 32 |  | <pre>.   .   P v ∃x.( Q(x) )</pre> | subproof E ∃ | 22 - 31 |
| 33 |  | <pre>.   ∃x.( P v Q(x) ) -> ( P v ∃x.( Q(x) ) )</pre> | subproof -> | 21 - 32 |
| 34 |  | <pre>.   ( P v ∃x.( Q(x) ) ) <-> ∃x.( P v Q(x) )</pre> | I <-> | 20, 33 |
| 35 |  | <pre>∃x.( T ) -> ( ( P v ∃x.( Q(x) ) ) <-> ∃x.( P v Q(x) ) )</pre> | subproof -> | 1 - 34 |

## ( P v ∀x.( Q(x) ) ) <-> ∀x.( P v Q(x) ) <a name="proof-5-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P v ∀x.( Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   P</pre> | Assume: -> |  |
| 3 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 4 | a | <pre>.   .   .   P</pre> | Reiterate | 2 |
| 5 | a | <pre>.   .   .   P v Q(a)</pre> | I v (L) | 4 |
| 6 |  | <pre>.   .   ∀x.( P v Q(x) )</pre> | subproof I ∀ [a -> x] | 3 - 5 |
| 7 |  | <pre>.   P -> ∀x.( P v Q(x) )</pre> | subproof -> | 2 - 6 |
| 8 |  | <pre>.   .   ∀x.( Q(x) )</pre> | Assume: -> |  |
| 9 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 10 | a | <pre>.   .   .   ∀x.( Q(x) )</pre> | Reiterate | 8 |
| 11 | a | <pre>.   .   .   Q(a)</pre> | E ∀ [x -> a] | 10, 9 |
| 12 | a | <pre>.   .   .   P v Q(a)</pre> | I v (R) | 11 |
| 13 |  | <pre>.   .   ∀x.( P v Q(x) )</pre> | subproof I ∀ [a -> x] | 9 - 12 |
| 14 |  | <pre>.   ∀x.( Q(x) ) -> ∀x.( P v Q(x) )</pre> | subproof -> | 8 - 13 |
| 15 |  | <pre>.   ∀x.( P v Q(x) )</pre> | E v | 7, 14, 1 |
| 16 |  | <pre>( P v ∀x.( Q(x) ) ) -> ∀x.( P v Q(x) )</pre> | subproof -> | 1 - 15 |
| 17 |  | <pre>.   ∀x.( P v Q(x) )</pre> | Assume: -> |  |
| 18 |  | <pre>.   .   P</pre> | Assume: -> |  |
| 19 |  | <pre>.   .   P v ∀x.( Q(x) )</pre> | I v (L) | 18 |
| 20 |  | <pre>.   P -> ( P v ∀x.( Q(x) ) )</pre> | subproof -> | 18 - 19 |
| 21 |  | <pre>.   .   ~ P</pre> | Assume: -> |  |
| 22 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 23 | a | <pre>.   .   .   ∀x.( P v Q(x) )</pre> | Reiterate | 17 |
| 24 | a | <pre>.   .   .   P v Q(a)</pre> | E ∀ [x -> a] | 23, 22 |
| 25 | a | <pre>.   .   .   ~ P -> Q(a)</pre> | Theorem: v to -> | 24 |
| 26 | a | <pre>.   .   .   ~ P</pre> | Reiterate | 21 |
| 27 | a | <pre>.   .   .   Q(a)</pre> | E -> | 25, 26 |
| 28 |  | <pre>.   .   ∀x.( Q(x) )</pre> | subproof I ∀ [a -> x] | 22 - 27 |
| 29 |  | <pre>.   .   P v ∀x.( Q(x) )</pre> | I v (R) | 28 |
| 30 |  | <pre>.   ~ P -> ( P v ∀x.( Q(x) ) )</pre> | subproof -> | 21 - 29 |
| 31 |  | <pre>.   P v ~ P</pre> | Theorem: excluded middle |  |
| 32 |  | <pre>.   P v ∀x.( Q(x) )</pre> | E v | 20, 30, 31 |
| 33 |  | <pre>∀x.( P v Q(x) ) -> ( P v ∀x.( Q(x) ) )</pre> | subproof -> | 17 - 32 |
| 34 |  | <pre>( P v ∀x.( Q(x) ) ) <-> ∀x.( P v Q(x) )</pre> | I <-> | 16, 33 |

## ∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) ) <a name="proof-5-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( P(x) v Q(x) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 3 | a | <pre>.   .   P(a) v Q(a)</pre> | Assume: E ∃ [x -> a] | 1 |
| 4 | a | <pre>.   .   .   P(a)</pre> | Assume: -> |  |
| 5 | a | <pre>.   .   .   ∃x.( P(x) )</pre> | I ∃ [a -> x] | 4, 2 |
| 6 | a | <pre>.   .   .   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | I v (L) | 5 |
| 7 | a | <pre>.   .   P(a) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | subproof -> | 4 - 6 |
| 8 | a | <pre>.   .   .   Q(a)</pre> | Assume: -> |  |
| 9 | a | <pre>.   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 8, 2 |
| 10 | a | <pre>.   .   .   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | I v (R) | 9 |
| 11 | a | <pre>.   .   Q(a) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | subproof -> | 8 - 10 |
| 12 | a | <pre>.   .   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | E v | 7, 11, 3 |
| 13 |  | <pre>.   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | subproof E ∃ | 2 - 12 |
| 14 |  | <pre>∃x.( P(x) v Q(x) ) -> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | subproof -> | 1 - 13 |
| 15 |  | <pre>.   ∃x.( P(x) ) v ∃x.( Q(x) )</pre> | Assume: -> |  |
| 16 |  | <pre>.   .   ∃x.( P(x) )</pre> | Assume: -> |  |
| 17 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 18 | a | <pre>.   .   .   P(a)</pre> | Assume: E ∃ [x -> a] | 16 |
| 19 | a | <pre>.   .   .   P(a) v Q(a)</pre> | I v (L) | 18 |
| 20 | a | <pre>.   .   .   ∃x.( P(x) v Q(x) )</pre> | I ∃ [a -> x] | 19, 17 |
| 21 |  | <pre>.   .   ∃x.( P(x) v Q(x) )</pre> | subproof E ∃ | 17 - 20 |
| 22 |  | <pre>.   ∃x.( P(x) ) -> ∃x.( P(x) v Q(x) )</pre> | subproof -> | 16 - 21 |
| 23 |  | <pre>.   .   ∃x.( Q(x) )</pre> | Assume: -> |  |
| 24 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 25 | a | <pre>.   .   .   Q(a)</pre> | Assume: E ∃ [x -> a] | 23 |
| 26 | a | <pre>.   .   .   P(a) v Q(a)</pre> | I v (R) | 25 |
| 27 | a | <pre>.   .   .   ∃x.( P(x) v Q(x) )</pre> | I ∃ [a -> x] | 26, 24 |
| 28 |  | <pre>.   .   ∃x.( P(x) v Q(x) )</pre> | subproof E ∃ | 24 - 27 |
| 29 |  | <pre>.   ∃x.( Q(x) ) -> ∃x.( P(x) v Q(x) )</pre> | subproof -> | 23 - 28 |
| 30 |  | <pre>.   ∃x.( P(x) v Q(x) )</pre> | E v | 22, 29, 15 |
| 31 |  | <pre>( ∃x.( P(x) ) v ∃x.( Q(x) ) ) -> ∃x.( P(x) v Q(x) )</pre> | subproof -> | 15 - 30 |
| 32 |  | <pre>∃x.( P(x) v Q(x) ) <-> ( ∃x.( P(x) ) v ∃x.( Q(x) ) )</pre> | I <-> | 14, 31 |

## ( ∀x.( P(x) ) v ∀x.( Q(x) ) ) -> ∀x.( P(x) v Q(x) ) <a name="proof-5-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( P(x) ) v ∀x.( Q(x) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 3 | a | <pre>.   .   ∀x.( P(x) ) v ∀x.( Q(x) )</pre> | Reiterate | 1 |
| 4 | a | <pre>.   .   .   ∀x.( P(x) )</pre> | Assume: -> |  |
| 5 | a | <pre>.   .   .   P(a)</pre> | E ∀ [x -> a] | 4, 2 |
| 6 | a | <pre>.   .   .   P(a) v Q(a)</pre> | I v (L) | 5 |
| 7 | a | <pre>.   .   ∀x.( P(x) ) -> ( P(a) v Q(a) )</pre> | subproof -> | 4 - 6 |
| 8 | a | <pre>.   .   .   ∀x.( Q(x) )</pre> | Assume: -> |  |
| 9 | a | <pre>.   .   .   Q(a)</pre> | E ∀ [x -> a] | 8, 2 |
| 10 | a | <pre>.   .   .   P(a) v Q(a)</pre> | I v (R) | 9 |
| 11 | a | <pre>.   .   ∀x.( Q(x) ) -> ( P(a) v Q(a) )</pre> | subproof -> | 8 - 10 |
| 12 | a | <pre>.   .   P(a) v Q(a)</pre> | E v | 7, 11, 3 |
| 13 |  | <pre>.   ∀x.( P(x) v Q(x) )</pre> | subproof I ∀ [a -> x] | 2 - 12 |
| 14 |  | <pre>( ∀x.( P(x) ) v ∀x.( Q(x) ) ) -> ∀x.( P(x) v Q(x) )</pre> | subproof -> | 1 - 13 |

# distributive-arrow <a name="distributive-arrow"></a>

## ∃x.( T ) -> ( ( P -> ∃x.( Q(x) ) ) <-> ∃x.( P -> Q(x) ) ) <a name="proof-6-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( T )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   P -> ∃x.( Q(x) )</pre> | Assume: -> |  |
| 3 |  | <pre>.   .   ∃x.( T )</pre> | Reiterate | 1 |
| 4 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 5 | a | <pre>.   .   .   T</pre> | Assume: E ∃ [x -> a] | 3 |
| 6 | a | <pre>.   .   .   .   P</pre> | Assume: -> |  |
| 7 | a | <pre>.   .   .   .   P -> ∃x.( Q(x) )</pre> | Reiterate | 2 |
| 8 | a | <pre>.   .   .   .   ∃x.( Q(x) )</pre> | E -> | 7, 6 |
| 9 | a, b | <pre>.   .   .   .   .   </pre> | term var: E ∃ |  |
| 10 | a, b | <pre>.   .   .   .   .   Q(b)</pre> | Assume: E ∃ [x -> b] | 8 |
| 11 | a, b | <pre>.   .   .   .   .   P</pre> | Reiterate | 6 |
| 12 | a, b | <pre>.   .   .   .   .   P -> Q(b)</pre> | I -> | 11, 10 |
| 13 | a, b | <pre>.   .   .   .   .   ∃x.( P -> Q(x) )</pre> | I ∃ [b -> x] | 12, 9 |
| 14 | a | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> | subproof E ∃ | 9 - 13 |
| 15 | a | <pre>.   .   .   P -> ∃x.( P -> Q(x) )</pre> | subproof -> | 6 - 14 |
| 16 | a | <pre>.   .   .   .   ~ P</pre> | Assume: -> |  |
| 17 | a | <pre>.   .   .   .   .   ~ Q(a)</pre> | Assume: -> |  |
| 18 | a | <pre>.   .   .   .   .   ~ P</pre> | Reiterate | 16 |
| 19 | a | <pre>.   .   .   .   ~ Q(a) -> ~ P</pre> | subproof -> | 17 - 18 |
| 20 | a | <pre>.   .   .   .   P -> Q(a)</pre> | Theorem: contrapositive | 19 |
| 21 | a | <pre>.   .   .   .   ∃x.( P -> Q(x) )</pre> | I ∃ [a -> x] | 20, 4 |
| 22 | a | <pre>.   .   .   ~ P -> ∃x.( P -> Q(x) )</pre> | subproof -> | 16 - 21 |
| 23 | a | <pre>.   .   .   P v ~ P</pre> | Theorem: excluded middle |  |
| 24 | a | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> | E v | 15, 22, 23 |
| 25 |  | <pre>.   .   ∃x.( P -> Q(x) )</pre> | subproof E ∃ | 4 - 24 |
| 26 |  | <pre>.   ( P -> ∃x.( Q(x) ) ) -> ∃x.( P -> Q(x) )</pre> | subproof -> | 2 - 25 |
| 27 |  | <pre>.   .   ∃x.( P -> Q(x) )</pre> | Assume: -> |  |
| 28 |  | <pre>.   .   .   P</pre> | Assume: -> |  |
| 29 |  | <pre>.   .   .   ∃x.( P -> Q(x) )</pre> | Reiterate | 27 |
| 30 | a | <pre>.   .   .   .   </pre> | term var: E ∃ |  |
| 31 | a | <pre>.   .   .   .   P -> Q(a)</pre> | Assume: E ∃ [x -> a] | 29 |
| 32 | a | <pre>.   .   .   .   P</pre> | Reiterate | 28 |
| 33 | a | <pre>.   .   .   .   Q(a)</pre> | E -> | 31, 32 |
| 34 | a | <pre>.   .   .   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 33, 30 |
| 35 |  | <pre>.   .   .   ∃x.( Q(x) )</pre> | subproof E ∃ | 30 - 34 |
| 36 |  | <pre>.   .   P -> ∃x.( Q(x) )</pre> | subproof -> | 28 - 35 |
| 37 |  | <pre>.   ∃x.( P -> Q(x) ) -> ( P -> ∃x.( Q(x) ) )</pre> | subproof -> | 27 - 36 |
| 38 |  | <pre>.   ( P -> ∃x.( Q(x) ) ) <-> ∃x.( P -> Q(x) )</pre> | I <-> | 26, 37 |
| 39 |  | <pre>∃x.( T ) -> ( ( P -> ∃x.( Q(x) ) ) <-> ∃x.( P -> Q(x) ) )</pre> | subproof -> | 1 - 38 |

## ( P -> ∀x.( Q(x) ) ) <-> ∀x.( P -> Q(x) ) <a name="proof-6-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P -> ∀x.( Q(x) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 3 | a | <pre>.   .   .   P</pre> | Assume: -> |  |
| 4 | a | <pre>.   .   .   P -> ∀x.( Q(x) )</pre> | Reiterate | 1 |
| 5 | a | <pre>.   .   .   ∀x.( Q(x) )</pre> | E -> | 4, 3 |
| 6 | a | <pre>.   .   .   Q(a)</pre> | E ∀ [x -> a] | 5, 2 |
| 7 | a | <pre>.   .   P -> Q(a)</pre> | subproof -> | 3 - 6 |
| 8 |  | <pre>.   ∀x.( P -> Q(x) )</pre> | subproof I ∀ [a -> x] | 2 - 7 |
| 9 |  | <pre>( P -> ∀x.( Q(x) ) ) -> ∀x.( P -> Q(x) )</pre> | subproof -> | 1 - 8 |
| 10 |  | <pre>.   ∀x.( P -> Q(x) )</pre> | Assume: -> |  |
| 11 |  | <pre>.   .   P</pre> | Assume: -> |  |
| 12 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 13 | a | <pre>.   .   .   ∀x.( P -> Q(x) )</pre> | Reiterate | 10 |
| 14 | a | <pre>.   .   .   P -> Q(a)</pre> | E ∀ [x -> a] | 13, 12 |
| 15 | a | <pre>.   .   .   P</pre> | Reiterate | 11 |
| 16 | a | <pre>.   .   .   Q(a)</pre> | E -> | 14, 15 |
| 17 |  | <pre>.   .   ∀x.( Q(x) )</pre> | subproof I ∀ [a -> x] | 12 - 16 |
| 18 |  | <pre>.   P -> ∀x.( Q(x) )</pre> | subproof -> | 11 - 17 |
| 19 |  | <pre>∀x.( P -> Q(x) ) -> ( P -> ∀x.( Q(x) ) )</pre> | subproof -> | 10 - 18 |
| 20 |  | <pre>( P -> ∀x.( Q(x) ) ) <-> ∀x.( P -> Q(x) )</pre> | I <-> | 9, 19 |

## ( ∃x.( P(x) ) -> ∃x.( Q(x) ) ) -> ( ∃x.( T ) -> ∃x.( P(x) -> Q(x) ) ) <a name="proof-6-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( P(x) ) -> ∃x.( Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ∃x.( T )</pre> | Assume: -> |  |
| 3 |  | <pre>.   .   .   ∃x.( P(x) )</pre> | Assume: -> |  |
| 4 |  | <pre>.   .   .   ∃x.( P(x) ) -> ∃x.( Q(x) )</pre> | Reiterate | 1 |
| 5 |  | <pre>.   .   .   ∃x.( Q(x) )</pre> | E -> | 4, 3 |
| 6 | a | <pre>.   .   .   .   </pre> | term var: E ∃ |  |
| 7 | a | <pre>.   .   .   .   Q(a)</pre> | Assume: E ∃ [x -> a] | 5 |
| 8 | a | <pre>.   .   .   .   .   P(a)</pre> | Assume: -> |  |
| 9 | a | <pre>.   .   .   .   .   Q(a)</pre> | Reiterate | 7 |
| 10 | a | <pre>.   .   .   .   P(a) -> Q(a)</pre> | subproof -> | 8 - 9 |
| 11 | a | <pre>.   .   .   .   ∃x.( P(x) -> Q(x) )</pre> | I ∃ [a -> x] | 10, 6 |
| 12 |  | <pre>.   .   .   ∃x.( P(x) -> Q(x) )</pre> | subproof E ∃ | 6 - 11 |
| 13 |  | <pre>.   .   ∃x.( P(x) ) -> ∃x.( P(x) -> Q(x) )</pre> | subproof -> | 3 - 12 |
| 14 |  | <pre>.   .   .   ~ ∃x.( P(x) )</pre> | Assume: -> |  |
| 15 |  | <pre>.   .   .   .   ~ ∃x.( ~ P(x) )</pre> | Assume: ~ |  |
| 16 | a | <pre>.   .   .   .   .   </pre> | term var: ~ ∃ |  |
| 17 | a | <pre>.   .   .   .   .   ∃x.( T )</pre> | Assume: ~ ∃ |  |
| 18 | a | <pre>.   .   .   .   .   T</pre> | Assume: ~ ∃ [x -> a] |  |
| 19 | a | <pre>.   .   .   .   .   .   P(a)</pre> | Assume: ~ |  |
| 20 | a | <pre>.   .   .   .   .   .   ∃x.( P(x) )</pre> | I ∃ [a -> x] | 19, 16 |
| 21 | a | <pre>.   .   .   .   .   .   ~ ∃x.( P(x) )</pre> | Reiterate | 14 |
| 22 | a | <pre>.   .   .   .   .   ~ P(a)</pre> | subproof ~ | 19 - 21 |
| 23 | a | <pre>.   .   .   .   .   ~ ∃x.( ~ P(x) )</pre> | Reiterate | 15 |
| 24 | a | <pre>.   .   .   .   .   ∃x.( ~ P(x) )</pre> | I ∃ [a -> x] | 22, 16 |
| 25 |  | <pre>.   .   .   .   ~ ∃x.( T )</pre> | subproof ~ ∃ | 16 - 24 |
| 26 |  | <pre>.   .   .   .   ∃x.( T )</pre> | Reiterate | 2 |
| 27 |  | <pre>.   .   .   ∃x.( ~ P(x) )</pre> | subproof ~ | 15 - 26 |
| 28 | a | <pre>.   .   .   .   </pre> | term var: E ∃ |  |
| 29 | a | <pre>.   .   .   .   ~ P(a)</pre> | Assume: E ∃ [x -> a] | 27 |
| 30 | a | <pre>.   .   .   .   ~ P(a) v Q(a)</pre> | I v (L) | 29 |
| 31 | a | <pre>.   .   .   .   P(a) -> Q(a)</pre> | Theorem: v to -> | 30 |
| 32 | a | <pre>.   .   .   .   ∃x.( P(x) -> Q(x) )</pre> | I ∃ [a -> x] | 31, 28 |
| 33 |  | <pre>.   .   .   ∃x.( P(x) -> Q(x) )</pre> | subproof E ∃ | 28 - 32 |
| 34 |  | <pre>.   .   ~ ∃x.( P(x) ) -> ∃x.( P(x) -> Q(x) )</pre> | subproof -> | 14 - 33 |
| 35 |  | <pre>.   .   ∃x.( P(x) ) v ~ ∃x.( P(x) )</pre> | Theorem: excluded middle |  |
| 36 |  | <pre>.   .   ∃x.( P(x) -> Q(x) )</pre> | E v | 13, 34, 35 |
| 37 |  | <pre>.   ∃x.( T ) -> ∃x.( P(x) -> Q(x) )</pre> | subproof -> | 2 - 36 |
| 38 |  | <pre>( ∃x.( P(x) ) -> ∃x.( Q(x) ) ) -> ( ∃x.( T ) -> ∃x.( P(x) -> Q(x) ) )</pre> | subproof -> | 1 - 37 |

## ∀x.( P(x) -> Q(x) ) -> ( ∀x.( P(x) ) -> ∀x.( Q(x) ) ) <a name="proof-6-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( P(x) -> Q(x) )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ∀x.( P(x) )</pre> | Assume: -> |  |
| 3 | a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 4 | a | <pre>.   .   .   ∀x.( P(x) -> Q(x) )</pre> | Reiterate | 1 |
| 5 | a | <pre>.   .   .   ∀x.( P(x) )</pre> | Reiterate | 2 |
| 6 | a | <pre>.   .   .   P(a) -> Q(a)</pre> | E ∀ [x -> a] | 4, 3 |
| 7 | a | <pre>.   .   .   P(a)</pre> | E ∀ [x -> a] | 5, 3 |
| 8 | a | <pre>.   .   .   Q(a)</pre> | E -> | 6, 7 |
| 9 |  | <pre>.   .   ∀x.( Q(x) )</pre> | subproof I ∀ [a -> x] | 3 - 8 |
| 10 |  | <pre>.   ∀x.( P(x) ) -> ∀x.( Q(x) )</pre> | subproof -> | 2 - 9 |
| 11 |  | <pre>∀x.( P(x) -> Q(x) ) -> ( ∀x.( P(x) ) -> ∀x.( Q(x) ) )</pre> | subproof -> | 1 - 10 |

# commutativity <a name="commutativity"></a>

## ∃x.( ∃y.( P(x,y) ) ) -> ∃y.( ∃x.( P(x,y) ) ) <a name="proof-7-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( ∃y.( P(x,y) ) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 3 | a | <pre>.   .   ∃y.( P(a,y) )</pre> | Assume: E ∃ [x -> a] | 1 |
| 4 | a, b | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 5 | a, b | <pre>.   .   .   P(a,b)</pre> | Assume: E ∃ [y -> b] | 3 |
| 6 | a, b | <pre>.   .   .   ∃x.( P(x,b) )</pre> | I ∃ [a -> x] | 5, 2 |
| 7 | a, b | <pre>.   .   .   ∃y.( ∃x.( P(x,y) ) )</pre> | I ∃ [b -> y] | 6, 4 |
| 8 | a | <pre>.   .   ∃y.( ∃x.( P(x,y) ) )</pre> | subproof E ∃ | 4 - 7 |
| 9 |  | <pre>.   ∃y.( ∃x.( P(x,y) ) )</pre> | subproof E ∃ | 2 - 8 |
| 10 |  | <pre>∃x.( ∃y.( P(x,y) ) ) -> ∃y.( ∃x.( P(x,y) ) )</pre> | subproof -> | 1 - 9 |

## ∀x.( ∀y.( P(x,y) ) ) -> ∀y.( ∀x.( P(x,y) ) ) <a name="proof-7-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( ∀y.( P(x,y) ) )</pre> | Assume: -> |  |
| 2 | b | <pre>.   .   </pre> | term var: I ∀ |  |
| 3 | b, a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 4 | b, a | <pre>.   .   .   ∀x.( ∀y.( P(x,y) ) )</pre> | Reiterate | 1 |
| 5 | b, a | <pre>.   .   .   ∀y.( P(a,y) )</pre> | E ∀ [x -> a] | 4, 3 |
| 6 | b, a | <pre>.   .   .   P(a,b)</pre> | E ∀ [y -> b] | 5, 2 |
| 7 | b | <pre>.   .   ∀x.( P(x,b) )</pre> | subproof I ∀ [a -> x] | 3 - 6 |
| 8 |  | <pre>.   ∀y.( ∀x.( P(x,y) ) )</pre> | subproof I ∀ [b -> y] | 2 - 7 |
| 9 |  | <pre>∀x.( ∀y.( P(x,y) ) ) -> ∀y.( ∀x.( P(x,y) ) )</pre> | subproof -> | 1 - 8 |

## ∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) ) <a name="proof-7-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( ∀y.( P(x,y) ) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 3 | a | <pre>.   .   ∃x.( ∀y.( P(x,y) ) )</pre> | Reiterate | 1 |
| 4 | a, b | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 5 | a, b | <pre>.   .   .   ∀y.( P(b,y) )</pre> | Assume: E ∃ [x -> b] | 3 |
| 6 | a, b | <pre>.   .   .   P(b,a)</pre> | E ∀ [y -> a] | 5, 2 |
| 7 | a, b | <pre>.   .   .   ∃x.( P(x,a) )</pre> | I ∃ [b -> x] | 6, 4 |
| 8 | a | <pre>.   .   ∃x.( P(x,a) )</pre> | subproof E ∃ | 4 - 7 |
| 9 |  | <pre>.   ∀y.( ∃x.( P(x,y) ) )</pre> | subproof I ∀ [a -> y] | 2 - 8 |
| 10 |  | <pre>∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) )</pre> | subproof -> | 1 - 9 |

## ∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) ) <a name="proof-7-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( ∀y.( P(x,y) ) )</pre> | Assume: -> |  |
| 2 | b | <pre>.   .   </pre> | term var: E ∃ |  |
| 3 | b | <pre>.   .   ∀y.( P(b,y) )</pre> | Assume: E ∃ [x -> b] | 1 |
| 4 | b, a | <pre>.   .   .   </pre> | term var: I ∀ |  |
| 5 | b, a | <pre>.   .   .   ∀y.( P(b,y) )</pre> | Reiterate | 3 |
| 6 | b, a | <pre>.   .   .   P(b,a)</pre> | E ∀ [y -> a] | 5, 4 |
| 7 | b, a | <pre>.   .   .   ∃x.( P(x,a) )</pre> | I ∃ [b -> x] | 6, 2 |
| 8 | b | <pre>.   .   ∀y.( ∃x.( P(x,y) ) )</pre> | subproof I ∀ [a -> y] | 4 - 7 |
| 9 |  | <pre>.   ∀y.( ∃x.( P(x,y) ) )</pre> | subproof E ∃ | 2 - 8 |
| 10 |  | <pre>∃x.( ∀y.( P(x,y) ) ) -> ∀y.( ∃x.( P(x,y) ) )</pre> | subproof -> | 1 - 9 |

# random <a name="random"></a>

## P -> ( ∀x.( ~ P ) -> ~ ∃x.( T ) ) <a name="proof-8-1"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ∀x.( ~ P )</pre> | Assume: -> |  |
| 3 | a | <pre>.   .   .   </pre> | term var: ~ ∃ |  |
| 4 | a | <pre>.   .   .   ∃x.( T )</pre> | Assume: ~ ∃ |  |
| 5 | a | <pre>.   .   .   T</pre> | Assume: ~ ∃ [x -> a] |  |
| 6 | a | <pre>.   .   .   ∀x.( ~ P )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   .   ~ P</pre> | E ∀ [x -> a] | 6, 3 |
| 8 | a | <pre>.   .   .   P</pre> | Reiterate | 1 |
| 9 |  | <pre>.   .   ~ ∃x.( T )</pre> | subproof ~ ∃ | 3 - 8 |
| 10 |  | <pre>.   ∀x.( ~ P ) -> ~ ∃x.( T )</pre> | subproof -> | 2 - 9 |
| 11 |  | <pre>P -> ( ∀x.( ~ P ) -> ~ ∃x.( T ) )</pre> | subproof -> | 1 - 10 |

## P -> ( ∃x.( T ) -> ~ ∀x.( ~ P ) ) <a name="proof-8-2"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   P</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ∃x.( T )</pre> | Assume: -> |  |
| 3 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 4 | a | <pre>.   .   .   T</pre> | Assume: E ∃ [x -> a] | 2 |
| 5 | a | <pre>.   .   .   .   ∀x.( ~ P )</pre> | Assume: ~ |  |
| 6 | a | <pre>.   .   .   .   ~ P</pre> | E ∀ [x -> a] | 5, 3 |
| 7 | a | <pre>.   .   .   .   P</pre> | Reiterate | 1 |
| 8 | a | <pre>.   .   .   ~ ∀x.( ~ P )</pre> | subproof ~ | 5 - 7 |
| 9 |  | <pre>.   .   ~ ∀x.( ~ P )</pre> | subproof E ∃ | 3 - 8 |
| 10 |  | <pre>.   ∃x.( T ) -> ~ ∀x.( ~ P )</pre> | subproof -> | 2 - 9 |
| 11 |  | <pre>P -> ( ∃x.( T ) -> ~ ∀x.( ~ P ) )</pre> | subproof -> | 1 - 10 |

## ∃x.( T ) -> ( ∀x.( ~ P ) -> ~ P ) <a name="proof-8-3"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( T )</pre> | Assume: -> |  |
| 2 |  | <pre>.   .   ∀x.( ~ P )</pre> | Assume: -> |  |
| 3 |  | <pre>.   .   ∃x.( T )</pre> | Reiterate | 1 |
| 4 | a | <pre>.   .   .   </pre> | term var: E ∃ |  |
| 5 | a | <pre>.   .   .   T</pre> | Assume: E ∃ [x -> a] | 3 |
| 6 | a | <pre>.   .   .   ∀x.( ~ P )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   .   ~ P</pre> | E ∀ [x -> a] | 6, 4 |
| 8 |  | <pre>.   .   ~ P</pre> | subproof E ∃ | 4 - 7 |
| 9 |  | <pre>.   ∀x.( ~ P ) -> ~ P</pre> | subproof -> | 2 - 8 |
| 10 |  | <pre>∃x.( T ) -> ( ∀x.( ~ P ) -> ~ P )</pre> | subproof -> | 1 - 9 |

## ∃x.( R ) -> R <a name="proof-8-4"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( R )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 3 | a | <pre>.   .   R</pre> | Assume: E ∃ [x -> a] | 1 |
| 4 |  | <pre>.   R</pre> | subproof E ∃ | 2 - 3 |
| 5 |  | <pre>∃x.( R ) -> R</pre> | subproof -> | 1 - 4 |

## R -> ∀x.( R ) <a name="proof-8-5"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   R</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: I ∀ |  |
| 3 | a | <pre>.   .   R</pre> | Reiterate | 1 |
| 4 |  | <pre>.   ∀x.( R )</pre> | subproof I ∀ [a -> x] | 2 - 3 |
| 5 |  | <pre>R -> ∀x.( R )</pre> | subproof -> | 1 - 4 |

## ∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R <a name="proof-8-6"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∃x.( Q(x) ^ ( Q(x) -> R ) )</pre> | Assume: -> |  |
| 2 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 3 | a | <pre>.   .   Q(a) ^ ( Q(a) -> R )</pre> | Assume: E ∃ [x -> a] | 1 |
| 4 | a | <pre>.   .   Q(a)</pre> | E ^ (L) | 3 |
| 5 | a | <pre>.   .   Q(a) -> R</pre> | E ^ (R) | 3 |
| 6 | a | <pre>.   .   R</pre> | E -> | 5, 4 |
| 7 |  | <pre>.   R</pre> | subproof E ∃ | 2 - 6 |
| 8 |  | <pre>∃x.( Q(x) ^ ( Q(x) -> R ) ) -> R</pre> | subproof -> | 1 - 7 |

## ( ∀x.( Q(x) ) ^ ∃x.( Q(x) -> R ) ) -> R <a name="proof-8-7"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( Q(x) ) ^ ∃x.( Q(x) -> R )</pre> | Assume: -> |  |
| 2 |  | <pre>.   ∀x.( Q(x) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ∃x.( Q(x) -> R )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 5 | a | <pre>.   .   Q(a) -> R</pre> | Assume: E ∃ [x -> a] | 3 |
| 6 | a | <pre>.   .   ∀x.( Q(x) )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   Q(a)</pre> | E ∀ [x -> a] | 6, 4 |
| 8 | a | <pre>.   .   R</pre> | E -> | 5, 7 |
| 9 |  | <pre>.   R</pre> | subproof E ∃ | 4 - 8 |
| 10 |  | <pre>( ∀x.( Q(x) ) ^ ∃x.( Q(x) -> R ) ) -> R</pre> | subproof -> | 1 - 9 |

## ( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) ) <a name="proof-8-8"></a>

| Line | Term var | Formula | Justification | Lines used |
| - | - | - | - | - |
| 1 |  | <pre>.   ∀x.( Q(x) ) ^ ∃x.( T )</pre> | Assume: -> |  |
| 2 |  | <pre>.   ∀x.( Q(x) )</pre> | E ^ (L) | 1 |
| 3 |  | <pre>.   ∃x.( T )</pre> | E ^ (R) | 1 |
| 4 | a | <pre>.   .   </pre> | term var: E ∃ |  |
| 5 | a | <pre>.   .   T</pre> | Assume: E ∃ [x -> a] | 3 |
| 6 | a | <pre>.   .   ∀x.( Q(x) )</pre> | Reiterate | 2 |
| 7 | a | <pre>.   .   Q(a)</pre> | E ∀ [x -> a] | 6, 4 |
| 8 | a | <pre>.   .   ∃x.( Q(x) )</pre> | I ∃ [a -> x] | 7, 4 |
| 9 |  | <pre>.   ∃x.( Q(x) )</pre> | subproof E ∃ | 4 - 8 |
| 10 |  | <pre>( ∀x.( Q(x) ) ^ ∃x.( T ) ) -> ∃x.( Q(x) )</pre> | subproof -> | 1 - 9 |

