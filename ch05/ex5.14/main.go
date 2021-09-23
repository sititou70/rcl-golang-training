package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
)

// types
type Station struct {
	gcd     string
	name    string
	address string
}
type StringSet map[string]bool
type Join map[string]StringSet

// main
func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: ./%s ORIGIN_STATION_NAME DESTINATION_STATION_NAME\n", os.Args[0])
		os.Exit(1)
	}

	// load data
	print("loading data...")
	//// load station
	stations := map[string]Station{}
	cdToGcd := map[string]string{}
	origPattern := os.Args[1]
	destPattern := os.Args[2]
	origCandidate := map[string]Station{}
	destCandidate := map[string]Station{}
	f, err := os.Open("assets/station20210312free.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		s := strings.Split(string(line), ",")
		cd := s[0]
		station := Station{s[1], s[2], s[3]}

		cdToGcd[cd] = station.gcd
		if cd == station.gcd {
			stations[station.gcd] = station
			if strings.Contains(station.name, origPattern) {
				origCandidate[station.gcd] = station
			}
			if strings.Contains(station.name, destPattern) {
				destCandidate[station.gcd] = station
			}
		}
	}
	f.Close()
	//// load join
	join := Join{}
	f, err = os.Open("assets/join20210312.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r = bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		s := strings.Split(string(line), ",")
		cd1 := s[0]
		cd2 := s[1]
		gcd1, ok := cdToGcd[cd1]
		if !ok {
			continue
		}
		gcd2, ok := cdToGcd[cd2]
		if !ok {
			continue
		}

		if join[gcd1] == nil {
			join[gcd1] = make(StringSet)
		}
		join[gcd1][gcd2] = true
		if join[gcd2] == nil {
			join[gcd2] = make(StringSet)
		}
		join[gcd2][gcd1] = true
	}
	f.Close()
	println("DONE")

	// select station
	if len(origCandidate) == 0 {
		panic("origin station not found!")
	}
	if len(destCandidate) == 0 {
		panic("destination station not found!")
	}
	origStation := promptStation(origCandidate, "Select Origin Station")
	destStation := promptStation(destCandidate, "Select Destination Station")

	// search path
	bfsFunc := func(path Path) []Path {
		gcd := path[len(path)-1]
		if gcd == destStation.gcd {
			println("path found:")
			for _, p := range path {
				fmt.Printf("%v ", stations[p].name)
			}
			println()
			os.Exit(0)
		}

		paths := []Path{}
		for joinGcd := range join[gcd] {
			newPath := make(Path, len(path))
			copy(newPath, path)
			paths = append(paths, append(newPath, joinGcd))
		}
		return paths
	}
	breadthFirst(bfsFunc, []Path{{origStation.gcd}})
	println("path not found...")
}

// utils
type Path []string

func breadthFirst(f func(path Path) []Path, worklist []Path) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			key := item[len(item)-1]
			if !seen[key] {
				seen[key] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func promptStation(candidate map[string]Station, label string) Station {
	s := []Station{}
	for _, station := range candidate {
		s = append(s, station)
	}
	sort.Slice(s, func(i, j int) bool {
		return len(s[i].name) < len(s[j].name)
	})
	prompt := promptui.Select{
		Label: label,
		Items: s,
	}
	i, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return s[i]
}
