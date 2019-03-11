package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Struct declaration to assimilate the player names and numerical counterparts
As per instruction, the player has an attribute list known as stats, so the way I grouped them
together was to make a 'player' type that can be applied to all players
*/
type player struct {
	/*
	fields of the same type can be declared on the same line. I seperated them
	out slightly to improve readability, the fields in the struct are in order with the lines of the
	input file, i.e. Hank Aaron 13245...etc
	*/
	firstname, lastname                string
	plate_apperances, at_bats, singles int
	doubles, triples, homeruns         int
	walks, hit_by_pitch                int
}
/*
Function to get the filename from the user, and give user the feedback on given path
If the user enters a valid path, then the function will carry on.
If the path is invalid, then an error message will display.
Use the bufio scanner to read the file
After opening and reading the file, convert each line to a single
list for calculations on that list later
*/



func pathway(path string) ([]string, error) {
	var err error = nil
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("This is not a valid path")
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return nil, errors.New("File cannot be read")
	}

	// Scan lines
	scanner.Split(bufio.ScanLines)
	var lines []string
	// var counter int
	for i := 0; scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}
	return lines, err
	
}
/*
A function to set and increase a counter for the iteration
through all lines in the file
The number of players is determind by the number of lines
since each line begins with a player name
*/
func line_scan(file string)  (int){
	filename, _ := os.Open(file)
	fileScan := bufio.NewScanner(filename)
	player_count := 0
	for fileScan.Scan(){
		player_count++
	}
	fmt.Println("Found in file are ----->: ", player_count, "Players")
	fmt.Println("|------------------------------------------------------------------------------------------------------------------------|")

	return player_count
}


/*
Main function
*/
func main() {
	/*
	Get user to enter the string name that is their input file.
	System will search directory for this file and open.
	Function pathway is run on file considering that there are no errors
	*/
	var user_entered string
	fmt.Print("Please enter the name of the file you would like to open that is in your current directory\n")
	fmt.Scan(&user_entered)
	lines, err := pathway(user_entered)
	if err != nil {
		log.Fatalf(err.Error())
	}
	/* 
	line count is the totaL lines in the file as 
	gathered from the function above, I'm using it here to create a 
	dynamic way to adjust the count below so that it always relies
	on the given filesize
	*/
	line_count := line_scan(user_entered)
	/*
	For the index between start and number of lines
	getRecord on that line at an iterative index increment
	*/
	for i := 0;  i < line_count; i++{
		Calc(lines[i])
		/*
		Added for spacing
		*/
		fmt.Println("|------------------------------------------------------------------------------------------------------------------------|")
	}
}

func Calc(in string) (player, error){
	var pl player
	words := strings.Fields(in)
	// Based on input, this shouldn't happen, but error checking anyway
	if len(words) != 10 {
		return pl, errors.New("Wrong number of fields parsed, some fields couldn't be calculated due to missing data. Pleease check file")
	}
	/*
	error check
	*/
	vs, err := getInts(words[2:])
	if err != nil {
		return pl, err
	}
	pl = player{words[0], words[1], vs[0], vs[1], vs[2],vs[3],vs[4],vs[5],vs[6],vs[7]}
	/*
	Names correspond with fields in the struct, only did this to make it a bit simpler
	*/
	firstname := pl.firstname
	lastname := pl.lastname
	singles := pl.singles
	doubles := pl.doubles
	triples := pl.triples
	homeruns := pl.homeruns
	at_bats := pl.at_bats
	walks := pl.walks
	hbp := pl.hit_by_pitch
	pa := pl.plate_apperances
	/*
	Batting average calculations and appropriate conversions into a decimal form
	All variables are based on formula and repective conversions to make averages
	*/
	batting_total := float32((singles + doubles + triples + homeruns))
	at_bats_conversion := float32(at_bats)
	batting_average := ((batting_total/at_bats_conversion))
	//!!! Figure out how to get all batting averages in this array, and do the average for them
	var store []float32
	store = append(store, batting_average)
	fmt.Println(store)
	// !!!
	
	/*
	Slugging calculations and appropriate conversions into a decimal form
	All variables are based on formula and repective conversions to make averages
	*/
	slugging_total := float32((singles + 2 * doubles + 3 * triples + 4 * homeruns))
	slugging := (slugging_total/at_bats_conversion)
	/*
	On base calculations and converstions into decimal form
	All variables are based on formula and repective conversions to make averages
	*/
	on_base := float32(singles + doubles + triples + homeruns + walks + hbp)
	plate_apperances_conversion := float32(pa)
	on_base_percentage := (on_base/plate_apperances_conversion)
	/*
	Printing the percentages out in the format specified
	*/
	fmt.Println(" Player: ",firstname, lastname, ":", "[", "Batting Average:",batting_average, "]", "[",
	"Slugging:", slugging, "]", "[" ,"On Base Percentage: ", on_base_percentage, "]")




	return pl, nil
}

func getInts (xs []string) ([]int, error) {
	var ys = make([]int, len(xs))
	/*
	Parse ints from file using string convert method,
	We are only given 8 stats, otherwise I'd make this dynamic as well
	*/
	for i := 0; i < 8; i++ {
		v, err := strconv.ParseInt(xs[i], 0, 64)
		if err != nil {
			return nil, errors.New("Could not convert a number")
		}
		ys[i] = int(v)
	}
	return ys, nil
}