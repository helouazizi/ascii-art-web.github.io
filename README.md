# ascii-art-web

This project is designed to accept a string and output formatted ASCII art into a specified text file. The user can customize the output using flags and banners.
`standard.txt` The primary functionality ensures that the output file is generated based on the specified command-line options.

## Features

- Output string as ASCII art into a specified text file.
- Supports customizable options and banners.
- Accepts a single string argument as an alternative mode of operation.
- Validates command-line inputs to ensure correct usage.

## Installation

Ensure you have Go installed on your machine. To clone the repository and navigate to the project directory, run:
`git clone <repository-url>`
`cd <project-directory>`

## Running the Program

After setting up, you can execute the program with the appropriate command as described in the Usage section.
` go run . [resultfilename] [STRING] [BANNER] `

## Note

The program checks for non-printable characters and will terminate with an error message if any are found.
It will also inform you if the incorrect number of command-line arguments is provided.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.
