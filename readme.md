### Basic understanding on russian:
`https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos-ru`

### Install and set proper path variables:
```https://golang.org/doc/install```


### Setup godog and run it:
``GO111MODULE=on go get github.com/cucumber/godog/cmd/godog@v0.9.0``

##### Trouble shooting godog
`https://github.com/cucumber/godog/issues/279`
Disregard. Fixed by setting export GO111MODULE=on, running go mod init, and then installing a specific version of godog with a command like go get github.com/cucumber/godog/cmd/godog@v0.8.1 as described in the notice in #253


Why GODOG - because it is friendly for Cucumber users.
Also checked:
https://github.com/smartystreets/goconvey <- Nice reporting vs repo haven't been updated for 9 months
https://github.com/onsi/ginkgo  <-- Most mature framework, more like specs descriptions
https://github.com/franela/goblin <-- Mocha style, mostly familiar for JS users

#### Adding binding for godod and go test

So how to do it is described in godog readme.

For us we create basic_test.go file which is responsible for collection of feature files
to run test from root use :
``go test -v --godog.format=pretty``