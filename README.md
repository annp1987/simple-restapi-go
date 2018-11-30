# simple-httpsever-restapi-go

How to test
===========
# get list staffs
$ curl -X GET http://127.0.0.1:8000/api
# show staff
$ curl -X GET http://127.0.0.1:8000/api/{id}
# create a staff
$ curl -X POST -d'{"Id": "3", "Name": "qwe"}' http://127.0.0.1:8000/api
# delete a staff
$ curl -X DELETE http://127.0.0.1:8000/api/{id}
