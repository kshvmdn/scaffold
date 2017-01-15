## scaffold

Create full directories using text-based template files.

### Installation

  - Use the following to install and run (assumes you have Go [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/install#testing)):

  ```sh
  $ go get github.com/kshvmdn/scaffold
  $ scaffold -template=<PATH/TO/TEMPLATE_FILE.txt> -directory=<PATH/TO/DIRECTORY>
  ```

  - You can also build from source if you prefer that:

  ```sh
  $ git clone https://github.com/kshvmdn/scaffold.git
  $ cd scaffold
  $ go build scaffold.go
  $ ./scaffold -template=<PATH/TO/TEMPLATE_FILE.txt> -directory=<PATH/TO/DIRECTORY>
  ```

### Usage

  - Run `--help` for the help menu.

    ```sh
    $ scaffold --help
    $ Usage of scaffold:
      -config string
          Config file
      -directory string
          Root directory (default ".")
    ```

  - Provide the script with a template file (the directory structure, see [`example.txt`](example.txt)) and an optional directory name.

    ```sh
    $ scaffold -template=example.txt -directory=~/Desktop/course
    ```

  - After running the above:

    ```sh
    $ tree ~/Desktop/course
    /Users/kashavmadan/Desktop/course/
    ├── assignments
    │   ├── a1
    │   │   └── a1.tex
    │   ├── a2
    │   │   └── a2.tex
    │   ├── a3
    │   │   └── a3.tex
    │   └── a4
    │       └── a4.tex
    ├── exam
    │   └── coverage.md
    ├── misc
    │   ├── previous-offerings
    │   │   ├── 2015
    │   │   └── 2016
    │   │       └── 2016_exam.docx
    │   └── textbook
    └── tests
        ├── t1
        └── t2

    14 directories, 6 files
    ```

#### Structure Specifications

  - Use a line-break and two **spaces** to indicate directory subcontents (each pair of two spaces represents a single level of nesting).
  - Directory names must end with a single forward slash (`/`). Files can be named anything.

### Contribute

This project is completely open source. Feel free to open an issue or submit a pull request.
