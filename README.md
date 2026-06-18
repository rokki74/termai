# TERMAI
A minimal cli application to interact with llms on internet from your terminal. 


## Focused on the bare minimal utility ask a question get it answered by nice llms.

Existing AI plugins rely on heavy Node.js dependencies and complex GUI overlays.
This tool takes a different approach: a pure Go binary that reads standard input, makes an HTTP request, and writes to standard output, treating AI generation as a simple communication from your terminal to the internet models.
Works seamlessly without installation of heavy runtimes or programs eg Node.js or Python or their backing works as a single binary on almost all machines.

## Why another Extra Tool 
After i found out about ddgr, i started switching into the tty(linux terminal) mostly never even using the gui at all then i faced one problem, i couldn't access the chatgpt or gemini, one of the very powerful models on internet yes the internet browsing i could mostly do using ddgr just by giving it search query and it leverages duduckgo servers to get me results back. So i tested some solutions out there.
So one of terminal tools for accessing ai models on internet duckaichat(a tool for duckduckgo ai), it was good except kept breaking due to error with NodeType in the python runtime due to improper installation setup. So i switched to tgpt 'gpt on shell in short' i didn't understand why it also failed i was too lazy to investigate as i needed something to work on the run immediately. 
In summary the most resilient of them all to me was gemini-cli and i tried it but i felt anxious everytime so i wanted to lock it away in isolation when using it so i tried firejail only to realise it needed some npm and nodejs environment to run then i added the --private flag to firejail still for some reason it wouldn't work so i just thought wait i just recently saw gemini advertising the gemini apikey creation for developer through vertex api i think so i went straight there pulled it then since i wanted to build a  fully independent tool i choose the best language for building it it was golang and that was how the idea for termai came up. Something which won't push configuration headaches to the user and comes fully independent while being simple enough to not expose even the tiniest detail of your to the internet and is old way with the model never being able to run behind the scenes and do irreversible actions on your file system.

## Visual Demo (also acts as summary for usage)


## Installation
git clone http://github.com/rokki74/termai.git

### Build the binary 
cd termai 

#### run this on your terminal
go build termai (this one builds automatically for your own machine)

##### Make sure you are running the command inside the termai directory use cd to move into it.
##### Can't build the binary, maybe don't want to install the needed golang program needed? Then access it at https://github.com/rokki74/termai/builds as i occasionally build it into the folder just so i avoid the inconvenience on some people just like you.

###### https://github.com/rokki74/termai/builds/windows --for windows .exe
###### https://github.com/rokki74/termai/builds/linux   --for linux
###### https://github.com/rokki74/termai/builds/mac     --for mac os

###### Once done just run wherever you put your executable into for example if you put it into /Desktop then you would want to run ./termai there 

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

## DEmo screenshots



#### Common Errors and Challenges
1. Forgetting to run termai install on the first setup (This brings up a window for you to paste in your api keys from either openai or gemini) Without which the program is just a piece of garbage in your computer so please remember 
2. The tool may output result that the model you choose is currently in spiked demand please switch to the next if you want for example gemini mostly has these limits but are fairer than openai's so preferred as the default for the tool 
3. The tool currently only supports two models as this was in my convinience but i shall add up more in the near future also rember the tool is open source
