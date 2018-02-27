// Calum Bell - 24016526 - Computer Science - 21/11/17
// Subtask One - PiOnTheDartboard (The Monte Carlo Problem)

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

/* Word Struct -
* This struct is used to map the array to a frequencies table
 */
type word struct {
	word      string
	frequency int // Number of times found
}

/* Keyword Struct -
 * The keyword struct is populated with user input and then converted to json
 * to be sent to the server
 */
type Keyword struct {
	Keyword string
}

// Stopwords to ignore during output sourced from - https://github.com/bbalet/stopwords
var stopwords = map[string]string{
	"a":            "",
	"about":        "",
	"above":        "",
	"across":       "",
	"after":        "",
	"afterwards":   "",
	"again":        "",
	"against":      "",
	"all":          "",
	"almost":       "",
	"alone":        "",
	"along":        "",
	"already":      "",
	"also":         "",
	"although":     "",
	"always":       "",
	"am":           "",
	"among":        "",
	"amongst":      "",
	"amoungst":     "",
	"amount":       "",
	"an":           "",
	"and":          "",
	"another":      "",
	"any":          "",
	"anyhow":       "",
	"anyone":       "",
	"anything":     "",
	"anyway":       "",
	"anywhere":     "",
	"are":          "",
	"around":       "",
	"as":           "",
	"at":           "",
	"back":         "",
	"be":           "",
	"became":       "",
	"because":      "",
	"become":       "",
	"becomes":      "",
	"becoming":     "",
	"been":         "",
	"before":       "",
	"beforehand":   "",
	"behind":       "",
	"being":        "",
	"below":        "",
	"beside":       "",
	"besides":      "",
	"between":      "",
	"beyond":       "",
	"bill":         "",
	"both":         "",
	"bottom":       "",
	"but":          "",
	"by":           "",
	"call":         "",
	"can":          "",
	"cannot":       "",
	"cant":         "",
	"co":           "",
	"con":          "",
	"com":          "",
	"could":        "",
	"couldnt":      "",
	"cry":          "",
	"de":           "",
	"describe":     "",
	"detail":       "",
	"do":           "",
	"done":         "",
	"down":         "",
	"due":          "",
	"during":       "",
	"each":         "",
	"eg":           "",
	"eight":        "",
	"either":       "",
	"eleven":       "",
	"else":         "",
	"elsewhere":    "",
	"empty":        "",
	"enough":       "",
	"etc":          "",
	"even":         "",
	"ever":         "",
	"every":        "",
	"everyone":     "",
	"everything":   "",
	"everywhere":   "",
	"except":       "",
	"few":          "",
	"fifteen":      "",
	"fify":         "",
	"fill":         "",
	"find":         "",
	"fire":         "",
	"first":        "",
	"five":         "",
	"for":          "",
	"former":       "",
	"formerly":     "",
	"forty":        "",
	"found":        "",
	"four":         "",
	"from":         "",
	"front":        "",
	"full":         "",
	"further":      "",
	"get":          "",
	"give":         "",
	"go":           "",
	"had":          "",
	"has":          "",
	"hasnt":        "",
	"have":         "",
	"he":           "",
	"hence":        "",
	"her":          "",
	"here":         "",
	"hereafter":    "",
	"hereby":       "",
	"herein":       "",
	"hereupon":     "",
	"hers":         "",
	"herself":      "",
	"him":          "",
	"himself":      "",
	"his":          "",
	"how":          "",
	"however":      "",
	"hundred":      "", // Custom Insertions
	"http":         "",
	"https":        "",
	"html":         "",
	"gif":          "",
	"jpg":          "",
	"utf8":         "",
	"s":            "",
	"ie":           "", // -----------------
	"if":           "",
	"in":           "",
	"inc":          "",
	"indeed":       "",
	"interest":     "",
	"into":         "",
	"is":           "",
	"it":           "",
	"its":          "",
	"itself":       "",
	"keep":         "",
	"last":         "",
	"latter":       "",
	"latterly":     "",
	"least":        "",
	"less":         "",
	"ltd":          "",
	"made":         "",
	"many":         "",
	"may":          "",
	"me":           "",
	"meanwhile":    "",
	"might":        "",
	"mill":         "",
	"mine":         "",
	"more":         "",
	"moreover":     "",
	"most":         "",
	"mostly":       "",
	"move":         "",
	"much":         "",
	"must":         "",
	"my":           "",
	"myself":       "",
	"name":         "",
	"namely":       "",
	"neither":      "",
	"never":        "",
	"nevertheless": "",
	"next":         "",
	"nine":         "",
	"no":           "",
	"nobody":       "",
	"none":         "",
	"noone":        "",
	"nor":          "",
	"not":          "",
	"nothing":      "",
	"now":          "",
	"nowhere":      "",
	"of":           "",
	"off":          "",
	"often":        "",
	"on":           "",
	"once":         "",
	"one":          "",
	"only":         "",
	"onto":         "",
	"or":           "",
	"other":        "",
	"others":       "",
	"otherwise":    "",
	"our":          "",
	"ours":         "",
	"ourselves":    "",
	"out":          "",
	"over":         "",
	"own":          "",
	"part":         "",
	"per":          "",
	"perhaps":      "",
	"please":       "",
	"put":          "",
	"rather":       "",
	"re":           "",
	"same":         "",
	"see":          "",
	"seem":         "",
	"seemed":       "",
	"seeming":      "",
	"seems":        "",
	"serious":      "",
	"several":      "",
	"she":          "",
	"should":       "",
	"show":         "",
	"side":         "",
	"since":        "",
	"sincere":      "",
	"six":          "",
	"sixty":        "",
	"so":           "",
	"some":         "",
	"somehow":      "",
	"someone":      "",
	"something":    "",
	"sometime":     "",
	"sometimes":    "",
	"somewhere":    "",
	"still":        "",
	"such":         "",
	"system":       "",
	"take":         "",
	"ten":          "",
	"than":         "",
	"that":         "",
	"the":          "",
	"their":        "",
	"them":         "",
	"themselves":   "",
	"then":         "",
	"thence":       "",
	"there":        "",
	"thereafter":   "",
	"thereby":      "",
	"therefore":    "",
	"therein":      "",
	"thereupon":    "",
	"these":        "",
	"they":         "",
	"thickv":       "",
	"thin":         "",
	"third":        "",
	"this":         "",
	"those":        "",
	"though":       "",
	"three":        "",
	"through":      "",
	"throughout":   "",
	"thru":         "",
	"thus":         "",
	"to":           "",
	"together":     "",
	"too":          "",
	"top":          "",
	"toward":       "",
	"towards":      "",
	"twelve":       "",
	"twenty":       "",
	"two":          "",
	"un":           "",
	"under":        "",
	"until":        "",
	"up":           "",
	"upon":         "",
	"us":           "",
	"very":         "",
	"via":          "",
	"was":          "",
	"we":           "",
	"well":         "",
	"were":         "",
	"what":         "",
	"whatever":     "",
	"when":         "",
	"whence":       "",
	"whenever":     "",
	"where":        "",
	"whereafter":   "",
	"whereas":      "",
	"whereby":      "",
	"wherein":      "",
	"whereupon":    "",
	"wherever":     "",
	"whether":      "",
	"which":        "",
	"while":        "",
	"whither":      "",
	"who":          "",
	"whoever":      "",
	"whole":        "",
	"whom":         "",
	"whose":        "",
	"why":          "",
	"will":         "",
	"with":         "",
	"within":       "",
	"without":      "",
	"would":        "",
	"www":          "",
	"yet":          "",
	"you":          "",
	"your":         "",
	"yours":        "",
	"yourself":     "",
	"yourselves":   ""}

