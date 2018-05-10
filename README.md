# x-server

token x server

## user

- create user

curl -s -X POST http://123.206.29.15:4100/api/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

- login passwd为passwd的 sha256 hash

curl -s -X POST http://123.206.29.15:4100/login -d 'userID=13810167616&passwd=0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e'

- user info

curl http://123.206.29.15:4100/api/v1/users/13810167616

- /users/updatepasswd/:userid/:old/:new

curl -s -X POST http://123.206.29.15:4100/api/v1/users/updatepasswd/13810167616/0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e/a11111

## account

- create account /account/create/:userid/:name/:password

curl -s -X POST http://123.206.29.15:4100/api/v1/account/create/13810167616/培华/a11111

return

0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c

- /account/list/:userid

curl http://123.206.29.15:4100/api/v1/account/list/13810167616

- /account/updatepwd/:addr/:password/:newpassword

curl -s -X POST http://123.206.29.15:4100/api/v1/account/updatepwd/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/a11111/b11111

curl -s -X POST http://123.206.29.15:4100/api/v1/account/updatepwd/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/b11111/a11111

- /account/recover/:addr/:newpassword

curl -s -X POST http://123.206.29.15:4100/api/v1/account/recover/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/a11111

## token

- deploy token /token/deploy

curl -s -X POST http://123.206.29.15:4100/api/v1/token/deploy -d 'address=0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c&pwd=a11111&name=fifu&symbol=FIFU&total=1000000000&desc=fifu deploy'

return 0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- balance /token/balance/:addr

curl -s -X POST http://123.206.29.15:4100/api/v1/token/balance/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c -d 'conaddrs=0x1dd80b503e3b5de5724fe204bc87fb5387b0470c,0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd,0xa6e40695a50fae5f934298124e241224d37fd8bb,0xb3804f741937595475bb7e8e46e7645613a5705a'

- transfer /token/transfer

curl -s -X POST http://123.206.29.15:4100/api/v1/token/transfer -d 'conaddr=0x1dd80b503e3b5de5724fe204bc87fb5387b0470c&from=0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c&to=0xeca4635f3fE81b4b8Cc6d40deFf99Eb8428C7BeD&amount=1000&pwd=a11111&memo=loan'

- /token/transfer/list/:tokenaddress/:address/:page

curl http://123.206.29.15:4100/api/v1/token/transfer/list/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/1

- /token/transfer/all/:tokenaddress/:page

curl http://123.206.29.15:4100/api/v1/token/transfer/all/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c/1

- /token/transfer/count/:tokenaddress

curl http://123.206.29.15:4100/api/v1/token/transfer/count/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- TokenInfo /token/info/:address

curl http://123.206.29.15:4100/api/v1/token/info/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- UpdateTokenWeight /token/weight/:address/:weight

curl -s -X POST http://123.206.29.15:4100/api/v1/token/weight/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c/100

- ListToken /token/list/:page

curl http://123.206.29.15:4100/api/v1/token/list/1

## view

- /view/create/:userid/:address/:tokenaddress

curl -s -X POST http://123.206.29.15:4100/api/v1/view/create/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/delete/:userid/:address/:tokenaddress

curl -s -X POST http://123.206.29.15:4100/api/v1/view/delete/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/info/:userid/:address/:tokenaddress

curl http://123.206.29.15:4100/api/v1/view/info/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/list/:userid/:address

curl http://123.206.29.15:4100/api/v1/view/list/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c

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
