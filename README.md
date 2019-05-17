# x-server

token x server

## user

- create user

curl -s -X POST http://47.89.15.157:13500/api/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

curl -s -X POST http://123.206.29.15:4100/api/v1/users/create -d 'userID=18810096114&passwd=a11111&name=sui&email=ministor@126.com'

- login passwd

curl -s -X POST http://47.89.15.157:13500/login -d 'userID=13810167616&passwd=a11111'

- user info

curl http://47.89.15.157:13500/api/v1/users/13810167616

- /users/updatepasswd/:userid/:old/:new

curl -s -X POST http://123.206.29.15:4100/api/v1/users/updatepasswd/13810167616/0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e/a11111

- /user/count

curl http://localhost:8080/api/v1/user/count

- /user/list/:search/:page

curl http://localhost:8080/api/v1/user/list/138/1

## adminuser

- create admin user

curl -s -X POST http://47.89.15.157:13500/api/admin/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

- login passwd

curl -s -X POST http://47.89.15.157:13500/api/admin/v1/login -d 'userID=13810167616&passwd=a11111'

- user info

curl http://localhost:8080/api/admin/v1/users/13810167616

- /users/updatepasswd/:userid/:old/:new

curl -s -X POST http://localhost:8080/api/admin/v1/users/updatepasswd/13810167616/a11111/a11112

## account

- create account /account/create/:userid/:name/:password

curl -s -X POST http://47.89.15.157:13500/api/v1/account/create/13810167616/13810167616/a11111

return

0x7Ab49510710dBE327bCcB705d9aEC8714E15674c

curl -s -X POST http://localhost:5100/api/v1/account/create/18810096114/18810096114/a11111

return : 0x38ffE5e6f20942DD4e2A105702DaadECa8b8c40f

- /account/list/:userid

curl http://localhost:8080/api/v1/account/list/13810167616

- /account/updatepwd/:addr/:password/:newpassword

curl -s -X POST http://123.206.29.15:4100/api/v1/account/updatepwd/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/a11111/b11111

curl -s -X POST http://localhost:8080/api/v1/account/updatepwd/0xC326E6afb88528Fdb3A39883b88aef7e2Ec066A4/b11111/a11111

- /account/recover/:addr/:newpassword

curl -s -X POST http://123.206.29.15:4100/api/v1/account/recover/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/a11111

## token

- deploy token /token/deploy

curl -s -X POST http://47.89.15.157:13500/api/v1/token/deploy -d 'userID=13810167616&pwd=a11111&name=歌手A&symbol=STT&total=1000000000&desc=歌手A发的Token'

return 0x37f3c9e63d7791df19313b9e9bbed05ce1af61c4

curl -s -X POST http://localhost:5100/api/v1/token/deploy -d 'userID=13810167616&pwd=a11111&name=歌手B&symbol=STA&total=1000000000&desc=歌手B发的Token'

return 0xbf747f0da136d4b41e156ec32f12aefa4466b779

- balance /token/balance/:addr

curl -s -X POST http://47.89.15.157:13500/api/v1/token/balance/0x7Ab49510710dBE327bCcB705d9aEC8714E15674c -d 'conaddrs=0x37f3c9e63d7791df19313b9e9bbed05ce1af61c4,0x1da146dafb1999feaca30f926f057505ebab6328'

- /token/count

curl http://localhost:8080/api/v1/token/count

- transfer /token/transfer

curl -s -X POST http://123.206.29.15:4100/api/v1/token/transfer -d 'conaddr=0x387127e92e95f492f84f28446cf542ad85d43bbf&from=0xBA168bB12b41d8c28b5D0266038dE1a387654032&to=0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0&amount=1000&pwd=a11111&memo=loan'

curl -s -X POST http://47.89.15.157:13500/api/v1/token/transfer -d 'conaddr=0x37f3c9e63d7791df19313b9e9bbed05ce1af61c4&from=0x7Ab49510710dBE327bCcB705d9aEC8714E15674c&to=0x38ffE5e6f20942DD4e2A105702DaadECa8b8c40f&amount=1000&pwd=a11111&memo=loan'

- /token/transfer/list/:tokenaddress/:address/:page

curl http://123.206.29.15:4100/api/v1/token/transfer/list/0x13e55998e931687c1a19d6281d58fb3622e5c6fc/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/1

- /token/transfer/all/:tokenaddress/:page

curl http://123.206.29.15:4100/api/v1/token/transfer/all/0x13e55998e931687c1a19d6281d58fb3622e5c6fc/1

- /token/transfer/count/:tokenaddress

curl http://123.206.29.15:4100/api/v1/token/transfer/count/0x13e55998e931687c1a19d6281d58fb3622e5c6fc

- /token/transfer/allcount

curl http://localhost:8080/api/v1/token/transfer/allcount

- /token/transfer/allsum

curl http://localhost:8080/api/v1/token/transfer/allsum

- TokenInfo /token/info/:address

curl http://123.206.29.15:4100/api/v1/token/info/0x13e55998e931687c1a19d6281d58fb3622e5c6fc

- UpdateTokenWeight /token/weight/:address/:weight

curl -s -X POST http://123.206.29.15:4100/api/v1/token/weight/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c/100

- ListToken /token/list/:search/:page

curl http://localhost:5100/api/v1/token/list//1

## view

- /view/create/:userid/:address/:tokenaddress

curl -s -X POST http://123.206.29.15:4100/api/v1/view/create/13810167616/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/0xd9a5883d56ce1cd11dd0c294c9d6daa261267a25

- /view/delete/:userid/:address/:tokenaddress

curl -s -X POST http://123.206.29.15:4100/api/v1/view/delete/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/info/:userid/:address/:tokenaddress

curl http://123.206.29.15:4100/api/v1/view/info/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/list/:userid/:address

curl http://123.206.29.15:4100/api/v1/view/list/13810167616/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0

- /view/balance/:userid/:address

curl http://123.206.29.15:4100/api/v1/view/balance/13810167616/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0

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

http -f GET localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0NDk0MzUsImlhdCI6MTU0MDQ0NTgzNSwic3ViIjoiMTM4MTAxNjc2MTYifQ.igJd5llGMqiDOoEMMrQ9qqpDVkUQuvmYkpDGEb-BdH4" "Content-Type: application/json"

http -f GET localhost:8080/api/auth/v1/hello "Content-Type: application/json"

http -v --json POST localhost:8080/api/auth/v1/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjUzMzM3MjQsImlhdCI6MTUyNTMzMDEyNCwic3ViIjoiMTM4MTAxNjc2MTYifQ.FG7PTl4QBhJ8VJxEB3Q94x7smPSygMZoY6zQQWDiZQs"
