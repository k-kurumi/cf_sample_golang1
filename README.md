# go log writer

require: go1.6 or later

for local
```
$ go get -u -v github.com/k-kurumi/cf_sample_golang1
$ glide install
$ PORT=3333 go run main.go
```

for cf
```
$ go get -u -v github.com/k-kurumi/cf_sample_golang1
$ glide install

## edit manifest.yml ##

$ cf push
```
