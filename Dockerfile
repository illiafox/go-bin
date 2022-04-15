# build stage
FROM golang:1.18-alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /server
RUN cd /server/cmd/server && go build -o bin

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /server/cmd/ /app/cmd/

# html files
COPY --from=build-env /server/shared/ /app/shared/

WORKDIR /app/cmd/server
ENTRYPOINT ./bin $ARGS