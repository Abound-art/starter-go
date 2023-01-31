# Main 

This package contains the scaffolding for what an algorithm's main method
can/will look like. It ingests inputs, performs a computation to produce
an image, and saves that image to a file - that's truly it.

The Docker file specifies how to take this go program and produce a docker
image with it - essentially bundling it so that it can run on ABOUND's 
servers.

To see how to run/deploy it, see the repository README.md