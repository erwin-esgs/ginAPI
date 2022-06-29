package Concurency

import (
	"fmt"
	//"net/http"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var wg sync.WaitGroup

	var arr []interface{}

	var getResult = func(arg1 ...int) {
		i := 0
		if len(arg1) > 0 {
			i = arg1[0]
		}
		// Calling Sleep method
		randomNumber := (rand.Intn(4-1) + 1)
		time.Sleep(time.Duration(randomNumber) * time.Second)

		// Printed after sleep is over
		//return strconv.Itoa(i) + " I sleep for " + strconv.Itoa(randomNumber) + " second"
		text := strconv.Itoa(i) + " I sleep for " + strconv.Itoa(randomNumber) + " second"
		fmt.Println(text)
		arr = append(arr, text)
		defer wg.Done()
	}

	wgNumber := 4
	wg.Add(wgNumber)

	for i := 0; i < wgNumber; i++ {
		go getResult(i)
	}
	wg.Wait()
	fmt.Println(c.Request.RequestURI)
	// result := http.Get()
	c.JSON(200, gin.H{
		"message": arr,
	})
}
