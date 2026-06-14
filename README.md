# TERMAI
A minimal cli application to interact with llms on internet from your terminal. 


## Focused on the bare minimal utility ask a question get it answered by nice llms.

Existing AI plugins rely on heavy Node.js dependencies and complex GUI overlays.
This tool takes a different approach: a pure Go binary that reads standard input, makes an HTTP request, and writes to standard output, treating AI generation as a simple communication from your terminal to the internet models.
Works seamlessly without installation of heavy runtimes or programs eg Node.js or Python or their backing works as a single binary on almost all machines.

## Why another Extra Tool 
After i found out about ddgr, i started switching into the tty mostly bever even loading the gui at all then i faced one problem, i couldn't access the chatgpt or gemini very powerful models on internet so i tested some solutions out there.
So one of then duckaichat, it was good except kept breaking due to error with NodeType so i switched to tgpt 'gpt on shell in short' i didn't understand why it also failed i was too lazy to investigate. 
In summary the most resilient of them all to me was gemini-cli and i tried it but i felt anxious everytime so i wanted to lock it away in isolation when using it so i tried firejail only to realise it needed some npm and nodejs environment to run then i added the --private flag to firejail still for some reason it wouldn't work so i just thought wait i just recently saw gemini advertising the gemini apikey creation for developer through vertex api i think so i went straight there pulled it then since i wanted to build a  fully independent tool i choose the best language for building it it was golang and that was how the idea for termai came up.

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


