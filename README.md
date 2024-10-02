# cmdr
cmdr is a CLI tool that helps you analyze your command history.

## Features

cmdr provides insights into your command-line usage patterns with the following features:

1. **Top Commands**: By default, cmdr displays the top 5 most frequently used commands along with their usage count.

   Example:
   ```
   $ cmdr
   You have used git 150 times
   You have used ls 120 times
   You have used cd 100 times
   You have used npm 80 times
   You have used docker 50 times
   ```

2. **Customizable Results**: Use the `--top` flag followed by a number to specify how many top commands you want to see.

   Example:
   ```
   $ cmdr --top=10
   ```
   This will display the top 10 most frequently used commands.

3. **Mistake Analysis**: Use the `-M` or `--mistakes` flag to see commands that you've attempted to run but don't exist or have failed.

   Example:
   ```
   $ cmdr --mistakes --top=2
   You have used gti 5 times but it does not exist
   You have used sl 3 times but it does not exist
   ```
4. **Include or Exclude Arguments**: Use the `--args` flag to include arguments. They are excluded by default.

   Example:
   ```
   $ cmdr --args
   You have used git push 10 times
   You have used ls -l 5 times
   ```

## Usage

Here are some example commands to get you started with cmdr:

1. View top 5 most used commands (default behavior):
   ```
   $ cmdr
   ```

2. View top 10 most used commands:
   ```
   $ cmdr --top=10
   ```

3. View commands that don't exist or have failed:
   ```
   $ cmdr -M
   ```
   or
   ```
   $ cmdr --mistakes --top=2
   ```

## Installation

To install cmdr, you'll need to have Go installed on your system. Follow these steps:

### Linux

1. Clone the repository:
   ```
   git clone https://github.com/BradyDouthit/cmdr.git
   cd cmdr
   ```

2. Build the project:
   ```
   go build -o cmdr
   ```

3. Add the compiled binary to your system's PATH:
   ```
   export PATH=$PATH:/path/to/cmdr
   ```

4. You can now run `cmdr` from anywhere in your terminal.

### Windows

1. Clone the repository:
   ```
   git clone https://github.com/BradyDouthit/cmdr.git
   cd cmdr
   ```

2. Build the project:
   ```
   go build -o cmdr.exe
   ```

3. Add the directory containing `cmdr.exe` to your system's PATH

4. You can now run `cmdr` from anywhere in your command prompt or PowerShell.

Note: Make sure you have Go installed and properly configured on your system before following these steps.

## Dependencies

cmdr uses the following dependencies:
- [Lipgloss](https://github.com/charmbracelet/lipgloss)

## Currently Supported Terminal Emulators
- `zsh`
- `bash`

## Contributing

[Add contribution guidelines here]

## License

[Add license information here]

