# x-server

token x server

## user

- create user

curl -s -X POST http://123.206.29.15:4100/api/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

curl -s -X POST http://123.206.29.15:4100/api/v1/users/create -d 'userID=18810096114&passwd=a11111&name=eric&email=ministor@126.com'

- login passwd

curl -s -X POST http://123.206.29.15:4100/login -d 'userID=13810167616&passwd=a11111'

- user info

curl http://123.206.29.15:4100/api/v1/users/13810167616

- /users/updatepasswd/:userid/:old/:new

curl -s -X POST http://123.206.29.15:4100/api/v1/users/updatepasswd/13810167616/0fb6c6c0b7621fb7bd6ff1e6fb656bc746e2254a4f671dee25c0ce3ddd9ccf3e/a11111

- /user/count

curl http://123.206.29.15:4100/api/v1/user/count

- /user/list/:search/:page

curl http://123.206.29.15:4100/api/v1/user/list/138/1

## adminuser

- create admin user

curl -s -X POST http://123.206.29.15:4100/api/admin/v1/users/create -d 'userID=13810167616&passwd=a11111&name=eric&email=ministor@126.com'

- login passwd

curl -s -X POST http://123.206.29.15:4100/api/admin/v1/login -d 'userID=13810167616&passwd=a11111'

- user info

curl http://123.206.29.15:4100/api/admin/v1/users/13810167616

- /users/updatepasswd/:userid/:old/:new

curl -s -X POST http://123.206.29.15:4100/api/admin/v1/users/updatepasswd/13810167616/a11111/a11112

## account

- create account /account/create/:userid/:name/:password

curl -s -X POST http://123.206.29.15:4100/api/v1/account/create/13810167616/13810167616/a11111

return

0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0

- /account/list/:userid

curl http://123.206.29.15:4100/api/v1/account/list/13810167616

- /account/updatepwd/:addr/:password/:newpassword

curl -s -X POST http://123.206.29.15:4100/api/v1/account/updatepwd/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/a11111/b11111

curl -s -X POST http://123.206.29.15:4100/api/v1/account/updatepwd/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/b11111/a11111

- /account/recover/:addr/:newpassword

curl -s -X POST http://123.206.29.15:4100/api/v1/account/recover/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/a11111

## token

- deploy token /token/deploy

curl -s -X POST http://123.206.29.15:4100/api/v1/token/deploy -d 'userID=13810167616&pwd=a11111&name=歌手C&symbol=STC&total=1000000000&desc=歌手C发的Token'

return 0x573d52f6547ca177384ad763e4435ce7b566b9f7

curl -s -X POST http://123.206.29.15:4100/api/v1/token/deploy -d 'userID=13810167616&pwd=a11111&name=歌手A&symbol=STA&total=1000000000&desc=歌手A发的Token'

return 0xc8d5357c4a0857f265a8b0d4185c00ae7e817954

- balance /token/balance/:addr

curl -s -X POST http://123.206.29.15:4100/api/v1/token/balance/0x946ED2768601cD9e232b1791dd883B93cd7D8CAc -d 'conaddrs=0x573d52f6547ca177384ad763e4435ce7b566b9f7,0xc8d5357c4a0857f265a8b0d4185c00ae7e817954,0xa6e40695a50fae5f934298124e241224d37fd8bb,0xb3804f741937595475bb7e8e46e7645613a5705a'

- /token/count

curl http://123.206.29.15:4100/api/v1/token/count

- transfer /token/transfer

curl -s -X POST http://123.206.29.15:4100/api/v1/token/transfer -d 'conaddr=0x387127e92e95f492f84f28446cf542ad85d43bbf&from=0xBA168bB12b41d8c28b5D0266038dE1a387654032&to=0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0&amount=1000&pwd=a11111&memo=loan'

- /token/transfer/list/:tokenaddress/:address/:page

curl http://123.206.29.15:4100/api/v1/token/transfer/list/0x13e55998e931687c1a19d6281d58fb3622e5c6fc/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/1

- /token/transfer/all/:tokenaddress/:page

curl http://123.206.29.15:4100/api/v1/token/transfer/all/0x13e55998e931687c1a19d6281d58fb3622e5c6fc/1

- /token/transfer/count/:tokenaddress

curl http://123.206.29.15:4100/api/v1/token/transfer/count/0x13e55998e931687c1a19d6281d58fb3622e5c6fc

- /token/transfer/allcount

curl http://123.206.29.15:4100/api/v1/token/transfer/allcount

- /token/transfer/allsum

curl http://123.206.29.15:4100/api/v1/token/transfer/allsum

- TokenInfo /token/info/:address

curl http://123.206.29.15:4100/api/v1/token/info/0x13e55998e931687c1a19d6281d58fb3622e5c6fc

- UpdateTokenWeight /token/weight/:address/:weight

curl -s -X POST http://123.206.29.15:4100/api/v1/token/weight/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c/100

- ListToken /token/list/:search/:page

curl http://123.206.29.15:4100/api/v1/token/list/138/1

## view

- /view/create/:userid/:address/:tokenaddress

curl -s -X POST http://123.206.29.15:4100/api/v1/view/create/13810167616/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0/0x13e55998e931687c1a19d6281d58fb3622e5c6fc

- /view/delete/:userid/:address/:tokenaddress

curl -s -X POST http://123.206.29.15:4100/api/v1/view/delete/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/info/:userid/:address/:tokenaddress

curl http://123.206.29.15:4100/api/v1/view/info/13810167616/0x20B5c492B53919e4bfC279C4a43CE350Bb7fDF7c/0x1dd80b503e3b5de5724fe204bc87fb5387b0470c

- /view/list/:userid/:address

curl http://123.206.29.15:4100/api/v1/view/list/13810167616/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0

- /view/balance/:userid/:address

curl http://123.206.29.15:4100/api/v1/view/balance/13810167616/0x12f91a58bf2714ec602f3c3b9841Ddf98478DFF0