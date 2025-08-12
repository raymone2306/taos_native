```
cd tests/examples/JDBC/taosdemo
mvn clean package -Dmaven.test.skip=true
# Create tables first, then insert data
java -jar target/taosdemo-2.0.1-jar-with-dependencies.jar -host <hostname> -database <db name> -doCreateTable true -superTableSQL "create table weather(ts timestamp, f1 int) tags(t1 nchar(4))" -numOfTables 1000 -numOfRowsPerTable 100000000 -numOfThreadsForInsert 10 -numOfTablesPerSQL 10 -numOfValuesPerSQL 100
# Insert data directly without creating tables
java -jar target/taosdemo-2.0.1-jar-with-dependencies.jar -host <hostname> -database <db name> -doCreateTable false -superTableSQL "create table weather(ts timestamp, f1 int) tags(t1 nchar(4))" -numOfTables 1000 -numOfRowsPerTable 100000000 -numOfThreadsForInsert 10 -numOfTablesPerSQL 10 -numOfValuesPerSQL 100
```

If you encounter the error Exception in thread "main" `java.lang.UnsatisfiedLinkError: no taos in java.library.path`, please check whether the TDengine client package is installed or TDengine is compiled and installed. If you are sure it is installed and still encounter this error, you can add `-Djava.library.path=/usr/lib` after the `java` command to specify the path to the shared library.
