# lang_bench

This a collection of code snippets that do the same thing: `lpush` a URL parameter into redis. To add a language, create a simple webapp that listens on port 3001. The app should then accept a single URL parameter, `test`, and then `lpush` its contents to the redis list titled `test_list`.

## Dependencies
`brew install go redis node maven`

## Starting the servers

### Ruby
```
bundle install
ruby app.rb
```

### Node
```
npm install
node app.js
```

### Go
First, create your go workspace:
```
mkdir ~/go
echo 'export GOPATH=~/go' >> ~/.zshrc
```
Install the redigo package:
`go get github.com/garyburd/redigo/redis`

Run the app:
`go run app.go`

### Java
```
cd java
mvn package jetty:run
```

## To benchmark
`./benchmark.sh`
