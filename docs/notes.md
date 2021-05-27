# Notes


## Key areas of focus during contract review

Of the security-relevant items for solidity contracts, the key issues surround race-conditions and recursion.

The order of operations is extremely strict.

The EVM can be halted at any point. Contracts can exploit recursion logic bugs in other contracts (reentry).

tl;dr: same issues as webapp and DB race conditions, but probably even easier to exploit.

## Tooling

ToB's manticore: https://github.com/trailofbits/manticore

## CTF-like challenges

ToB's manticore challeges: https://github.com/trailofbits/manticore-examples

Openzeppelin's ethernaut wargame: https://ethernaut.openzeppelin.com/

