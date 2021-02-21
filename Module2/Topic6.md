1. Boolean type: The possible values of this type are predefined constants true and false. Example of it can be `var a bool = true`

2. Numerical Type:

    1. Integers and floating point numbers

        1. Go has architecture dependent types such as int, uint, uintptr. They have appropriate length for the machine on which the program runs.

        2. An int is a default signed type, which implies it takes a size of 32 bit on a 32 bit machine and 64 bit on a 64 bit machine. 

        3. We have int8,int16,int32,int64, similarly we have for uimt (unsigned int) and then we have the float values.
    
3. The unicode package has some useful functions for testing characters. Suppose, we have a character named ch. Then the following methods might be useful

    1. unicode.isLetter(ch)
    2. unicode.isSpace(ch)
        