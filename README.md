# ChatGPT Reddit Bot

This is a stupid little terminal Reddit bot that retrieves
new posts from a specified subreddit and generates responses using ChatGPT.

All it does is get `N` new posts from a specified subreddit and then get a ChatGPT response.
It then prints things to the terminal... that's it, nothing special. Just some tinkering

## Prerequisites

- Install Go 1.22
- Get yourself an OpenAI API key.

## Usage

To run the program, first set `OPEN_AI_KEY` in your environment:

```bash
export OPEN_AI_KEY=<your_key>
```

then execute the following command:

```bash
go run main.go
```

## Configuration

You can modify the following parameters in the main function:

- `Subreddit`: Specify the subreddit from which to retrieve posts.
- `N`: Number of posts to retrieve.
- `Random`: Set to true to retrieve posts randomly.
