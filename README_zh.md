##  Gdao Test Demo [[English](https://github.com/donnie4w/gdaodemo/blob/master/README.md)]

###### 这是Gdao的测试demo程序，已经打包sqlite测试数据库hstest.db文件,并生成元素数据。demo程序可以直接运行，除了需要测试读写分离或多数据源操作，其他Test默认操作hstest.db数据库数据，可以直接运行，查看数据操作结果。

##### demo程序测试以下几个方面

1. gdao的映射文件数据操作
2. gdao的事务，存储过程，批处理，序列化等操作
3. gdao接口的CRUD函数使用
4. gdaoCache 缓存接口使用
5. gdaoSlave数据读写分离绑定与移除操作
6. gdaoMapper SQL文件映射与接口调用操作

### Gdao的详细使用说明请查看使用文档：https://tlnet.top/gdaodoc

## 以下是dome概述

### 生成数据库表映射文件

######  代码构建工具下载：https://tlnet.top/download

##### 以window环境为例

1.  生成配置文件: gdao.json

```bash
//生成配置文件
win110_gdao.exe init
```
2.  修改给gdao.json的数据库连接,以mysql为例

```text
  "dbtype": "mysql",
  "dbhost": "localhost",
  "dbport": 3306,
  "dbname": "hstest",
  "dbuser": "root",
  "dbpwd": "123456",
  "package": "dao",
```
3.  执行数据文件生成命令
```bash
win110_gdao.exe -c gdao.json
```
##### 执行结果，生成数据库表的对应文件

```text
dao/Hstest.go
dao/Hstest1.go
dao/Hstest2.go
```

4.  配置数据源，sqlite数据文件hstest.db已经打包在gdaodome目录中，直接调用Open函数读入文件即可。为了多数据源操作方便，这里采用配置统一模式读入

##### sqlite.json

```json
{
  "dbtype": "sqlite",
  "dbhost": "",
  "dbport":0,
  "dbname": "hstest.db",
  "dbuser": "",
  "dbpwd": ""
}
```

5.  Gdao设置数据源

##### 初始化函数init中，设置gdao数据源

```go
func init() {
	if db, err := getDataSource("sqlite.json"); err == nil {
		gdao.Init(db, gdao.SQLITE)
		gdao.SetLogger(true)  // 测试环境中，打开日志打印
		logger.Info("datasource init")
	}
}
```

6.  映射文件的基础操作

```go
//查询
func TestSelect(t *testing.T) {
    hs := dao.NewHstest()
    hs.Where(hs.Id.EQ(10))
    h, _ := hs.Select(hs.Id, hs.Value, hs.Rowname)
    logger.Debug(h)
}

//更新
func TestUpdate(t *testing.T) {
    hs := dao.NewHstest()
    hs.SetRowname("hello10")
    hs.Where(hs.Id.EQ(10))
    hs.Update()
}   

//新增
func TestInsert(t *testing.T) {
    hs := dao.NewHstest()
    hs.SetValue("hello123")
    hs.SetLevel(12345)
    hs.SetBody([]byte("hello"))
    hs.SetRowname("hello1234")
    hs.SetUpdatetime(time.Now())
    hs.SetFloa(123456)
    hs.SetAge(123)
    hs.Insert()
}

//批处理
func TestBatch(t *testing.T) {
    hs := dao.NewHstest2()
    hs.SetAge(100)
    hs.SetName("www")
    hs.SetCreatetime(time.Now())
    hs.SetFloa(1.1)
    hs.AddBatch()
    
    hs.SetAge(1000)
    hs.SetName("wwww")
    hs.SetCreatetime(time.Now())
    hs.SetFloa(1.11)
    hs.AddBatch()
    hs.ExecBatch()
}  

//序列化
func Test_serialize(t *testing.T) {
    hs := dao.NewHstest2()
    hs.Limit(1)
    hs1, _ := hs.Select()
    bs, _ := hs1.Encode()
    logger.Debug("encode len(bs):", len(bs))
    logger.Debug(hs1)
    logger.Debug("----------Encode-----------")
    hs2 := dao.NewHstest2()
    hs2.Decode(bs)
    logger.Debug(hs2)
    logger.Debug("----------Decode-----------")
}

//读写分离
func TestSlave(t *testing.T) {
    //配置并获取另一个数据源作为备库数据源，这里使用mysql数据库
    mysql, _ := getDataSource("mysql.json")
    //绑定备库与Hstest1，Hstest1读取操作将指向mysql
    gdaoSlave.BindClass[dao.Hstest1](mysql, gdao.MYSQL)
    hs := dao.NewHstest1()
    hs.Where(hs.Id.Between(0, 5))
    hs.OrderBy(hs.Id.Desc())
    hs.Limit(3)
    if hslist, err := hs.Selects(); err == nil {
        for _, hs := range hslist {
            logger.Debug(hs)    
        }
    }
}

//数据缓存
func TestCache(t *testing.T) {
    //绑定Hstest 使用 数据缓存，缓存时效为300秒
    gdaoCache.BindClass[dao.Hstest]()  
    hs := dao.NewHstest()
    hs.Where((hs.Id.Between(0, 2)).Or(hs.Id.Between(10, 15)))
    hs.Limit(1)
    if hs, err := hs.Select(); err == nil {
        logger.Debug(hs)
    }
    logger.Debug("----------------------Set Cache----------------------")
    logger.Debug()
    //时效未过期时，相同条件SQL直接返回缓存数据
    hs = dao.NewHstest()
    hs.Where((hs.Id.Between(0, 2)).Or(hs.Id.Between(10, 15)))
    hs.Limit(1)
    if hs, err := hs.Select(); err == nil {
        logger.Debug(hs)
    }
    logger.Debug("----------------------Get Cache----------------------")
    logger.Debug()
    //移除Hstest的缓存绑定，再次读取操作不再读取缓存库
    gdaoCache.RemoveClass[dao.Hstest]()
    hs = dao.NewHstest()
    hs.Where((hs.Id.Between(0, 2)).Or(hs.Id.Between(10, 15)))
    hs.Limit(1)
    if hs, err := hs.Select(); err == nil {
        logger.Debug(hs)
    }
    logger.Debug("----------------------No Use Cache----------------------")
}

//事务
func Test_transaction(t *testing.T) {
    //获取事务对象
    tx, _ := gdao.NewTransaction()
    hs := dao.NewHstest2()
    //使用事务
    hs.UseTransaction(tx)
    hs.SetAge(100)
    hs.SetName("www")
    hs.Where(hs.Id.EQ(1))
    hs.Update()
    
    //事务对象支持执行SQL的CURD操作
    tx.ExecuteUpdate("update hstest set age=? where id=?", 100, 1)
    tx.Rollback()
    
    //检查是否回滚成功
    fmt.Println(gdao.ExecuteQueryBean("select * from hstest2 where id=?", 1))
    fmt.Println(gdao.ExecuteQueryBean("select * from hstest where id=?", 1))
}
```

7.  gdao基础CURD接口

```go
//查询
func TestSelect(t *testing.T) {
    bean, _ := gdao.ExecuteQueryBean("select id,value,rowname from hstest where id=?", 10)
    logger.Debug(bean)

    hstest, _ := gdao.ExecuteQuery[dao.Hstest]("select id,value,rowname from hstest where id=?", 10)
    logger.Debug(hstest)
}

//更新
func Test_update(t *testing.T) {
    gdao.ExecuteUpdate("update hstest set age=? where id=?", 100, 1)
}

//删除
func Test_delete(t *testing.T) {
    gdao.ExecuteUpdate("delete from hstest2  where id=?", 1)
}
```