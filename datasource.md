Go 使用 database/sql 包连接到不同的数据库。根据需要，选择相应的驱动程序，并配置正确的 Data Source Name (DSN)

##### 以下为一些常见数据库的数据源设置以及连接数据库示例：

**说明**：DSN（Data Source Name）是在 Go 中使用 database/sql 包连接数据库时的一个重要配置项。DSN 是一个字符串，它包含了连接到特定数据库所需的必要信息，比如用户名、密码、主机地址、端口、数据库名称等

### MySQL
MySQL 的 DSN(Data Source Name) 格式如下：
```
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...]
```

- **username**: MySQL 用户名。
- **password**: MySQL 密码（可选）。
- **protocol**: 连接协议，如 `tcp`、`unix` 等。
- **address**: 主机地址和端口，如 `127.0.0.1:3306`。
- **dbname**: 要连接的数据库名称。
- **param1=value1**: 连接参数，如 `charset=utf8mb4`。

示例：
```go
dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

需要安装 MySQL 驱动程序：

```sh
go get -u github.com/go-sql-driver/mysql
```

在代码中连接 MySQL 数据库：

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

}
```


### PostgreSQL
PostgreSQL 的 DSN(Data Source Name) 格式如下：
```
user=user password=password dbname=mydb sslmode=disable host=localhost port=5432
```

- **user**: PostgreSQL 用户名。
- **password**: PostgreSQL 密码。
- **host**: 主机地址。
- **port**: 端口号。
- **dbname**: 要连接的数据库名称。
- **sslmode**: SSL 连接模式，默认为 `disable`。

示例：
```go
dsn := "user=user password=password dbname=mydb sslmode=disable host=localhost port=5432"
```
安装 PostgreSQL 驱动程序：

```sh
go get -u github.com/lib/pq
```

在代码中连接 PostgreSQL 数据库：

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func main() {
    dsn := "user=username password=password dbname=mydb sslmode=disable host=localhost port=5432"
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

}
```


### SQLite
SQLite 的 DSN(Data Source Name) 格式很简单，通常是文件路径。

示例：
```go
dsn := "./mydatabase.db"
```
安装 SQLite 驱动程序：

```sh
go get -u github.com/mattn/go-sqlite3
```

在代码中连接 SQLite 数据库：

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    dsn := "./mydatabase.db"
    db, err := sql.Open("sqlite3", dsn)
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

}
```

### SQL Server
SQL Server 的 DSN(Data Source Name) 格式如下：
```
sqlserver://user:password@host:port?database=dbname
```

- **user**: SQL Server 用户名。
- **password**: SQL Server 密码。
- **host**: 主机地址。
- **port**: 端口号。
- **dbname**: 要连接的数据库名称。

示例：
```go
dsn := "sqlserver://user:password@localhost:1433?database=testdb"
```

安装 SQL Server 驱动程序：

```sh
go get -u github.com/denisenkom/go-mssqldb
```

在代码中连接 SQL Server 数据库：

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/denisenkom/go-mssqldb"
)

func main() {
    dsn := "sqlserver://username:password@localhost:1433?database=dbname"
    db, err := sql.Open("sqlserver", dsn)
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()
}
```


### Oracle
Oracle 的 DSN(Data Source Name) 格式如下：
```
dsn := "user="username" password="password" connectString="host:port/dbname"
```

- **user**: Oracle 用户名。
- **password**: Oracle 密码。
- **host**: 主机地址。
- **port**: 端口号。
- **dbname**: 要连接的数据库名称。

示例：
```go
dsn := "user="root" password="123456" connectString="localhost:1521/dbname""
```

安装 Oracle 驱动程序

```sh
go get -u github.com/godror/godror
```

在代码中连接 Oracle 数据库

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    _ "github.com/godror/godror"
)

func main() {
    dsn := `user="root" password="123456" connectString="localhost:1521/dbname"`
    db, err := sql.Open("godror", dsn)
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()
	
}
```