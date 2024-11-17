package main

import (
	"fmt"
	"os/exec"
	"sync"
	
)

type Animal struct {
	Name              string
	HighSpeed         int
	Size              string
	ClimbTree         bool
	RecognizeDiseases bool
}

func (a *Animal) printInfo() {
	
	fmt.Printf("Animal: %s\n", a.Name)
	fmt.Printf("HighSpeed: %d km/h\n", a.HighSpeed)
	fmt.Printf("Size: %s\n", a.Size)
	fmt.Printf("Can Climb Trees: %v\n", a.ClimbTree)
	fmt.Printf("Can Recognize Diseases: %v\n\n", a.RecognizeDiseases)
}

func animalRoutine(animal *Animal, wg *sync.WaitGroup, notifWG *sync.WaitGroup) {
	defer wg.Done()

	animal.printInfo()

	notifWG.Add(1)
	go func() {
		defer notifWG.Done()
		sendNotification(animal.Name, fmt.Sprintf("Information about %s has been loaded!", animal.Name))
	}()
}

func sendNotification(title, message string) {
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`New-BurntToastNotification -Text "%s", "%s"`, title, message))

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func main() {
	animals := []Animal{
		{"Lion", 80, "Large", true, true},
		{"Elephant", 40, "Huge", false, false},
		{"Monkey", 30, "Medium", true, true},
		{"Cheetah", 120, "Medium", false, false},
	}

	var wg, notifWG sync.WaitGroup

	for i := range animals {
		wg.Add(1)
		go animalRoutine(&animals[i], &wg, &notifWG)
	}

	
	wg.Wait()
	
	notifWG.Wait()

	
	fmt.Println("All animal information loaded and notifications sent.")
}
