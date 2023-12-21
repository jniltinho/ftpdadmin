.PHONY: clean build


# https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj
# https://github.com/go-gorm/gen/tree/master/tools/gentool
# https://gorm.io/gen/index.html
# https://gorm.io/gen/gen_tool.html
# https://github.com/xxjwxc/gormt


DSN = proftpd:proftpd@tcp(localhost:3306)/proftpd?charset=utf8mb4&parseTime=True&loc=Local


all: install-gentool run-gentool

install-gentool:
	go install gorm.io/gen/tools/gentool@latest

run-gentool:
	gentool -dsn "$(DSN)" -tables "users,groups,quotatallies"
	#gentool -c "./gen.tool"


run-gormt:
	curl -L# https://github.com/xxjwxc/gormt/releases/download/v2.1.gorm/gormt_linux.zip -o gormt.tar.gz
	tar -xzf gormt.tar.gz && rm gormt.tar.gz
	chmod +x gormt
	./gormt -H=127.0.0.1 -d=proftpd -p=proftpd -u=proftpd --port=3306 -F=true