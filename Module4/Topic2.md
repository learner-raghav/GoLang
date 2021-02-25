1. The defer keyword allows us to postpone the execution of a statement or a function until 
    the end of the enclosing function.
   
2. Defer executes something when the enclosing function returns. This happens after every return even when an 
    error occurs in the midst of executing the function.
   
3. The defer resembles the finally block in OO-Languages as Java, C#; in most cases,
    it also serves to free up allocated resources. This is a very good usecase of the defer statement.
   
4. The defer function allows us to guarantee that certain clean up tasks 
    are performed before we return from a function. Example of those tasks can be
    
    1. Closing a file stream
    2. Unlocking a locked resource.
    3. Prinitng a footer in a resource.
    4. Closing a database connection
    
5. 