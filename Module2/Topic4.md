1. Variables contain data and data can be of different types or different data types in Go. Go is a statically typed language. It means the compiler must know, the types of variables, either because it was explicitly indicated or because the compiler can infer the type from the code context.

2. A type defines a set of values and the set of operations that can take place on those values. Here is an overview of some categories or types.

        elementary - int,float,bool,string
        structured - struct,array,slice,map,channel

3. Sometimes in various programming languages, we need to convert a value into value of another type. This is called as type casting. Go does not allow implicit conversion, which means Go never does such a conversion by itself. The conversion must explicitly be done as valueOfTypeB = typeB(valueOfTypeA)

4. A value that cannot be changed by the program is called a constant. This data can only be of type boolean, number or string. We can define a constant in Go using the const keyword.

