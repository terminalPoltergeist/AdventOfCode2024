package day01

import (
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInputs() [2][1000]int {
    var idLists [2][1000]int

    inputs, err := os.ReadFile("inputs/day01.txt")
    if err != nil {
        println(err.Error())
    }

    lines := strings.Split(string(inputs), "\n")
    lines = lines[:len(lines)-1] // remove empty "line" due to final \n

    for idx, line := range lines {
      regex := regexp.MustCompile(`^(\d+?)\s+(\d+?)$`)
      matches := regex.FindStringSubmatch(line)
      
      first, _ := strconv.Atoi(matches[1])
      second, _ := strconv.Atoi(matches[2])

      idLists[0][idx] = first
      idLists[1][idx] = second
    }

    return idLists
}

func sortIds(ids [2][1000]int) [2][1000]int {
    // bubble sort implementation
    for i := range ids[0] {
        sorting := false
        for j := 0; j < (len(ids[0]) - i - 1); j++ {
            if ids[0][j] > ids[0][j+1] {
                tmp := ids[0][j+1]
                ids[0][j+1] = ids[0][j]
                ids[0][j] = tmp
                sorting = true
            }
        }
        
        if !sorting {
            break
        }
    }

    for i := range ids[1] {
        sorting := false
        for j := 0; j < (len(ids[1]) - i - 1); j++ {
            if ids[1][j] > ids[1][j+1] {
                tmp := ids[1][j+1]
                ids[1][j+1] = ids[1][j]
                ids[1][j] = tmp
                sorting = true
            }
        }
        
        if !sorting {
            break
        }
    }

    return ids
}

func sumPairs(ids [2][1000]int) int {
    sum := 0
    for idx := range ids[0] {
        distance := int(math.Abs(float64(ids[0][idx] - ids[1][idx])))
        sum += distance
    }
    return sum
}

func similarityScore(ids [2][1000]int) int {
    score := 0
    for _, id1 := range ids[0] {
        count := 0
        for _, id2 := range ids[1] {
            if id2 == id1 {
                count++
            }
        }
        score += id1 * count
    }

    return score
}

func Run() (any, error) {
    ids := parseInputs()
    sortedIds := sortIds(ids)
    return [2]int{sumPairs(sortedIds), similarityScore(sortedIds)}, nil
}
