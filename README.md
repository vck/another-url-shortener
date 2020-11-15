# What

YET ANOTHER URL SHORTENER

IDEA: 

generate n length randomized string/letters, set it as the key for the url
where the url here act as the value of {randomize_key: url}

GOAL:

append the data in-memory and for n interval of time, write the dictionary into a json file
then use it if the server crash, then load it all over to memory, use this routine. there's 
going to be a routine to check wheater a backup file is existed or not. load it to the memory if
exist, otherwise build empty dictionary to use. 


# How

- Run `make run`

- Sort `curl localhost:8080/sort?url=www.google.com`

- Redirect to sorted: `localhost:8080/<KEY>`

# Todo

- Add benchmark for sorting algorithm
- Redirect still not working [BUG]
- Write db into a file [FEATURE]
- Lookup from filedb [FEATURE]
- Add routine to check if filedb exists [FEATURE]
