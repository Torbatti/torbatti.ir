---
title: زبان برنامه نویسی سی
---

- pick your Context model: (callbacks: libxev/libuv, stackful coroutines: [green/os]-threads, stackless 
coroutines: [zig/rust]-async, generators/state-machines)
- pick your Scheduling model: (preemptive: *threads, cooperative: the others)
- pick your IO model: (blocking: os-threads, non-blocking: event loop)


Notice that zig/rust async is only one part in this. Zig is not it's stdlib (and neither is Rust's its ecosystem). 


https://github.com/Cloudef/zig-aio

https://github.com/mitchellh/libxev



https://www.youtube.com/watch?v=Ul8OO4vQMTw


Just to layer on to what @Protty is suggesting, if you are a new(er) programmer and you are looking to understand this more, I would suggest reading some books on the subject. Get yourself a really strong foundation. 

Since Zig is a new language, I don't think any books exist for it specifically, but you can lean on other more developed language ecosystems and the knowledge is transferrable. It is really all the same thing under the hood, the abstraction layer provided by the language is the only thing that changes.

For example, Rust has this great book called "Asynchronous Programming in Rust" by Carl Frederik Samson. It covers the difference between async and parallelism, callbacks vs green threads, os-backed queues, all that good stuff. It is a Rust book to be sure, but you can ignore the Rust specific information. And I think he even walks you through implementing your own event queue from scratch.

Going even further, you can read some classic books like "Operating Systems: Three Easy Pieces" which has a section on concurrency.


threads for parallelism (run things on cores) is different from threads for concurrency (run multiple blocking things) 
Protty
 — 
08/11/2024 8:09 AM
theres libraries that implement it like zigcoro, but you dont need it necessarily: can always use callbacks in the style of libuv/libxev/etc.