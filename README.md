# Ask Program

## Overview
The `ask` program allows you to query different large language models (LLMs) directly from the command line. By default, it uses the Anthropic LLM, but you can specify GPT if desired.

## Installation
1. Copy the script to `/usr/local/bin`:
    ```sh
    sudo cp ask-<os>-<arch> /usr/local/bin/
    sudo chmod +x /usr/local/bin/ask
    ```
If `/usr/local/bin` is not in your path, place it somewhere else!

2. Or install with `go install github.com/andyantrim/ask`

## Usage
The `ask` program can be used with the following syntax:
```sh
ask [-llm=LLM] "your question here"
```
- `-llm`: Optional parameter to specify the LLM to use. Defaults to `claude`. Accepted values are `gpt` and `claude`.

### Examples
1. Using the default Anthropic LLM:
    ```sh
    ask "What is the capital of France?"
    ```

2. Using GPT-4o-mini LLM:
    ```sh
    ask -llm=gpt "What is the capital of France?"
    ```

Note: Ensure you have the necessary permissions and environment settings to run the script.

## Contribution
Feel free to contribute to this program by submitting issues or pull requests to the repository.

## License
This project is licensed under the MIT License.
