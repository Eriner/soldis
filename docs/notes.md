# Notes

Of the security-relevant items for solidity contracts, the key issues surround race-conditions and recursion.

The order of operations is extremely strict.

The EVM can be halted at any point. Contracts can exploit recursion logic bugs in other contracts (reentry).

tl;dr: same issues as webapp and DB race conditions, but probably even easier to exploit.
