# Lorenz

This package is given as an example of a simple algorithm for converting a
set of input parameters into art. It is a simple implementation of the [Lorenz
System](https://en.wikipedia.org/wiki/Lorenz_system), which is a co-creation of
Margaret Hamilton, Ellen Fetter and Edward Lorenz, first used in 1963 to model
atmospheric convection.

The actual operation of this model is totally unimportant for the purposes of
an artist who is hoping to create an algorithm on ABOUND. All that is needed to
understand is that this package describes the inputs that it wants to take from
the configuraiton JSON (the `Config` struct), and offers a method to convert
the configuration into an image (the `Run` method). In this way it offers a
template that you can copy and hollow out, replacing it with the logic that
produces your art.
