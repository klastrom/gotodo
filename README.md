# todo app

## 概要

このアプリはシンプルなTODOリストアプリです。Goで作成され、Dockerコンテナとして実行されます。

## 使用方法

### Dockerイメージのビルド

以下のコマンドを使用してDockerイメージをビルドします：
sh
docker build -t gotodo:latest

### Dockerコンテナの実行

以下のコマンドを使用してDockerコンテナを実行します：
sh
docker run -p 8080:8080 gotodo:latest


### TODOアイテムの操作

#### TODOリストを取得する

以下のコマンドを使用してTODOリストを取得します：
sh
curl http://localhost:8080/todos


#### TODOアイテムを追加する

以下のコマンドを使用してTODOアイテムを追加します：
sh
curl -X POST http://localhost:8080/todos/add -H "Content-Type: application/json" -d '{"task": "Buy milk"}'


#### TODOアイテムを削除する

以下のコマンドを使用してTODOアイテムを削除します：
sh
curl -X POST http://localhost:8080/todos/delete -H "Content-Type: application/json" -d '{"id": 1}'


## エンドポイント

- `GET /todos` - TODOリストを取得します。
- `POST /todos/add` - 新しいTODOアイテムを追加します。リクエストボディにJSON形式でタスクを指定します。
- `POST /todos/delete` - TODOアイテムを削除します。リクエストボディにJSON形式でIDを指定します。

