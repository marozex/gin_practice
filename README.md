## 参照
udemy Gin入門

## 環境構築手順
```
go mod init gin_practice // go.mod作成
go get -u github.com/gin-gonic/gin
vi main.go // 公式のサンプルコードを貼り付け
go run main.go
curl localhost:8080/ping // pongが表示さればok
```
### air導入
```
go install github.com/air-verse/air@latest
vi ~/.zshrc //export PATH="go env GOPATHで確認したパス/bin:$PATH"を追記
source ~/.zshrc

air init //.air.tomlが作成されればok
air // ローカルでサーバー起動し、ホットリロード有効となる
```