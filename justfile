# build project
build: clean
     go build -o ./bin/trackingsrv -ldflags="-s -w" cmd/*

# build & run server
run: build
    ./bin/trackingsrv tracking --host 127.0.0.1 -p 5000

# clean build directory
clean:
     @[ -d "./bin" ] && rm -r ./bin && echo "bin directory cleaned" || true

# build and compress binary
upx: build
    upx --best --lzma bin/trackingsrv

#build docker image
image tag:
    docker buildx build --build-arg GITHUB_TOKEN="$GITHUB_TOKEN" --tag {{tag}} .