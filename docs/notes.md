# Notes

This is a collection of notes I am taking on this project to remember the things I set to later and not struggle multiple times on the same bug because I forgot to write how I solved my bug. 

## UTF-8 incident

Basically the problem is me myself and I. I am really dumb and don't know how to read...

Here is the prototype for `loadFontEx`

```c
Font LoadFontEx(const char *fileName, int fontSize, int *codepoints, int codepointCount);
```

And here is how I use it

```c
rl.LoadFontEx(*settings.UI.FontFamily, 100, nil),
```

See it now ? The number of arguments for both functions is not the same. I don't use `codepointCount` because why should I use something I don't understand right ? Rather than reading what a codepoint is, I just said fuck that, I won't put it in. Later struggling understanding why I cannot use non-ascii chars. 

So for myself later, a codepoint is basically a fancy way to say the texture for a char in a font. So `codepointCount` means the number of chars I want to use. If I don't put it, it will load only ascii and replace everything else with question marks. 

Now why aren't I just implementing complete UTF-8 in my text editor ? In golang, `strings` are weird and made of runes. However, when I want to get the length of a string, with `len()` it gives the number of chars and `chars` != `runes` so all my calculations are wrong and I need to replace `len()` everywhere. Moreover, UTF-8 some chars are bigger or smaller and this could be a big problem and need some debugging. UTF-8 support will be needed but right now(28.11.2024), it is not a priority. I also want a nice way of finding how many codepoints I want to import and this seems not so fun. 

So now remember, everything is here for a reason. So if I don't get something which doesn't seem important, guess what IT FUCKING IS IMPORTANT. 
