# Hiccup
---
Take advantage of Burp scope files and use them in the shell during recon and information gathering.

The goal of this tool is to help speed up recon, by removing out of scope sites early on in the process, without 
having to set up or apply custom regex to each output.
By taking advantage of the HackerOne export option (though not perfect) this provides a nice workflow to clean up 
target lists.


## Features

- Match based on include, or exclude lists
- Stdin based input - friendly to scripts
- Protocol matching, or just host matching