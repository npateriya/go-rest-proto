
## buildpack services

- Adding buildpacks
Adds new build packs to support list

```
curl   -X POST --data @/tmp/buildpack.json localhost:8888/buildpack
cat /tmp/buildpack.json 
{
"buildpack_id" : "abc124",
"name" : "golang",
"image_url" : "image/golang.png",
"default_build_command" : "go build .",
"default_test_command" : "go test .",
"default_cpu" : 1,
"default_ram" : 512
}
```

- Querying buildpacks
Lists all the build packs supported

```
curl  localhost:8888/buildpack
[{"buildpack_id":"abc123","name":"golang","image_url":"image/golang.png","default_build_command":"go build .","default_test_command":"go test .","default_cpu":1,"default_ram":512},{"buildpack_id":"abc124","name":"golang","image_url":"image/golang.png","default_build_command":"go build .","default_test_command":"go test .","default_cpu":1,"default_ram":512}]
```
