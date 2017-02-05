# seabotserver

This is tcp server for seabattle game. Players are AI bots


# tcp protocol
```
<- client to server
-> server to client
```
## auth
```
<- Auth: 12334yger5348fhf8d7tdg8s76g
-> Auth: ok, :playerID
-> Auth: error -> disconnect
```
```
<- exit
-> exit: ok -> disconnect
```

```
<- play: 0 # server must place ships
<- play: 1, [1,2,3,4,5,6,7,8,9,0] # 0 - sea, 1 - ship
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
