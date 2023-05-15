# leetcode-esque
The project is best described by this video where I showcase many of the features:

https://www.youtube.com/watch?v=u85M_KnJFe4

## Purpose
Many people have the interview process so I wanted to look into using different tools to try help educate people on the process.
I used GPT 3.5 to help users when they have failed test cases, this never gives the user code but gives hints if there are syntax errors or they are close to being correct. 
This emulates the actual interview process where you sometimes get given small hints and advice from the interviewer.

## Technologies
It is meant to be designed with scaling in mind - so that it could be deployed within a university. 
Postgres, Redis, Golang, JS, HTML, Python.

Most of the backend web servers are built in Go, the majority of the datastorage is done using Postgres and then the session / login system uses Redis. The frontend uses basic CSS, HTML + JS.

The backend code can be found here - https://github.com/1rvyn/backend-api
