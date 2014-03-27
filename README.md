# lang_bench

## Dependencies
`brew install go redis node`

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

## To benchmark
`./benchmark.sh`
