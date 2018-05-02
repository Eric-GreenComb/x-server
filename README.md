# x-server

token x server

## user

- create user

curl -s -X POST http://localhost:8080/api/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

- user info

curl http://localhost:8080/api/v1/users/13810167616

## jwt

- Login API:

http -v --json POST localhost:8080/login username=13810167616 password=0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e

- Refresh token API:

http -v -f GET localhost:8080/api/auth/v1/refresh_token "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUyNTA0NTUsImlkIjoiMTM4MTAxNjc2MTYiLCJvcmlnX2lhdCI6MTUyNTI0Njg1NX0.5Te4JxQ3Dmb6SsjBrzNe56gyQi8fO3NzPlvvqGOgd3Y"  "Content-Type: application/json"

- Hello world

http -f GET localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUyNTA0NTUsImlkIjoiMTM4MTAxNjc2MTYiLCJvcmlnX2lhdCI6MTUyNTI0Njg1NX0.5Te4JxQ3Dmb6SsjBrzNe56gyQi8fO3NzPlvvqGOgd3Y" "Content-Type: application/json"

http -v --json POST localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUyNTA0NTUsImlkIjoiMTM4MTAxNjc2MTYiLCJvcmlnX2lhdCI6MTUyNTI0Njg1NX0.5Te4JxQ3Dmb6SsjBrzNe56gyQi8fO3NzPlvvqGOgd3Y"
