# go-jump

`go-jump` is a cli utility to jump to recently visited directories.
It hooks into your shell to count all visitied directories.

The database structure is very simple:

```sh
Counter     Path
1337        /home/me/Code
```

## Usage
Quick integration in `zsh`:

```sh
function j() {
    if [ "$1" != "" ]
    then
        foundPath=`go-jump $1`
        cd $foundPath
    else
        echo "provide path"
    fi
}

chpwd_go_jump() {
    go-jump add $PWD
}
add-zsh-hook chpwd chpwd_go_jump
```

After that simply use `j documents` to search for `documents` and jump into the folder if
it already exists in the database.
