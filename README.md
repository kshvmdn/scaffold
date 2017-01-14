## scaffold

Bootstrap full directories with text-based template files.

### Installation

  ```
  $ go get github.com/kshvmdn/scaffold
  $ scaffold --config=<PATH/TO/CONFIG_FILE> --directory=<PATH/TO/DIRECTORY>
  ```

  - You can also build from source if you prefer that:

  ```sh
  $ git clone https://github.com/kshvmdn/scaffold.git
  $ go build scaffold.go
  $ ./scaffold --config=<PATH/TO/CONFIG_FILE> --directory=<PATH/TO/DIRECTORY>
  ```

### Usage

  - Run `--help` for a help menu.

    ```sh
    $ scaffold --help
    $ Usage of scaffold:
      -config string
          Config file
      -directory string
          Root directory (default ".")
    ```

  - Provide the script with a path to your configuration file (the directory structure, see `school.txt` below) and an optional directory name. The script will print a line that you'll have to copy and paste into your shell (if you're interested in why, see [#1](https://github.com/kshvmdn/scaffold/issues/1), _help wanted with this!_).

    ```sh
    $ scaffold --config=school.txt --directory=~/Desktop/csc263
    mkdir -p ~/Desktop/csc263/{assignments/{a1,a2,a3,a4,tests/{t1,t2,},exam,misc/{textbook,previous-offerings,},},}; touch ~/Desktop/csc263/{assignments/{a1/{a1.tex,},a2/{a2.tex,},a3/{a3.tex,},a4/{a4.tex,},tests/{,,},exam/{coverage.md,},misc/{,,},},}
    ```

  - After copying and pasting the above:

    ```sh
    $ tree ~/Desktop/csc263
    /Users/kashavmadan/Desktop/csc263/
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
    │   └── textbook
    └── tests
        ├── t1
        └── t2

    12 directories, 5 files
    ```

  - The `school.txt` file that was used:

    ```txt
    assignments/
      a1/
        a1.tex
      a2/
        a2.tex
      a3/
        a3.tex
      a4/
        a4.tex
    tests/
      t1/
      t2/
    exam/
      coverage.md
    misc/
      textbook/
      previous-offerings/
    ```

#### Structure Specifications

  - Use a line-break and 2 **spaces** to represent directory contents.

### Contribute

  This project is completely open source. Feel free to open an issue or submit a pull request.
