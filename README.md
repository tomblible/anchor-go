# anchor-go

![logo](logo.png)

## usage

```bash
anchor-go --src=/path/to/idl.json
```

Generated Code will be generated and saved to `./generated/`.

## TODO

- [x] instructions
- [x] accounts
- [x] types
- [ ] events
- [ ] errors
- [ ] handle tuple types
- [ ] constants

## Future Development

TBD

## what is anchor-go?

`anchor-go` generates Go clients for [Solana](https://solana.com/) programs (smart contracts) written using the [anchor](https://github.com/project-serum/anchor) framework.

## what is anchor?

Link: https://github.com/project-serum/anchor

```
Anchor is a framework for Solana's Sealevel runtime providing several convenient developer tools for writing smart contracts.
```

## I have an anchor program; how do I generate a Go client for it? (step by step)

### example 1: metaplex nft candy machine
```convert
anchor idl convert  old.json >new.json
```

```bash
make swap-proxy
```

Note
----

- anchor-go is in active development, so all APIs are subject to change.
- This code is unaudited. Use at your own risk.
