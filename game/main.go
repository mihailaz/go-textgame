package main

import (
	"bufio"
	"fmt"
	"game/player"
	"game/world"
	"os"
	"strings"
)

/*
	код писать в этом файле
	наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/
	initGame()
	command := prompt("хм, стоит осмотреться")
	for command != "выход" {
		answer := handleCommand(command)
		command = prompt(answer)
	}
}

var mainPlayer player.Player

func initGame() {
	/*
		эта функция инициализирует игровой мир - все комнаты
		если что-то было - оно корректно перезатирается
	*/
	startLocation := world.Create()
	mainPlayer = player.Player{
		CurrentLocation: startLocation,
		Actions:         player.Actions,
		TaskList:        player.Create(),
	}
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return mainPlayer.Do(command)
}

func prompt(label string) string {
	var result string
	r := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fprint(os.Stderr, label+"\n")
		if err != nil {
			panic(err)
		}
		result, _ = r.ReadString('\n')
		if result != "" {
			break
		}
	}
	return strings.TrimSpace(result)
}
