# AMIDA KUJI (WINDOWS DLL)

[![MIT License](https://img.shields.io/badge/Licence-MIT-blue.svg)](./LICENSE) [![Build Status](https://travis-ci.org/NaniteFactory/amidakuji.svg?branch=master)](https://travis-ci.org/NaniteFactory/amidakuji) [![Go Report Card](https://goreportcard.com/badge/github.com/nanitefactory/amidakuji)](https://goreportcard.com/report/github.com/nanitefactory/amidakuji) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/20cc2ee6013d4f5fae11817fa5fb1a6b)](https://www.codacy.com/app/NaniteFactory/amidakuji?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=NaniteFactory/amidakuji&amp;utm_campaign=Badge_Grade) [![Coverage Status](https://coveralls.io/repos/github/NaniteFactory/amidakuji/badge.svg?branch=master)](https://coveralls.io/github/NaniteFactory/amidakuji?branch=master)

> Ghost Leg (Chinese: 畫鬼腳), known in Japan as Amidakuji (阿弥陀籤, "Amida lottery", so named because the paper was folded into a fan shape resembling Amida's halo) or in Korea as Sadaritagi (사다리타기, literally "ladder climbing"), is a method of lottery designed to create random pairings between two sets of any number of things, as long as the number of elements in each set is the same. This is often used to distribute things among people, where the number of things distributed is the same as the number of people. For instance, chores or prizes could be assigned fairly and randomly this way.

(Explanation from [Wikipedia](https://en.wikipedia.org/wiki/Ghost_Leg))

- - -

## Examples

| [GIF1](examples/user_conf_sample6.json) | [GIF2](examples/user_conf_sample3.json) |
|-----------------------------------------|-----------------------------------------|
| ![1](examples/1.gif)                    | ![2](examples/2.gif)                    |

| GIF3                 |
|----------------------|
| ![3](examples/3.gif) |

- - -

## Control

|       | Action                  | Input       |
|-------|-------------------------|-------------|
| 🔀 	| Shuffle                 | 1           |
| ▶️ 	| Find path               | 2           |
| ⏩ 	| Find path (faster)      | 3           |
| ⏸ 	| Pause                   | Space       |
| ⬆️➡️⬇️⬅️ 	| Move camera around      | Arrow keys  |
| ↩️ 	| Rotate camera           | Enter       |
| 🔭 	| Zoom in and out         | Scroll      |
| 🎇 	| Firework                | Left click  |
| 🔬 	| Inspection              | Right click |
| 🔁 	| Toggle full screen mode | Tab         |

- - -

## Build & Run

You may want to use MSYS or MSYS2 to run bash commands on Windows systems.

1. Have Go installed properly

```bash
$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
$ echo "PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> ~/.bashrc
$ echo "PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> ~/.bash_profile
```

2. go get

```bash
$ go get -v github.com/nanitefactory/amidakuji/...
$ go get -v -u github.com/go-bindata/go-bindata/...
```

3. make

```bash
$ cd $GOPATH/src/github.com/nanitefactory/amidakuji/
$ make
```


That's it!

- - -

#### Build only command

```bash
$ make build
```

- - -

#### Run only command

```bash
$ rundll32.exe amidakuji.dll Test
```

- - -

## Clean-up

```bash
$ make clean
```

- - -

## External sources

#### Library
- [Pixel](https://github.com/faiface/pixel/tree/7cff3ce3aed80129b7b1dd57e63439426e11b6ee)
- [Beep](https://github.com/faiface/beep/tree/63cc6fbbac46dba1a03e55f0ebc965d6c82ca8e1)
- [GLFW 3.2](https://github.com/go-gl/glfw/tree/513e4f2bf85c31fba0fc4907abd7895242ccbe50/v3.2/glfw)
- [dialog](https://github.com/sqweek/dialog/tree/2f9d9e5dd848a3bad4bdd0210c73bb90c13a3791)

#### Music
- [Night Tempo - Pure Present](https://nighttempo.bandcamp.com/album/pure-present) - [08 Kikuchi Momoko - Night Cruising (Night Tempo 100% Pure Remastered)](https://nighttempo.bandcamp.com/track/kikuchi-momoko-night-cruising-night-tempo-100-pure-remastered-2)
- [Night Tempo - Pure Present](https://nighttempo.bandcamp.com/album/pure-present) - [19 Takeuchi Mariya - Plastic Love (Night Tempo 100% Pure Remastered)](https://nighttempo.bandcamp.com/track/takeuchi-mariya-plastic-love-night-tempo-100-pure-remastered-3)

#### Image
- [Gophers...](https://github.com/egonelbre/gophers/tree/dfb1bc3e6092179bd80d2e4156a8d32dba484cc9)

#### Font
- [나눔바른고딕 (NanumBarunGothic.ttf)](https://hangeul.naver.com/2017/nanum)

