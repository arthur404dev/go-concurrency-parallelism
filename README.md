# Golang Concurrency vs Parallelism

## Single Threaded vs Multi-Threaded Programs

Imagine You have to mine some ores using gophers, you'll have **Gary** as our worker:

![gopher](https://miro.medium.com/max/700/1*bFlCApzWW8EYVmSAnXcWYA.jpeg)

A common way of performing this task on single Threaded applications is by using **Gary** through all the stages of the mining, like this:
![gary working](https://miro.medium.com/max/700/1*ocFND1VTSp89syQdtvestg.jpeg)

That's fine and all, but it really doesn't take advantage of maybe assigining different tasks to different workers, like this:

![Gary Jane Peter](https://miro.medium.com/max/700/1*TAzVDPM6qAZI90yPLkvI7g.jpeg)

That's what we call concurrent programming, defined by:

> In computer science, concurrency is the ability of **different parts** or **units** of a program, algorithm, or problem to be **executed out-of-order** or at the same time simultaneously partial order, **without affecting** the final outcome.

While parallelism is defined by:

>The term Parallelism refers to **techniques to make programs faster** by performing **several computations at the same time**. This requires hardware with **multiple** processing units. In many cases the sub-computations are of the same structure, but this is not necessary.

## **Case Study** : URL Status Checker
---

1. Application Basic Logic:

```mermaid
flowchart LR
    A[URL Status Checker] ---> B[http request]
    B -.-> C[http://x-team.com]
    B -.-> D[http://github.com]
    B -.-> E[http://stackoverflow.com]
    B -.-> F[http://google.com]
    subgraph URLS
    C
    D
    E
    F
    end
```
2. This program has a **natural** blocking architecture:
```mermaid
flowchart LR
C[http://x-team.com]
D[http://github.com]
E[http://stackoverflow.com]
F[http://google.com]
subgraph URLS
C
D
E
F
end
C --> |wait|D
D --> |wait|E
E --> |wait|F

URLS --> A[need to receive the http response before moving to the next URL]
```
3. Why not try to optimize this?
```mermaid
flowchart 
C[http://x-team.com]
D[http://github.com]
E[http://stackoverflow.com]
F[http://google.com]

C -->|make request| X[return status]
D -->|make request| Y[return status]
E -->|make request| Z[return status]
F -->|make request| W[return status]

subgraph URLS
C
D
E
F
end

```
4. Behind the Scenes
```mermaid
flowchart TB
subgraph Routines
A
B
C
D
E
W
Z
J
X
Y
end
A[CPU Core]<-->B[GO Scheduler]
B <-.-> C[Go Routine]
B <-.-> |pause/unpause|D[Go Routine]
B <-.-> E[Go Routine]
X[CPU Core]<-->Y[GO Scheduler]
Y <-.-> W[Go Routine]
Y <-.-> Z[Go Routine]
Y <-.-> J[Go Routine]
subgraph Running Program
A1[Main Routine]-->C1(Created when program is launched)
B1[Child Go Routine]-->D1
B2[Child Go Routine]-->D1
B3[Child Go Routine]-->D1
D1(Created by the go keyword)
end

```
5. Why the child routines didn't run?
```mermaid
gantt
    title Time -->
    dateFormat DD-HH
    axisFormat %L
    section Program
    main routine:active,crit,mr,01-00,5h
    program exit: crit,after mr,4h
    child go routine:active,mr2,01-02,6h
    child go routine:active,mr2,01-02,7h
    child go routine:active,mr2,01-02,5h
```
6. The plan with Channels
```mermaid
flowchart TB
A[Main Routine] <--> C{Channel}
subgraph Routines
D[Child Go Routine]
E[Child Go Routine]
F[Child Go Routine]
end
D <-.-> C
E <-.-> C
F <-.-> C
```