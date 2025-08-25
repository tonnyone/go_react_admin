brew install postgresql@14

```
echo 'export PATH="/usr/local/opt/postgresql@15/bin:$PATH"' >> ~/.zshrc
```


```
rm -rf /opt/homebrew/var/postgresql@14/ 
initdb --locale=C -E UTF-8 /opt/homebrew/var/postgresql@14
```

```
pg_ctl -D '/opt/homebrew/var/postgresql@14' -l logfile start
```


```
brew services start postgresql@14
brew services stop postgresql@14
```

```
psql postgres
```


```
CREATE DATABASE demodb ENCODING 'UTF8' LC_COLLATE='en_US.UTF-8' LC_CTYPE='en_US.UTF-8' TEMPLATE=template0;
CREATE USER dbuser WITH PASSWORD '123456';
GRANT ALL PRIVILEGES ON DATABASE demodb TO dbuser;
CREATE TABLE test (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    attributes JSONB
);
```


```
psql -h 127.0.0.1 -p 5432 -U dbuser -d demodb
```


## 参考
https://sqlconjuror.com/mysql-and-postgresql-equivalent-commands/
http://www.postgres.cn/docs/14/index.html
https://www.sjkjc.com/postgresql/psql-commands/