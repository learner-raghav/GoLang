1. Go has basically 3 conditional or branching structures. They are the following

    1. if-else - if init,cond {}
    2. switch-case - switch(cond) {case --}
    3. select

2. Testing support in functions

    1. Sometimes the functions in Go are defined such that they return two results. One is the value and the other is the status of the execution. For example - a function will return a value and true in case of successful execution. Whereas it will return a value (probably nil) and false in case of an unsuccessful execution.

    2. Instead of true and false, an error-variable can be returned. In case of successful execution, the error is nil. Otherwise, it contains the error information. It is then obvious to test the execution with an if statement because of its notation, this is often called the comma, ok pattern.

                `v,ok = sample_function(parameter)`
    
    3. The value goes to v, and the ok parameter hold the status of the error during the execution. If there is no error, sample_function will return true, otherwise it will return an error value.

    4. `an,err = strconv.Atoi(orig)`, if it is convertible to inetger, it returns true or it shows nil.

3. Switch - Case construct.

    1. switch statement is used when we don't necessarily want multiple if-else blocks. We have multiple cases and then the default block follows. If we have a fallthrough, it will still check the next case.