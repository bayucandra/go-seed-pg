# Getting started guide

- Copy and rename template **.env.example** to **.env**
- Adjust **.env** accordingly
- `SOURCE_DATA` inside **.env** is path of your data seed which contain CSV exported data from PostgreSQL table
- So, create a directory based on your `SOURCE_DATA` variable
- Sort the table based on relationship dependencies
- File name template is: `<xxx>.<xxx>_table_name.csv` where xxx is 3 digit number to sort the order of table seeding
- For example:
    - 001.001_table_user_group.csv
    - 002.001_table_user_permission.csv
    - 003.001_table_user.csv
- Exported CSV must still contain fields/columns name at first row