/* Main function -
 * Handles the HTTP request and creates the parseManager to break down the response
 */
func main() {

	url := "http://127.0.0.1:8080/search/bing"

	/* ----- Get User Input ----- */
	inputtedWord := ""
	fmt.Println("URL:>", url)
	fmt.Println("Please enter a keyword to search for: ")
	fmt.Scanln(&inputtedWord)
	/* --------------------------- */

	// Use JSON.marshall to convert the struct to JSON format

	/* ----- Convert To JSON ----- */

	// Convert our inputtedWord String to a JSON object
	key := Keyword{inputtedWord}
	// Convert that struct to JSON
	buf, err := json.Marshal(key)
	/* Handle any Errors */
	if err != nil {
		log.Fatal(err) // Throw fatal error & print to console
	}

	var jsonStr = []byte(buf) // Convert to format which the server accepts
	/*-------------------*/

	fmt.Printf("%s\n", buf)

	/* --------------------------- */

	/* ----- Request from Server ----- */
	// Create the request with the JSON object we made
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json") // Set Header to JSON-type

	client := &http.Client{}    // Create a 'client' from the http library
	resp, err := client.Do(req) // Use the .Do function to send the request

	/* Handle any Errors */
	if err != nil {
		fmt.Printf("There was an error connecting to the server", err)
		panic(err)
	}
	/*-------------------*/
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	/* ---------------------------------- */

	/* File I/O */
	if resp.Status == "200 OK" {
		fmt.Printf("\nPrinting output to sampleresponse.txt!")
		responseFile, _ := os.OpenFile("sampleresponse.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // create file if not existing, and append to it. 666 = permission bits rw-/rw-/rw-
		responseFile.WriteString(string(body))
		defer responseFile.Close()
		parseManager(string(body), inputtedWord)
	} else {
		fmt.Printf("\nServer request failed! Attempting to use local file instead.\n\n")
		byteResponse, err := ioutil.ReadFile("sampleresponse.txt") // just pass the file name

		if err != nil {
			fmt.Print(err)

		}
		stringResponse := string(byteResponse)
		parseManager(stringResponse, inputtedWord)
	}

	/* -------- */

	// Create a parse manager to handle the response

	// Keep the process alive -
	for {

	}
}

/* parseManager  -
 * Initialises the parsers and channels to run the analysis in parallel
 * Provides the user with the ability to see the raw response
 * Listens on each channel for the response of the parsers
 * Combines the results of each parser in the finalArray
 * Sorts the final array into a map in descending frequency order
 */
func parseManager(body string, keyword string) {

	/* ----- Initialise variables ----- */

	textString := strings.ToLower(body)           // put all words in body to lowercase
	re := regexp.MustCompile("\\w+")              // * Define a regex expression to remove whitespace and punctuation
	textArray := re.FindAllString(textString, -1) // Take all valid words from the string and put them in textArray

	const noParsers = 3        // Number of threads, can be dynamic
	workToDo := len(textArray) // Find out how many words we need to process
	var allocatedWork = workToDo / noParsers
	var leftoverWork = workToDo % noParsers
	var finalMap map[string]int = map[string]int{} // parser results are combined into this map
	var finalTable []word                          // Stores the sorted word structs

	seeRawResponse := 'N' // Stores user input

	/* ------------------------------- */

	/* - Initialise array to store parser I/O channels - */

	// Initialise channel from parser to parseManager containing the frequency map - [word : frequency]
	var outputChannels [noParsers]chan map[string]int
	// Initialise channel from parseManager to parser containing the array slice of work
	var inputChannels [noParsers]chan []string

	/* ------------------------------------------------- */

	/* ----- Get User Input & Print Information ----- */

	fmt.Printf("\nparseManager was born!")
	fmt.Printf("\nparseManager recieved %v words related to %v\n", workToDo, keyword)
	fmt.Printf("\nDo you want to see the raw response? Y/N\n\n")
	fmt.Scanln(&seeRawResponse) // Read from the console

	if seeRawResponse == 78 {
		fmt.Println(body)
	}

	fmt.Printf("\n-------------------------------")
	fmt.Printf("\n|-------Initialising Data------|")
	fmt.Printf("\n-------------------------------")
	fmt.Printf("\nNumber of Parsers: %v", noParsers)
	fmt.Printf("\nTotal Work to Complete: %v", workToDo)
	fmt.Printf("\nAllocated Work: %v", allocatedWork)
	fmt.Printf("\nRemainder Work: %v", leftoverWork)
	fmt.Printf("\n-------------------------------\n\n")

	/* ---------------------------------------------- */

	/* ----- Generate Channels, Parsers and assign wo ----- */
	/* Generate and send each parser their section of work */

	/* Calculate Time */

	i := 0
	for i < noParsers {
		// Make and store the channels for this Parser
		inputChannels[i] = make(chan []string)
		outputChannels[i] = make(chan map[string]int)

		fmt.Printf("\nCreating Parser #%v", i)
		go parser(i, inputChannels[i], outputChannels[i])

		if i == (noParsers - 1) { // If we are the last parser, recieve any 'leftover work'
			fmt.Printf("\nSending Slice [%v:%v] to #%v", (allocatedWork * i), (allocatedWork*i)+allocatedWork+leftoverWork, i)
			work := textArray[(allocatedWork * i) : ((allocatedWork*i)+allocatedWork)+leftoverWork]
			inputChannels[i] <- work
		} else {
			fmt.Printf("\nSending Slice [%v:%v] to #%v", (allocatedWork * i), (allocatedWork*i)+allocatedWork, i)
			work := textArray[(allocatedWork * i):((allocatedWork * i) + allocatedWork)] //fmt.Printf("\nSending %v words to Parser %v", len(work), i)
			inputChannels[i] <- work
		}
		i += 1
	}

	/* ------------------------------------------- */

	/* ---------- Listen and combine completed work ---------- */
	/* Listen on each channel and add results to a final array */

	start := time.Now()

	count := 0
	for count < noParsers {
		select {
		case newMap := <-outputChannels[0]:
			{ // Listen on the first channel
				for k, v := range newMap {
					finalMap[k] += v
				}
				count += 1
			}
		case newMap := <-outputChannels[1]:
			{ // Listen on the second channel
				for k, v := range newMap {
					finalMap[k] += v
				}
				count += 1
			}
		case newMap := <-outputChannels[2]:
			{ // Listen on the third channel
				for k, v := range newMap {
					finalMap[k] += v
				}
				count += 1
			}
		}
	}

	timeAfterProcess := time.Now()
	elapsedPostProcess := timeAfterProcess.Sub(start)
	fmt.Printf("\nTime Taken to process into frequency map: %v", elapsedPostProcess)

	/* ----------------------------------------------------- */

	/* ---------- Sort Results into Table ------------ */

	fmt.Printf("\nResults Generated!\n")

	for k, v := range finalMap {
		finalTable = append(finalTable, word{k, v})
	}

	sort.Slice(finalTable, func(i, j int) bool {
		return finalTable[i].frequency > finalTable[j].frequency
	})

	for _, word := range finalTable {
		fmt.Printf("%s, %d\n", word.word, word.frequency)
	}

	/* ------------------------------------------------- */

	/* ---------------- Screen Output ------------------ */

	timeAfterSort := time.Now()
	elapsedPostSort := timeAfterSort.Sub(start)
	totalTime := elapsedPostSort + elapsedPostProcess
	// var timePerWord float64
	// timePerWord = (totalTime)/ workToDo)

	fmt.Printf("\n|---------------------------------------------|")
	fmt.Printf("\n|----------------- Completed! ----------------|")
	fmt.Printf("\n|---------------------------------------------|")
	fmt.Printf("\n| Search Term: %v                    |", keyword)
	fmt.Printf("\n|---------------------------------------------|")
	fmt.Printf("\n| Number of Parsers: %v                        |", noParsers)
	fmt.Printf("\n| Total Work to Complete: %v               |", workToDo)
	fmt.Printf("\n| Allocated Work: %v                       |", allocatedWork)
	fmt.Printf("\n| Remainder Work: %v                           |", leftoverWork)
	fmt.Printf("\n|----------------- Time Taken ----------------|")
	fmt.Printf("\n| Time Taken to process into frequency map:   |\n| %v                                 |", elapsedPostProcess)
	fmt.Printf("\n| Time Taken to sort into table: %v  |", elapsedPostSort)
	fmt.Printf("\n| Total Time Taken: %v               |", totalTime)
	//fmt.Printf("\n| Time Taken per word: %v               |", timePerWord)
	fmt.Printf("\n|---------------------------------------------|")
	fmt.Printf("\n|------- Top Ten Results (- Stop Words)-------|")
	fmt.Printf("\n|---------------------------------------------|")

	numberOfResults := 0
	for i := 0; numberOfResults < 10; i++ {
		_, exists := stopwords[finalTable[i].word]
		if exists == true {
		} else { // Start at 1
			fmt.Printf("\n| %v : %v", finalTable[i].word, finalTable[i].frequency)
			numberOfResults++
		}
	}
	fmt.Printf("\n|---------------------------------------------|\n\n")

	/* ------------------------------------------------- */

}

/* parser  -
 * Takes an arraySlice from parseManager via the todo channel
 * Converts the array to a frequencies map of word structs
 * Sends the map back to parseManager via the done channel
 */
func parser(id int, todo <-chan []string, done chan map[string]int) {
	fmt.Printf("\nParser #%v is now running!", id)
	var wordsToParse = <-todo // Recieve work from parseManager
	fmt.Printf("\nParser #%v has %v words to process: \n", id, len(wordsToParse))

	// Frequency table - each word has an associated value - [word, frequency]
	frequencies := make(map[string]int)
	// range string slice gives index, word pairs

	// For each word in our slice, set word to the word at index
	// Range gives each word and the index its at, index isn't used so we use '_'
	for _, word := range wordsToParse {

		// Get the current value for this word from the map
		_, freq := frequencies[word]

		// If this value exists, add one
		if freq == true {
			frequencies[word] += 1
		} else { // Start at 1
			frequencies[word] = 1
		}
	}
	// fmt.Print(frequencies)
	done <- frequencies
}
