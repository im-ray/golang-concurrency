# golang-concurrency
Experiments with Golang Concurrency
Testing The Repo Setup

Race Condition , Mutual Exclusion, Channel
```
go run -race .
```
#### Channels::

A means of allowing communication to and from a GoRoutine  

channels can be buffered, or unbuffered  

Once you're done with a channel, you must close it  

Channels typically only accept a given type or interface  


#### The Sleeping Barber Problem by Dijkstra in 1965 Enforces to Understand Concurrency  

if multiple condition matches in select statement it choose one of them randomly


