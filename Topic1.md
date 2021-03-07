SOLID design principles

1. The single responsibility principle. A class should have one and only one reason to change.
    A code that has a single responsibility has the fewest reasons to change.
   There should be less coupling and less cohesion. A good example to explain this is 
   the ReaderWriter interface. We should use Reader when required and Writer separately.
   
2. Software entities should be Open for extension and closed for modification. As we can see in the case of embedding, we are open to 
    the use of the year field, but when it comes to modification of the method, a different method is called.
   
3. Liskov substitution principle basically says that if s is a subtype of T, then whereever, we use T, we can replace it by S
    For example - because all types can now be represented as byte array.
   type Reader interface {
    Read(buf []bytes) (n int, err error)   
}
   
4. No client should depend on the methods they don't use. We shoudl define more fine grained methods.
    func Save(rwc io.ReadWriteCloser,doc *Document)
   
5. The principle of dependency inversion suggests that high level modules should not 
    depend on low level modules, Both should depend on abstractions.
   

CRUX

1. The single responsibility principle encourage to structure the functions, types, 
    and methods into packages and should serve single purpose.
   
2. The liskov susbtrituition principle encourages us to express the dependencies between your package in terms of interfaces.
    and not concrete types!!
   
3. Dependency inversion basically helps in moving the knowledge of things our package depends on from compile time
to run time with the help of abstractions (interfaces)