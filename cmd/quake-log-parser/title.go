package main

import "fmt"

func startCLI() {
	// Initial message
	fmt.Println(`
 ___    _   _      _      _  __  _____     ___   ___   ___        _      ____    _____   _   _      _    
 / _ \  | | | |    / \    | |/ / | ____|   |_ _| |_ _| |_ _|      / \    |  _ \  | ____| | \ | |    / \   
| | | | | | | |   / _ \   | ' /  |  _|      | |   | |   | |      / _ \   | |_) | |  _|   |  \| |   / _ \  
| |_| | | |_| |  / ___ \  | . \  | |___     | |   | |   | |     / ___ \  |  _ <  | |___  | |\  |  / ___ \ 
 \__\_\  \___/  /_/   \_\ |_|\_\ |_____|   |___| |___| |___|   /_/   \_\ |_| \_\ |_____| |_| \_| /_/   \_\
 _        ___     ____     ____       _      ____    ____    _____   ____                                 
| |      / _ \   / ___|   |  _ \     / \    |  _ \  / ___|  | ____| |  _ \                                
| |     | | | | | |  _    | |_) |   / _ \   | |_) | \___ \  |  _|   | |_) |                               
| |___  | |_| | | |_| |   |  __/   / ___ \  |  _ <   ___) | | |___  |  _ <                                
|_____|  \___/   \____|   |_|     /_/   \_\ |_| \_\ |____/  |_____| |_| \_\     
	`)

	fmt.Println("Welcome to the Quake III Arena Log Parser CLI")
}
