package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Command struct {
	Keys []string
	Args []string
	Desc string
	Fn   func(args []string) error
}

type Addiction struct {
	Name       string
	StreakGoal int // days
	StartedAt  time.Time
}

var Commands []Command = []Command{}
var saveFile = "addictions.json"

func init() {
	Commands = []Command{
		{
			Keys: []string{"help", "h"},
			Args: []string{},
			Desc: "Shows all available commands",
			Fn:   Help,
		},
		{
			Keys: []string{"list", "ls"},
			Args: []string{},
			Desc: "Lists all addictions and their streaks",
			Fn:   List,
		},
		{
			Keys: []string{"add", "a"},
			Args: []string{},
			Desc: "Adds a new addiction",
			Fn:   Add,
		},
		{
			Keys: []string{"remove", "rm"},
			Args: []string{},
			Desc: "Removes an addiction",
			Fn:   Remove,
		},
		{
			Keys: []string{"reset", "rs"},
			Args: []string{},
			Desc: "Resets an addiction streak",
			Fn:   Reset,
		},
		{
			Keys: []string{"live", "lv"},
			Args: []string{},
			Desc: "Displays all addictions and their streaks in real-time with an animated progress bar",
			Fn:   Live,
		},
	}
}

func SaveAddiction(addiction Addiction) error {
	// save addiction to file (json array of addictions)
	// if file does not exist, create it
	// if file exists, append addiction to it

	file, err := os.OpenFile(saveFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	var addictions []Addiction

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() > 0 {
		err = json.NewDecoder(file).Decode(&addictions)
		if err != nil {
			return err
		}
	}

	for _, a := range addictions {
		if strings.ToLower(addiction.Name) == strings.ToLower(a.Name) {
			return fmt.Errorf("Addition with name '%s' already exists", addiction.Name)
		}
	}

	addictions = append(addictions, addiction)

	file.Seek(0, 0)
	file.Truncate(0)
	err = json.NewEncoder(file).Encode(addictions)
	if err != nil {
		return err
	}

	return nil
}

func UpdateAddiction(addiction Addiction) error {
	// update addiction in file

	file, err := os.OpenFile(saveFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	var addictions []Addiction

	err = json.NewDecoder(file).Decode(&addictions)
	if err != nil {
		return err
	}

	for i, a := range addictions {
		if strings.ToLower(addiction.Name) == strings.ToLower(a.Name) {
			addictions[i] = addiction
			break
		}
	}

	file.Seek(0, 0)
	file.Truncate(0)
	err = json.NewEncoder(file).Encode(addictions)
	if err != nil {
		return err
	}

	return nil
}

func ResetAddiction(name string) error {
	// reset addiction streak

	file, err := os.OpenFile(saveFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	var addictions []Addiction

	err = json.NewDecoder(file).Decode(&addictions)
	if err != nil {
		return err
	}

	for i, addiction := range addictions {
		if strings.ToLower(addiction.Name) == strings.ToLower(name) {
			addiction.StartedAt = time.Now()
			addictions[i] = addiction
			break
		}
	}

	file.Seek(0, 0)
	file.Truncate(0)
	err = json.NewEncoder(file).Encode(addictions)
	if err != nil {
		return err
	}

	return nil
}

func LoadAddictions() ([]Addiction, error) {
	// load all addictions from file
	// if file does not exist, return empty array

	file, err := os.OpenFile(saveFile, os.O_RDONLY, 0755)
	if err != nil {
		return []Addiction{}, nil
	}
	defer file.Close()

	var addictions []Addiction

	err = json.NewDecoder(file).Decode(&addictions)
	if err != nil {
		return nil, err
	}

	return addictions, nil
}

func RemoveAddiction(name string) error {
	// remove addiction from file

	file, err := os.OpenFile(saveFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	var addictions []Addiction

	err = json.NewDecoder(file).Decode(&addictions)
	if err != nil {
		return err
	}

	for i, addiction := range addictions {
		if strings.ToLower(addiction.Name) == strings.ToLower(name) {
			addictions = append(addictions[:i], addictions[i+1:]...)
			break
		}
	}

	file.Seek(0, 0)
	file.Truncate(0)
	err = json.NewEncoder(file).Encode(addictions)
	if err != nil {
		return err
	}

	return nil
}
