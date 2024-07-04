# Liner

This fork adds minimal support for vi editing to
[Liner](https://github.com/peterh/liner). Projects that make use of Liner have
[users](https://github.com/open-policy-agent/opa/issues/1503) that may be
looking for [vi editing functionality](https://github.com/go-delve/delve/issues/3384). The modifications
in this repository offer basic vi editing behavior.

<details>
    <summary>code statistics</summary>

```
===============================================================================
 Language            Files        Lines         Code     Comments       Blanks
===============================================================================
 Go                     17         3211         2666          272          273
 Markdown                1           87            0           75           12
===============================================================================
 Total                  18         3298         2666          347          285
===============================================================================
```

</details>

We implement basic, single-action, vi editing functions. At present, we do not
support number keys to execute multiple operations, a repeat key, or undo actions.

Users may expect a method to visually discern between vi **normal** and **insert** modes.
Though we do not support changing cursor behavior, however, we can send terminal escape
codes change the appearance of text in the terminal.

## Vi Keys Supported

| key | function                                     |
|-----|----------------------------------------------|
| r   | replace character under cursor               |
| i   | enter insert mode                            |
| I   | enter insert mode at beginning of line       |
| a   | enter insert mode one character to the right |
| A   | enter insert mode at end of line             |
| ^ H | move cursor to the beginning of line         |
| $ L | move cursor to the end of line               |
| h   | move cursor one character to the left        |
| l   | move cursor one character to the right       |
| w   | move cursor to the next word                 |
| W   | move cursor to the next Word                 |
| e   | move cursor to the end of next word          |
| E   | move cursor to the end of next Word          |
| b   | move cursor to the previous word             |
| x   | delete the character under cursor            |
| X   | delete the character before cursor           |
| d   | delete the next word                         |
| D   | delete from cursor to end of line            |
| C   | change text from cursor to end of line       |
| j   | next matching history                        |
| k   | previous matching history                    |
| p   | paste from yank buffer                       |
| ~   | change character case                        |


<details>
    <summary>default keys supported</summary>

| key                    | function                                                     |
| ---------------------- | ------                                                       |
| Ctrl-A, Home           | move cursor to the beginning of line                         |
| Ctrl-E, End            | move cursor to the end of line                               |
| Ctrl-B, Left           | move cursor one character to the left                        |
| Ctrl-F, Right          | move cursor one character to the right                       |
| Ctrl-Left, Alt-B       | move cursor to the previous word                             |
| Ctrl-Right, Alt-F      | move cursor to the next word                                 |
| Ctrl-D, Del            | delete character under cursor if line is not empty           |
| Ctrl-D                 | end of file/exit if line is empty                            |
| Ctrl-C                 | abort input                                                  |
| Ctrl-L                 | clear screen                                                 |
| Ctrl-T                 | transpose previous character with current character          |
| Ctrl-H, BackSpace      | delete character before cursor                               |
| Ctrl-W, Alt-BackSpace  | delete word leading up to cursor                             |
| Alt-D                  | delete word following cursor                                 |
| Ctrl-K                 | delete from cursor to end of line                            |
| Ctrl-U                 | delete from start of line to cursor                          |
| Ctrl-P, Up             | previous match from history                                  |
| Ctrl-N, Down           | next match from history                                      |
| Ctrl-R                 | reverse search history (Ctrl-S forward, Ctrl-G cancel)       |
| Ctrl-Y                 | paste from yank buffer (Alt-Y to paste next yank instead)    |
| Tab                    | next completion                                              |
| Shift-Tab              | (after Tab) previous completion                              |

</details>

## Vi Functions

```
EnableViMode(bool)      // enable or disable vi editing functionality
EnableViPrompt(bool)    // allow liner prompt to change based on current vi mode
SetViMode(vimode)       // explicitly se the vi mode (normal, insert, replace)
SetViNormalStyle(style) // style used when in vi normal mode
SetViInsertStyle(style) // style used when in vi insert mode
```
