#!/bin/bash

# brew --prefix postgresql
# brew --prefix libpq

# 配置参数
DB_NAME="demodb"
DB_USER="dbuser"
DB_HOST="localhost"
DB_PORT="5432"
OUTPUT_FILE="db_dump.sql"

# 导出所有表结构和数据
PGPASSWORD="123456" pg_dump -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -F p -c -O > $OUTPUT_FILE

echo "导出完成，文件位置：$OUTPUT_FILE"