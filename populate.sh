#!/usr/bin/fish
http PUT :8080/todos title=test1 text="testing" due_at="2023-09-26T13:13:13Z"
http PUT :8080/todos title=test2 text="more testing" due_at="2023-09-27T13:13:13Z"

