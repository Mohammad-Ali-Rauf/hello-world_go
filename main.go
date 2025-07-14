package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	red   = "\033[31m"
	green = "\033[32m"
	blue  = "\033[34m"
	bold  = "\033[1m"
	reset = "\033[0m"
)

func main() {

	if runtime.GOOS == "windows" {
		fmt.Printf(red+bold+"Detected OS: %s — access denied.\n"+reset, runtime.GOOS)
		fmt.Println(red + bold + "I don't talk to babies 🍼" + reset)
		os.Exit(1)
	} else if runtime.GOOS == "linux" && strings.Contains(os.Getenv("WSL_DISTRO_NAME"), "Ubuntu") {
		fmt.Printf(red+bold+"Detected OS: %s — access denied.\n"+reset, runtime.GOOS)
		fmt.Println(red + bold + "Trying to play smart and decieve me? Good riddance. ❌" + reset)
		os.Exit(1)
	} else if runtime.GOOS == "darwin" {
		fmt.Printf(red+bold+"Detected OS: %s — access denied.\n"+reset, runtime.GOOS)
		fmt.Println(red + bold + "You paid $2000 to run a terminal with butterfly keys? Tragic." + reset)
		os.Exit(1)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var name string
	fmt.Print(blue + "Enter your name, reality check incoming: " + reset)
	fmt.Scanln(&name)

	fmt.Printf("\n%sAnalyzing %s's 'achievements'...%s\n", bold, name, reset)
	time.Sleep(1 * time.Second)

	insults := []string{
		"Still thinks HTML is a programming language ❌",
		"Couldn't hack a microwave ❌",
		"Still licks balls of those who say AI is a conscious being  ❌",
		"Still uses printf instead of logging ❌",
		"Acts tough on GitHub, cries during merge conflicts ❌",
		"Printed 'Hello World'? ✅ (barely counts)",
	}

	r.Shuffle(len(insults), func(i, j int) {
		insults[i], insults[j] = insults[j], insults[i]
	})

	for _, insult := range insults {
		fmt.Println(red + "• " + insult + reset)
		time.Sleep(700 * time.Millisecond)
	}

	fmt.Printf("\n%s%s... You just printed your insult.\n", bold, name)
	fmt.Println("Calm down, your only credential in life is Certified Momma's Boy. 💀" + reset)
}
