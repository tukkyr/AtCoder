# AtCoder のサンプル取得

## oj ツールを使う

[oj](https://github.com/online-judge-tools/oj)

## go のテストケースとして実行する

### download

```sh
oj d -f '%i.%e' -d cases <target url>
```

### test

```sh
oj t -c 'go run .' -d cases
```

### commit

#### windows

```sh
cat main.go | clip.exe
```

#### mac

```sh
cat main.go | pdcopy
```
