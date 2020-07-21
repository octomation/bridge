> # 🌉 bridge
>
> A bridge between the protobuf and the brief formats.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]

## 💡 Idea

```bash
$ bridge service.proto > service.brief
```

A full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/bridge
```

### Binary

```bash
$ curl -sSfL https://raw.githubusercontent.com/kamilsk/bridge/master/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/kamilsk/bridge/master/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/kamilsk/bridge@latest
# or use egg tool
$ egg tools add github.com/kamilsk/bridge@latest
```

> [egg][] is an `extended go get`.

### Bash and Zsh completions

```bash
$ bridge completion bash > /path/to/bash_completion.d/bridge.sh
$ bridge completion zsh  > /path/to/zsh-completions/_bridge.zsh
# or autodetect
$ source <(bridge completion)
```

> See `kubectl` [documentation](https://kubernetes.io/docs/tasks/tools/install-kubectl/#enabling-shell-autocompletion).

## 🤲 Outcomes

### 👨‍🔬 Research

- [github.com/3rf/codecoroner](https://github.com/3rf/codecoroner)
- [github.com/tsenart/deadcode](https://github.com/tsenart/deadcode)
- [gitlab.com/opennota/check](https://gitlab.com/opennota/check)
- [honnef.co/go/tools](https://github.com/dominikh/go-tools)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/kamilsk/bridge
[build.icon]:       https://travis-ci.com/kamilsk/bridge.svg?branch=master
[coverage.page]:    https://codeclimate.com/github/kamilsk/bridge/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/c275f8f35b5b164b18cd/test_coverage
[design.page]:      https://www.notion.so/octolab/bridge-c4220e653415488c847823c88a31dc18?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/kamilsk/bridge
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
