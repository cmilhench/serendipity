package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cmilhench/serendipity"
)

var revision = "latest"

func main() {
	global := flag.NewFlagSet("serendipity", flag.ExitOnError)
	global.Usage = func() {
		fmt.Fprintf(global.Output(), "usage: %s <command> [<args>]\n", global.Name())
		global.PrintDefaults()
	}
	global.Parse(os.Args[1:])
	if len(global.Args()) == 0 {
		global.Usage()
		os.Exit(1)
	}

	log.SetFlags(log.LstdFlags)

	switch global.Args()[0] {
	case "person":
		personCommand()
	default:
		fmt.Printf("%q is not valid command.\n", global.Args()[0])
		os.Exit(2)
	}
}

func personCommand() {
	command := flag.NewFlagSet("person", flag.ExitOnError)
	command.Usage = func() {
		fmt.Fprintf(command.Output(), "usage: %s %s <command> [<args>]\n", os.Args[0], command.Name())
		fmt.Fprintf(command.Output(), "Creates random user profiles\n")
		command.PrintDefaults()
	}
	count := command.Uint("n", 1, "Number of people")
	seed := command.String("s", "", "Random seed")
	command.Parse(os.Args[2:])
	if len(command.Args()) != 0 {
		command.Usage()
		os.Exit(1)
	}

	outputCSV(*count, *seed)
}

func outputJson(count uint, seed string) {
	h := md5.New()
	h.Write([]byte(seed))
	num := binary.BigEndian.Uint64(h.Sum(nil))

	r := serendipity.New()
	r.Seed(int64(num))

	var p *serendipity.PersonInfo

	fmt.Print("[\n")
	for i := 0; i < int(count); i++ {
		p = r.Person()
		buf, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}
		if i != 0 {
			fmt.Print(",\n")
		}
		fmt.Printf("%s", buf)
	}
	fmt.Print("\n]")
}

func outputCSV(count uint, seed string) {
	h := md5.New()
	h.Write([]byte(seed))
	num := binary.BigEndian.Uint64(h.Sum(nil))

	r := serendipity.New()
	r.Seed(int64(num))

	w := csv.NewWriter(os.Stdout)

	if err := w.Write(toHeaders(r.Person())); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}
	for i := 0; i < int(count); i++ {
		if err := w.Write(toSlice(r.Person())); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
func toHeaders(*serendipity.PersonInfo) []string {
	return []string{
		"sub", "preferred_username", "name", "given_name", "family_name", "middle_name", "nickname",
		"profile", "picture", "website", "email", "email_verified", "gender", "zoneinfo", "locale",
		"phone_number", "phone_number_verified", "address_id", "birthdate", "updated_at", "created_at",
	}

}
func toSlice(record *serendipity.PersonInfo) []string {
	values := make([]string, 21)
	if record == nil {
		return values
	}
	values[0] = record.Sub
	values[1] = record.Username
	values[2] = record.Name
	values[3] = record.GivenName
	values[4] = record.FamilyName
	values[5] = record.MiddleName
	values[6] = record.Nickname
	values[7] = record.Profile
	values[8] = record.Picture
	values[9] = record.Website
	values[10] = record.Email
	values[11] = fmt.Sprintf("%v", record.EmailVerified)
	values[12] = fmt.Sprintf("%s", record.Gender)
	values[13] = record.ZoneInfo
	values[14] = record.Locale
	values[15] = record.PhoneNumber
	values[16] = fmt.Sprintf("%v", record.PhoneNumberVerified)
	values[17] = ""
	values[18] = record.Birthday.String()
	values[19] = record.Updated
	values[20] = time.Now().UTC().String()

	return values
}
