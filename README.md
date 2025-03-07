﻿# Alfred

**Alfred** is a lightweight tool that automatically watches for changes in your Go files and triggers specified command in `alfred_config.json` under key value  `build_command`. Designed to streamline your development process, Alfred ensures your build stays updated with every modification.

## Features

- Monitors a directory for changes in `.go` files.
- Automatically runs `build_command` (specified build command) when a change is detected.
- Excludes irrelevant files like `.git` and `.exe`.
- Configurable build output via `alfred_config.json`.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your system.
- Clone this repository and navigate to the project folder.

### Installation

1. Clone the repository and build the executable:
   ```bash
   git clone https://github.com/yourusername/alfred.git
   cd alfred
   go build -o alfred
   ```

2. Add the generated `alfred` executable to your system's PATH:
   - **Linux/macOS**:
     ```bash
     export PATH=$PATH:/path/to/alfred
     ```
   - **Windows**:
     Add the directory containing `alfred.exe` to your environment variables' PATH.

---

### Usage

1. Navigate to the project directory where you want Alfred to watch for changes:
   ```bash
   cd /path/to/your/project
   ```

2. Run the Alfred executable:
   ```bash
   alfred
   ```

3. To customize the behavior, create an `alfred_config.json` file in your project directory:
   ```json
   {
       "build_name": "your_custom_build_name",
       "build_version": "1.0.0",
       "build_command": "go run main.go",
       "Watch_files": [".go"]
   }
   ```

4. When Alfred detects changes in `.go` files, it will automatically:
   - Build the project.
   - Use the custom executable name specified in your `alfred_config.json` (if present).
   - Default to the name `alfred` if no custom configuration is provided.

---

## File Structure

- `main.go` - Core logic to monitor files and trigger build commands.
- `logger.go` - Configures logging using Zap.
- `config_parser.go` - Parses the `alfred_config.json` configuration file.
- `alfred_config.json` - Configuration file for build settings.

## Configuration

You can customize the build command by editing the `alfred_config.json` file:
```json
   {
       "build_name": "your_custom_build_name",
       "build_version": "1.0.0",
       "build_command": "go run main.go",
       "Watch_files": [".go"]
   }
```

## Example

1. Modify `example.go` in the directory.
2. Alfred detects the change and runs:
   ```bash
   go build -o alfred
   ```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

