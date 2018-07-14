# GoToBooX project

Description: The project is a golang based web application that is a book-crossing automated system. It allows to manage online member accounts, book reservation and their usage tracking.

A barebones Go app, which is deployed on Heroku.

This application [is here](https://gotoboox.herokuapp.com/) - check it out.

## Running Locally

Make sure you have changed the port in the router file. Comment section will help you.


## Documentation

To use the application the following system env must be set

    export POSTGRES_URL="postgres://niuaznefoznzqh:33f9db3d3723c0a337c18e6f0c599d358765159048ab0c4ec5a1d28969616854@ec2-54-217-214-68.eu-west-1.compute.amazonaws.com:5432/dcclgdtqr61bti"
    export PORT=8080

Later... maybe


## For the team

Every change in master will affect heroku's remote server. So be careful

## Mockgen

To generate a mock, for example for UsersRepo,use:

    mockgen -destination=src/mocks/mock_usersRepo.go -package=mocks github.com/metalscreame/GoToBoox/src/dataBase/repository UserRepository


To pass code through metalinter use:

    gometalinter.v2 --vendor --tests --skip=mocks --exclude='_gen.go' --exclude='docs.go' --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml