
# Left distributive

| Outer | Inner | Forward | Backward | Left | Right |
| - | - | - | - | - | - |
| ^ | ^ |  |  | P ^ ( Q ^ R ) | ( P ^ Q ) ^ ( P ^ R ) |
| ^ | v |  |  | P ^ ( Q v R ) | ( P ^ Q ) v ( P ^ R ) |
| ^ | -> |  | P: F, Q: T, R: T | P ^ ( Q -> R ) | ( P ^ Q ) -> ( P ^ R ) |
| ^ | <-> |  | P: F, Q: T, R: T | P ^ ( Q <-> R ) | ( P ^ Q ) <-> ( P ^ R ) |
| v | ^ |  |  | P v ( Q ^ R ) | ( P v Q ) ^ ( P v R ) |
| v | v |  |  | P v ( Q v R ) | ( P v Q ) v ( P v R ) |
| v | -> |  |  | P v ( Q -> R ) | ( P v Q ) -> ( P v R ) |
| v | <-> |  |  | P v ( Q <-> R ) | ( P v Q ) <-> ( P v R ) |
| -> | ^ |  |  | P -> ( Q ^ R ) | ( P -> Q ) ^ ( P -> R ) |
| -> | v |  |  | P -> ( Q v R ) | ( P -> Q ) v ( P -> R ) |
| -> | -> |  |  | P -> ( Q -> R ) | ( P -> Q ) -> ( P -> R ) |
| -> | <-> |  |  | P -> ( Q <-> R ) | ( P -> Q ) <-> ( P -> R ) |
| <-> | ^ | P: F, Q: T, R: F |  | P <-> ( Q ^ R ) | ( P <-> Q ) ^ ( P <-> R ) |
| <-> | v |  | P: F, Q: T, R: F | P <-> ( Q v R ) | ( P <-> Q ) v ( P <-> R ) |
| <-> | -> |  | P: F, Q: T, R: T | P <-> ( Q -> R ) | ( P <-> Q ) -> ( P <-> R ) |
| <-> | <-> | P: F, Q: T, R: F | P: F, Q: T, R: T | P <-> ( Q <-> R ) | ( P <-> Q ) <-> ( P <-> R ) |

# Right distributive

| Outer | Inner | Forward | Backward | Left | Right |
| - | - | - | - | - | - |
| ^ | ^ |  |  | ( P ^ Q ) ^ R | ( P ^ R ) ^ ( Q ^ R ) |
| ^ | v |  |  | ( P v Q ) ^ R | ( P ^ R ) v ( Q ^ R ) |
| ^ | -> |  | P: T, Q: T, R: F | ( P -> Q ) ^ R | ( P ^ R ) -> ( Q ^ R ) |
| ^ | <-> |  | P: T, Q: T, R: F | ( P <-> Q ) ^ R | ( P ^ R ) <-> ( Q ^ R ) |
| v | ^ |  |  | ( P ^ Q ) v R | ( P v R ) ^ ( Q v R ) |
| v | v |  |  | ( P v Q ) v R | ( P v R ) v ( Q v R ) |
| v | -> |  |  | ( P -> Q ) v R | ( P v R ) -> ( Q v R ) |
| v | <-> |  |  | ( P <-> Q ) v R | ( P v R ) <-> ( Q v R ) |
| -> | ^ | P: T, Q: F, R: F |  | ( P ^ Q ) -> R | ( P -> R ) ^ ( Q -> R ) |
| -> | v |  | P: T, Q: F, R: F | ( P v Q ) -> R | ( P -> R ) v ( Q -> R ) |
| -> | -> |  | P: T, Q: T, R: F | ( P -> Q ) -> R | ( P -> R ) -> ( Q -> R ) |
| -> | <-> | P: T, Q: F, R: F | P: T, Q: T, R: F | ( P <-> Q ) -> R | ( P -> R ) <-> ( Q -> R ) |
| <-> | ^ | P: T, Q: F, R: F |  | ( P ^ Q ) <-> R | ( P <-> R ) ^ ( Q <-> R ) |
| <-> | v |  | P: T, Q: F, R: F | ( P v Q ) <-> R | ( P <-> R ) v ( Q <-> R ) |
| <-> | -> |  | P: T, Q: T, R: F | ( P -> Q ) <-> R | ( P <-> R ) -> ( Q <-> R ) |
| <-> | <-> | P: T, Q: F, R: F | P: T, Q: T, R: F | ( P <-> Q ) <-> R | ( P <-> R ) <-> ( Q <-> R ) |

# Associative

| Op | Forward | Backward | Left | Right |
| - | - | - | - | - |
| ^ |  |  | ( P ^ Q ) ^ R | P ^ ( Q ^ R ) |
| v |  |  | ( P v Q ) v R | P v ( Q v R ) |
| -> |  | P: F, Q: T, R: F | ( P -> Q ) -> R | P -> ( Q -> R ) |
| <-> |  |  | ( P <-> Q ) <-> R | P <-> ( Q <-> R ) |

# Transitivity

| Op | Forward | Backward | Left | Right |
| - | - | - | - | - |
| ^ |  | P: T, Q: F, R: T | ( P ^ Q ) ^ ( Q ^ R ) | P ^ R |
| v | P: F, Q: T, R: F | P: T, Q: F, R: F | ( P v Q ) ^ ( Q v R ) | P v R |
| -> |  | P: T, Q: F, R: T | ( P -> Q ) ^ ( Q -> R ) | P -> R |
| <-> |  | P: T, Q: F, R: T | ( P <-> Q ) ^ ( Q <-> R ) | P <-> R |

# Commutativity

| Op | Forward | Backward | Left | Right |
| - | - | - | - | - |
| ^ |  |  | P ^ Q | Q ^ P |
| v |  |  | P v Q | Q v P |
| -> | P: F, Q: T | P: T, Q: F | P -> Q | Q -> P |
| <-> |  |  | P <-> Q | Q <-> P |
