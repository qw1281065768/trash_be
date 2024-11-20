
pkill -f 'gorest'

go build -x -o bin/gorest main.go

sh ./bin/start.sh