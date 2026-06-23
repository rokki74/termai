package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"github.com/sashabaranov/go-openai"
)

func main(){
	verbosity := flag.Bool("v", false, "-v for opting between verbosrre output and non verbose, get more information when verbose selected")
	interactivity := flag.Bool("i", false, "-i to opt into interactive mode or out whereby you won't have to retype the ./termai command before running the tool everytime")
	targetEngine := flag.String("target", "gemini", "-target Use this to choose between the models you want to use for example -target='chatgpt'")
	query := flag.String("search", "", "-search='your query' use to search your query directly and return immediately or get output out at once without any other functionality of the code. Useful when working with long chained pipelines and you need the result as just what the model answers back with no additional tool's information.")
	installCmd := flag.NewFlagSet("install", flag.ExitOnError)

	if *verbosity{
		log.Println("Getting the home directory for saving configs into file")
	}
	homedir, err := os.UserHomeDir()
	if err != nil{
		log.Fatalf("Could not get the home directory. Error: %v", err)
	}

	configFile := path.Join(homedir, "termaiconfig.txt")
	availableModels := []string{"GEMINI", "CHATGPT"}
	if len(os.Args)>1 && os.Args[1] =="install"{
		installCmd.Parse(os.Args[2:])
		if *verbosity{
			log.Println("Running initial setup")
		}
	  install(configFile, availableModels)
	}

	if *verbosity{
			log.Println("Configuring..")
		}
  config(configFile, availableModels)

	if *verbosity{
		log.Println("Reading the config file")
	}
 models, found := readConfigFile(configFile)
 if !found{
	 log.Fatal("Cannot find api keys, make sure to run termai -install before using the tool!")
 }

	var client *openai.Client
	var req openai.ChatCompletionRequest
	switch *targetEngine{
		case "gemini": 
			apikey, found := models["GEMINI"]
			if !found{
				log.Fatal("Gemini's key not found, try running termai --install to fix the problem")
			}
			config := openai.DefaultConfig(apikey)
			config.BaseURL = "https://generativelanguage.googleapis.com/v1beta/openai"

			client = openai.NewClientWithConfig(config)

			req = openai.ChatCompletionRequest{
				 Model: "gemini-3.5-flash",
				 Stream: true,
			}
			if *verbosity{
				log.Println("Using gemini as the target model")
			}
		case "chatgpt":
			apikey, found := models["CHATGPT"]
			if !found{
				log.Fatal("Chatgpt's key not found, try running termai --install to fix the problem")
			}
			config := openai.DefaultConfig(apikey)
			config.BaseURL = "https://generativelanguage.googleapis.com/v1beta"

			client = openai.NewClientWithConfig(config)

			req = openai.ChatCompletionRequest{
				 Model: "gemini-3.5-flash",
				 Stream: true,}	
			if *verbosity{
				log.Println("Using chatgpt as the target model")
			}

		default: 
		   log.Fatalf("There doesn't exist such a model in support like %v", *targetEngine)
		}

	if *verbosity{
		log.Println("Ready to run..")
	}
 if isPiped(){
	 runAi(client, &req, query)
	 return
 }else if *interactivity{
 runInteractive(client, &req)
	return
 }
}

func isPiped() bool{
	fileInfo, err := os.Stdin.Stat()
	if err != nil{
		return false
	}

	return (fileInfo.Mode() & os.ModeCharDevice) ==0
}

func runInteractive(client *openai.Client, req *openai.ChatCompletionRequest){
 fmt.Println("WELCOME TO TERMAI HELPS YOU ACCESS INTERNET MODELS eg GPT, GEMINI\n pipe, ask interactively or one line command to provide input, use --help to find out the correct flags to use for example ./termai --help")

 for{
	 fmt.Println("/termai>")
	 bytes, err := io.ReadAll(os.Stdin)
	 if err !=nil{
		 log.Fatalf("Error encountered reading input, %v", err)
	 }

	 prompt := string(bytes)
	 runAi(client, req, &prompt)
 }
}

func runAi(client *openai.Client, req *openai.ChatCompletionRequest, prompt *string){
	ctx := context.Background()

  req.Messages = []openai.ChatCompletionMessage{
			 {
				 Role: openai.ChatMessageRoleUser,
				 Content: *prompt,
			 },
		}

	response, err := client.CreateChatCompletion(ctx, *req)
	if err != nil{
		log.Fatalf("Error occured with client, Error: %v\n", err)
	}

	 fmt.Println(response.Choices[0].Message.Content)
}

func config(configFile string, availableModels []string){
   _, err := os.Stat(configFile)
   if os.IsNotExist(err){
		 log.Println("Didn't find previous configs, prompting for install ..")
	 install(configFile, availableModels)
	 }

	 log.Println("Found a configuration file do you wish to override and setup new apikeys? (yes/no)")
	 var input string
	 fmt.Scan(&input)

	 if strings.ToLower(input)=="yes"{
		 install(configFile, availableModels) 
	 }
}

func install(configFile string, availableModels []string){
	fmt.Println("CONFIGURING TERMAI. Follow prompts to continue")
	fmt.Println("Enter the api keys in provided fields")

	models := make(map[string]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for i :=0; i<len(availableModels); i++{
		fmt.Printf("%v API KEY: ", availableModels[i])	
    
		var prompt string
		if scanner.Scan(){
       prompt = scanner.Text()

		   models[availableModels[i]] = strings.TrimSpace(prompt)
		}

		if err := scanner.Err(); err != nil{
			log.Printf("Error occurred reading in the api key for model:  %v, Error: %v", availableModels[i], err)
		  continue
		}
	}

	written := writeConfigFile(configFile, models)

	if !written{
		log.Fatalf("Could not write config file %v", configFile)
	}
}

func readConfigFile(configFile string)(map[string]string, bool){
  f, err := os.Open(configFile)	
	if err != nil{
		log.Fatalf("Error reading file. Error: %v", err)
	}
  defer f.Close()

	scanner := bufio.NewScanner(f)

	models := make(map[string]string, 0)
	for scanner.Scan(){
		line := scanner.Text()
		mKey, mValue, found := strings.Cut(line, "=")
		if !found{
			continue
		}

		models[mKey] = mValue
	}
 
	return models, true
}

func writeConfigFile(configFile string, models map[string]string) bool{
  f, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) 
	if err != nil{
		return false
	}
  defer f.Close() 

	for key, value := range models{
	  fmt.Fprintf(f, "%s=%s\n", strings.ToUpper(key), value)	
	}

	return true
}

