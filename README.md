# go-ORM
## 项目描述
go-ORM 是一个学习型项目，实现了一个简单的 ORM 框架。

## 功能实现
go-ORM 目前实现了的功能有：
* 表的创建、删除、迁移
* 记录的 CRUD 操作，查询条件的链式操作
* 单一主键的设置
* 钩子功能
* 事务功能
* 简单的数据库迁移(仅支持字段的新增与删除，不支持字段类型变更)

## 目录结构
```bash
$ tree
.
├── clause
│   ├── clause.go
│   ├── clause_test.go
│   └── generator.go
├── dialect
│   ├── dialect.go
│   ├── sqlite3.go
│   └── sqlite3_test.go
├── go.mod
├── goorm.go
├── goorm_test.go
├── go.sum
├── LICENSE
├── log
│   ├── log.go
│   └── log_test.go
├── README.md
├── schema
│   ├── schema.go
│   └── schema_test.go
└── session
    ├── hooks.go
    ├── hooks_test.go
    ├── raw.go
    ├── raw_test.go
    ├── record.go
    ├── record_test.go
    ├── table.go
    ├── table_test.go
    └── transaction.go

5 directories, 25 files
```