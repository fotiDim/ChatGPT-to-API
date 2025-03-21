# ChatGPT-to-API
Create a fake API using ChatGPT's website

> ## IMPORTANT
> You will not get free support for this repository. This was made for my own personal use and documentation will continue to be limited as I don't really need documentation. You will find more detailed documentation in the Chinese docs by a contributor.

**API endpoint: http://127.0.0.1:8080/v1/chat/completions.**

[中文文档（Chinese Docs）](https://github.com/xqdoo00o/ChatGPT-to-API/blob/master/README_ZH.md)
## Setup
    
### Authentication

After 2024-04-02, authentication is optional because there is no need for authentication for gpt-3.5.

You have 2 ways to set credentials:

#### from file

Access token and PUID(only for PLUS account) retrieval has been automated by [OpenAIAuth](https://github.com/xqdoo00o/OpenAIAuth/) with account email & password.

`accounts.txt` - A list of accounts separated by new line 

Format:
```
email:password
...
```

#### from environment

Alternatively, set `OPENAI_ACCOUNTS="user1:pass1;user2:pass2;user3:pass3"` environment variable.`

---

All authenticated access tokens and PUID will be stored in `access_tokens.json` and auto-renewed after 24 hours.

Caution! please use unblocked IP for authentication. First login to `https://chatgpt.com/` to check IP availability if you can.

---

### HAR file pool

Currently logged in account, using the GPT-4 model and most GPT-3.5 models, you need to configure a HAR file (file with .har suffix) to complete captcha verification.

  1. Use a chromium-based browser (Chrome, Edge) to open the browser developer tools (F12), switch to the Network tab, and check the **preserve log** option.

  2. Log in to `https://chatgpt.com/`, create a new chat and select the GPT-4 model, enter any text, switch to the GPT-3.5 model, and enter any text.

  3. Click the Export HAR button under the Network tab to export the file `chatgpt.com.har` and place it in the `harPool` folder of the same level as this program.

### API Authentication (Optional)

#### from file

Custom API keys for this fake API, just like OpenAI api

`api_keys.txt` - A list of API keys separated by new line

Format:
```
sk-123456
88888888
...
```

#### from environment

Set the `API_KEYS="sk-123456,88888888"` environment variable.

## Getting set up
```  
git clone https://github.com/xqdoo00o/ChatGPT-to-API
cd ChatGPT-to-API
go build
./freechatgpt
```

### Environment variables
  - `SERVER_HOST` - Set to 127.0.0.1 by default
  - `SERVER_PORT` - Set to 8080 by default
  - `ENABLE_HISTORY` - Set to false by default
  - `OPENAI_ACCOUNTS` - The accounts you use for OpenAI
  - `API_KEYS` - The API keys for clients to authenticate against your own API
### Files (Optional)
  - `proxies.txt` - A list of proxies separated by new line

    ```
    http://127.0.0.1:8888
    ...
    ```
  - `access_tokens.json` - A JSON array of access tokens for cycling (Alternatively, send a PATCH request to the [correct endpoint](https://github.com/xqdoo00o/ChatGPT-to-API/blob/master/docs/admin.md))
    ```
    {"account1":{token:"access_token1", puid:"puid1"}, "account2":{token:"access_token2", puid:"puid2"}...}
    ```
  - `cookies.json` - A JSON that stores login cookies. If the OpenAI account is logged in with a third party such as Google, you can add a third-party account (also suitable for first-party account) and any password in `accounts.txt`. Modify this file as follows to login account.
    ```
    {
        "third party username": [
            {
                "Name": "__Secure-next-auth.session-token",
                "Value": "After logging into a third-party account on browser，the value of __Secure-next-auth.session-token in cookies",
                "Path": "/",
                "Domain": "",
                "Expires": "0001-01-01T00:00:00Z",
                "MaxAge": 0,
                "Secure": true,
                "HttpOnly": true,
                "SameSite": 2,
                "Unparsed": null
            }
        ]
    }
    ```

## Admin API docs
https://github.com/xqdoo00o/ChatGPT-to-API/blob/master/docs/admin.md

## API usage docs
https://platform.openai.com/docs/api-reference/chat
