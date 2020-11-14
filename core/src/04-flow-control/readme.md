![](/assets/gologo.png)

# The Go programming Language - Flow control

controlling the flow of programs in Go is much the same as most high level 3rd and 4th gen languages, conditional statements in `if`, `else if` and `else` blocks and the use of loops.

## Conditional statements

An `if` statement allows you to to execute a specific block of code only if the conditions are met. The `else if` block allows us to specify further conditions to be evaluated if the original condition is not met and the closing `else` block acts like a default execution block and the selected course of action where all specific conditionals are not met. It is worth noting a that a conditional construct can have zero, one or many `else if` blocks.

Let's see an `if`, `else if`, `else` block.

![](/core/src/04-flow-control/assets/401-basic-if.png)

## Loops

Computers are great at repeating tasks and an important aspect of program control flow is the use of loops to perform tasks that should be repeated under controlling conditions. Two main types of loop are the `for loop` which i useful when we know upfront how many times we would like to repeat a particular task. The second is the while loop which is is more of a _keep repeating this task while a condition exists that needs us to execute it further. Once that condition no longer holds stop looping_.

Let's see a really trivial `for loop` example. While the program does nothing useful it covers all the key elements pf a for loop.

![](/core/src/04-flow-control/assets/402-basic-for.png)

## The while loop in go

There is NO `while` loop or keyword in Go as is found in many/most other languages. We can however achieve the same result with a `for loop` that is rephrased. See example below for how we handle entry conditions, stop conditions and controller mechanisms.

![](/core/src/04-flow-control/assets/403-while-in-go.png)

## Break and Continue

in our loops we have additional controlling features such as `break` and `continue`. As the name suggests `break` allows us to break out of a loop completely. On the other hand `continue` allows us to conditionally check in the iteration of a loop for a condition where we might want to avoid the typical action can be bypassed and the loop proceeds to the next iteration.

![](/core/src/04-flow-control/assets/404-break-continue.png)

## label statements

Labels are used in `break`, `continue` and `goto` statements. Defined labels must be used or there will be a compilation time error. Labels will also not collide with identifiers that are not a label. The scope of a label is the body of the function it was declared in, this excludes any nested function. The primary use of a label is when terminating enclosing outer loops.

![](/core/src/04-flow-control/assets/405-labels.png)

## The goto statement

A goto statement can operate much like a for loop and the goto is very much a relic of traditional programming languages that is often left out of new languages. The caveat is that in the goto/label block no new variables that were not already in play can be introduced. See below for a simple example operating like a `for` loop.

![](/core/src/04-flow-control/assets/406-goto.png)

## The switch statement

The switch statement is an elegant syntax for handling cases instead of an `if, else if, else` block and adds clarity where there are many cases which would make the if-else construct a bit noisy or messy.

![](/core/src/04-flow-control/assets/407-switch.png)

## Scope - block, file & package

It's also worth noting the differences in scope in a Go program, we can have block, file and package scopes.

![](/core/src/04-flow-control/assets/04-scope.png)
