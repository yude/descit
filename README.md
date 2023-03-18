# descit
🤩 ChatGPT describes your error 

## Install & Update
* Please make sure Go's bin folder is in your `$PATH`.
```
go install github.com/yude/descit@HEAD
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