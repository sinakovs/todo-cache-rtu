package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTodoFromFile() ([]todo, error) {
	var todos = []todo{}

	file, err := os.Open(dataPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := todoDecode(scanner.Text())
		todos = append(todos, line)
	}

	return todos, nil
}

func todoDecode(text string) todo {
	var todo todo

	line := strings.Split(string(text), "::")
	completed, err := strconv.ParseBool(line[2])
	if err != nil {
		fmt.Println("Error parsing str to bool", err)
		return todo
	}

	todo.ID = line[0]
	todo.Item = line[1]
	todo.Completed = completed

	return todo
}

func findTodoById(id string) (*todo, error) {
	for i, t := range Todo {
		if t.ID == id {
			return &Todo[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getShardsLen(m ShardMap) int {
	var length int
	for _, s := range m {
		length = length + len(s.data)
	}

	return length
}

func showAllShards(m ShardMap) []map[string]*todo {
	var all []map[string]*todo
	for _, s := range m {
		all = append(all, s.data)
	}

	return all
}
