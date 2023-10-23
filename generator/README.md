## Generator

![diamond-square algorithm](./docs/example_5.jpeg "results of the diamond-square algorithm")

## Use

```shell
docker build -t diamond-square:latest .
docker run --name generator diamond-square -size 1025 -path /project/map.jpeg
docker cp generator:/project/map.jpeg .
```
