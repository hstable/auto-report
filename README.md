每日自动上报

## Usage

```shell
go run main.go -u your-username -p your-password -e your-email
```

## Docker

构建镜像：

```shell
docker build -t auto-report .
```

运行：

```shell
 docker run -e USERNAME=your-username \
 -e PASSWORD=your-password \
 -e EMAIL=your-email \
 auto-report
```

## License

auto-report is licensed by an MIT license as can be found in the LICENSE file.

