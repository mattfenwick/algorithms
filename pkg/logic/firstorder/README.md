# First-order Logic

This is an implementation of classical [first-order logic](https://en.wikipedia.org/wiki/First-order_logic):

 - negation/contradiction:
   - double negation: `~ ~ P => P`
   - law of the excluded middle for all propositions: `P v ~ P`
   - unrestricted reductio ad absurdum
 - first-order: forall and existential quantifiers
 - introduction and elimination rules for each connective

Proofs are presented in the Fitch style, using subproofs and reiterations.

The domain [is *not* assumed to be empty](https://en.wikipedia.org/wiki/First-order_logic#Empty_domains).
Proofs which require a non-empty domain must explicitly state that hypothesis.

# Rules

## Connectives

| Connective | Introduction | Elimination |
| - | - | - |
| ^ | P, Q => P ^ Q | P ^ Q => P; P ^ Q => Q |
| v | P => P v Q; Q => P v Q | P -> R, Q -> R, P v Q => R |
| -> | P, Q => P -> Q | P -> Q, P => Q |
| ~ | subproof: P => ~P | subproof: ~P => P |
| <-> | P <-> Q => P -> Q; P <-> Q => Q -> P | P -> Q, Q -> P => P <-> Q |
| ∃ | `P(a)` => `∃x.P(x)` | [E-∃ subproof](#existential-elimination) |
| ∀ | [I-∀ subproof](#forall-introduction) | `∀x.P(x)` => `P(a)` |

## Subproofs

A temporary hypothesis is assumed for a nested scope.  Deductions are then carried out.
Results from enclosing scopes may only be used in subproofs if explicitly reiterated.

Reiteration brings a deduction from an enclosing scope into a subproof.

The hypothesis and following step results may not be used outside of their subproof.

Term variables are used by forall and existential subproofs to refer to arbitary instances of the domain as part of predicate formulas.  Term variables may not be used outside of their subproof, nor may they be allowed to escape.

This table summarizes the proof types:

 - preconditions: what must already be in the enclosing scope
 - hypothesis: a formula assumed to be true for the scope of the proof
 - last line: used to determine the result.  How exactly it's used depends on the proof type
 - result: a formula added as true to the *enclosing* scope after completion of the subproof

| Proof type | Term var | Preconditions | Hypothesis | Last line | Result |
| - | - | - | - | - | - |
| Arrow | - | - | `P` | `Q` | `P -> Q` |
| Contradiction | - | - | `P` | `Q` ... `~ Q` | `~ P` |
| ∃ elimination | `a` | `∃x.P(x)` | `P(a)` | `Q` | `Q` |
| ∃ contradiction | `a` | - | `∃x.P(x)` | `Q` ... `~ Q` | `~∃x.P(x)` |
| ∀ introduction | `a` | - | - | `P(a)` | `∀x.P(x)` |

### Implication

Starting from hypothesis `P`, `Q` is derived.  Therefore `P -> Q` is added to the enclosing scope.

### Contradiction

Starting from hypothesis `P`, either `~ P` or both of `Q` and `~ Q` are derived.  Therefore `~ P`
is added to the enclosing scope.

The last line must be a formula which contradicts a previous formula in the same proof.
For example, if the subproof ends with `~ Q`, then either `Q` or
`~ ~ Q` must have previously been asserted in the subproof.

### Existential elimination <a name="existential-elimination"></a>

Precondition: the existential formula must already be present in the enclosing scope.

A term var, for example `a`, is added to the subproof's term var scope.
For the existential formula `∃x.P(x)`, the formula `P(a)` will be added to the subproof's scope.

Reiterations *may not* mention the term var.

Note that `P(a)` is the result of substituting `a` for every bound instance of `x` in the existential formula.

The last line *may not* use the term var `a`.  Justification:

 - term vars may not be allowed to escape their scope
 - the result of an existential elimination subproof is its last line
 - therefore, using the term var in the last line would allow it to escape its scope

#### Existential contradiction

Handy shortcut for assuming an existential which will lead to a contradiction.

Basically, a combination of the "contradiction" and "existential elimination" subproofs
for convenience.

### Forall introduction

A term var, for example `a`, is added to the subproof's term var scope.

Reiterations *may not* mention the term var.

For the last line formula `P(a)` and bound var `x`, `∀x.P(x)` will be added to the subproof's scope.

Note that `∀x.P(x)` is the result of substituting `x` for every instance of `a` in the last line's formula.

## Results

### Distributive property

Explanation: 

 - left distribution: `P -> ( Q ^ R ) <-> ( ( P -> Q ) ^ ( P -> R ) )`
   - `->` is the outer operator
   - `^` is the inner operator
   - forwards: from left side of `<->` to right side
   - backwards: from right side of `<->` to left side
 - right distribution: `( P ^ Q ) v R <-> ( ( P v R ) ^ ( Q v R ) )`
 - anti distribution: `( P v Q ) -> R <-> ( ( P -> R) ^ ( Q -> R ) )`

| Inner | Outer | Left | Right |
| - | - | - | - |
| ^ | ^ | Both | Both |
| ^ | v | Both | Both |
| ^ | -> | Both | Anti |
| v | ^ | Both | Both |
| v | v | Both | Both |
| v | -> | Both | Anti |
| -> | ^ | Forwards | ? |
| -> | v | Both | ? |
| -> | -> | Both | Forwards |

Quantifiers:

TODO
