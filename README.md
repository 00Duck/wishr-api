# Wishr
An open source, self-hosted wish list app for groups and families.

### Database setup
Run `mysql_secure_installation` and configure as needed.

`mysql -u root` and enter the following:

```sql
create database wishr;
create user 'wishr'@'localhost' identified by 'password';
grant all privileges on wishr.* to 'wishr'@'localhost';
flush privileges;
```