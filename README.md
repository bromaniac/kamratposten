A super simple image board intended for < 10 users.

Goals:
- No database (data on disk)
- Serialize data to JSON
- No anonymous posting
- Hardcoded users to start with
- One static binary that can be rsynced to a server
- No client side rendering, server only (old school and simple)
- No attention mining
- No likes, no faves, no polls, no stars, no claps, no counts, no up/downvotes
- Flat threads, no nesting
- Ideally no (or minimal) configuration needed
- Should run fine on a micro VPS or RPI