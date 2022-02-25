<img src="https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png" alt="back-to-the-future-v2" width=400>

<p align="Left">
  <p align="Left">
    <a href="https://github.com/borischen0203/litclock/actions/workflows/go.yml"><img alt="GitHub release" src="https://github.com/borischen0203/litclock/actions/workflows/go.yml/badge.svg?logo=github&style=flat-square"></a>
  </p>
</p>

# Literal clock
This `litclock` command-line tool mainly converts numeric time to human friendly text.

Numeric Time -> Human Friendly Text:
- 1:00 -> One o'clock
- 13:05 -> Five past one

Service link:
https://github.com/borischen0203/litclock-service

# Features
- `litclock` command: Be able to convert numeric time to human text.


# How to use

## On macOS via Homebrew:
Step1:
```bash
brew tap borischen0203/litclock
```
Step2:
```bash
brew install litclock
```

## Run in Docker:
Required
- Install docker

### Run process
Step1: Pull docker image(borischen0203/litclock)
```bash
docker pull borischen0203/litclock
```
Step2:  Run docker image as below command
```bash
docker run -it --rm borischen0203/litclock
```

### Docker run demo
```bash
# Display the current time in the human text without input parameter
$ docker run -it --rm borischen0203/litclock
$ Seven past two

# Display the the human text with input numeric time
$ docker run -it --rm borischen0203/litclock 15:40
$ Twenty to four
```

## Run in Local:

Required
- Install go(version >= 1.6)
- Install `make` cli(https://formulae.brew.sh/formula/make)
```bash
brew install make
```

### Run process
Step1: Clone the repo
```bash
git clone https://github.com/borischen0203/litclock.git
```
Step2: Use `make` to execute makefile run test and build
```bash
make all
```
Step3: Execute build file with or without parameter
```bash
./litclock
```
```bash
./litclock 15:40
```
### Local run demo
```bash
# Display the current time in the human text without input parameter
$ ./litclock
$ Seven past two

# Display the human text with input numeric time
$ ./litclock 15:40
$ Twenty to four
```

## Tech stack
- Golang
- Cobra
- Docker
- make
- Github actions
- shell

## Todo:
- [X] Release cli on Homebrew