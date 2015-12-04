package cli

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"
)

func TestCommandDoNotIgnoreFlags(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	test := []string{"blah", "blah", "-break"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action:      func(_ *Context) {},
	}
	err := command.Run(c)

	expect(t, err.Error(), "flag provided but not defined: -break")
}

func TestCommandIgnoreFlags(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	test := []string{"blah", "blah", "-break"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:            "test-cmd",
		Aliases:         []string{"tc"},
		Usage:           "this is for testing",
		Description:     "testing",
		Action:          func(_ *Context) {},
		SkipFlagParsing: true,
	}
	err := command.Run(c)

	expect(t, err, nil)
}

func TestCommandYamlFileTest(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")
	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *Context) {
			val := c.Int("test")
			expect(t, val, 15)
		},
		UseYaml: true,
		Flags:   []Flag{IntFlag{Name: "test"}},
	}
	err := command.Run(c)

	expect(t, err, nil)
}

func TestCommandYamlFileTestGlobalEnvVarWins(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	os.Setenv("THE_TEST", "10")
	defer os.Setenv("THE_TEST", "")
	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *Context) {
			val := c.Int("test")
			expect(t, val, 10)
		},
		UseYaml: true,
		Flags:   []Flag{IntFlag{Name: "test", EnvVar: "THE_TEST"}},
	}
	err := command.Run(c)

	expect(t, err, nil)
}

func TestCommandYamlFileTestSpecifiedFlagWins(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	test := []string{"test-cmd", "--load", "current.yaml", "--test", "7"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *Context) {
			val := c.Int("test")
			expect(t, val, 7)
		},
		UseYaml: true,
		Flags:   []Flag{IntFlag{Name: "test"}},
	}
	err := command.Run(c)

	expect(t, err, nil)
}

func TestCommandYamlFileTestDefaultValueFileWins(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *Context) {
			val := c.Int("test")
			expect(t, val, 15)
		},
		UseYaml: true,
		Flags:   []Flag{IntFlag{Name: "test", Value: 7}},
	}
	err := command.Run(c)

	expect(t, err, nil)
}

func TestCommandYamlFileFlagHasDefaultGlobalEnvYamlSetGlobalEnvWins(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("test", 0)
	ioutil.WriteFile("current.yaml", []byte("test: 15"), 0666)
	defer os.Remove("current.yaml")

	os.Setenv("THE_TEST", "11")
	defer os.Setenv("THE_TEST", "")

	test := []string{"test-cmd", "--load", "current.yaml"}
	set.Parse(test)

	c := NewContext(app, set, nil)

	command := Command{
		Name:        "test-cmd",
		Aliases:     []string{"tc"},
		Usage:       "this is for testing",
		Description: "testing",
		Action: func(c *Context) {
			val := c.Int("test")
			expect(t, val, 11)
		},
		UseYaml: true,
		Flags:   []Flag{IntFlag{Name: "test", Value: 7, EnvVar: "THE_TEST"}},
	}
	err := command.Run(c)

	expect(t, err, nil)
}
