module gostudy/gin

go 1.13

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191029031824-8986dd9e96cf
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191028164358-195ce5e7f934
)

require (
	github.com/gin-contrib/sessions v0.0.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.11
	github.com/kr/pretty v0.1.0 // indirect
)
