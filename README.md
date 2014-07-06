public-ip
=========

Return public ip via api, written in Go. I use it to display my public ip address in tmux status bar.

Hosted on Heroku and accesible via [http://ip.vladimirm.com/](http://ip.vladimirm.com/)

Response types:

| Url                          | Format | Response               |
| :--------------------------- | :----- | :--------------------- |
| http://ip.vladimirm.com/     | text   | 203.0.113.13           |
| http://ip.vladimirm.com/json | json   | { "ip": "203.0.113.13" |
