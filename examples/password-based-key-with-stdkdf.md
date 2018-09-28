# password-based private key

Since Curve25519 accepts any 32 bytes as a private key, we just need a tool to generate pseudorandom
32 bytes reliably and deterministically from user input - enter password-based key derivation.

I built such a tool with a very easy user interface, which wraps the winner of the
[Password Hashing Competition](https://password-hashing.net/):
[Argon2](https://github.com/P-H-C/phc-winner-argon2/blob/master/argon2-specs.pdf). It's called
`stdkdf` and can be found at [github.com/ansemjo/stdkdf](https://github.com/ansemjo/stdkdf).

## idea

I was trying to wrap my head around using password-derived symmetric keys for mass-encryption of
messages or files, which could later be decrypted with only the password originally used to derive
said key. Specifically, I wanted to integrate this function directly into
[aenker](https://github.com/ansemjo/aenker), so that I would not have to store the used key
derivation method in the file on disk, which would allow the file to consist of only
pseudorandom-looking parts, effectively storing a long string of pseudorandom data on disk.

This approach presents difficulties in handling the nonces though: a random nonce per message would
break the precomputability of the password-based key, as different nonces lead to different keys.
But using a fixed nonce could potentially weaken low-entropy password - or would need to be stored
seperately for later use.

And then it hit me: what I was trying to do was essentially public-key cryptography! Even better:
with public key cryptography I would not need to keep the precomputed value secret at all! As
described above, Curve25519 ...

- accepts any 32 bytes as a private key
- can calculate the public key from the private key
- and can use that public key to derive many shared secrets for encryption
- owner can derive private key with password-based kdf
- owner only needs private key and ephemeral public key to decrypt
- ephemral public key is essentially pseudorandom data, so could be used (with additional hashing)
  as a nonce

## commandline

Integration into `aenker` is still _todo™_ at this point but the concept looks like this on the
commandline:

Generate a public key from a password:

    $ stdkdf -salt aenker -cost hard | curvekey pubkey | tee public
    Enter password:
    G1e2l4T/roWRAOqbgAzMreD7fKX0X4hLLkWILWGS+nw=

Distribute this public key to the person encrypting, who then derives ephemeral shared keys:

    $ curvekey shared < public
    shared secret:
      4XRK1IZdqq9e1XIGdgGX5ElX/fr6NKxgoigb0w6wEww=
    ephemeral public key:
      Q0C8L8FsjdxbmDfqukqh5KUlgwSo0QHyHj0hGUZQ+Cg=

Use the shared secret to encrypt data, concatenate with the ephemeral public key and send back.

The owner can then compute the shared secret aswell and decrypt the data:

    $ stdkdf -salt aenker -cost hard | curvekey shared --key /dev/stdin --peer Q0C8L8FsjdxbmDfqukqh5KUlgwSo0QHyHj0hGUZQ+Cg=
    Enter password:
    shared secret:
      4XRK1IZdqq9e1XIGdgGX5ElX/fr6NKxgoigb0w6wEww=
