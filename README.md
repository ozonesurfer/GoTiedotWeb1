This is a sample MVC website/web server written in Go (golang) and driven by Howard Guo's Tiedot database. Tiedot was written entirely in Go.

# Dependencies

To install the necessary packages, issue the following commands:

<em>go get github.com/QLeelulu/goku</em>

<em>go get github.com/HouzuoGuo/tiedot</em> 

or 

<em>go get loveoneanother.at/tiedot</em>

The most recent build of Tiedot might not be at loveoneanother.at/tiedot, so try the first command first.

# Configuration

You will need to change the DATABASE_DIR setting in <em>/src/gotiedotweb/config.go</em> and make sure the new path exists.

# Compilation And Execution

Add this git's path to the GOPATH environment variable, then go to its root and issue

<em>go build</em>

That should generate an executable suitable for your operating system. Linux users might need to add parameters. This website/web server is visible at http://localhost:8000 .