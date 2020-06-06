# Generic OTP Service
Just a project to learn GO and a thought experiment around OTP generation.

## Prerequisites

Just some notes. To get this working:

1. Well, you obviously need GoLang installed.
2. Assuming you mustered the intelligence to install GO, you will need to run `go mod download` to install dependencies for this project. If it was not immediately obvious, do that in the root directory of the project.

## Starting - Local

To start local, just run `go run main.go` and it should work. If not, well, I guess I screwed up.

## Starting - Docker

Remember to create your own `Dockerfile` based on the template provided.  

If you have `node.js` installed, you can use the scripts in `package.json` to build and start the `docker` image. In order to do so, simply run these commands in sequence:

1. `npm run docker:build`
1. `npm run docker:start`

If you don't have `node.js`, well, just copy the docker commands in the `package.json` file.

By the way, not having `docker` installed is one way for this not to work, if that was not immediately obvious.

### Notes
1. To update swagger, run `swag init` on the root directory. If you do this without getting the dependencies installed, you need to get your head checked.
    1. If you have `node.js` installed and want to automate this, you can run `npm run swagger:init` to achieve the same effect. Link this up with your build automation tool or IDE, if you cannot call the aforementioned command directly. 
