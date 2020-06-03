# Generic OTP Service
Just a project to learn GO and a thought experiment around OTP generation.

## Prerequisites

Just some notes. To get this working:

1. Well, you obviously need GoLang installed.
2. Assuming you mustered the intelligence to install GO, you will need to run `go mod download` to install dependencies for this project. If it was not immediately obvious, do that in the root directory of the project.

## Starting - Local

To start local, just run `go run main.go` and it should work.

## Starting - Docker

If you have `node.js` installed, you can use the scripts in `package.json` to build and start the docker image. In order to do so, simply run these commands in sequence:

1. `npm run docker:build`
1. `npm run docker:start`

If you don, well, just copy the docker commands in the `package.json` file.
