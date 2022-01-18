# go-wordle-solver

Helps solve [wordle](https://www.powerlanguage.co.uk/wordle/) puzzles with brute force
by using [this English word frequency list](https://www.kaggle.com/rtatman/english-word-frequency).

Tip: You can [do past wordles on this site](https://metzger.media/games/wordle-archive/?levels=select).

## Usage

Currently non-interactive. You have to change [main.go](main.go) to set the `initialConditions`
to match the conditions you've learned by trying guesses.

I suggest starting with `rates`.

When the hints come back, edit `main.go` to hardcode what you learned into `initialConditions`, then:

```
go run main.go
```

It will print another guess. Enter that on the wordle site, add the new conditions to the code,
and repeat the process as needed.

## TODO

* Make it interactive so you don't have to change the code each time like some sort of caveman
* Allow user to see top 5 next words to try, and their frequency expressed as a percentage
