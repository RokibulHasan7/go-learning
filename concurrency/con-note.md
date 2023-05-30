# Concurrency Notes

- Concurrency refers to a programming language's ability to deal with lots of things at once.
- Parallelism need multiple CPUs.
    - Parallelism is like 2 car in 2 Roads.
    - Concurrency is like 2 car in 1 road but in 2 Len.
- Channels are one of the more convenient ways to send and receive notifications.
- There are a couple different types of channels:
  - Unbuffered channel: Unbuffered channels require both the sender and receiver to be present to be successful operations. It requires a goroutine to read the data, otherwise, it will lead to deadlock. By default, channels are unbuffered.
  - Buffered channel: Buffered channels have the capacity to store values for future processing. The sender is not blocked until it becomes full and it doesn't necessarily need a reader to complete the synchronization with every operation.
- Goroutines are lightweight threads of execution used in Go to support concurrency.
- Declaration of channels based on directions
  1. Bidirectional channel : chan T
  2. Send only channel: chan <- T
  3. Receive only channel: <- chan T
- 