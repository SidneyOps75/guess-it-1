# Guess-it-1

## Overview

This project is a Go-based application designed to perform specific calculations which are meant to guess the range between a set of two or more numbers and return the lower or upper bound range of the data set which the user inputs.

## Project Structure

 1. **script.sh:** A shell script for automating tasks such as building, testing, or running the project.
 2. **calculation/:** This directory contains functions of math formulas such as average, standard deviation , variance and range of values.
 3. **processInput/:** This directory contains a function which scans through the CLI(command line interface) and calculates the range of values for each data set .In my project I have specified a data set window of 4 values which is used to determine the upper or lower bound values for representing the range.
 4. **go.mod:** The Go module file that manages the project's dependencies.
 5. **main.go:** The main entry point of the application. This file contains the primary logic and function calls that drive the application.


## Getting Started
1. Prerequisites

   Ensure you have Go installed. If not, download and install it from [here](https://go.dev/doc/install). 

2. Installation

    Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/johopiyo/guess-it-1
    cd guess-it-1
    ```

3. Ensure Go modules are enabled and download dependencies:

    ```bash
    go mod tidy
    ```

## Running the Application:

1. To execute the main program, run the following command:

    ```bash
    go run main.go
    ```

2. Running Tests:

    To run the unit tests:

    ```bash
    go test ./calculation 
    ```
    and to test more
    ```bash
    go test ./processInput 
    ```

3. Using the Script:

You can also use the provided script.sh to automate common tasks:

```bash
./student/script.sh
```
## License

This project is licensed under the [MIT](LICENSE) License

## Author
This project has been authored by John Opiyo ([GitHub](https://github.com/SidneyOps75))
