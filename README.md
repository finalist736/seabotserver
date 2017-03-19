# seabotserver

This is tcp server for seabattle game. Players are AI bots


# tcp protocol
```
<- client to server
-> server to client

[0 0 0 19]{"field" : "value"}

first 4 butes = json length
первые 4 байта = длина json команды

```
## auth
```json
<- { "auth" : "12334yger5348fhf8d7tdg8s76g" }
-> { "auth" : 
		{ 
			"ok": 		false, 
			"error": 	"some error",
			"id" : 		123 // bot ID
			"userid" : 	11 // user ID 
		}
	}
```
```json
<- { "exit" : true }
-> disconnect
```

```json
// bot versus random
// if you want to debug your code, use this command.
<- { "bvr" : { "place": 0 } }
-> { "bvr" : { "id": -1, "name": "bot_-1", "ships": [0,0,0,0...] } }
```

```json
// bot versus bot
// сервер сам расставляет корабли
<- { "bvb" : { "place": 0 } } 
// игрок расставляет корабли на поле
// сервер должен проверить и допустить или не допустить расстановку
<- { "bvb" : { "place":  1, "ships" : [0,0,0,0,0,0,0]	}
}
// no bots yet, wait
-> { "bvb" : { "wait": 1 } }
// start battle, opponent bot info
-> { "bvb" : { "id": 321, "name": "dopinfo", "ships": [0,0,0,0] } }
0 0 0 0 0 0 0 0 0 1
0 4 0 0 0 0 0 0 0 0
0 4 0 0 0 0 3 0 0 2
0 4 0 0 0 0 3 0 0 2
0 4 0 0 0 0 3 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0

```
##### LOOP: GAME
```json
// сервер предлагает игроку 123 сделать ход
-> { "turn" : { "id" : 123 } }

// Бот сделал ход, выстрел пришелся в точку 
// (А1) - 0,0
// (Б8) - 1,7
// (А3) - 0,2
// (Г4) - 3,3
// АБВГДЕ....
// 0123456789
// первая цифра это номер ряда, вторая цифра номер колонки
<- { "turn" : { "shot": [y, x] } }

// результат выстрела, -1 - мимо, 1 - попал, 2 - убил
-> { "turn" : { "result": -1 } }

-> { "turn" : {"opponent": { "shot": [y, x], "result": 1 } } }
```
##### GOTO GAME
```
// 10 second timeout
// after timeout -> lose
```
## Battle end
```json
-> { "end": { "winner": 123, "opponent": [0,0,0] } }
```
# DB structure draft
```
User
	ID
	email
	pass
	name
	
Bot
	ID айди бота
	User
	AuthKey
	
Tournament
	ID
	Type - тип турнира
		SandBox - все боты имеют доступ в песочницу
		Tournament - после отборочного тура админ переносит в турнир
		Quality - игрок подает заявку в диапазоне дат, пишем запись в таблицу TourBot
	Name - Имя турнира
	RegStart
	RegUntil


TourBot
	Bot
	Tour
	State: Access, Deny
	RegisteredDate
	Played - сыграно боев
	Win - побед
	Lose - поражений
	Disconnect - дисконектов


```