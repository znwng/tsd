# Tmux Session Deployer

Define and deploy tmux sessions from a single `.hcl` file.

## Installation

- Clone the repository:
    ```sh
    git clone https://github.com/znwng/tsd
    cd tsd
    ```

- Build binary
    ```sh
    go build -o tsd ./cmd/tsd
    ```

## Usage

Demo `.hcl` file:
```hcl
session "demo" {
    window "window1_name" {
        run = [
            "echo runs_commands",
            "echo again"
        ]
    }

    window "window2_name" {
        run = [
            "echo runs_commands_in_window2"
        ]
    }

    window "window3_name" {
        send = [
            "sends_commands"
        ]
    }

    focus = "window1_name"
}
```

- `session`: Defines the session name
- `window`: Used to create windows with custom names
- `run`: Array of commands to run. Each line must contain one command enclosed in double quotes.
- `send`: Just puts in the command without actually running it. Useful if you want the command to be ready when you open the session.
- `focus`: Specifies the window that appears when you attach to the session.

> For more info run `tsd --help`

Current version only consists of features I needed.
