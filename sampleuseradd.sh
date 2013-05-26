#!/bin/bash
curl 127.0.0.1:8080/users --data '{"Action":"NEW", "Token":"thisismytoken", "Value":{"UUID":"11229324","FirstName":"Jeromy","LastName":"Johnson","Resume":"None Yet","Skills":["Golang","C","Genius"],"Major":"Computer Science","Minors":[],"Interests":["Females","Programming"],"Email":"me@jero.my"}}'
