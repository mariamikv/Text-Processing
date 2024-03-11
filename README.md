
# Text Processing with Strategy Design Pattern in Go

This project demonstrates the implementation of the strategy design pattern in Go for various text processing tasks. It allows you to choose and apply different algorithms interchangeably through a common interface.

## Features

- Implements the TextProcessor interface to define a standard way of handling text processing tasks.
- Provides concrete strategy implementations for:
    - Sorting (e.g., Bubble Sort, Insertion Sort)
    - Regular expression filtering
    - (You can add more strategies as needed)
- Utilizes a TextProcessingContext to hold the chosen strategy and facilitates processing text based on the chosen strategy.


## Getting Started:

1. Clone this repository:
```bash
git clone https://github.com/mariamikv/Text-Processing.git
```

2. Install dependencies (if any):
```bash
go mod download
```

3. Run the Program:
```bash
go run main.go
```

## Acknowledgements

 - [Strategy Desing Pattern](https://refactoring.guru/design-patterns/strategy)

