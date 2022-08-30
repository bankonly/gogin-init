init:
	go get -u github.com/gin-gonic/gin
	go get go.mongodb.org/mongo-driver/mongo
	go get github.com/joho/godotenv
	go get golang.org/x/crypto/bcrypt
	go get -d github.com/bankonly/goutils

dev:
	gomon main.go

start:
	go run main.go

push:
	git add .
	git commit -m "$m"
	git push