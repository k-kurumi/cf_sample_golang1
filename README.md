# go log writer

go1.6 or later

## Usage

1. install glide
https://github.com/Masterminds/glide/releases

2. clone and install libs

        $ cd $GOPATH/src
        $ git clone https://github.com/k-kurumi/cf_sample_golang1
        $ cd cf_sample_golang1

        $ glide install

3. run or cf push

        $ PORT=8888 go run main.go

        ## or

        $ cf push -f manifest.yml
