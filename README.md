To run this website, you need follow command(Put your MySQL password into x, if you didn't install MySQL server, you need to install one):
$env:DBPASS='x'

Run the main.go:
go run main.go

If you don't know how to create the database in the MySQL, you can do this, open MySQL server:
mysql -u root -p

Enter your password, then type this command(you can use other name, but you need to change the DBName in the main.go too):
create database usersshark;

Then use this command:
use usersshark;

Then this:
source table.sql;

Then you can test it, type command:
select * from users;

Then you will get this:
+---------+------------+------------------+
| user_id | username   | password         |
+---------+------------+------------------+
|       1 | firstuser  | firstuser        |
+---------+------------+------------------+

Then you are good to go, you can go to localhost:8080 to see this website, to do this, you need type the URL:localhost:8080 in the webbrowser(this website only you and people in your network can see).

Please click on raw button to see the this file.
