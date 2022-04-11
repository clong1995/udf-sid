# 什么是UDF？
UDF是mysql的一个拓展接口，UDF（Userdefined function）可翻译为用户自定义函数，  
这个是用来拓展Mysql的技术手段。  
# UDF可以做什么？
MySQL的内置函数虽然丰富，但毕竟不能满足所有人的需要，  
有时候我们需要对表中的数据进行一些处理而内置函数不能满足需要的时候，  
就需要对MySQL进行一些扩展。  
比如生成分布式递增ID，编解码，数据加解密，数据转换，发起http请求等。
# UDF优势和劣势。
**优势**  
1. 数据库和业务程序解耦。  
2. sql语句可以直接调用UDF函数完成数据处理，减少额外和业务程序的传输，提高性能。  
3. UDF本身的兼容性很好，并且比存储过程和存储方法具有更高的执行效率，同时支持聚集函数。  

**劣势**  
1. UDF需要单独开发，一般由C语言开发，需要注意内存泄漏。
2. 数据库升级需要重新注册UDF函数（数据库一般很少升级）。

# SnowFlake，什么是分布式递增ID？
**传统自增id**  
通常由数据库自己维护，业务复杂情况下，会有id冲突，无法分库分表，更无法分布式使用。  
**uuid**  
长度太长一般为36位字符串，占据空间大，UUID一般是无序，数据库无法使用插入优化。  
用于分布式系统中有ID冲突风险。  
**SnowFlake**  
是twitter发明，用于生成递增的无符号整数的分布式数字id，转换成整型长度为19位的数字。  
大厂基于SnowFlake的开源项目。
百度 UIDGenertor，美团 Leaf，腾讯 Seqsvr等。  
算法单机每秒内理论上最多可以生成1000*(2^12)，也就是409.6万个ID。

# 这个项目是什么？解决什么问题？
**痛点**
1. SnowFlake的id是大整形，前端如js、node无法接收。
2. 采用SnowFlake的字符串形式，没有了数据库插入优势。
3. 在后端程序中做转换，写法复杂不优雅且易遗漏。

**解决**  
本项目通过golang语言，开发mysql的SnowFlake的编解码库，用于彻底把SnowFlake分层。  
编写前端后端，一律使用SnowFlake的字符串形式。  
插入数据库和取出数据的时候，本库在数据库层面转换id为大整形。

# 安装
**查看plugin地址**  
`show variables like '%plugin_dir%'`

**将 sid.so 放到 plugin_dir 目录下**
**重启数据库**  
mac:  
`brew services restart mysql`  
linux:  
`service mysql restart`  

**创建函数**  
解码字符串为bigint数字id  
```sql
CREATE FUNCTION DSID RETURNS INTEGER SONAME 'sid.so';
```
编码bigint为字符串id  
```sql
CREATE FUNCTION ESID RETURNS STRING SONAME 'sid.so';
```
**删除函数**  
```sql
DROP FUNCTION DSID;
```
```sql
DROP FUNCTION ESID;
```

# 使用方法
**将数字转为字符串**
```sql
SELECT ESID(1499287156466651136) AS id;
-- 输出 4tRpRfVhPab
```
**将字符串转为数字**
```sql
SELECT DSID('4tRpRfVhPab') AS id;
-- 输出 1499287156466651136
```
 
