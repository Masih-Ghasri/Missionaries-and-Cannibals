# Missionaries and Cannibals Game

This repository contains a console-based implementation of the "Missionaries and Cannibals" puzzle written in Go.

## Game Description
The "Missionaries and Cannibals" puzzle is a classic river crossing game. The objective is to transport all missionaries and cannibals from the left side of the river to the right side using a boat, while adhering to specific rules to ensure everyone's safety.

### Rules
1. The boat can carry at most two people at a time.
2. Cannibals must never outnumber missionaries on either side of the river; otherwise, the cannibals will eat the missionaries.
3. The boat cannot cross the river without at least one person on board.

### Objective
Successfully move all three missionaries and three cannibals to the right side of the river without violating the rules.

## How to Play
The game is played in the terminal. You will be prompted to enter the number of missionaries and cannibals to move between the riverbanks.

### Example Gameplay
1. Move 1 missionary and 1 cannibal to the right.
2. Move 1 cannibal back to the left.
3. Move 2 missionaries to the right.
4. Move 1 missionary back to the left.
5. Move 1 missionary and 1 cannibal to the right.
6. Repeat similar moves until all are safely transported to the right side.

## Setup and Installation
### Prerequisites
- Go 1.18 or later

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/Masih-Ghasri/Missionaries-and-Cannibals.git
   git clone https://gitlab.com/Ma_salsa/missionaries-and-cannibals.git
   cd missionaries-cannibals-game
   ```
2. Run the game:
   ```bash
   go run main.go
   ```

## CI/CD Pipelines
This repository includes CI/CD pipelines for both **GitHub Actions** and **GitLab CI/CD** to ensure code quality and automated testing.

### GitHub Actions
The GitHub Actions workflow (`.github/workflows/go.yml`) ensures that the Go code is linted and tested on every push and pull request.

### GitLab CI/CD
The `.gitlab-ci.yml` file runs the same set of tests and lints on GitLab repositories.

## Contributing
We welcome contributions! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments
Thanks to the community for inspiring this project!
