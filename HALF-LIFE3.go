/*package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	soundFile := "/Users/daniel/Documents/HalfLifeIntro.mp3"

	cmd := exec.Command("afplay", soundFile)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	title := "HALF-LIFE 3λ: ЯOЯЯIM || MIRROR"
	text := "Hello, Gordon."
	text2 := "You may not know where you are,"
	text3 := "But I assure you, Mister Freeman,"
	text4 := " you're home."

	text5 := "You've been looking for me your whole life, Gordon."
	text6 := "I created it all."
	text7 := "They call me"
	text8 := "..."
	text9 := " the G-man. I take many names, but I prefer Euler. Call me ∑, as I am the 'Summation'.\nLike the definition of ∑, I am the sum of all."
	text10 := "Let me show you something."
	text11 := `
	.                                            .
	*   .                  .              .        .   *          .
 .         .                     .       .           .      .        .
	   o                             .                   .
		.              .                  .           .
		 0     .
				.          .                 ,                ,    ,
.          \          .                         .
	 .      \   ,
  .          o     .                 .                   .            .
	.         \                 ,             .                .
			  #\##\#      .                              .        .
			#  #O##\###                .                        .
  .        #*#  #\##\###                       .                     ,
	   .   ##*#  #\##\##               .                     .
	 .      ##*#  #o##\#         .                             ,       .
		 .     *#  #\#     .                    .             .          ,
					 \          .                         .
____^/\___^--____/\____O______________/\/\---/\___________---______________
  /\^   ^  ^    ^                  ^^ ^  '\ ^          ^       ---
		--           -            --  -      -         ---  __       ^
  --  __                      ___--  ^  ^                         --  __`
	// text12 := "Recognize it? Yes, you must be very familiar with Xen by now."
	// text13 :=


	fmt.Println("Introducing...")
	time.Sleep(3 * time.Second)
	fmt.Println("From the minds of VALVE || EVLAV")
	fmt.Println()
	fmt.Println()
	time.Sleep(3 * time.Second)
	typewritereffect(title)
	time.Sleep(2 * time.Second)

	clearScreen()
	time.Sleep(2 * time.Second)
	typewritereffect(text)
	fmt.Println()
	time.Sleep(1 * time.Second)
	typewritereffect(text2)
	fmt.Println()
	time.Sleep(1 * time.Second)
	typewritereffect(text3)
	time.Sleep(3 * time.Second)
	typewritereffectSlow(text4)
	time.Sleep(1 / 2 * time.Second)
	clearScreen()

	time.Sleep(1 * time.Second)
	typewritereffect(text5)
	fmt.Println()
	time.Sleep(1 * time.Second)
	typewritereffect(text6)
	fmt.Println()
	typewritereffect(text7)
	typewritereffectMedium(text8)
	time.Sleep(1 + (1/2)*time.Second)
	typewritereffect(text9)
	time.Sleep(2 * time.Second)
	typewritereffect(text10)
	time.Sleep(3 * time.Second)
	clearScreen()
	fmt.Println()
	typewritereffectSuperFast(text11)
	fmt.Println()
	fmt.Println()


}

func typewritereffect(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(100 * time.Millisecond) // Adjust the duration as needed
	}
}

func typewritereffectSlow(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(500 * time.Millisecond) // Adjust the duration as needed
	}
}

func typewritereffectMedium(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(300 * time.Millisecond) // Adjust the duration as needed
	}
}
func typewritereffectSuperFast(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(5 * time.Millisecond) // Adjust the duration as needed
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported OS, cannot clear screen")
	}
}
*/

