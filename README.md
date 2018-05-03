# x-server

token x server

## user

- create user

curl -s -X POST http://localhost:8080/api/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

- user info

curl http://localhost:8080/api/v1/users/13810167616

## token

- deploy token /token/deploy/:name/:symbol/:keystore/:pwd

curl -s -X POST http://localhost:8080/api/v1/token/deploy -d 'name=fifu&symbol=FIFU&keystore=/home/eric/go/src/github.com/Eric-GreenComb/parity-poa/parity0/keys/FiFuPoA/UTC--2018-05-03T08-41-40.750244197Z--87163ed732947d8e400ebc2772f1f1837b5e0ffa&pwd=a11111'

- balance /token/balance

http://localhost:8080/api/v1/token/balance?conaddr=0xfBE97c0324C44824F1d1c26669Dc590162cB9f16&addr=0x87163ED732947D8E400eBC2772F1F1837B5e0Ffa

- transfer /token/transfer

curl -s -X POST http://localhost:8080/api/v1/token/transfer -d 'conaddr=0xfBE97c0324C44824F1d1c26669Dc590162cB9f16&to=0xeca4635f3fE81b4b8Cc6d40deFf99Eb8428C7BeD&keystore=/home/eric/go/src/github.com/Eric-GreenComb/parity-poa/parity0/keys/FiFuPoA/UTC--2018-05-03T08-41-40.750244197Z--87163ed732947d8e400ebc2772f1f1837b5e0ffa&amount=1000&pwd=a11111'

## badger

- SetBadgerKey /badger/set/:key/:value

curl -s -X POST http://localhost:8080/api/v1/badger/set/name/eric

- SetBadgerKeyTTL /badger/setwithttl/:key/:value

curl -s -X POST http://localhost:8080/api/v1/badger/setwithttl/name/eric/5

- GetBadgerKey /badger/get/:key

curl http://localhost:8080/api/v1/badger/get/name

## jwt

- Login API:

http -v --json POST localhost:8080/login userID=13810167616 passwd=0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e

- Refresh token API:

http -v -f GET localhost:8080/api/auth/v1/refresh_token "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUzMzM3MjQsImlhdCI6MTUyNTMzMDEyNCwic3ViIjoiMTM4MTAxNjc2MTYifQ.FG7PTl4QBhJ8VJxEB3Q94x7smPSygMZoY6zQQWDiZQs"  "Content-Type: application/json"

- Hello world

http -f GET localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUzMzM3MjQsImlhdCI6MTUyNTMzMDEyNCwic3ViIjoiMTM4MTAxNjc2MTYifQ.FG7PTl4QBhJ8VJxEB3Q94x7smPSygMZoY6zQQWDiZQs" "Content-Type: application/json"

http -v --json POST localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUzMzM3MjQsImlhdCI6MTUyNTMzMDEyNCwic3ViIjoiMTM4MTAxNjc2MTYifQ.FG7PTl4QBhJ8VJxEB3Q94x7smPSygMZoY6zQQWDiZQs"
