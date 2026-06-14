# TERMAI
A minimal cli application to interact with llms on internet from your terminal. 


## Focused on the bare minimal utility ask a question get it answered by nice llms.

Existing AI plugins rely on heavy Node.js dependencies and complex GUI overlays.
This tool takes a different approach: a pure Go binary that reads standard input, makes an HTTP request, and writes to standard output, treating AI generation as a simple communication from your terminal to the internet models.
Works seamlessly without installation of heavy runtimes or programs eg Node.js or Python or their backing works as a single binary on almost all machines.

## Installation
git clone http://github.com/rokki74/termai.git

### Build the binary 
cd termai 

go build termai

##### Can't build the binary, maybe don't want to install the needed golang program needed? Then access it at https://github.com/rokki74/termai/builds/tool_binary as i occasionally build it into the folder just so i avoid the inconvenience on some people just like you.

### Run 
termai

##### run termai --help to see how to use it further and flags etc

### Add termai to system wide commands if needed 
#### mac os 
1. Move the binary to the system path
sudo cp termai /usr/local/bin/


Use mv instead to move thus: sudo mv termai /usr/local/bin

2. Grant execution permissions
sudo chmod +x /usr/local/bin/termai

# 3. Fix the macOS Gatekeeper warning (avoids the "Developer cannot be verified" popup)
xattr -d com.apple.quarantine /usr/local/bin/nvim-ai-filter  

#### linux 
1. Move or copy the compiled binary to the system path
sudo cp termai /usr/local/bin/

Use mv instead to move thus: sudo mv termai /usr/local/bin

2. Ensure it has execution permissions
sudo chmod +x /usr/local/bin/termai

#### windows
1. Open system environment variables and add the path to termai into it.
2. Or although i won't give good view in this, You can manually move the binary to C:\Windows\System32\ 

## Seek Help 
You can always ask for help from me or other programmers('yes literally anyone of them will help or refer you to a better person in the case') you might know by forwarding them the program's github link with the question you are asking. 
Please don't forget to mention crucial information like the name of the program and the repository and if possible try reading this README.md or using termai --help if you have had it installed.
I Japheth will always be happy to help, I might not respond quickly please be a little bit patient.


