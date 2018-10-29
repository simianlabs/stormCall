# stormCall

Simple lambda function in Go to call StackStorm webhook and sent info from AwS IoT buttom.

### Setting up Go Lambda

https://docs.aws.amazon.com/lambda/latest/dg/getting-started.html


### Build

1. Get aws-lambda-go libraries: https://github.com/aws/aws-lambda-go
```
$ go get -v github.com/aws/aws-lambda-go
```
2. Modify placeholders in main.go
3. Build for linux
```
$ GOOS=linux GOARCH=amd64 go build -o main main.go
```
4. Pack binaries
```
$ zip main.zip main
```

### Disable SSL check

Add this line to main function if you need to skip ssl verification

```go
// Disable ssl verification if you dont have valid certificate
http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
```


### Setting IoT deployment.

Full tutorial available: https://simianlabs.io/2018/10/17/infrastructure-on-click/

### Maintainer
Stefan Monko || smonko@simianlabs.io  
`Simian Labs` - (https://github.com/simianlabs)  
http://simianlabs.io || sl@simianlabs.io