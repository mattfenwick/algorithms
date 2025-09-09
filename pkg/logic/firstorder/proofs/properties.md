
# Left distributive

| Outer | Inner | Forward | Backward | Left | Right |
| - | - | - | - | - | - |
| ^ | ^ |  |  | P ^ ( Q ^ R ) | ( P ^ Q ) ^ ( P ^ R ) |
| ^ | v |  |  | P ^ ( Q v R ) | ( P ^ Q ) v ( P ^ R ) |
| ^ | -> |  | P: F, Q: T, R: T | P ^ ( Q -> R ) | ( P ^ Q ) -> ( P ^ R ) |
| v | ^ |  |  | P v ( Q ^ R ) | ( P v Q ) ^ ( P v R ) |
| v | v |  |  | P v ( Q v R ) | ( P v Q ) v ( P v R ) |
| v | -> |  |  | P v ( Q -> R ) | ( P v Q ) -> ( P v R ) |
| -> | ^ |  |  | P -> ( Q ^ R ) | ( P -> Q ) ^ ( P -> R ) |
| -> | v |  |  | P -> ( Q v R ) | ( P -> Q ) v ( P -> R ) |
| -> | -> |  |  | P -> ( Q -> R ) | ( P -> Q ) -> ( P -> R ) |

# Right distributive

| Outer | Inner | Forward | Backward | Left | Right |
| - | - | - | - | - | - |
| ^ | ^ |  |  | ( P ^ Q ) ^ R | ( P ^ R ) ^ ( Q ^ R ) |
| ^ | v |  |  | ( P v Q ) ^ R | ( P ^ R ) v ( Q ^ R ) |
| ^ | -> |  | P: T, Q: T, R: F | ( P -> Q ) ^ R | ( P ^ R ) -> ( Q ^ R ) |
| v | ^ |  |  | ( P ^ Q ) v R | ( P v R ) ^ ( Q v R ) |
| v | v |  |  | ( P v Q ) v R | ( P v R ) v ( Q v R ) |
| v | -> |  |  | ( P -> Q ) v R | ( P v R ) -> ( Q v R ) |
| -> | ^ | P: T, Q: F, R: F |  | ( P ^ Q ) -> R | ( P -> R ) ^ ( Q -> R ) |
| -> | v |  | P: T, Q: F, R: F | ( P v Q ) -> R | ( P -> R ) v ( Q -> R ) |
| -> | -> |  | P: T, Q: T, R: F | ( P -> Q ) -> R | ( P -> R ) -> ( Q -> R ) |
