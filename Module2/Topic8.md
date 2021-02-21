1. Go gives the programmers control over which data structure is a pointer and which is not, however, calculations with pointer values are not supported in Go programs.

2. By giving the programmer control over basic memory layout, Go provides the ability to control the size of a given collection of Data structures, the number of allocations, the memory access patterns, all of which are important for building systems that perform well.

3. Pointers in Go

    1. Programs store values in memory and each memory block (or word) has an address, which is usually represented as a hexadecimal number, like 0x6b0820

    2. Go has an & operator which when placed before a variable, gives the memory address of that variable.

    3. The address of  a particular memory location can be stored in a special data type called as a pointer. In the above case it is referred to as int. So, i1 is denoted by *int. `var intP *int` and then we do `intP = &i1`, intP stores the memory address of i1 which means it points to the location of i1. In other words, it references the variable i1.

    4. The size of a pointer variable is 4 bytes on a 32 bit machine and 8 bytes on a 64 bit machine, regardless of the size of the value they point to. Using a pointer to refer to a value is called an indirection. 

    5. Calculations over pointer concept does not exist in Go and hence make it memory safe. 

    6. When the program has to work with variables, that occupy a lot of memory, working with pointers can reduce memory usage and increase efficiency.

    7. Because of pointer indirection, sometimes using pointers can make the code unnecarrily complex. 