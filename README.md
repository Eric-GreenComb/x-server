# x-server

token x server

## jwt

- Login API:

http -v --json POST localhost:8080/login username=admin password=admin

- Refresh token API:

http -v -f GET localhost:8080/api/auth/v1/refresh_token "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUxMDY0NTUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUyNTEwMjg1NX0.9pn5wi0qQv3aTG6acwzAyPrw6Dq8kS-qpu7nZbEudL8"  "Content-Type: application/json"

- Hello world

http -f GET localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUxMDY0NTUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUyNTEwMjg1NX0.9pn5wi0qQv3aTG6acwzAyPrw6Dq8kS-qpu7nZbEudL8"  "Content-Type: application/json"

- Authorization

http -v --json POST localhost:8080/login username=test password=test

http -f GET localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUxMDcwODEsImlkIjoidGVzdCIsIm9yaWdfaWF0IjoxNTI1MTAzNDgxfQ.Iss3mj8Z0ynfZSgEAqHepB-yMdYesqv9lxwOdXMBt50"  "Content-Type: application/json"
