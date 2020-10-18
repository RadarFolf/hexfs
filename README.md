
![logo-Artboard-1-0-4x.png](https://i.postimg.cc/YSXZmBDk/logo-Artboard-1-0-4x.png)

# hexFS

Do-it-yourself modern file host software, written in Golang. An excellent, efficient, and sleek alternative to JavaScript uploaders.

### What can you use it for?

- Your screenshots from ShareX or MagicCap
- Having *your* own platform to store *your* files on, not someone else's. Maybe you don't trust them. Or maybe you have a cool domain you want to run this under.

### About the project

- Store data about files (SHA256, original IP, etc.) with MongoDB
- Delegation of files to external Google Cloud Storage with encryption at-rest. 
- Simple access system - either make it private, share a standard key with friends which allows them to upload, not delete, or open to the public (though the latter isn't recommended!).
- Bulk delete files by IP, SHA256 hash and individual IDs. 
- Built-in file extension filter--protect your users from malicious extensions.
- Final executable is small - only about 16 MB in size.
- Rate limiting and max capacity watch (to limit the amount of files that can be stored), handled by Redis
- It's not written in JavaScript. 
- The logo is very cool. :^)

### Some key notes

- hexFS cannot download files on your behalf from the Internet. This is deliberately a security decision.
- hexFS will run completely fine if not containerized, but it's still *recommended*! You could use tmux or systemd to run it. Just throwing out ideas.
- There are no fancy plugins or extensions, or frontend template for that matter. Make them yourself.

### Run

- Put your key in conf/ as "key.json"
- Put your config in conf/ as "config.yml" using "conf/example.yml" as a reference

Make sure to bind hexFS port (3030) to other ports on your system. Here's an example of how you would run it, after building the image.

`sudo docker container run -d -p 127.0.0.1:3030:3030 --name hexfs hexfs`

### Support

[discord server](https://discord.gg/F7RBKh2).