##  Gdao Test Demo [[中文文档](https://github.com/donnie4w/gdaodemo/blob/master/README_zh.md)]

###### This is a test demo program for Gdao, including a packaged SQLite test database file `hstest.db`, with generated data. The demo program can be run directly. Except for testing read-write separation or multiple data source operations, other tests default to operating on the `hstest.db` database data and can be run directly to see the data operation results.

##### The demo program tests the following aspects:

1. Data operations using Gdao's mapping files
2. Transactions, stored procedures, batch processing, and serialization operations in Gdao
3. CRUD functions using Gdao interfaces
4. Using the GdaoCache caching interface
5. Binding and removing read-write separation with GdaoSlave
6. SQL file mapping and interface calls with GdaoMapper

### For detailed usage instructions, please refer to the documentation: https://tlnet.top/gdaoendoc

## Overview of the Demo

### Generating Database Table Mapping Files

###### Download the code construction tool: https://tlnet.top/download

##### Example for Windows environment

1. Generate the configuration file: `gdao.json`

```bash
// Generate configuration file
win110_gdao.exe init
```

2. Modify the database connection in `gdao.json`, using MySQL as an example

```json
{
  "dbtype": "mysql",
  "dbhost": "localhost",
  "dbport": 3306,
  "dbname": "hstest",
  "dbuser": "root",
  "dbpwd": "123456",
  "package": "dao"
}
```

3. Execute the data file generation command

```bash
win110_gdao.exe -c gdao.json
```

##### The result will be files corresponding to the database tables:

```text
dao/Hstest.go
dao/Hstest1.go
dao/Hstest2.go
```

4. Configure the data source. The SQLite data file `hstest.db` is already packaged in the `gdaodemo` directory. You can directly call the `Open` function to read the file. For ease of multiple data source operations, we use a unified configuration mode to read in.

##### `sqlite.json`

```json
{
  "dbtype": "sqlite",
  "dbhost": "",
  "dbport": 0,
  "dbname": "hstest.db",
  "dbuser": "",
  "dbpwd": ""
}
```

5. Setting the data source in Gdao

##### In the initialization function `init`, set the Gdao data source

```go
func init() {
	if db, err := getDataSource("sqlite.json"); err == nil {
		gdao.Init(db, gdao.SQLITE)
		gdao.SetLogger(true)  // Enable logging in test environment
		logger.Info("datasource init")
	}
}
```

6. Basic operations with mapping files

```go
// Query
func TestSelect(t *testing.T) {
    hs := dao.NewHstest()
    hs.Where(hs.Id.EQ(10))
    h, _ := hs.Select(hs.Id, hs.Value, hs.Rowname)
    logger.Debug(h)
}

// Update
func TestUpdate(t *testing.T) {
    hs := dao.NewHstest()
    hs.SetRowname("hello10")
    hs.Where(hs.Id.EQ(10))
    hs.Update()
}   

// Insert
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

// Batch processing
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

// Serialization
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

// Read-write separation
func TestSlave(t *testing.T) {
    // Configure and get another data source as a replica data source, here using MySQL
    mysql, _ := getDataSource("mysql.json")
    // Bind the replica to Hstest1, Hstest1 read operations will point to MySQL
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

// Data caching
func TestCache(t *testing.T) {
    // Bind Hstest to use data caching, cache duration is 300 seconds
    gdaoCache.BindClass[dao.Hstest]()  
    hs := dao.NewHstest()
    hs.Where((hs.Id.Between(0, 2)).Or(hs.Id.Between(10, 15)))
    hs.Limit(1)
    if hs, err := hs.Select(); err == nil {
        logger.Debug(hs)
    }
    logger.Debug("----------------------Set Cache----------------------")
    logger.Debug()
    // If the cache has not expired, the same condition SQL returns cached data directly
    hs = dao.NewHstest()
    hs.Where((hs.Id.Between(0, 2)).Or(hs.Id.Between(10, 15)))
    hs.Limit(1)
    if hs, err := hs.Select(); err == nil {
        logger.Debug(hs)
    }
    logger.Debug("----------------------Get Cache----------------------")
    logger.Debug()
    // Remove the cache binding for Hstest, subsequent read operations will not use the cache
    gdaoCache.RemoveClass[dao.Hstest]()
    hs = dao.NewHstest()
    hs.Where((hs.Id.Between(0, 2)).Or(hs.Id.Between(10, 15)))
    hs.Limit(1)
    if hs, err := hs.Select(); err == nil {
        logger.Debug(hs)
    }
    logger.Debug("----------------------No Use Cache----------------------")
}

// Transactions
func Test_transaction(t *testing.T) {
    // Get transaction object
    tx, _ := gdao.NewTransaction()
    hs := dao.NewHstest2()
    // Use transaction
    hs.UseTransaction(tx)
    hs.SetAge(100)
    hs.SetName("www")
    hs.Where(hs.Id.EQ(1))
    hs.Update()
    
    // The transaction object supports executing SQL CRUD operations
    tx.ExecuteUpdate("update hstest set age=? where id=?", 100, 1)
    tx.Rollback()
    
    // Check if rollback succeeded
    fmt.Println(gdao.ExecuteQueryBean("select * from hstest2 where id=?", 1))
    fmt.Println(gdao.ExecuteQueryBean("select * from hstest where id=?", 1))
}
```

7. Basic CRUD interfaces in Gdao

```go
// Query
func TestSelect(t *testing.T) {
    bean, _ := gdao.ExecuteQueryBean("select id, value, rowname from hstest where id=?", 10)
    logger.Debug(bean)
    
    hstest, _ := gdao.ExecuteQuery[dao.Hstest]("select id, value, rowname from hstest where id=?", 10)
    logger.Debug(hstest)
}

// Update
func Test_update(t *testing.T) {
    gdao.ExecuteUpdate("update hstest set age=? where id=?", 100, 1)
}

// Delete
func Test_delete(t *testing.T) {
    gdao.ExecuteUpdate("delete from hstest2 where id=?", 1)
}
```