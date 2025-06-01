# SPACE-BEACON

**space-beacon** is a protocol.

## Concepts

With **space-beacon** there is a concept of a **space-base** and a **space-command**.

You can think of a **space-base** as a computer.
The purpose of the **space-base** computer is to run some _software_.
A **space-command** computer would tell the **space-base** computer what _software_ to install and run.

You can think of a **space-command** as a computer that would manage all the **space-base**s.

Typically, there is a one-to-many relationship between a **space-command** and **space-base**s.
I.e., typically, one **space-command** managers many **space-base**s.

## Protocol

The **space-beacon** protocol works within a local network.

By default, each **space-base** sends UDP messages on the **mutlicast** IP address `239.83.80.67` (`0xEF535043`) using the UDP port `21328` (`0x5350`).

Each of the **space-base**s are sending this messags so that a **space-command** can automagically discover it, so that the **space-command** can (potentially) start managing it (the **space-base**).

And, of course, by default, a **space-command** if listening for UDP messages on the **mutlicast** IP address `239.83.80.67` (`0xEF535043`) with the UDP port `21328` (`0x5350`).

## Message

The message that each **space-base** sends is:

```golang
"SPACE/0.1" + "\n" + // <--- magic
"DOROOD"    + "\n" + // <--- message type
""          + "\n"   // <--- second "\n" in a row to specify the the end of the message
```
