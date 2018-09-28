# curvekey

Perform elliptic-curve Diffie-Hellmann key exchange over [Curve25519](https://cr.yp.to/ecdh.html) on
the commandline.

# install

    go get -u github.com/ansemjo/curvekey

# synopsis

Generate a new key:

    $ curvekey new -k my.sec -p my.pub

Obtain a peer's public key and obtain a shared secret:

    $ curvekey dh -p peer.pub -k my.sec
    shared secret:
      edSr8Lko65qj4wk62S8/31iex4FFGDJyD5AAB9pjnq4=

Your partner can obtain the same secret with his secret key:

    $ curvekey dh -p my.pub -k peer.sec
    shared secret:
      edSr8Lko65qj4wk62S8/31iex4FFGDJyD5AAB9pjnq4=

For an ephemeral key omit the secret key:

    $ curvekey dh -p peer.pub
    shared secret:
      MblUjj06Rh8Q/V0HoPl6nrFavXmCEuHChG5HM0hdnJ4=
    ephemeral public key:
      Emqq9yeFdhhspTW6aldqWWXSOpLxHDL4kVZUqMlKYAE=

Transmit the ephemeral public key to your peer:

    $ curvekey dh -k peer.sec -p Emqq9yeFdhhspTW6aldqWWXSOpLxHDL4kVZUqMlKYAE=
    shared secret:
      MblUjj06Rh8Q/V0HoPl6nrFavXmCEuHChG5HM0hdnJ4=

Consult the command help at any time:

    $ curvekey --help
    $ curvekey help shared
    ...

# examples

Other usage examples and scenarios can be found in [examples](examples/):

- [password-based](examples/password-based-key-with-stdkdf.md) keys generated with
  [ansemjo/stdkdf](https://github.com/ansemjo/stdkdf)

# warning

This is **unauthenticated** Diffie-Hellmann. You should

- transmit your public keys over secure channels
- post them at a public place controlled by you
- verify their integrity over a second channel
- sign them by other means, e.g. GnuPG

Otherwise you may be susceptible to a man-in-the-middle attack.

And anyway: this was mainly intended as a thought experiment and another fun commandline tool to go
with [ansemjo/stdkdf](https://github.com/ansemjo/stdkdf) and
[ansemjo/aenker](https://github.com/ansemjo/aenker). Do not rely on this.

# license

Copyright (c) 2018 Anton Semjonov Licensed under the [MIT License](LICENSE)
