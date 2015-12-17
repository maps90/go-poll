# About go-poll
go-poll is an attempt to make polling api micro-service using following technologies :
* [Go Lang].
* [Redis].

## HOW TO: INSTALL
* Dowload and [install GO].
* on Linux setup GOPATH to your .bashrc/.zshrc
```
echo 'export GOPATH=$PATH:$HOME/{your-working-directory}' >> ~/.bashrc
```
* on windows or osx, [checkout golang documentation].

## HOW TO: INSTALL
```
$> go install
$> go build
```

## HOW TO: Run
```
$> ./go-poll
```

## API list
* GET candidates
* GET candidate/:id
* POST votes
* GET votes

## Check API
```
$> curl -i "{host}:{@port}/api/v1/candidates"
$> curl -i "{host}:{@port}/api/v1/candidate/1"
```

[Redis]:http://redis.io/
[Go]:https://golang.org/
[checkout golang documentation]:https://golang.org/doc/install#osx
