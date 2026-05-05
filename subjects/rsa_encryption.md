## Exercise 07: Pretty Bad Privacy

Implement RSA encryption/decryption.

Subcommands:
1. `gen-keys <pub> <priv>`: Generate prime numbers p, q and derive keys.
2. `encrypt <pub>`: Encrypt stdin.
3. `decrypt <priv>`: Decrypt stdin.

Formula: `encrypt(m) = m^E % M`
