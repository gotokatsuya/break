# break

## Getting project

```
$ cd $GOPATH/src/github.com/gotokatsuya
$ git clone https://github.com/gotokatsuya/break.git
$ cd break
```

## Database

Create `dev_break.db` and `test_break.db`.
```
$ mysql -u root
CREATE DATABASE `dev_break` DEFAULT CHARSET utf8mb4;
CREATE DATABASE `test_break` DEFAULT CHARSET utf8mb4;
```

## Install & Run

```
$ brew install glide
$ brew install direnv
$ cd $GOPATH/src/github.com/gotokatsuya/break
$ direnv allow
$ make install
$ make run
```
