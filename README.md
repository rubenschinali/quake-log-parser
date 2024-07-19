# Quake 3 Arena Logger

Welcome to the Quake 3 Arena Logger! This project aims to parse and log game data from Quake 3 Arena logs for analysis and insights.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [JSON Output](#json-output)
- [Screen Output](#screen-output)
- [Contributing](#contributing)
- [License](#license)

## Introduction
The Quake 3 Arena Logger is designed to help gamers and developers analyze gameplay by parsing logs generated during Quake 3 Arena matches. 

## Features
- Parse and log game events from Quake 3 Arena log files
- Interactive CLI for specifying log files or folders
- Outputs parsed data in JSON format or printed to screen
- Supports concurrent processing of various log files 
- Customizable data extraction and formatting

## Installation

### Prerequisites
- [Golang](https://golang.org/doc/install) (version 1.22 or higher)

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/rubenschinali/quake-log-parser.git
   cd quake-log-parser
   ```

2. Build the project:
   ```bash
   go build -o quake-log-parser ./cmd/quake-log-parser
   ```

## Usage

### Command-Line Interface
The CLI program allows you to interactively input file names or specify a folder with log files.

#### Example Command
```bash
./quake-log-parser/qgames.log
```

## JSON Output
Saves to disk in JSON format the summary of the game Players, Kills, and Kill Methods.

```json
{
  "output_format": "json",
  "log_level": "info",
  "filters": {
    "player": "player_name",
    "event_type": "kill"
  }
}
```

## Screen Output
Outputs to the screen the summary of the game Players, Kills, and Kill Methods.

Example: 
```
    Game ID: GAMEID
    Game Details:
    Total Kills: 10 
    Players:
        Player1
        Player2
        Player3
        Player4
    Kills:
        Player1: 4 kills
        Player2: 2 kills
        Player3: 0 kills
        Player4: 2 kills
    Kill Methods:
        MOD_ROCKET: 3 kills
        MOD_RAILGUN: 2 kills
        MOD_SHOTGUN: 3 kills
        MOD_ROCKET_SPLASH: 2 kills
```

## Contributing
Contributions are welcome. If you have ideas for new features or improvements, please open an issue or submit a pull request.

### Steps to Contribute
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a pull request

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to replace placeholders like `yourusername` with your actual GitHub username and customize the content according to your project's details. Let me know if you need any further adjustments!