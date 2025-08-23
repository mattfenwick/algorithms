# Propositional Logic

# Rules

## Connectives

| Connective | Introduction | Elimination |
| - | - | - |
| ^ | P, Q => P ^ Q | P ^ Q => P; P ^ Q => Q |
| v | P => P v Q; Q => P v Q | P -> R, Q -> R, P v Q => R |
| -> | P, Q => P -> Q | P -> Q, P => Q |
| ~ | P => ~~P | ~~P => P |
| <-> | P <-> Q => P -> Q; P <-> Q => Q -> P | P -> Q, Q -> P => P <-> Q |
| ∃ | TODO | TODO |
| ∀ | TODO | TODO |

## Subproofs

A temporary hypothesis is assumed for a nested scope.  Deductions are then carried out.
Results from enclosing scopes may only be used if explicitly reiterated.  There are two
outcomes from a subproof: Implication, and Contradiction.

Implication: let's call the hypothesis `P` and the result of the last line of the subproof
`Q`.  Then the implication `P -> Q` is added to the *enclosing* scope.

Contradiction: let's call the hypothesis `P`.  At some point, the subproof asserts `Q`; later,
the last line asserts `~ Q`.  This is a contradiction, and so `~ P` is added to the *enclosing*
scope.

## Special subproof rules

Reiterate: brings a deduction from an enclosing scope into a subproof.  Results from
enclosing scopes must be imported to be used.

### Restrictions

The hypothesis and following step results may not be used outside of their subproof.

A subproof contradiction must have as its last line an assertion which contradicts a previous
assertion in the subproof.  For example, if the subproof ends with `~ Q`, then either `Q` or
`~ ~ Q` must have previously been asserted in the subproof.
