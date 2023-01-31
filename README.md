# ABOUND Starter Repo - Golang

This is a starter repository for an ABOUND algorithm written in golang.

* Unsure on what ABOUND is? Check out https://abound.art
* Looking for another language? Check out our other options for starter repos [here](https://abound.art/artists)
* New to Golang Modules? Make sure to understand the basics [here](https://go.dev/blog/using-go-modules).

## What is in this repo

This repo includes all of the scaffolding to execute a function that takes
in arbitrary input data (that can be JSON serialized) and produce an image
(either a PNG or SVG). In short, this repo does everything except implement your
art algorithm, which will generally look like this:

```
type MyAlgoConfig struct {
    Seed int `json:"seed"`
    // ... any arguments/parameters your algo needs here
}

func (config *MyAlgoConfig) Run() (image.Image, error) {
    // Your code here which generates the image from the config.
}
```

It also includes an example art function (a lorenz attractor)
which stands in for the code you would write to implement an art algorithm.

## How to run locally / test your code

```
./lorenz/main/test/run.sh
```

will generate a piece of art at `/lorenz/main/test/output.png`
from the input parameters in `lorenz/main/test/input.json`.
This is just a thin wrapper around `go run ./lorenz/main/main.go`,
which sets the appropriate environment variables to consume `input.json` and produce
(`output.png`). 

You can copy the lorenz package and its contents to get started implementing
your algorithm, and then run `./myalgo/main/test/run.sh` to test it out locally
using your test input configuration.

## Packaging for deployment

From the root of the repository, run 

```
docker build -f lorenz/main/Dockerfile .
```

## Deploying on ABOUND 

Head to https://abound.art/artists for the most recent instructions on how to upload
your algorithm once it is written. Make sure to read through the constraints carefully
to make sure that your algorithm conforms to them prior to submission.

Once you're ready to upload, tag the binary with this command


```
docker build -t [Tag given by ABOUND] -f [Your Algo]/main/Dockerfile .
```