/*package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	soundFile := "/Users/daniel/Documents/HalfLifeIntro.mp3"

	// Start playing the music in the background
	playMusic(soundFile)

	// Your text output
	displayText()
}

func playMusic(soundFile string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", soundFile)
	case "darwin":
		cmd = exec.Command("afplay", soundFile)
	default:
		fmt.Println("Unsupported OS for playing music")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
	}

}
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var musicCmd *exec.Cmd

func main() {

	stopMusic()

	// Your text output
	displayText()

}

func playMusic(soundFile string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", soundFile)
	case "darwin":
		cmd = exec.Command("afplay", soundFile)
	default:
		fmt.Println("Unsupported OS for playing music")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	musicCmd = cmd // Store the music command for stopping later
}

func stopMusic() {
	switch runtime.GOOS {
	case "windows":
		exec.Command("taskkill", "/f", "/im", "afplay.exe").Run() // Replace with appropriate Windows command
	case "darwin":
		exec.Command("pkill", "afplay").Run()
	default:
		fmt.Println("Unsupported OS for stopping music")
	}
}

func displayText() {
	soundFile := "/Users/daniel/Documents/HalfLifeIntro.mp3"
	soundFile2 := "/Users/daniel/Documents/HL-VagueVoices.mp3"
	guitar := "/Users/daniel/Documents/DIABOLICALGUITAR.mp3"

	title := "\033[33mHALF-LIFE 3λ: MIRROR || ЯOЯЯIM"
	text := "\033[34mHello, Gordon."
	text2 := "You may not know where you are,"
	text3 := "But I assure you, Mister Freeman,"
	text4 := " you're home."

	text5 := "You've been looking for me your whole life, Gordon."
	text6 := "\033[31mI created it all."
	text7 := "\033[34mThey call me"
	text8 := "..."
	text9 := " the G-man. I take many names, but I prefer Euler. Call me ∑, as I am the 'Summation'.\nLike the definition of ∑, I am the sum of all.\033[34m"
	text10 := "Let me show you something.\033[34m"
	text11 := `
	.                                            .
	*   .                  .              .        .   *          .
 .         .                     .       .           .      .        .
	   o                             .                   .
		.              .                  .           .
		 0     .
				.          .                 ,                    
.                    .                         .
	 .         ,
  .          o     .                 .                   .            .
	.                          ,             .                .
			      .                              .        .
			              .                o        .
  .                           .                     ,
	   .                .                     .
	 .            .                             ,       .
		 .       .                    .             .          ,
					           .                         .
____^/\___^--____/\____O______________/\/\-__,-/\___________---__,--/^\___
  /\^   ^  ^    ^                  ^^ ^  '\ ^          ^       ---
		--           -            --  -      -         ---  __       ^
  --  __                      ___--  ^  ^                         --  __`
	text12 := "\033[34mRecognize it? Yes, you must be very familiar with Xen by now."
	text13 := "It was left in ruins after your visit, Gordon."
	text14 := "Which might I add was not a very enjoyable visit for you either, was it?"
	text15 := "You are a bad person, Mister Freeman."
	text16 := " But I believe every bad man is worthy of the universe."
	text17 := "A previous hire was unable to complete the tasks provided. I trust in you to do your job."
	text18 := "This position may not be what you wanted, but I have forced you in the past,\nand you have done what you needed."
	text19 := "Mister Freeman"
	text20 := "..."
	text21 := " are you ready?\033[34m"

	playMusic(soundFile)
	fmt.Println("Introducing...")
	time.Sleep(3 * time.Second)
	fmt.Println("From the minds of VALVE || EVLAV")
	fmt.Println()
	fmt.Println()
	time.Sleep(3 * time.Second)
	typewritereffect(title)
	time.Sleep(4 * time.Second)
	stopMusic()

	clearScreen()
	playMusic(soundFile2)
	time.Sleep(5 * time.Second)
	typewritereffect(text)
	fmt.Println()
	time.Sleep(4 * time.Second)
	typewritereffect(text2)
	fmt.Println()
	time.Sleep(1 * time.Second)
	typewritereffect(text3)
	time.Sleep(3 * time.Second)
	typewritereffectSlow(text4)
	time.Sleep(4 * time.Second)
	clearScreen()

	time.Sleep(1 * time.Second)
	typewritereffect(text5)
	fmt.Println()
	time.Sleep(1 * time.Second)
	typewritereffect(text6)
	fmt.Println()
	typewritereffect(text7)
	typewritereffectMedium(text8)
	time.Sleep(1 + (1/2)*time.Second)
	typewritereffect(text9)
	time.Sleep(2 * time.Second)
	fmt.Println()
	typewritereffect(text10)
	time.Sleep(3 * time.Second)
	clearScreen()
	fmt.Println()
	typewritereffectSuperFast(text11)
	fmt.Println()
	fmt.Println()
	typewritereffect(text12)
	time.Sleep(3 * time.Second)
	fmt.Println()
	typewritereffect(text13)
	fmt.Println()
	time.Sleep(2 * time.Second)
	typewritereffect(text14)
	time.Sleep(3 * time.Second)
	stopMusic()
	fmt.Println()
	clearScreen()
	typewritereffect(text15)
	time.Sleep(3 * time.Second)
	playMusic(guitar)
	typewritereffect(text16)
	fmt.Println()
	time.Sleep(4 * time.Second)
	typewritereffect(text17)
	time.Sleep(3 * time.Second)
	fmt.Println()
	typewritereffect(text18)
	time.Sleep(3 * time.Second)
	clearScreen()
	fmt.Println()
	typewritereffect(text19)
	typewritereffectMedium(text20)
	time.Sleep(2 * time.Second)
	typewritereffect(text21)
	fmt.Println()
	fmt.Println()
	credits()

}

func credits() {
	time.Sleep(2 * time.Second)
	fmt.Println("CODE BY: DANIEL C.")
	time.Sleep(3 * time.Second)
	fmt.Println("STORY BY: DANIEL C.")
	time.Sleep(2 * time.Second)
	fmt.Println("LEAD DEVELOPER: YOU GUESSED IT, DANIEL C.")
	time.Sleep(2 * time.Second)
	fmt.Println("PROGRAMMED IN GOLANG")
	time.Sleep(3 * time.Second)
	fmt.Println("NOT AFFILIATED WITH VALVE IN ANY WAY")
	time.Sleep(2 * time.Second)
	fmt.Println("\033[33mHALF-LIFE 3λ: MIRROR || ЯOЯЯIM")
}

func typewritereffect(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(100 * time.Millisecond)
	}
}

func typewritereffectSlow(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(500 * time.Millisecond)
	}
}

func typewritereffectMedium(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(300 * time.Millisecond)
	}
}
func typewritereffectSuperFast(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(2 * time.Millisecond)
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported OS, cannot clear screen")
	}
}
