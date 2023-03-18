# descit
🤩 ChatGPT describes your error 

## Install
```
go install github.com/yude/descit@latest
```

## Setup
* Open `~/.config/descit/.env` and write your OpenAI's token.
    * You can get it from [here](https://platform.openai.com/account/api-keys). (requires OpenAI account)
    ```
    TOKEN=your_token_goes_here_thx
    ```

## Use
```bash
descit <your_command>
# Example
# descit gcc main.c
```

## License
MIT License.