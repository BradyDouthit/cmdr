# Trendify
Trendify is a CLI tool that helps you analyze your command history.

## Features

Trendify provides insights into your command-line usage patterns with the following features:

1. **Top Commands**: By default, Trendify displays the top 5 most frequently used commands along with their usage count.

   Example:
   ```
   $ trendify
   You have used git 150 times
   You have used ls 120 times
   You have used cd 100 times
   You have used npm 80 times
   You have used docker 50 times
   ```

2. **Customizable Results**: Use the `--top` flag followed by a number to specify how many top commands you want to see.

   Example:
   ```
   $ trendify --top=10
   ```
   This will display the top 10 most frequently used commands.

3. **Mistake Analysis**: Use the `-M` or `--mistakes` flag to see commands that you've attempted to run but don't exist or have failed.

   Example:
   ```
   $ trendify --mistakes
   You have used gti 5 times but it does not exist
   You have used sl 3 times but it does not exist
   ```

## Usage

Here are some example commands to get you started with Trendify:

1. View top 5 most used commands (default behavior):
   ```
   $ trendify
   ```

2. View top 10 most used commands:
   ```
   $ trendify --top=10
   ```

3. View commands that don't exist or have failed:
   ```
   $ trendify -M
   ```
   or
   ```
   $ trendify --mistakes --top=2
   ```

## Installation

To install Trendify, you'll need to have Go installed on your system. Follow these steps:

### Linux

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/trendify.git
   cd trendify
   ```

2. Build the project:
   ```
   go build -o trendify
   ```

3. Move the binary to a directory in your PATH:
   ```
   sudo mv trendify /usr/local/bin/
   ```

4. You can now run `trendify` from anywhere in your terminal.

### Windows

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/trendify.git
   cd trendify
   ```

2. Build the project:
   ```
   go build -o trendify.exe
   ```

3. Add the directory containing `trendify.exe` to your system's PATH

4. You can now run `trendify` from anywhere in your command prompt or PowerShell.

Note: Make sure you have Go installed and properly configured on your system before following these steps.

## Dependencies

Trendify uses the following dependencies:
- [Lipgloss](https://github.com/charmbracelet/lipgloss)

## Contributing

[Add contribution guidelines here]

## License

[Add license information here]

