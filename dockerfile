###############################
#  HOW TO RUN:
#  Use this command from within the Go project directory:
#       docker build -t opendi-api .
###############################


###############################
# --BUILD THE API BINARY--
###############################

# This uses a temporary image and container. These won't stick around.

FROM ubuntu:20.04 AS build

WORKDIR /

COPY . .

RUN apt-get update

# software-properties-common is required for add-apt-repository to work
RUN apt-get install -y software-properties-common

# Add the golang-backports repo to get the correct version of Go
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get install -y golang-1.23
ENV PATH="/usr/lib/go-1.23/bin:${PATH}"

# Sanity check our mod file
RUN go mod download && go mod verify

# Build the program
RUN go build -o opendi-api .


###############################
# --CREATE THE RUNNABLE CONTAINER--
###############################

# Using a really small starting image here.
# If you need to debug, switch this to
#   FROM debain:bullseye-slim
# and run shell:
#   docker run -it opendi-api sh
FROM gcr.io/distroless/base-debian11

WORKDIR /

# Only copy the built executable
COPY --from=build /opendi-api .

VOLUME /db-data

EXPOSE 8080

CMD ["./opendi-api"]


###############################
#  HOW TO RUN THE RESULT:
#  On Docker Desktop, click the Play button to run the image. In Optional settings, set Host port in the Ports section to 8080
#  For Docker command line, use this command:
#       docker run -p 8080:8080 opendi-api
###############################