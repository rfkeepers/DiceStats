# DiceStats
A collection of roll result probabilities for a variety of dice rolling customs.

This repo is actually a pretty shameful exhibit of hacky code that lacks documentation, lacks testing, showcases terrible patterns of duplication with little or no mind to elegant reuse, is barely readable, and has no care for performance.  Hell, you can't even compile the whole thing, because they all contain a main(), among other collisions; you have to run each file at a time.

To give that some justification, perhaps it would make more sense for you to know that I wrote each of these at the drop of a hat on play.golang.org, usually when I had a sudden curiosity to explore a given rolling mechanic.  It was only later that I decided to port my history of playground links into a single repo.  The code isn't here to show off my coding capabilities.  This is just an archive of something that I infrequently use when testing designs.