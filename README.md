This is slapped and bodged together, don't expect anything.

## Configuration
Non-platform environment:
```
NITTER_DOMAIN
```

### Getting Twitter creds
1. Make an account https://developer.twitter.com/en/portal/
1. Click `View products`, get `Elevated` access.
1. `Projects & Apps` → `Project 1` → `FoxNitter`
1. `User authentication settings` →
   1. Enable `OAuth 1.0a`
   1. App permissions → `Read, Write, Direct messages` (apparently you can't retweet without DMs‽)
   1. `Callback` can be anything
1. `Keys and tokens` → `Consumer keys` (ENV: CLIENT_*)
1. `Authentication Tokens` → `Generate` (ENV: ACCESS_*)

Environment needed:
```
CLIENT_ID
CLIENT_SECRET
ACCESS_TOKEN
ACCESS_SECRET
```

### Getting Telegram creds
1. You can get a Telegram bot API key from [@botfather](https://t.me/botfather) (env: `TELEGRAM_KEY`)
1. Get channel ID by running the bot without `TELEGRAM_ID`

Environment needed:
```
TELEGRAM_KEY
TELEGRAM_ID
```
