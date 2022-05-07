#==============================================
#   A docker file that you can just drop in a 
#   golang main package to generate a minified docker (or podman) image.
#   
#   Usage :
#   cd yourdir
#   docker build -t myimage .
#   docker run myimage
#
#   (c) xavier268@github.com
#   2022-05-07 : initial version
#===============================================


# First, we prepare an image to compile the program ...
FROM golang:alpine as COMPILER

WORKDIR /build
RUN apk update && apk upgrade
# disabling CGO
ENV CGO_ENABLED="0"     
RUN go env

# make sure you add any useful file or folder that is needed to build by adding to the instruction below !
COPY *.go go.mod go.sum  ./
RUN go build -v -o out .
RUN ls -lah out

# Then, we copy just the compiled executable into an empty image.
FROM scratch

COPY --from=COMPILER /build/out  /out

# You may want to copy other files, if any is needed for runtime ..
# COPY ....
# NB: you need this form of CMD to be sure not to call sh, that is not available here...
CMD ["/out"]






