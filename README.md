# Snake Game

Go言語で作成されたターミナルベースのスネークゲームです。

## 概要

このプロジェクトは、Go言語とtcellライブラリを使用して作成されたクラシックなスネークゲームです。ターミナル上で動作し、キーボード入力でスネークを操作してエサを食べながら成長させることができます。

## 必要条件

- Go 1.20以上
- ターミナル環境

## インストール

1. リポジトリをクローンします：

```bash
git clone https://github.com/ryota1119/snake_game.git
cd snake_game
```

1. 依存関係をインストールします：

```bash
go mod download
```

## 実行方法

プロジェクトのルートディレクトリで以下のコマンドを実行します：

```bash
go run cmd/snake_game/main.go
```

または、バイナリをビルドして実行：

```bash
go build -o snake_game cmd/snake_game/main.go
./snake
```

## 操作方法

- **矢印キー**: スネークの移動方向を変更
- **q**: ゲーム終了
- **ESC**: ゲーム終了

## ゲームルール

- スネークを操作してエサ（`*`）を食べます
- エサを食べるたびにスネークが成長し、スコアが増加します
- 壁や自分の体に衝突するとゲームオーバーです
- できるだけ多くのエサを食べて高スコアを目指しましょう

## プロジェクト構造

```plaintext
snake_game/
├── cmd/
│   └── snake/
│       └── main.go          # エントリーポイント
├── internal/
│   └── snake/
│       ├── food.go          # エサの管理
│       ├── game.go          # ゲームロジック
│       ├── input.go         # 入力処理
│       ├── render.go        # 描画処理
│       ├── snake.go         # スネークの管理
│       └── types.go         # 型定義
├── go.mod                   # Goモジュール定義
├── go.sum                   # 依存関係のチェックサム
└── README.md               # このファイル
```

## 技術スタック

- **言語**: Go
- **ターミナルUI**: tcell
- **アーキテクチャ**: モジュラー設計

## 開発者

このプロジェクトは学習目的で作成されました。
# snake_game
