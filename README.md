<h1>Basic Go application</h1>
<p>This is a skeleton project for a Go application, which captures the popular and tried-and-tested approach build techniques.</p>
<p>Benefits of this structure:</p>
<ul>
<li>It gives a clean separation between Go and non-Go assets. All
the Go code we write will live exclusively under the cmd and
internal directories, leaving the project root free to hold non-Go
assets like UI files, makefiles and module definitions (including
our go.mod file). This can make things easier to manage when it
comes to building and deploying your application in the future.</li>
<li>It scales really nicely if you want to add another executable
application to your project. For example, you might want to add a
CLI (Command Line Interface) to automate some administrative
tasks in the future. With this structure, you could create this CLI
application under cmd/cli and it will be able to import and reuse
all the code you’ve written under the internal directory.</li>
<li>By creating a custom SnippetModel type and implementing
methods on it we’ve been able to make our model a single, neatly
encapsulated object, which we can easily initialize and then pass
to our handlers as a dependency. Again, this makes for easier to
maintain, testable code.</li>
<li>Because the model actions are defined as methods on an object
— in our case SnippetModel — there’s the opportunity to create
an interface and mock it for unit testing purposes.</li>
<li>Total control over which database is used at
runtime, just by using the -dsn command-line flag.</li>
</ul>
