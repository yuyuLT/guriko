# docker-golang-boilerplate

Docker template project for golang developer

## How to develop
1. Write codes in {ProjectRoot}/app dir.

    - Set main script name to main.go.

    - Command go mod get or go mod tidy is unnecessary.

      (Only write import section in codes)

1. Copy .env.template to .env and configure it.

1. Execute ./init.sh shell script.

1. Execute ./run.sh shell script.

1. Modify codes in {ProjectRoot}/app dir.

1. Return to step 4.

## How to build production
1. Set ENVIRONMENT variable to "prod" in .env file.

1. Execute ./init.sh shell script.

1. Execute ./run.sh shell script.

## Run.sh script's sub command
- none

    - if ENVIRONMENT is set to prod

        build binary and run it in the container

    - else ENVIRONMENT is set to something other than prod

        run main.go in the container
- air

    - Independent of ENVIRONMENT

        hot reload main.go
