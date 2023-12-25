all: gotool
	@go build -v .   # make: 生成二进制文件
clean:  # 做一些清理：删除二进制文件，删除vim swp文件
	rm -f apiserver
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:  # 格式化代码和源码静态检查
	gofmt -w .
	go tool vet . |& grep -v vendor;true
ca:  # 生成证书
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"

help: # 打印帮助文档
	@echo "make - compile the source code"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"

.PHONY: clean gotool ca help

# 上面的Makefile文件中，PHONY
#是个伪目标，形式上是一个目标，但是不需要依赖，伪目标一般只是为了执
#行目标下面的命令（比如clean就是伪目标）。@放在行首，表示不打印此行。默认在编译的过程中，会把此
#行的展开效果字符串打印出来。