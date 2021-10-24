<img src="https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png" alt="back-to-the-future-v2" width=300>

<p align="Left">
  <p align="Left">
    <a href="https://github.com/borischen0203/litclock/actions/workflows/go.yml"><img alt="GitHub release" src="https://github.com/borischen0203/litclock/actions/workflows/go.yml/badge.svg?logo=github&style=flat-square"></a>
  </p>
</p>

# Literal clock
This `litclock` command-line tool mainly converts numeric time to human friendly text. Fox example, 5:05 -> Five past five.

# Features
- `litclock` command: Be able to convert numeric time to human text.


# How to use

## Docker:
```bash
# Step1: docker pull
docker pull borischen0203/litclock

# Step2: docker run
docker run -it --rm borischen0203/litclock

# Or docker run + numeric time
docker run -it --rm borischen0203/litclock 15:40
```
### Docker run demo
```bash
# Display the current time in the human text without input parameter
> docker run -it --rm borischen0203/litclock
> Seven past two

# Display the current time in the human text with input parameter
> docker run -it --rm borischen0203/litclock 15:40
> Twenty to four
```

## Local:

### pre:
- Install go(version >= 1.6)
- Install make cli(https://formulae.brew.sh/formula/make)

```bash
# clone a repo
git clone https://github.com/borischen0203/litclock.git

# Use `make` to execute makefile run test and build
make all

# Execute with or without parameter
./litclock
./litclock 15:40
```
### Local run demo
```bash
# Display the current time in the human text without input parameter
> ./litclock
> Seven past two

# Display the current time in the human text with input parameter
> ./litclock 15:40
> Twenty to four
```

## Tech stack
- Golang
- Cobra
- Docker
- make
- Github actions
- shell

### Todo:
- [ ] Release cli on Homebrew
```bash
brew tap borischen0203/litclock
brew install litclock
```