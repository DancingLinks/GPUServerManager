# Build docker file

mkdir build
cp -r web build
cp -r static build
cp Dockerfile build
if [ ! -d build/static/server.json  ];then
  touch build/static/server.json
fi

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/go_build ./main

cd build

docker build -t dancinglinks/gpu-manager .

cd ..
rm -rf build
