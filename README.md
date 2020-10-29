# yet another-url-shortener

- generate n length randomized string/letters, set it as the key for the url
where the url here act as the value

=> {randomize_key: url}

append the data in-memory and for n interval of time, write the dictionary into a json file
then use it if the server crash, then load it all over to memory, use this routine. there's 
going to be a routine to check wheater a backup file is existed or not. load it to the memory if
exist, otherwise build empty dictionary to use. 
