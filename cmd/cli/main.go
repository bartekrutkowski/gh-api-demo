package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v42/github"
	"github.com/jedib0t/go-pretty/v6/table"
)

// config struct holds the application config
type config struct {
	username     string
	format_raw   bool
	format_table bool
}

// application struct holds GitHub client pointer, config struct pointer
// and application logger pointer
type application struct {
	client *github.Client
	config *config
	logger *log.Logger
}

// parseFlags function accepts a pointer to a config struct, reads the flags passed
// to the application from commandline, parses them and stores them in the config struct
func parseFlags(conf *config) {
	flag.StringVar(&conf.username, "acc", "", "GitHub account username (required)")
	flag.BoolVar(&conf.format_raw, "raw", false, "Display full output in raw JSON format (optional)")
	flag.BoolVar(&conf.format_table, "table", false, "Display stripped output in table format (optional)")
	flag.Parse()
}

// newApp function sets up standard logger, application config, parses the flags
// and returns pointer to application struct
func newApp() (app *application) {
	// Initialize application wide logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Initialize GitHub client
	client := github.NewClient(nil)

	// Initialize config struct
	var conf config

	// Read the commandline flags
	parseFlags(&conf)

	// Check if username was passed, print out usage and exit with error if not
	if conf.username == "" {
		flag.PrintDefaults()

		logger.Fatalln("username argument can't be empty or missing")
	}

	// Check if formatting flags were both set at once and exit with error if they were
	if conf.format_raw && conf.format_table {
		flag.PrintDefaults()

		logger.Fatalln("raw and table format output options are mutually exclusive")
	}

	// Initialie application struct
	app = &application{
		client: client,
		config: &conf,
		logger: logger,
	}

	return app
}

// getGistsByUser receiver function gets a context pointer and fetches a list
// of public gists for a user stored in the app.conf struct, returning it
// as a slice of github.Gist or an error
func (app application) getGistsByUser(ctx context.Context) ([]*github.Gist, error) {
	gists, _, err := app.client.Gists.List(ctx, app.config.username, nil)
	if err != nil {
		return nil, err
	}

	return gists, nil
}

// printGists receiver function gets a slice of pointers to github.Gist
// and prints them out depending on the values of app.config.format_raw and
// app.config.format_table struct fields
func (app application) printGists(gists []*github.Gist) {
	// Check if there are any gists to display
	if len(gists) == 0 {
		app.logger.Printf("No gists found for user %s\n", app.config.username)

		return
	}

	// Dump the full gist object if format_raw is set
	if app.config.format_raw {
		text, _ := json.MarshalIndent(gists, "", "	")
		fmt.Printf("%s\n", text)

		return
	}

	// Display formatted table with chosen fields if format_table is set
	if app.config.format_table {
		t := table.NewWriter()

		t.SetOutputMirror(os.Stdout)
		t.SetTitle(fmt.Sprintf("GitHub Gists for user %s", app.config.username))
		t.AppendHeader(table.Row{"#", "URL", "Description", "Date"})

		for i, gist := range gists {
			t.AppendRow(table.Row{
				i + 1, gist.GetHTMLURL(), gist.GetDescription(), gist.CreatedAt,
			})
		}

		t.Render()

		return
	}

	// Display list of gist links if neither format option is set
	for _, gist := range gists {
		fmt.Printf("%s\n", gist.GetHTMLURL())
	}
}

func main() {
	// Initialize application
	app := newApp()

	// Initialize context
	ctx := context.Background()

	// Get the gists for the user
	gists, err := app.getGistsByUser(ctx)
	if err != nil {
		app.logger.Fatalln(err)
	}

	// Print the gists accordingly to cli flags
	app.printGists(gists)
}
