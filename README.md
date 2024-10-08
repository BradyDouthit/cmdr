# cmdr
`cmdr` is a CLI tool with a goal of helping develpers understand their CLI usage in order to improve efficiency.

> [!NOTE]  
>  I hope to make this a tool developers like myself can use regularly to improve CLI usage a little every day. After all, we use the terminal so much we should be good at it. That said, I don't have much time for projects like this so it is pretty basic as of now but contributions are welcome (no strict rules or anything like that just make a PR)!

## Examples

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

3. **Most Commonly Invalid Commands**: Use the `-I` or `--invalid` flag to see commands that you've attempted to run but don't exist or have failed.

   Example:
   ```
   $ cmdr --invalid --top=2
   You have used gti 5 times but it does not exist
   You have used sl 3 times but it does not exist
   ```
4. **Valid Commands**: Use the `-V` or `--valid` flag to see commands that you've attempted to run but don't exist or have failed.

   Example:
   ```
   $ cmdr --valid
   You ran go 143 times
   You ran ls 124 times
   You ran clear 105 times
   You ran cd 96 times
   You ran git 67 times
   ```
5. **Include or Exclude Arguments**: Use the `--args` flag to include arguments. They are excluded by default.

   Example:
   ```
   $ cmdr --args
   You have used git push 10 times
   You have used ls -l 5 times
   ```
6. **Combine Flags**: You can combine flags to get the desired output.

   Example:
   ```
   $ cmdr --args --valid --top=3 
   You ran go run . 45 times
   You ran cd .. 21 times
   You ran npm run dev 14 times
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
Feel free to contribute to cmdr by submitting a pull request or opening an issue. Your contributions and ideas are welcome!
