1. A function with no parameters is called a niladic function in go. In Go, function 
    overloading is not allowed as it creates a lot of overhead at the compile time and hence reduces 
   performance.
   
2. All the basic data types in go are passed by value (struct,int,bool,string etc) whereas 
    maps,slices,interfaces and channels are passed by reference in go.
   
3. When we have Named return variables, we necessarily need not return values instead those
    values are automatically returned if the values are named.
   
4. If the last parameter of a function is followed by ...type, this indicates that the function
    can deal with a variable number of parameters of that type, possibly also a 0. Such 
   functions are called as variadic functions.
   
