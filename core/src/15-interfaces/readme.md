![](/assets/gologo.png)

# The Go programming Language - Interfaces

Interfaces are a core concept of Go and can be a complicated topic. They are everywhere in the standard library and are an abstract concept. So what exactly is an interface? It's a collection of method signatures that an object or named type can implement. Interfaces define the behaviour and can achieve polymorphism however they are not `generics` like can be found in other languages and the call for true `generics` in Go is one of the hotly fought subjects that has been running for years now.

You may think of it as a classification system, in order for a named type to be of, or adhere to a specific interface it must implement the methods of that interface. It is therefore preferred in go to hae small/tight interfaces and object or named type can honour several than have massive ones that increase the cognitive complexity and reduce possible re-usage.

## basic interface implementation

![](/core/src/15-interfaces/assets/1501-interfaces.png)

## polymorphic tendencies with interfaces

![](/core/src/15-interfaces/assets/1502-polymorphic.png)

## type assertions and type switches

![](/core/src/15-interfaces/assets/1503-assertions.png)

In Go we cannot have an interface implement another interface as a form of inheritance. Instead we can embed an interface in another, so a peer extension more than a parent-child relationship. We can see an example of this below where we are embedding interfaces and the interface is only complete once we have actually implemented a method we're not using. \_this leads to a principle of only having what you truly need in an interface as implementing methods just for the sake of compliance is bad practice. A better practice is to have a variable implement multiple interfaces than ghost implementing pieces of larger interfaces.

![](/core/src/15-interfaces/assets/1504-embedded.png)
