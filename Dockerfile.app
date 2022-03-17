ARG APP_NAME=fiber-starter-project

FROM registry.gitlab.com/ansidev/docker:golang-builder-latest AS builder

FROM registry.gitlab.com/ansidev/docker:golang-prod-latest as production

ENTRYPOINT ["/app/fiber-starter-project"]
