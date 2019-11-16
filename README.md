# Getting started guide

## Go-seed-pg

**Go-seed-pg** is database table seeding for PostgreSQL. It is mainly using CSV file as data source.

Following is steps to seed the database:

- Configuration setting
    
    Mainly, configuration is passed trough Environment variables. There are 2 option regarding this matter:
    
    1. By **.env** file. You can copy and rename template **.env.example** to **.env**. Or by **.env-custom** then pass it by : `go-seed-pg .env-custom`
    2. Directly configure Environment variables in your OS ( helpful for CI/CD things ). If you using linux or OSX, should be something like `export VAR_NAME=value`. Please check list of variables need to sets at .env-example
    
- Adjust **.env** or **environment variables** ( depend on **Configuration setting** you choose) accordingly
- `SOURCE_DATA` inside **.env** is path of your data seed which contain CSV exported data from PostgreSQL table
- So, create a directory based on your `SOURCE_DATA` variable
- Sort the table based on relationship dependencies
- File name template is: `<xxx>.<xxx>_table_name.csv` where xxx is 3 digit number to sort the order of table seeding
- For example:
    - 001.001_table_user_group.csv
    - 002.001_table_user_permission.csv
    - 003.001_table_user.csv
- CSV data rules:
    - Exported CSV must still contain fields/columns name at first row
    - Null values must be exported as `null` not in quote or empty quote
    
CSV Example:

```
"uid","email","""password""","name","role_id","group_id","is_active","is_approved","created_at"
"1","admin@example.com","","BOT",null,null,"true","true","2019-07-30 19:00:15"
```
