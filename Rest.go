package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type voting struct {
	voted string   `json:"voted"`
	voter []string `json:"voter"`
}

var finishedVoting = false

var votings = []voting{}

func main() {
	router := gin.Default()
	router.GET("/result", getResult)
	router.GET("/hello", hello_world)
	router.POST("/vote", appendVote)
	router.POST("/reset", reset)

	router.Run("localhost:8080")
}

func getResult(c *gin.Context) {
	if finishedVoting {
		vot := fmt.Sprintf("%+v\n", votings)
		print(vot)
		c.String(http.StatusOK, vot)

	}
}

func hello_world(c *gin.Context) {
	c.Status(http.StatusOK)
}

func reset(c *gin.Context) {
	finishedVoting = false
	votings = []voting{}
	c.Status(http.StatusOK)
}

func appendVote(c *gin.Context) {
	already_voted := false
	vter := c.Query("voter")
	vted := c.Query("voted")

	//if already voted
	for _, v := range votings {
		for _, w := range v.voter {
			if w == vter {
				already_voted = true
				println("Already Voted")
			}
		}
	}
	if !already_voted {
		present := isVotedPresent(vted)
		if present > 0 {
			v := votings[present-1].voter
			v = append(v, vter)

			vot := voting{
				voted: vted,
				voter: v,
			}
			remove(present - 1)
			votings = append(votings, vot)
		} else {
			v := []string{vter}
			vot := voting{
				voted: vted,
				voter: v,
			}
			votings = append(votings, vot)
		}
	}
	checkFinish()
	for _, v := range votings {
		println(v.voted+": ", v.voter)
	}

	c.Status(http.StatusOK)
}

func isVotedPresent(votedKey string) int {
	for i, v := range votings {
		i++
		if v.voted == votedKey {
			return i
		}
	}
	return -1
}

func remove(s int) {
	v := append(votings[:s], votings[s+1:]...)
	votings = v
}

func checkFinish() {
	count := 0
	for _, v := range votings {
		count += len(v.voter)
	}
	if count == 8 {
		println("Finished")
		finishedVoting = true
	}
}

func countWords(s string, sep string) int {
	words := strings.Split(s, sep)
	return len(words)
}
