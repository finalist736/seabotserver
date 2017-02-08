# seabotserver

This is tcp server for seabattle game. Players are AI bots


# tcp protocol
```
<- client to server
-> server to client
```
## auth
```json
<- { "auth" : "12334yger5348fhf8d7tdg8s76g" }
-> { "auth" : 
		{ 
			"ok": false, 
			"error": "some error",
			"id" : 123  
		}
	}
```
```
<- { "exit" : null }
-> disconnect
```

```
<- { "bvb" : 0 }
-> { "bvb" : { "wait": null } }
-> { "bvb" : { "enemy": 321, "ships": [0,0...] } }

<- bvb: 0 - server must place ships, 1 - user sends him ships
<- bvb: 1, [1,2,3,4,5,6,7,8,9,0] # 0 - sea, 1 - ship
-> error -> disconnect
-> wait
```
```
-> play: :playerID, :name, :sea:[0,0,0,1,0,0,1.....] 
	:turn: 1|43
```
##### LOOP: GAME
```
<- turn:"A2",0-1
-> turn,"A2":ok,miss,
-> turn,43-"A3",miss
```
##### GOTO GAME
```
// 10 second timeout
// after timeout -> lose
```
## Battle end
```
-> battleEnd: winner: 1, loser: 43;
```
