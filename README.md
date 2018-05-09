# x-server

token x server

## user

- create user

curl -s -X POST http://localhost:8080/api/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

- login passwd为passwd的 sha256 hash

curl -s -X POST http://localhost:8080/login -d 'userID=13810167616&passwd=0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e'

- user info

curl http://localhost:8080/api/v1/users/13810167616

- /users/updatepasswd/:userid/:old/:new

curl -s -X POST http://localhost:8080/api/v1/users/updatepasswd/13810167616/3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea/a11111

## account

- create account /account/create/:userid/:name/:password

curl -s -X POST http://localhost:8080/api/v1/account/create/13810167616/培华/a11111

return

0x946ED2768601cD9e232b1791dd883B93cd7D8CAc

- /account/list/:userid

curl http://localhost:8080/api/v1/account/list/13810167616

- get account /account/info/:address

curl http://localhost:8080/api/v1/account/info/0x6f2330c4a2cea74a35ac3ae5ec04308a50ae2a60

- /account/updatepwd/:addr/:password/:newpassword

curl -s -X POST http://localhost:8080/api/v1/account/updatepwd/0xBB89978Cc97f30661f635CbF015639C078D31523/a11111/b11111

curl -s -X POST http://localhost:8080/api/v1/account/updatepwd/0xBB89978Cc97f30661f635CbF015639C078D31523/b11111/a11111

- /account/recover/:addr/:newpassword

curl -s -X POST http://localhost:8080/api/v1/account/recover/0xBB89978Cc97f30661f635CbF015639C078D31523/a11111

## token

- deploy token /token/deploy

curl -s -X POST http://localhost:8080/api/v1/token/deploy -d 'address=0x5C0417A3a20F928e1Ea115E0A5c185d0879E153F&pwd=a11111&name=fifu2&symbol=FIFU2&total=1000000000&desc=fifu2 deploy'

return 0x09e86ffe4333212f20f7ec958a166e8fdb0c6aa5

0xe3d032720bddcce775cd66c1d864b7030a733f5f

0xa6e40695a50fae5f934298124e241224d37fd8bb

0xb3804f741937595475bb7e8e46e7645613a5705a

- balance /token/balance/:addr

curl -s -X POST http://localhost:8080/api/v1/token/balance/0x5C0417A3a20F928e1Ea115E0A5c185d0879E153F -d 'conaddrs=0xe3d032720bddcce775cd66c1d864b7030a733f5f,0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd,0xa6e40695a50fae5f934298124e241224d37fd8bb,0xb3804f741937595475bb7e8e46e7645613a5705a'

curl -s -X POST http://localhost:8080/api/v1/token/balance/0xeca4635f3fE81b4b8Cc6d40deFf99Eb8428C7BeD -d 'conaddrs=0x09e86ffe4333212f20f7ec958a166e8fdb0c6aa5,0x70066930c500dbc07517b11e0393dc260f7296db'

curl -s -X POST http://localhost:8080/api/v1/token/balance/0xeca4635f3fE81b4b8Cc6d40deFf99Eb8428C7BeD -d 'conaddrs=0x09e86ffe4333212f20f7ec958a166e8fdb0c6aa5,0x70066930c500dbc07517b11e0393dc260f7296db'

- transfer /token/transfer

curl -s -X POST http://localhost:8080/api/v1/token/transfer -d 'conaddr=0xb3804f741937595475bb7e8e46e7645613a5705a&from=0x5C0417A3a20F928e1Ea115E0A5c185d0879E153F&to=0xeca4635f3fE81b4b8Cc6d40deFf99Eb8428C7BeD&amount=1000&pwd=a11111&memo=loan'

- /token/transfer/list/:tokenaddress/:address/:page

curl http://localhost:8080/api/v1/token/transfer/list/0x09e86ffe4333212f20f7ec958a166e8fdb0c6aa5/0x3c5ffa487ea89a36d3f05166bba15b959e315a59/1

- /token/transfer/all/:tokenaddress/:page

curl http://localhost:8080/api/v1/token/transfer/all/0x09e86ffe4333212f20f7ec958a166e8fdb0c6aa5/1

- /token/transfer/count/:tokenaddress

curl http://localhost:8080/api/v1/token/transfer/count/0x09e86ffe4333212f20f7ec958a166e8fdb0c6aa5

- TokenInfo /token/info/:address

curl http://localhost:8080/api/v1/token/info/0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd

- UpdateTokenWeight /token/weight/:address/:weight

curl -s -X POST http://localhost:8080/api/v1/token/weight/0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd/100

- ListToken /token/list/:page

curl http://localhost:8080/api/v1/token/list/1

- ListTokenTransfer /token/transfer/list/:address/:page

curl http://localhost:8080/api/v1/token/transfer/list/0x3c5ffa487ea89a36d3f05166bba15b959e315a59/1

curl http://localhost:8080/api/v1/token/transfer/list/0xeca4635f3fE81b4b8Cc6d40deFf99Eb8428C7BeD/1

## view

- /view/create/:userid/:address/:tokenaddress

curl -s -X POST http://localhost:8080/api/v1/view/create/13810167616/0x6f2330C4A2ceA74a35ac3AE5ec04308A50Ae2A60/0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd

- /view/delete/:userid/:address/:tokenaddress

curl -s -X POST http://localhost:8080/api/v1/view/delete/13810167616/0x6f2330C4A2ceA74a35ac3AE5ec04308A50Ae2A60/0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd

- /view/info/:userid/:address/:tokenaddress

curl http://localhost:8080/api/v1/view/info/13810167616/0x6f2330C4A2ceA74a35ac3AE5ec04308A50Ae2A60/0x95d9d1f3c47f49ca6201a6ec8c0431310c16a2fd

- /view/list/:userid/:address

curl http://localhost:8080/api/v1/view/list/13810167616/0x6f2330C4A2ceA74a35ac3AE5ec04308A50Ae2A60

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
