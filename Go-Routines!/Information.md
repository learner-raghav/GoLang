-- Concurrency And Parallelism!!

1. Concurrency, by definition, is the ability to break down a computer program or algorithm
    into individual parts, which can be executed independently. The final outcome of a concurrent program is the
    same as that if a program which has been executed sequentially.
   
2. Concurrency is about dealing with lots of things at once, Whereas parallelism is 
    about doing lots of things at once.
   
3. Parallelism is when we break up the task into subtasks and execute them simultaneously. Each
    of the subtasks is independent, and may or may not be related. In short, we carry out 
   many computations at the same time in parallelism.
   
4. Parallel computing helps us solve large problems efficiently by using more than one CPU
    to execute multiple computations at the same time, which saves time in case of large 
    datasets.
   
5. However, Parallel programming is hard to achieve as we need to ensure the independence
    of tasks when it comes to dividing the problem and sharing the data.
   
6. Concurrency is about structure, Parallelism is about structure. We design and dtructure a bug problem
    into smaller problems which can be solved independently. So concurrecny has to be there for parallelism to exist.
   
7. In conculsion, concurrency is about composing a solution to a problem whicle parallelism 
    is solving the problem by running things in parallel. Remember concurrency does not imply parallelism
    It is just a way to structure and design tasks independently so that we can use parallelism to make
    them efficient.
   
-- Concurrent Sequential Memory!!

1. Concurrent processes have to operate individually but with a shared data source. The
    concurrent sequentail memory 