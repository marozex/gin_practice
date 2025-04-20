## 環境構築手順
go mod init gin_practice // go.mod作成
go get -u github.com/gin-gonic/gin
vi main.go // 公式のサンプルコードを貼り付け
go run main.go
curl localhost:8080/ping