
# Monad based err handling in golang

This repository contains example of handling errors with help of monads. 
Follow commits and their messages to see this approach.

## Related materials

* https://awalterschulze.github.io/blog/post/monads-for-goprogrammers/ - 
article, which presented approach is mostly based on. 
* https://github.com/awalterschulze/goderive - 
I used this project to generate compose functions
* https://speakerdeck.com/rebeccaskinner/monadic-error-handling-in-go - 
presentation describes more generic approach, 
without a need to implement compose functions. 
However, the generic method (WrapEither) is not available in mentioned 
[Gofpher](https://github.com/asteris-llc/gofpher) library.
Moreover, using generic methods has many drawbacks, which are summarised 
[here](https://speakerdeck.com/rebeccaskinner/monadic-error-handling-in-go?slide=79). 
I finally abandoned idea to implement it. 
