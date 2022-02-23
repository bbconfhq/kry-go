# kry-go
KRY written in Go.

## Description
This project provides simple platform for Problem-Solving contest.

Originally, it was written in Kotlin, but moved into Golang.

Its name is combination of main contributor's first letter. (**K**i-dong, **R**yo, **Y**un)

Compiling submitted codes will work in under [sandbox](https://github.com/blurfx/sandbox).

## Getting started
### Prerequisites
* Go (>= 1.17)
* docker, docker-compose

### Execute
1. Clone project
```shell
$ git clone https://github.com/gwanryo/kry-go.git
```
2. Install dependencies
```shell
$ go get && go get -u
```
3. Edit .env.sample and save to .env
4. (Only non-docker) Configure your DBMS properly, then run `main.go`
```shell
$ go run main.go
```
5. (Only docker) Run `docker-compose`
```shell
$ docker-compose up --build
```
6. (Optional) Use `swag` if you want regenerate swagger document
```shell
$ swag init --parseDependency --parseInternal
```

## Help
Any advice will welcome, please report on [Issues](https://github.com/gwanryo/kry-go/issues)

## License
This project is licensed under the MIT License - see the LICENSE for more details

## Acknowledgments
Thanks for any inspiration, code reviews, opinion, etc.
* [blurfx](https://github.com/blurfx)
* [murry](https://github.com/murry2018)
* [YunGoon](https://github.com/dsyun96)
