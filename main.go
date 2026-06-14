package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"github.com/sashabaranov/go-openai"
)

func main() {
	fmt.Println("WELCOME TO TERMAI")
  fmt.Println("Using gemini 2.5 flash and GPT4o mini as available options for target engines...")
  
	sigChan := make(chan os.Signal, 1)
  signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	targetEngine := "gemini" 
	
	var client *openai.Client
	var modelName string
  
	mdls := make(map[string]string)

	homedir, err := os.UserHomeDir()
	if err != nil{
		fmt.Printf("Error occurred during setup cannot get home directory, error: %v\n",err)
	}
	configFileName := "termaiconfig.txt"
	configFilePath := path.Join(homedir, configFileName)
  setup(mdls, configFilePath);

	fmt.Println("Reading saved configuration..")
	models := make(map[string]string) 
	if err := readConfigFile(models, configFilePath); err != nil{
		log.Fatal("Unable to read the config file, may have been deleted!")
	}

	log.Println("Read the configs successfully!\n Waiting for llm's response...")
	if targetEngine == "gemini" {
		log.Println("Using selected model Gemini...")
		gem, ok := models["GEMINI"]
		if !ok{
			fmt.Println("key GEMINI key not found in map")
			return
		}
		config := openai.DefaultConfig(gem)
		// Reroute the OpenAI client to Google's compatibility endpoint
		config.BaseURL = "https://generativelanguage.googleapis.com/v1beta/openai/v1"
		client = openai.NewClientWithConfig(config)
		modelName = "gemini-2.5-flash"
	} else {

		log.Println("Using selected model Chatgpt...")
		gpt, ok := models["CHATGPT"]
		if !ok{
			fmt.Println("key CHATGPT key not found in map")
			return
		}
		config := openai.DefaultConfig(gpt)
		client = openai.NewClientWithConfig(config)
		modelName = openai.GPT4o
	}

	log.Println("Type in your prompt: ")

	for{
		  select{
			case <-sigChan:
				fmt.Println("Bye!")
				os.Exit(0)
			default:
				runAi(*client, modelName)
			}
	}	
}

func runAi(client openai.Client, modelName string){
	   var prompt string;

			fmt.Print(">")
      reader := bufio.NewScanner(os.Stdin)
			if reader.Scan(){
				prompt = reader.Text()
			}

			if err := reader.Err(); err != nil{
				fmt.Println("Bye!")
				os.Exit(0)
			}

			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model: modelName,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: prompt, 
						},
					},
				},
			)

			if err != nil {
				log.Printf("ChatCompletion error: %v\n", err)
				return
			}

			fmt.Println("===============RESPONSE============")
			fmt.Println(resp.Choices[0].Message.Content)

}

func setup(mdls map[string]string, configFilePath string){

	fmt.Printf("using config file: %v for configurations.\n", configFilePath)
	

  configureApiKeys(mdls, configFilePath)
}

func configureApiKeys(mdls map[string]string, configFilePath string){
	_, err := os.Stat(configFilePath)
  if err !=nil && errors.Is(err, os.ErrNotExist){
		fmt.Println("api keys not configured\n Please paste in your api keys if this is your first time")
		if _, e := os.Create(configFilePath); e !=nil{
			log.Fatalf("Error occurred creating the configuration file. error: %v\n", e)
		}	
    
		setupModels(mdls, configFilePath)
	}else{
			fmt.Println("Found an older configuration do you wish to change the saved keys? (yes/no)")
		  var input string;
			fmt.Scan(&input)

			if strings.ToLower(input) == "yes"{
			    setupModels(mdls, configFilePath)
       }
		}
}

func setupModels(modelsMap map[string]string, configFilePath string){
    fmt.Println("please paste each api to their respective provider, leave out if you wish to skip the model") 

		models := [2]string{"GEMINI", "CHATGPT"}
		for i := 0; i<2; i++{
				fmt.Printf("%v KEY: \n", models[i])
				reader := bufio.NewReader(os.Stdin)
				 
				input, err := reader.ReadString('\n')
				if err !=nil{
					log.Printf("skipping %v's key..\n", models[i])
				}

        modelsMap[models[i]] = strings.TrimSpace(input)
		}

	
		if saveModelKeys(modelsMap, configFilePath){
			fmt.Println("Almost there, final touches...")
		}else{
			log.Fatal("Could not save the api keys, please check again..")
		}
	}

	func saveModelKeys(modelsMap map[string]string, configFilePath string)bool{


		file, err := os.OpenFile(configFilePath, os.O_WRONLY, 0644)
		if err != nil{
			fmt.Printf("Failure opening config file, %v\n", err)
			return false
		}
    defer file.Close()

		var written int = 0;
		for mName := range modelsMap{
			  mKey := modelsMap[mName]
				n, err := fmt.Fprintf(file, "%s=%s\n", mName,mKey)
				if err !=nil{
					fmt.Printf("failed to save config for model %v due to error: %v\n", mName, err)
					continue
				}

        written +=n;
		}
    
		if written >0{
			return true
		}
		return false
	}

func readConfigFile(mdls map[string]string, configFilePath string) error{
   f, err := os.Open(configFilePath)

	 if err != nil{
		 fmt.Printf("Failed to read configuration file due to error: %v, skip if it is the first time loading the tool...\n", err)
		 
		 return nil
	 }
   defer f.Close()

   scanner := bufio.NewScanner(f)

	 for scanner.Scan(){
		 line := strings.TrimSpace(scanner.Text())
		 if line ==""||len(line)==0{
			 continue
		 }

		 parts := strings.SplitN(line, "=", 2)
		 if len(parts)==2{
			 mName := strings.TrimSpace(parts[0])
			 mKey := strings.TrimSpace(parts[1])

			 mdls[mName]=mKey
		 }
	 }

	 return nil
}
