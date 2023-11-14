# message queue go

```sql
mysql
> CREATE
DATABASE web_payments;
       CREATE
USER 'web_payments_user'@'localhost' IDENTIFIED WITH mysql_native_password BY 'EXAMPLE_PASSWORD';
       GRANT ALL PRIVILEGES ON web_payments.* TO
'web_payments_user'@'localhost';
       FLUSH
PRIVILEGES;
```

```sql
mysql
>
CREATE TABLE payments
(
    payment_id     BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    payment_date   DATETIME,
    first_name     VARCHAR(50),
    last_name      VARCHAR(50),
    payment_mode   VARCHAR(255),
    payment_ref_no VARCHAR(255),
    amount         DECIMAL(17, 4)
) ENGINE = InnoDB;
```